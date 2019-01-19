package txbuilder

import (
	"fmt"
	"github.com/baymax19/js2go/cosmos-sdk/types"
	"github.com/baymax19/js2go/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/crypto/keys/hd"
	"github.com/cosmos/go-bip39"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"strings"
)

const (
	defaultBIP39Passphrase = ""
)

type BaseReq struct {
	TxEncoder     types.TxEncoder
	AccountNumber uint64 `json:"account_number"`
	Sequence      uint64 `json:"sequence"`
	Gas           uint64 `json:"gas"`
	ChainID       string `json:"chain_id"`
	Memo          string `json:"memo"`
	Fee           string `json:"fee"`
}

func NewBaseReq(accountNumber, sequence, gas uint64, chainID, memo, fee string) *BaseReq {
	return &BaseReq{
		AccountNumber: accountNumber,
		Sequence:      sequence,
		Gas:           gas,
		ChainID:       chainID,
		Memo:          memo,
		Fee:           fee,
	}
}

func (bldr BaseReq) WithTxEncoder(txEncoder types.TxEncoder) BaseReq {
	bldr.TxEncoder = txEncoder
	return bldr
}

func (bldr BaseReq) BuildAndSign(seed string, msgs []types.Msg) ([]byte, error) {

	msg, err := bldr.Build(msgs)
	if err != nil {
		return nil, err
	}

	return bldr.MakeSignUsingSeed(seed, msg)
}
func (bldr BaseReq) Build(msgs []types.Msg) (StdSignMsg, error) {
	chainID := bldr.ChainID
	if chainID == "" {
		return StdSignMsg{}, errors.Errorf("chain ID required but not specified")
	}

	fee := types.Coin{}
	if bldr.Fee != "" {
		parsedFee, err := types.ParseCoin(bldr.Fee)
		if err != nil {
			return StdSignMsg{}, err
		}

		fee = parsedFee
	}
	return StdSignMsg{
		ChainID:       bldr.ChainID,
		AccountNumber: bldr.AccountNumber,
		Sequence:      bldr.Sequence,
		Memo:          bldr.Memo,
		Msgs:          msgs,
		Fee:           auth.NewStdFee(bldr.Gas, fee),
	}, nil
}

func (bldr BaseReq) MakeSignUsingSeed(mnemonic string, msg StdSignMsg) ([]byte, error) {
	sign, err := makeSign(mnemonic, msg)
	if err != nil {
		return nil, err
	}

	stdTx := auth.NewStdTx(msg.Msgs, msg.Fee, []auth.StdSignature{sign}, msg.Memo)

	return bldr.TxEncoder(stdTx)
}

func makeSign(mnemonic string, msg StdSignMsg) (auth.StdSignature, error) {

	words := strings.Split(mnemonic, " ")

	if len(words) != 12 && len(words) != 24 {
		err := fmt.Errorf("recovering only works with 12 word (fundraiser) or 24 word mnemonics, got: %v words", len(words))
		return auth.StdSignature{}, err
	}

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, defaultBIP39Passphrase)
	if err != nil {
		return auth.StdSignature{}, err
	}

	masterPriv, ch := hd.ComputeMastersFromSeed(seed)

	derivedPriv, err := hd.DerivePrivateKeyForPath(masterPriv, ch, hd.FullFundraiserPath)
	if err != nil {
		panic(err)
	}
	privKey := secp256k1.PrivKeySecp256k1(derivedPriv)

	sigBytes, err := privKey.Sign(msg.Bytes())
	if err != nil {
		return auth.StdSignature{}, nil
	}

	return auth.StdSignature{
		PubKey:    secp256k1.PrivKeySecp256k1(derivedPriv).PubKey(),
		Signature: sigBytes,
	}, nil
}

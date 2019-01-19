package keybase

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/mintkey"
	"github.com/cosmos/go-bip39"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"strings"
)

func CreateKey(name, password, mnemonic string) (info Info, err error) {

	words := strings.Split(mnemonic, " ")
	if len(words) != 12 && len(words) != 24 {
		err = fmt.Errorf("recovering only works with 12 word (fundraiser) or 24 word mnemonics, got: %v words", len(words))
		return
	}

	seed, err := bip39.NewSeedWithErrorChecking(mnemonic, defaultBIP39Passphrase)
	if err != nil {
		return
	}

	info, err = persistDerivedKey(seed, password, name, hd.FullFundraiserPath)

	return
}

func persistDerivedKey(seed []byte, password, name, path string) (info Info, err error) {

	masterPriv, ch := hd.ComputeMastersFromSeed(seed)

	derivedPriv, err := hd.DerivePrivateKeyForPath(masterPriv, ch, path)
	if err != nil {
		panic(err)
	}

	info = writeLocalKey(secp256k1.PrivKeySecp256k1(derivedPriv), name, password)

	return
}

func writeLocalKey(priv tmcrypto.PrivKey, name, passpharse string) Info {

	privKeyArmor := mintkey.EncryptArmorPrivKey(priv, passpharse)
	pubKey := priv.PubKey()
	info := newLocalInfo(name, pubKey, privKeyArmor)

	return info
}

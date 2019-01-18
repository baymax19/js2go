package types

import (
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/encoding/amino"
	"github.com/tendermint/tmlibs/bech32"
)

type AccAddress []byte

func AccAddressFromBech32(address string) (AccAddress, error) {
	bz, err := GetFromBech32(address, "cosmos")
	if err != nil {
		panic(err)
	}
	return AccAddress(bz), nil
}

func (aa AccAddress) String() string {
	bech32Str, err := bech32.ConvertAndEncode("cosmos", aa.Bytes())
	if err != nil {
		panic(err)
	}
	return bech32Str
}

func (aa AccAddress) Bytes() []byte {
	return aa
}

func PubKeyFromBytes(pubkey crypto.PubKey) string {

	PubkeyString, err := bech32.ConvertAndEncode("cosmospub", pubkey.Bytes())
	if err != nil {
		panic(err)
	}
	return PubkeyString
}

func PubKeyFromBech32String(pubkey string) crypto.PubKey {
	bz, err := GetFromBech32(pubkey, "cosmospub")
	if err != nil {
		panic(err)
	}

	pubKey, err := cryptoAmino.PubKeyFromBytes(bz)
	if err != nil {
		panic(err)
	}
	return pubKey
}

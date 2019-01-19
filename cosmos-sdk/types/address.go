package types

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (aa AccAddress) Equals(aa2 AccAddress) bool {
	if aa.Empty() && aa2.Empty() {
		return true
	}

	return bytes.Compare(aa.Bytes(), aa2.Bytes()) == 0
}

func (aa AccAddress) Empty() bool {
	if aa == nil {
		return true
	}

	aa2 := AccAddress{}
	return bytes.Compare(aa.Bytes(), aa2.Bytes()) == 0
}

func (aa AccAddress) Marshal() ([]byte, error) {
	return aa, nil
}

func (aa *AccAddress) Unmarshal(data []byte) error {
	*aa = data
	return nil
}

func (aa AccAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(aa.String())
}

func (aa *AccAddress) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}

	aa2, err := AccAddressFromBech32(s)
	if err != nil {
		return err
	}

	*aa = aa2
	return nil
}

func (aa AccAddress) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(fmt.Sprintf("%s", aa.String())))
	case 'p':
		s.Write([]byte(fmt.Sprintf("%p", aa)))
	default:
		s.Write([]byte(fmt.Sprintf("%X", []byte(aa))))
	}
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

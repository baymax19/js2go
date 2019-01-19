package keys

import (
	"github.com/baymax19/js2go/cosmos-sdk/crypto/keybase"
	"github.com/baymax19/js2go/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	"github.com/gopherjs/gopherjs/js"
)

const (
	defaultEntropySize = 256
)

type KeyOutput struct {
	*js.Object
	Name    string `js:"name"`
	Address string `js:"address"`
	PubKey  string `js:"pub_key"`
	Seed    string `js:"seed"`
}

func CreateKey(name, password string) *js.Object {

	entropy, err := bip39.NewEntropy(defaultEntropySize)
	if err != nil {
		panic(err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy[:])
	if err != nil {
		panic(err)
	}

	info, err := keybase.CreateKey(name, password, mnemonic)
	if err != nil {
		panic(err)
	}

	return writeInfo(info, mnemonic)
}

func writeInfo(info keybase.Info, mnemonic string) *js.Object {

	data := &KeyOutput{Object: js.Global.Get("Object").New()}

	data.Address = types.AccAddress(info.GetAddress()).String()
	data.PubKey = types.PubKeyFromBytes(info.GetPubKey())
	data.Name = info.GetName()
	data.Seed = mnemonic

	return data.Object
}

package keybase

import (
	"github.com/baymax19/js2go/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

const (
	defaultBIP39Passphrase = ""
)

type Info interface {
	GetType() string
	GetAddress() types.AccAddress
	GetPubKey() crypto.PubKey
	GetName() string
}

type localInfo struct {
	Name         string        `json:"name"`
	PubKey       crypto.PubKey `json:"pub_key"`
	PrivKeyArmor string        `json:"priv_key_armor"`
}

var _ Info = &localInfo{}

func newLocalInfo(name string, pubKey crypto.PubKey, privKeyArmor string) Info {
	return &localInfo{
		Name:         name,
		PubKey:       pubKey,
		PrivKeyArmor: privKeyArmor,
	}
}

func (info *localInfo) GetType() string { return "local" }

func (info *localInfo) GetAddress() types.AccAddress { return info.GetPubKey().Address().Bytes() }

func (info *localInfo) GetPubKey() crypto.PubKey { return info.PubKey }

func (info *localInfo) GetName() string { return info.Name }

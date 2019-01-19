package types

import "github.com/baymax19/js2go/codec"

var Cdc = codec.New()

func init() {
	codec.RegisterCrypto(Cdc)
}

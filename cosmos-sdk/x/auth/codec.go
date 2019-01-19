package auth

import "github.com/baymax19/js2go/codec"

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(StdTx{}, "auth/StdTx", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
	codec.RegisterCrypto(msgCdc)
}

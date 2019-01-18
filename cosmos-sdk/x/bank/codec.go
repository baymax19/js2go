package bank

import (
	"github.com/baymax19/js2go/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSend{}, "cosmos-sdk/Send", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}

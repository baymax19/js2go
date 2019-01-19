package client

import (
	"github.com/baymax19/js2go/codec"
	"github.com/baymax19/js2go/cosmos-sdk/types"
	"github.com/baymax19/js2go/cosmos-sdk/x/auth"
)

func GetTxEncoder(cdc *codec.Codec) (encoder types.TxEncoder) {
	if encoder == nil {
		encoder = auth.DefaultTxEncoder(cdc)
	}
	return
}

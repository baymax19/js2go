package main

import (
	"github.com/baymax19/js2go/cosmos-sdk/client/keys"
	"github.com/baymax19/js2go/cosmos-sdk/types"
	"github.com/baymax19/js2go/cosmos-sdk/x/auth"
	"github.com/baymax19/js2go/cosmos-sdk/x/bank"
	"github.com/baymax19/js2go/cosmos-sdk/x/bank/cli"
	jtypes "github.com/baymax19/js2go/types"
	"github.com/gopherjs/gopherjs/js"
)

var cdc = jtypes.Cdc

func main() {
	auth.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	types.RegisterCodec(cdc)


	js.Module.Get("exports").Set("createKey", keys.CreateKey)
	js.Module.Get("exports").Set("sendCoins", cli.SendCoins)
}

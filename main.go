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


	//seed := "sound coral chimney claim humor peasant reward vanish desk trouble army door shallow insect fence typical ice tonight change dust reduce bracket ancient embark"
	//data := cli.SendCoins("cosmos1v0m40792sx0cf69elugcqqxmqg3rdy7ra0j9kl", "cosmos1h7w6g8k8d2qflesnzyqap50zvldzvtgdmk7a6t", "1STAKE", seed)
	//fmt.Println(data)

}

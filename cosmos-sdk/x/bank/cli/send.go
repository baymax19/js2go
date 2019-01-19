package cli

import (
	"encoding/base64"
	"github.com/baymax19/js2go/cosmos-sdk/types"
	"github.com/baymax19/js2go/cosmos-sdk/x/auth"
	"github.com/baymax19/js2go/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/baymax19/js2go/cosmos-sdk/x/bank"
	jscodec "github.com/baymax19/js2go/types"
)

func SendCoins(from, to, amount, seed string) string {

	fromAddr, err := types.AccAddressFromBech32(from)
	if err != nil {
		panic(err)
	}

	toAddr, err := types.AccAddressFromBech32(to)
	if err != nil {
		panic(err)
	}

	coins, err := types.ParseCoins(amount)
	if err != nil {
		panic(err)
	}

	msg := bank.CreateMsg(fromAddr, toAddr, coins)
	baseReq := txbuilder.NewBaseReq(2, 6, 200000, "sentinel-vpn", "", "0STAKE").WithTxEncoder(auth.DefaultTxEncoder(jscodec.Cdc))

	txBytes, err := baseReq.BuildAndSign(seed, []types.Msg{msg})
	if err != nil {
		panic(err)
	}

	data := base64.StdEncoding.EncodeToString(txBytes)
	return data
}

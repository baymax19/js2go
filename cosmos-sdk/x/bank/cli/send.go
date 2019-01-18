package cli

import (
	"github.com/baymax19/js2go/cosmos-sdk/types"
	"github.com/baymax19/js2go/cosmos-sdk/x/bank"
	"github.com/gopherjs/gopherjs/js"
)

func SendCoins(from, to, amount, seed string) *js.Object {

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
	return &js.Object{}
}

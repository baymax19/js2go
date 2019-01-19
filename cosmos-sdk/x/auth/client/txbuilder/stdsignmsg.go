package txbuilder

import (
	"github.com/baymax19/js2go/cosmos-sdk/types"
	"github.com/baymax19/js2go/cosmos-sdk/x/auth"
)

type StdSignMsg struct {
	ChainID       string      `json:"chain_id"`
	AccountNumber uint64      `json:"account_number"`
	Sequence      uint64      `json:"sequence"`
	Fee           auth.StdFee `json:"fee"`
	Msgs          []types.Msg `json:"msgs"`
	Memo          string      `json:"memo"`
}

func (msg StdSignMsg) Bytes() []byte {
	return auth.StdSignBytes(msg.ChainID, msg.AccountNumber, msg.Sequence, msg.Fee, msg.Msgs, msg.Memo)
}

package bank

import (
	"encoding/json"
	"github.com/baymax19/js2go/cosmos-sdk/types"
)

const MsgRoute = "bank"

type MsgSend struct {
	Inputs  []Input  `json:"inputs"`
	Outputs []Output `json:"outputs"`
}

func NewMsgSend(in []Input, out []Output) MsgSend {
	return MsgSend{Inputs: in, Outputs: out}
}

func CreateMsg(from, to types.AccAddress, coins types.Coins) types.Msg {
	input := NewInput(from, coins)
	output := NewOutput(to, coins)
	msg := NewMsgSend([]Input{input}, []Output{output})
	return msg
}

func (msg MsgSend) Route() string { return MsgRoute }
func (msg MsgSend) Type() string  { return "send" }

func (msg MsgSend) ValidateBasic() error { return nil }

func (msg MsgSend) GetSignBytes() []byte {
	var inputs, outputs []json.RawMessage
	for _, input := range msg.Inputs {
		inputs = append(inputs, input.GetSignBytes())
	}
	for _, output := range msg.Outputs {
		outputs = append(outputs, output.GetSignBytes())
	}
	b, err := msgCdc.MarshalJSON(struct {
		Inputs  []json.RawMessage `json:"inputs"`
		Outputs []json.RawMessage `json:"outputs"`
	}{
		Inputs:  inputs,
		Outputs: outputs,
	})
	if err != nil {
		panic(err)
	}
	return b
}

func (msg MsgSend) GetSigners() []types.AccAddress {
	addrs := make([]types.AccAddress, len(msg.Inputs))
	for i, in := range msg.Inputs {
		addrs[i] = in.Address
	}
	return addrs
}

type Input struct {
	Address types.AccAddress `json:"address"`
	Coins   types.Coins      `json:"coins"`
}

func NewInput(addr types.AccAddress, coins types.Coins) Input {
	input := Input{
		Address: addr,
		Coins:   coins,
	}
	return input
}

func (in Input) GetSignBytes() []byte {
	bin, err := msgCdc.MarshalJSON(in)
	if err != nil {
		panic(err)
	}
	return bin
}

func (in Input) ValidateBasic() error { return nil }

type Output struct {
	Address types.AccAddress `json:"address"`
	Coins   types.Coins      `json:"coins"`
}

func NewOutput(addr types.AccAddress, coins types.Coins) Output {
	output := Output{
		Address: addr,
		Coins:   coins,
	}
	return output
}

func (out Output) GetSignBytes() []byte {
	bin, err := msgCdc.MarshalJSON(out)
	if err != nil {
		panic(err)
	}
	return bin
}

func (out Output) ValidateBasic() error { return nil }

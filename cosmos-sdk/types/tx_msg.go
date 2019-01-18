package types

type Msg interface {
	Route() string

	Type() string

	ValidateBasic() error

	GetSignBytes() []byte

	GetSigners() []AccAddress
}

type Tx interface {
	GetMsgs() []Msg

	ValidateBasic() error
}

type TxDecoder func(txBytes []byte) (Tx, err error)

type TxEncoder func(tx Tx) ([]byte, error)

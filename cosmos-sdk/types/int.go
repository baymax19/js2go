package types

import (
	"math/big"
	"math/rand"
)

// Int wraps integer with 256 bit range bound
// Checks overflow, underflow and division by zero
// Exists in range from -(2^255-1) to 2^255-1
func newIntegerFromString(s string) (*big.Int, bool) {
	return new(big.Int).SetString(s, 0)
}

func equal(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) == 0 }

func gt(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) == 1 }

func lt(i *big.Int, i2 *big.Int) bool { return i.Cmp(i2) == -1 }

func add(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Add(i, i2) }

func sub(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Sub(i, i2) }

func mul(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Mul(i, i2) }

func div(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Div(i, i2) }

func mod(i *big.Int, i2 *big.Int) *big.Int { return new(big.Int).Mod(i, i2) }

func neg(i *big.Int) *big.Int { return new(big.Int).Neg(i) }

func random(i *big.Int) *big.Int { return new(big.Int).Rand(rand.New(rand.NewSource(rand.Int63())), i) }

type Int struct {
	i *big.Int
}

// BigInt converts Int to big.Int
func (i Int) BigInt() *big.Int {
	return new(big.Int).Set(i.i)
}

// NewInt constructs Int from int64
func NewInt(n int64) Int {
	return Int{big.NewInt(n)}
}

// NewIntFromBigInt constructs Int from big.Int
func NewIntFromBigInt(i *big.Int) Int {
	if i.BitLen() > 255 {
		panic("NewIntFromBigInt() out of bound")
	}
	return Int{i}
}

// NewIntFromString constructs Int from string
func NewIntFromString(s string) (res Int, ok bool) {
	i, ok := newIntegerFromString(s)
	if !ok {
		return
	}
	// Check overflow
	if i.BitLen() > 255 {
		ok = false
		return
	}
	return Int{i}, true
}

// ZeroInt returns Int value with zero
func ZeroInt() Int { return Int{big.NewInt(0)} }

// Int64 converts Int to int64
// Panics if the value is out of range
func (i Int) Int64() int64 {
	if !i.i.IsInt64() {
		panic("Int64() out of bound")
	}
	return i.i.Int64()
}

// IsInt64 returns true if Int64() not panics
func (i Int) IsInt64() bool {
	return i.i.IsInt64()
}

// IsZero returns true if Int is zero
func (i Int) IsZero() bool {
	return i.i.Sign() == 0
}

// Sign returns sign of Int
func (i Int) Sign() int {
	return i.i.Sign()
}

// Equal compares two Ints
func (i Int) Equal(i2 Int) bool {
	return equal(i.i, i2.i)
}

// GT returns true if first Int is greater than second
func (i Int) GT(i2 Int) bool {
	return gt(i.i, i2.i)
}

// LT returns true if first Int is lesser than second
func (i Int) LT(i2 Int) bool {
	return lt(i.i, i2.i)
}

// Add adds Int from another
func (i Int) Add(i2 Int) (res Int) {
	res = Int{add(i.i, i2.i)}
	// Check overflow
	if res.i.BitLen() > 255 {
		panic("Int overflow")
	}
	return
}

// AddRaw adds int64 to Int
func (i Int) AddRaw(i2 int64) Int {
	return i.Add(NewInt(i2))
}

// Sub subtracts Int from another
func (i Int) Sub(i2 Int) (res Int) {
	res = Int{sub(i.i, i2.i)}
	// Check overflow
	if res.i.BitLen() > 255 {
		panic("Int overflow")
	}
	return
}

// SubRaw subtracts int64 from Int
func (i Int) SubRaw(i2 int64) Int {
	return i.Sub(NewInt(i2))
}

// Mul multiples two Ints
func (i Int) Mul(i2 Int) (res Int) {
	// Check overflow
	if i.i.BitLen()+i2.i.BitLen()-1 > 255 {
		panic("Int overflow")
	}
	res = Int{mul(i.i, i2.i)}
	// Check overflow if sign of both are same
	if res.i.BitLen() > 255 {
		panic("Int overflow")
	}
	return
}

// MulRaw multipies Int and int64
func (i Int) MulRaw(i2 int64) Int {
	return i.Mul(NewInt(i2))
}

// Div divides Int with Int
func (i Int) Div(i2 Int) (res Int) {
	// Check division-by-zero
	if i2.i.Sign() == 0 {
		panic("Division by zero")
	}
	return Int{div(i.i, i2.i)}
}

// DivRaw divides Int with int64
func (i Int) DivRaw(i2 int64) Int {
	return i.Div(NewInt(i2))
}

// Mod returns remainder after dividing with Int
func (i Int) Mod(i2 Int) Int {
	if i2.Sign() == 0 {
		panic("division-by-zero")
	}
	return Int{mod(i.i, i2.i)}
}

// ModRaw returns remainder after dividing with int64
func (i Int) ModRaw(i2 int64) Int {
	return i.Mod(NewInt(i2))
}

// Neg negates Int
func (i Int) Neg() (res Int) {
	return Int{neg(i.i)}
}

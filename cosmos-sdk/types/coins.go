package types

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Coin struct {
	Denom  string `json:"denom"`
	Amount Int    `json:"amount"`
}

func NewCoin(denom string, amount Int) Coin {
	if amount.LT(ZeroInt()) {
		panic(fmt.Sprintf("negative coin amount: %v\n", amount))
	}

	return Coin{
		Denom:  denom,
		Amount: amount,
	}
}

func NewInt64Coin(denom string, amount int64) Coin { return NewCoin(denom, NewInt(amount)) }

func (coin Coin) String() string { return fmt.Sprintf("%v%v", coin.Amount, coin.Denom) }

func (coin Coin) SameDenomAs(other Coin) bool { return coin.Denom == other.Denom }

func (coin Coin) IsZero() bool { return coin.Amount.IsZero() }

func (coin Coin) IsGTE(other Coin) bool { return coin.SameDenomAs(other) && (!coin.Amount.LT(other.Amount)) }

func (coin Coin) IsLT(other Coin) bool { return coin.SameDenomAs(other) && coin.Amount.LT(other.Amount) }

func (coin Coin) IsEqual(other Coin) bool { return coin.SameDenomAs(other) && (coin.Amount.Equal(other.Amount)) }

func (coin Coin) Plus(coinB Coin) Coin {
	if !coin.SameDenomAs(coinB) {
		panic(fmt.Sprintf("invalid coin denominations; %s, %s", coin.Denom, coinB.Denom))
	}

	return Coin{coin.Denom, coin.Amount.Add(coinB.Amount)}
}

func (coin Coin) Minus(coinB Coin) Coin {
	if !coin.SameDenomAs(coinB) {
		panic(fmt.Sprintf("invalid coin denominations; %s, %s", coin.Denom, coinB.Denom))
	}

	res := Coin{coin.Denom, coin.Amount.Sub(coinB.Amount)}
	if !res.IsNotNegative() {
		panic("negative count amount")
	}

	return res
}

func (coin Coin) IsPositive() bool { return coin.Amount.Sign() == 1 }

func (coin Coin) IsNotNegative() bool { return coin.Amount.Sign() != -1 }

type Coins []Coin

func (coins Coins) String() string {
	if len(coins) == 0 {
		return ""
	}

	out := ""
	for _, coin := range coins {
		out += fmt.Sprintf("%v,", coin.String())
	}
	return out[:len(out)-1]
}

func (coins Coins) IsValid() bool {
	switch len(coins) {
	case 0:
		return true
	case 1:
		return coins[0].IsPositive()
	default:
		if !(Coins{coins[0]}).IsValid() {
			return false
		}

		lowDenom := coins[0].Denom
		for _, coin := range coins[1:] {
			if coin.Denom <= lowDenom {
				return false
			}
			if !coin.IsPositive() {
				return false
			}

			lowDenom = coin.Denom
		}

		return true
	}
}

func (coins Coins) Plus(coinsB Coins) Coins { return coins.safePlus(coinsB) }

func (coins Coins) safePlus(coinsB Coins) Coins {
	sum := ([]Coin)(nil)
	indexA, indexB := 0, 0
	lenA, lenB := len(coins), len(coinsB)

	for {
		if indexA == lenA {
			if indexB == lenB {
				return sum
			}

			return append(sum, removeZeroCoins(coinsB[indexB:])...)
		} else if indexB == lenB {
			return append(sum, removeZeroCoins(coins[indexA:])...)
		}

		coinA, coinB := coins[indexA], coinsB[indexB]

		switch strings.Compare(coinA.Denom, coinB.Denom) {
		case -1: // coin A denom < coin B denom
			if !coinA.IsZero() {
				sum = append(sum, coinA)
			}

			indexA++

		case 0: // coin A denom == coin B denom
			res := coinA.Plus(coinB)
			if !res.IsZero() {
				sum = append(sum, res)
			}

			indexA++
			indexB++

		case 1: // coin A denom > coin B denom
			if !coinB.IsZero() {
				sum = append(sum, coinB)
			}

			indexB++
		}
	}
}

func (coins Coins) Minus(coinsB Coins) Coins {
	diff, hasNeg := coins.SafeMinus(coinsB)
	if hasNeg {
		panic("negative coin amount")
	}

	return diff
}

func (coins Coins) SafeMinus(coinsB Coins) (Coins, bool) {
	diff := coins.safePlus(coinsB.negative())
	return diff, !diff.IsNotNegative()
}

func (coins Coins) IsAllGT(coinsB Coins) bool {
	diff, _ := coins.SafeMinus(coinsB)
	if len(diff) == 0 {
		return false
	}

	return diff.IsPositive()
}

func (coins Coins) IsAllGTE(coinsB Coins) bool {
	diff, _ := coins.SafeMinus(coinsB)
	if len(diff) == 0 {
		return true
	}

	return diff.IsNotNegative()
}

func (coins Coins) IsAllLT(coinsB Coins) bool { return coinsB.IsAllGT(coins) }

func (coins Coins) IsAllLTE(coinsB Coins) bool { return coinsB.IsAllGTE(coins) }

func (coins Coins) IsZero() bool {
	for _, coin := range coins {
		if !coin.IsZero() {
			return false
		}
	}
	return true
}

func (coins Coins) IsEqual(coinsB Coins) bool {
	if len(coins) != len(coinsB) {
		return false
	}

	coins = coins.Sort()
	coinsB = coinsB.Sort()

	for i := 0; i < len(coins); i++ {
		if coins[i].Denom != coinsB[i].Denom || !coins[i].Amount.Equal(coinsB[i].Amount) {
			return false
		}
	}

	return true
}

func (coins Coins) Empty() bool { return len(coins) == 0 }

func (coins Coins) AmountOf(denom string) Int {
	switch len(coins) {
	case 0:
		return ZeroInt()

	case 1:
		coin := coins[0]
		if coin.Denom == denom {
			return coin.Amount
		}
		return ZeroInt()

	default:
		midIdx := len(coins) / 2 // 2:1, 3:1, 4:2
		coin := coins[midIdx]

		if denom < coin.Denom {
			return coins[:midIdx].AmountOf(denom)
		} else if denom == coin.Denom {
			return coin.Amount
		} else {
			return coins[midIdx+1:].AmountOf(denom)
		}
	}
}

func (coins Coins) IsPositive() bool {
	if len(coins) == 0 {
		return false
	}

	for _, coin := range coins {
		if !coin.IsPositive() {
			return false
		}
	}

	return true
}

func (coins Coins) IsNotNegative() bool {
	if len(coins) == 0 {
		return true
	}

	for _, coin := range coins {
		if !coin.IsNotNegative() {
			return false
		}
	}

	return true
}

func (coins Coins) negative() Coins {
	res := make([]Coin, 0, len(coins))

	for _, coin := range coins {
		res = append(res, Coin{
			Denom:  coin.Denom,
			Amount: coin.Amount.Neg(),
		})
	}

	return res
}

func removeZeroCoins(coins Coins) Coins {
	i, l := 0, len(coins)
	for i < l {
		if coins[i].IsZero() {
			coins = append(coins[:i], coins[i+1:]...)
			l--
		} else {
			i++
		}
	}

	return coins[:i]
}

func (coins Coins) Len() int           { return len(coins) }
func (coins Coins) Less(i, j int) bool { return coins[i].Denom < coins[j].Denom }
func (coins Coins) Swap(i, j int)      { coins[i], coins[j] = coins[j], coins[i] }

var _ sort.Interface = Coins{}

func (coins Coins) Sort() Coins {
	sort.Sort(coins)
	return coins
}

var (
	reDnm  = `[[:alpha:]][[:alnum:]]{2,15}`
	reAmt  = `[[:digit:]]+`
	reSpc  = `[[:space:]]*`
	reCoin = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reAmt, reSpc, reDnm))
)

func ParseCoin(coinStr string) (coin Coin, err error) {
	coinStr = strings.TrimSpace(coinStr)

	matches := reCoin.FindStringSubmatch(coinStr)
	if matches == nil {
		return Coin{}, fmt.Errorf("invalid coin expression: %s", coinStr)
	}

	denomStr, amountStr := matches[2], matches[1]

	amount, ok := NewIntFromString(amountStr)
	if !ok {
		return Coin{}, fmt.Errorf("failed to parse coin amount: %s", amountStr)
	}

	return Coin{denomStr, amount}, nil
}

func ParseCoins(coinsStr string) (coins Coins, err error) {
	coinsStr = strings.TrimSpace(coinsStr)
	if len(coinsStr) == 0 {
		return nil, nil
	}

	coinStrs := strings.Split(coinsStr, ",")
	for _, coinStr := range coinStrs {
		coin, err := ParseCoin(coinStr)
		if err != nil {
			return nil, err
		}
		coins = append(coins, coin)
	}

	coins.Sort()

	if !coins.IsValid() {
		return nil, fmt.Errorf("parseCoins invalid: %#v", coins)
	}

	return coins, nil
}

package currency

import (
	"encoding/json"
	"math/big"
)

type Fiat Amount

// String returns an "amount string" with amount precision.
func (a Fiat) String() string {
	val := int64(a.Round())
	return StringFromInt64(val, FiatPrecision)
}

// UnmarshalJSON implementation of the `json.Unmarshaller` interface.
func (a *Fiat) UnmarshalJSON(data []byte) error {
	var amount Amount

	err := json.Unmarshal(data, &amount)
	if err != nil {
		return err
	}

	*a = Fiat(amount)
	return nil
}

// MarshalJSON implementation of the `json.Marshaller` interface.
func (a Fiat) MarshalJSON() ([]byte, error) {
	str := a.String()
	return []byte(str), nil
}

// Convert converts fiat amount into `Coin` by passed price.
func (a Fiat) Convert(price Price) ConversionResult {
	//	fiat / price = coins
	result, _ := new(big.Float).Quo(big.NewFloat(float64(a)),
		big.NewFloat(float64(price))).Int64()
	coins := Coin(result)

	return ConversionResult{
		Coins: coins,
		Fiat:  a,
		Price: price,
	}
}
func (a Fiat) Round() Fiat {
	f := Amount(a).Float()
	am := FromFloat(BankRound(f, FiatPrecision))
	return Fiat(am)
}

// GetPercent calculates the percentage value from the sum and rounds it up.
func (a Fiat) GetPercent(percent int64) Fiat {
	amount := Amount(a)
	amount = AmountPercent(amount, percent)
	return Fiat(amount)
}

package currency

import (
	"encoding/json"
	"math/big"
)

type Coin Amount

// String returns an "amount string" with amount precision.
func (a Coin) String() string {
	return StringFromInt64(int64(a), CoinPrecision)
}

// UnmarshalJSON implementation of the `json.Unmarshaller` interface.
func (a *Coin) UnmarshalJSON(data []byte) error {
	var amount Amount

	err := json.Unmarshal(data, &amount)
	if err != nil {
		return err
	}

	*a = Coin(amount)
	return nil
}

// MarshalJSON implementation of the `json.Marshaller` interface.
func (a Coin) MarshalJSON() ([]byte, error) {
	str := a.String()
	return []byte(str), nil
}

func (a Coin) Convert(price Price) ConversionResult {
	//	coins * price = fiat
	//	fiat to fixed precision
	//	fiat / price = fixed coins
	result, _ := new(big.Float).Mul(big.NewFloat(float64(a)), big.NewFloat(float64(price))).Int64()
	fiat := Fiat(result).Round()

	return fiat.Convert(price)
}

// GetPercent calculates the percentage value from the sum and rounds it up.
func (a Coin) GetPercent(percent int64) Coin {
	amount := Amount(a)
	amount = AmountPercent(amount, percent)
	return Coin(amount)
}

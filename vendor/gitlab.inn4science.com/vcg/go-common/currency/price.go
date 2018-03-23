package currency

import (
	"encoding/json"
	"math/big"
)

type Price float64

// String returns an "price string" with price precision.
func (price Price) String() string {
	var res big.Rat
	res.SetFloat64(float64(price))
	return res.FloatString(PricePrecision)
}

// UnmarshalJSON implementation of the `json.Unmarshaller` interface.
func (price *Price) UnmarshalJSON(data []byte) error {
	var amount float64

	err := json.Unmarshal(data, &amount)
	if err != nil {
		return err
	}

	*price = Price(amount)
	return nil
}

// MarshalJSON implementation of the `json.Marshaller` interface.
func (price Price) MarshalJSON() ([]byte, error) {
	str := price.String()
	return []byte(str), nil
}

func (price *Price) Round() Price {
	f := BankRound(float64(*price), PricePrecision)
	*price = Price(f)
	return *price
}

func (price *Price) Float() float64 {
	return float64(*price)
}

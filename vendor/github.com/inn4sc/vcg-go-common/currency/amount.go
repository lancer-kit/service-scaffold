package currency

import (
	"encoding/json"
	"fmt"
	"math/big"
)

const (
	One             = 100000000
	AmountPrecision = 8
	CoinPrecision   = 8
	FiatPrecision   = 2
	PricePrecision  = 4
)

// Amount is a type for coins or fiats values.
type Amount int64

// Parse returns `Amount` parsed from string or parse error.
func Parse(str string) (Amount, error) {
	var f, one, result big.Float

	_, ok := f.SetString(str)
	if !ok {
		return Amount(0), fmt.Errorf("cannot parse amount: %s", str)
	}

	one.SetInt64(One)
	result.Mul(&f, &one)

	i, _ := result.Int64()
	return Amount(i), nil
}

// FromFloat converts float64 value to the `Amount`.
func FromFloat(val float64) Amount {
	var one, mult, bigVal big.Rat
	one.SetInt64(One)
	bigVal.SetFloat64(val)
	mult.Mul(&bigVal, &one)

	res, _ := mult.Float64()
	return Amount(int64(res))
}

// StringFromInt64 returns an "amount string" from the provided raw int64 value `v`.
func StringFromInt64(val int64, precision int) string {
	var bigVal, one, result big.Rat
	bigVal.SetInt64(val)
	one.SetInt64(One)
	result.Quo(&bigVal, &one)

	return result.FloatString(precision)
}

// Int casts amount to int64.
func (a Amount) Int() int64 {
	return int64(a)
}

// Float casts amount to float64.
func (a Amount) Float() float64 {
	var bigVal, one, result big.Rat
	bigVal.SetInt64(int64(a))
	one.SetInt64(One)
	result.Quo(&bigVal, &one)
	f, _ := result.Float64()
	return f
}

// String returns an "amount string" with amount precision.
func (a Amount) String() string {
	return StringFromInt64(int64(a), AmountPrecision)
}

// UnmarshalJSON implementation of the `json.Unmarshaller` interface.
func (a *Amount) UnmarshalJSON(data []byte) error {
	var fVal float64
	err := json.Unmarshal(data, &fVal)
	if err == nil {
		*a = FromFloat(fVal)
		return nil
	}
	var sVal string
	err = json.Unmarshal(data, &sVal)
	if err == nil {
		*a, err = Parse(sVal)
		return err
	}

	return fmt.Errorf("invalid amount value: %s", string(data))
}

// MarshalJSON implementation of the `json.Marshaller` interface.
func (a Amount) MarshalJSON() ([]byte, error) {
	str := a.String()
	return []byte(str), nil
}

// GetPercent calculates the percentage value from the sum and rounds it up.
func (a Amount) GetPercent(percent int64) Amount {
	return AmountPercent(a, percent)
}

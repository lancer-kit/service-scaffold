package currency

import (
	"strings"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestCoin_String(t *testing.T) {
	coin := Coin(42123456789)
	str := coin.String()
	strs := strings.Split(str, ".")
	if len(strs) != 2 {
		t.Error("len(strs) != 2")
		t.Fail()
	}
	if len(strs[1]) != CoinPrecision {
		t.Error("len(strs[2]) != CoinPrecision")
		t.Fail()
	}
}

func TestCoin_Convert(t *testing.T) {
	coin := Coin(101039260)
	price := Price(0.3464)
	fiat := Fiat(FromFloat(0.35))

	conversion := coin.Convert(price)
	fmt.Println(conversion.Coins)
	assert.Equal(t, fiat.String(), conversion.Fiat.String())
}

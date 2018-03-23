package currency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiat_Round(t *testing.T) {
	testVal := map[Fiat]Fiat{
		Fiat(FromFloat(0.3464)): Fiat(FromFloat(0.35)),
		Fiat(FromFloat(0.3364)): Fiat(FromFloat(0.34)),
		Fiat(FromFloat(0.425)):  Fiat(FromFloat(0.42)),
		Fiat(FromFloat(0.415)):  Fiat(FromFloat(0.42)),
		Fiat(FromFloat(0.412)):  Fiat(FromFloat(0.41)),
	}

	for val, result := range testVal {
		res := val.Round()
		assert.Equal(t, result, res)
	}

}

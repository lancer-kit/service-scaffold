package currency

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	var testData = []struct {
		str    string
		amount Amount
		err    error
	}{
		{
			str:    "0",
			amount: 0,
			err:    nil,
		},
		{
			str:    "1",
			amount: 1 * One,
			err:    nil,
		},
		{
			str:    "0.9",
			amount: 90000000,
			err:    nil,
		},
		{
			str:    "0.99999999",
			amount: 99999999,
			err:    nil,
		},
		{
			str:    "0.999999895999",
			amount: 99999989,
			err:    nil,
		},
		{
			str:    "12.4",
			amount: 1240000000,
			err:    nil,
		},
		{
			str:    "123456789.986754321",
			amount: 12345678998675432,
			err:    nil,
		},
		{
			str:    "1nb4l1d",
			amount: 0,
			err:    fmt.Errorf("cannot parse amount: %s", "1nb4l1d"),
		},
		{
			str:    "13,6",
			amount: 0,
			err:    fmt.Errorf("cannot parse amount: %s", "13,6"),
		},
	}

	var am Amount
	var err error
	for _, td := range testData {
		am, err = Parse(td.str)
		assert.Equal(t, td.err, err)
		assert.Equal(t, td.amount, am)
	}
}

func TestFromFloat(t *testing.T) {
	var testData = []struct {
		float  float64
		amount Amount
	}{
		{
			float:  0,
			amount: 0,
		},
		{
			float:  1,
			amount: 1 * One,
		},
		{
			float:  0.9,
			amount: 90000000,
		},
		{
			float:  12.4,
			amount: 1240000000,
		},
		{
			float:  0.99999999,
			amount: 99999999,
		},
		{
			float:  0.999999995,
			amount: 99999999,
		},
		{
			float:  5.999999895999,
			amount: 599999989,
		},
		{
			float:  123456789.986754321,
			amount: 12345678998675432,
		},
	}

	var am Amount
	for _, td := range testData {
		am = FromFloat(td.float)
		assert.Equal(t, td.amount, am)
	}
}

func TestAmount_String(t *testing.T) {
	var testData = []struct {
		str    string
		amount Amount
	}{
		{
			str:    "0.00000000",
			amount: 0,
		},
		{
			str:    "1.00000000",
			amount: 1 * One,
		},
		{
			str:    "0.90000000",
			amount: 90000000,
		},
		{
			str:    "0.99999999",
			amount: 99999999,
		},
		{
			str:    "0.99999989",
			amount: 99999989,
		},
		{
			str:    "12.40000000",
			amount: 1240000000,
		},
		{
			str:    "123456789.98675432",
			amount: 12345678998675432,
		},
	}

	for _, td := range testData {
		assert.Equal(t, td.str, td.amount.String())
	}
}

func TestAmount_GetPercent(t *testing.T) {
	var testData = []struct {
		amount  Amount
		result  Amount
		percent int64
	}{
		{
			amount:  100,
			percent: 5,
			result:  5,
		},
		{
			amount: 1 * One,
		},
		{
			amount:  123456789,
			percent: 23,
			result:  28395062,
		},
		{
			amount:  60,
			percent: 1,
			result:  1,
		},
	}
	for _, td := range testData {
		assert.Equal(t, td.result, td.amount.GetPercent(td.percent))
	}
}

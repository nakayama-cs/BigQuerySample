package to

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFloat64Parts(t *testing.T) {
	cases := []struct {
		IntPart  int64
		FracPart int32
		Str      string
	}{
		{0, 0, "0.0"},
		{0, 0, "0.00"},
		{0, 0, "0.000"},
		{0, 0, "0.0000"},
		{0, 0, "0.00000"},
		{0, 0, "0"},
		{1, 5000, "1.5"},
		{1, 500, "1.05"},
		{1, 50, "1.005"},
		{1, 5, "1.0005"},
		{1, 0, "1.00005"},
	}
	for _, c := range cases {
		intPart, fracPart, err := parseFloat64Parts(c.Str)
		if !assert.NoError(t, err, c.Str) {
			continue
		}
		assert.Equal(t, []any{c.IntPart, c.FracPart}, []any{intPart, fracPart}, c.Str)
	}
}

func TestToFloat64(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		cases := []struct {
			IntPart  int64
			FracPart int32
			Expect   float64
		}{
			{0, 0, 0.0},
			{123, 4567, 123.4567},
			{0, 1, 0.1},
			{1, 0, 1.0},
			{123, 45678, 123.4567},
		}

		for _, c := range cases {
			actual, err := ToFloat64(c.IntPart, c.FracPart)
			if !assert.NoError(t, err) {
				continue
			}
			assert.Equal(t, c.Expect, actual)
		}
	})
}

func TestToFloat64Parts(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		cases := []struct {
			InFloat      float64
			ExIntegral   int64
			ExFractional int32
		}{
			{123.4567, 123, 4567},
			{123.456789, 123, 4567},
			{0.12345, 0, 1234},
			{12345.0, 12345, 0},
			{0.0, 0, 0},
			{0, 0, 0},
			{-123, -123, 0},
			{-123.4567, -123, 4567},
		}

		for _, c := range cases {
			outIntegral, outFractional, err := ToFloat64Parts(c.InFloat)
			if !assert.NoError(t, err) {
				continue
			}
			assert.Equal(t, c.ExIntegral, outIntegral)
			assert.Equal(t, c.ExFractional, outFractional)
		}
	})
}

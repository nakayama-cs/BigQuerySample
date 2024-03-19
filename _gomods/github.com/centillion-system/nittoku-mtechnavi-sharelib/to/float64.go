package to

import (
	"strconv"
	"strings"
)

func ToFloat64(intPart int64, fracPart int32) (float64, error) {
	intPartStr := strconv.FormatInt(intPart, 10)
	fracPartStr := strconv.FormatInt(int64(fracPart), 10)

	// 小数点5桁以下は切り捨て
	switch chars := []rune(fracPartStr); {
	case len(chars) > 4:
		fracPartStr = string(fracPartStr[:4])
	default:
		fracPartStr = string(fracPartStr + strings.Repeat("0", 4-len(chars)))
	}

	amountStr := intPartStr + "." + fracPartStr
	if xAmount, err := strconv.ParseFloat(amountStr, 64); err != nil {
		return 0.0, err
	} else {
		return xAmount, nil
	}
}
func ToFloat64Parts(amount float64) (int64, int32, error) {
	amountStr := strconv.FormatFloat(amount, 'f', -1, 64)

	i, f, err := parseFloat64Parts(amountStr)
	if err != nil {
		return 0, 0, err
	}
	return i, f, nil
}

func parseFloat64Parts(s string) (int64, int32, error) {
	parts := strings.SplitN(s, ".", 2)
	if len(parts) == 1 {
		// "." がないパターン
		parts = append(parts, "0")
	}
	intPartStr, fracPartStr := parts[0], parts[1]
	// 小数点5桁以下は切り捨て
	switch chars := []rune(fracPartStr); {
	case len(chars) > 4:
		fracPartStr = string(fracPartStr[:4])
	default:
		fracPartStr = string(fracPartStr + strings.Repeat("0", 4-len(chars)))
	}

	intPart, err := strconv.ParseInt(intPartStr, 10, 64)
	if err != nil {
		return 0, 0, err
	}
	fracPart, err := strconv.ParseInt(fracPartStr, 10, 32)
	if err != nil {
		return 0, 0, err
	}
	return intPart, int32(fracPart), nil
}

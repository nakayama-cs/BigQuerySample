package to

import (
	"fmt"
	"mtechnavi/sharelib/constants"
	pb "mtechnavi/sharelib/protobuf"
)

func ToDistantAmountCSV(amount *pb.DistantAmount) (string, float64, error) {
	var nameOptionUnitCode string
	var distantAmount float64

	if v, ok := constants.Csv_OutputName_value[amount.Unit]; !ok {
		return "", 0, fmt.Errorf("unit not found amount.unit:%d", amount.Unit)
	} else {
		nameOptionUnitCode = v
	}

	if v, err := ToFloat64(amount.IntegralAmount, amount.FractionalAmount); err != nil {
		return "", 0, err
	} else {
		distantAmount = v
	}

	return nameOptionUnitCode, distantAmount, nil
}

func ToDistantAmount(nameOptionUnitCode string, distantAmountValue float64) (*pb.DistantAmount, error) {
	distantAmount := &pb.DistantAmount{}

	if nameOptionUnitCode != "" {
		if v, ok := constants.Csv_OutputName_name[nameOptionUnitCode]; !ok {
			return nil, fmt.Errorf("unit code not found code:%s", nameOptionUnitCode)
		} else {
			distantAmount.Unit = v
		}
	} else {
		distantAmount.Unit = pb.DistantAmount_Unit(0) // 単位コードが未指定の場合は、デフォルト値（ゼロ値）を割り当てる
	}

	if i, f, err := ToFloat64Parts(distantAmountValue); err != nil {
		return nil, err
	} else {
		distantAmount.IntegralAmount = i
		distantAmount.FractionalAmount = f
	}
	return distantAmount, nil
}

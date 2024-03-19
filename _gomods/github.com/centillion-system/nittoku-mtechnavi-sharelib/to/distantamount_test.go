package to

import (
	pb "mtechnavi/sharelib/protobuf"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func ToDistantAmountCSV(amount *pb.DistantAmount) (string, float64, error) {
// func ToDistantAmount(nameOptionUnitCode string, distantAmountValue float64) (*pb.DistantAmount, error) {

func TestToDistantAmountCSV(t *testing.T) {
	cases := []struct {
		inAmount      *pb.DistantAmount
		exAmountUnit  string
		exAmountValue float64
	}{
		{&pb.DistantAmount{Unit: pb.DistantAmount_MILLI_METER, IntegralAmount: 123, FractionalAmount: 4567}, "3", 123.4567},
		{&pb.DistantAmount{Unit: pb.DistantAmount_CENTI_METER, IntegralAmount: -123, FractionalAmount: 4567}, "2", -123.4567},
		{&pb.DistantAmount{Unit: pb.DistantAmount_METER, IntegralAmount: 0, FractionalAmount: 4567}, "1", 0.4567},
		{&pb.DistantAmount{Unit: pb.DistantAmount_MICRO_METER, IntegralAmount: 0, FractionalAmount: 0}, "4", 0.0},
		{&pb.DistantAmount{Unit: pb.DistantAmount_NANO_METER, IntegralAmount: 123, FractionalAmount: 456789}, "5", 123.4567},
		{&pb.DistantAmount{Unit: pb.DistantAmount_NONE, IntegralAmount: 0, FractionalAmount: 0}, "", 0.0},
	}
	for _, c := range cases {
		outAmountUnit, outAmountValue, err := ToDistantAmountCSV(c.inAmount)
		if !assert.NoError(t, err) {
			continue
		}
		assert.Equal(t, []any{c.exAmountUnit, c.exAmountValue}, []any{outAmountUnit, outAmountValue})
	}
}

func TestToDistantAmount(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		cases := []struct {
			inUnitCode    string
			inAmountValue float64
			exAmount      *pb.DistantAmount
		}{
			{"3", 123.4567, &pb.DistantAmount{Unit: pb.DistantAmount_MILLI_METER, IntegralAmount: 123, FractionalAmount: 4567}},
			{"2", -123.4567, &pb.DistantAmount{Unit: pb.DistantAmount_CENTI_METER, IntegralAmount: -123, FractionalAmount: 4567}},
			{"1", 0.4567, &pb.DistantAmount{Unit: pb.DistantAmount_METER, IntegralAmount: 0, FractionalAmount: 4567}},
			{"4", 0.0, &pb.DistantAmount{Unit: pb.DistantAmount_MICRO_METER, IntegralAmount: 0, FractionalAmount: 0}},
			{"5", 123.456789, &pb.DistantAmount{Unit: pb.DistantAmount_NANO_METER, IntegralAmount: 123, FractionalAmount: 4567}},
			{"", 0.0, &pb.DistantAmount{Unit: pb.DistantAmount_NONE, IntegralAmount: 0, FractionalAmount: 0}},
		}

		for _, c := range cases {
			outAmount, err := ToDistantAmount(c.inUnitCode, c.inAmountValue)
			if !assert.NoError(t, err) {
				continue
			}
			assert.Equal(t, c.exAmount, outAmount)
		}
	})
}

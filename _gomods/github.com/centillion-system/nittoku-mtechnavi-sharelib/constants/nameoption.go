package constants

//go:generate generate-nameoption-constants -o nameoption_gen.go nameoption.tsv
//go:generate gofmt -w nameoption_gen.go

import (
	pb "mtechnavi/sharelib/protobuf"
)

var (
	// {systemName} => enum
	NameOption_SystemName_A0000027_name = map[string]pb.DistantAmount_Unit{
		// NameOption側にはNONEは用意しないためマッピング不要
		// NameOption_SystemName_A0000027_B00: pb.DistantAmount_NONE,
		NameOption_SystemName_A0000027_B01: pb.DistantAmount_METER,
		NameOption_SystemName_A0000027_B02: pb.DistantAmount_CENTI_METER,
		NameOption_SystemName_A0000027_B03: pb.DistantAmount_MILLI_METER,
		NameOption_SystemName_A0000027_B04: pb.DistantAmount_MICRO_METER,
		NameOption_SystemName_A0000027_B05: pb.DistantAmount_NANO_METER,
	}
	// enum => {systemName}
	NameOption_SystemName_A0000027_value = map[pb.DistantAmount_Unit]string{}

	// {Csv_OutputName_name} => enum
	Csv_OutputName_name = map[string]pb.DistantAmount_Unit{
		// NameOption側にはNONEは用意しないためマッピング不要
		"":  pb.DistantAmount_NONE,
		"1": pb.DistantAmount_METER,
		"2": pb.DistantAmount_CENTI_METER,
		"3": pb.DistantAmount_MILLI_METER,
		"4": pb.DistantAmount_MICRO_METER,
		"5": pb.DistantAmount_NANO_METER,
	}
	// enum => {Csv_OutputName}
	Csv_OutputName_value = map[pb.DistantAmount_Unit]string{}
)

func init() {
	for k, v := range NameOption_SystemName_A0000027_name {
		NameOption_SystemName_A0000027_value[v] = k
	}
	for k, v := range Csv_OutputName_name {
		Csv_OutputName_value[v] = k
	}
}

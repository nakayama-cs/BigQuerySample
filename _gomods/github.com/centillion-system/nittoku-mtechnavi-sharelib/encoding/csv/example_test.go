package csv_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mtechnavi/sharelib/encoding/csv"
	pb "mtechnavi/sharelib/encoding/csv/testdata/protobuf"
	"mtechnavi/sharelib/text/encoding"
	"os"

	"github.com/hashicorp/go-multierror"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Example_read() {
	file, err := os.Open("./testdata/example.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := encoding.NewReader(file, encoding.ShiftJIS)

	reader := csv.NewReader(r)
	reader.Comma = ','
	reader.Comment = '#'

	for {
		var record pb.SampleRecord
		if err := reader.Read(&record); errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", jsonc(&record))
	}

	// Output:
	// {"user_id":"U001","email":"user001@example.com","display_name":"取込 太郎","password":"","age":10,"weight":0,"tags":["default001","default002"],"enable_single_sign_on":false}
	// {"user_id":"U002","email":"user002@example.com","display_name":"取込 次郎","password":"","age":20,"weight":0.1,"tags":["a","b"],"enable_single_sign_on":true}
	// {"user_id":"U003","email":"user003@example.com","display_name":"取込 三郎","password":"","age":30,"weight":0.01,"tags":["b","c"],"enable_single_sign_on":true}
	// {"user_id":"U004","email":"user004@example.com","display_name":"取込 四郎","password":"","age":40,"weight":0.001,"tags":["default001","default002"],"enable_single_sign_on":true}
}

func Example_readAll() {
	file, err := os.Open("./testdata/example.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := encoding.NewReader(file, encoding.ShiftJIS)

	reader := csv.NewReader(r)
	reader.Comma = ','
	reader.Comment = '#'

	var records []*pb.SampleRecord
	if err := reader.ReadAll(&records); err != nil {
		panic(err)
	}
	for _, record := range records {
		fmt.Printf("%v\n", jsonc(record))
	}

	// Output:
	// {"user_id":"U001","email":"user001@example.com","display_name":"取込 太郎","password":"","age":10,"weight":0,"tags":["default001","default002"],"enable_single_sign_on":false}
	// {"user_id":"U002","email":"user002@example.com","display_name":"取込 次郎","password":"","age":20,"weight":0.1,"tags":["a","b"],"enable_single_sign_on":true}
	// {"user_id":"U003","email":"user003@example.com","display_name":"取込 三郎","password":"","age":30,"weight":0.01,"tags":["b","c"],"enable_single_sign_on":true}
	// {"user_id":"U004","email":"user004@example.com","display_name":"取込 四郎","password":"","age":40,"weight":0.001,"tags":["default001","default002"],"enable_single_sign_on":true}
}

func ExampleError_read() {
	file, err := os.Open("./testdata/error.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := encoding.NewReader(file, encoding.ShiftJIS)

	reader := csv.NewReader(r)
	reader.Comma = ','
	reader.Comment = '#'

	for {
		var record pb.SampleRecord
		err := reader.Read(&record)
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			switch e := err.(type) {
			case *csv.Error:
				fmt.Println("---")
				for _, v := range e.Causes {
					switch e := v.(type) {
					case *csv.Error:
						fmt.Printf("Error:\n\tColumn=%d, Line=%d\n\tMessage=%v\n",
							e.Column, e.Line, e.Error(),
						)
					}
				}
			case *multierror.Error:
				fmt.Println("---")
				for _, v := range e.WrappedErrors() {
					switch e := v.(type) {
					case *csv.Error:
						fmt.Printf("Error:\n\tColumn=%d, Line=%d\n\tMessage=%v\n",
							e.Column, e.Line, e.Error(),
						)
					}
				}
			default:
				panic(err)
			}
		}
		fmt.Printf("%v\n", jsonc(&record))
	}

	// Output:
	// ---
	// Error:
	// 	Column=5, Line=2
	// 	Message=CSVError Line:2,Column:5,Cause:[CSVError Line:0,Column:0,Cause:[strconv.ParseInt: parsing "error10": invalid syntax],ErrorLevel:ERROR,MessageName:E0000004,MessageArgs:[]],ErrorLevel:ERROR,MessageName:E1000002,MessageArgs:[AGE]
	// Error:
	// 	Column=6, Line=2
	// 	Message=CSVError Line:2,Column:6,Cause:[CSVError Line:0,Column:0,Cause:[strconv.ParseFloat: parsing "error0": invalid syntax],ErrorLevel:ERROR,MessageName:E0000004,MessageArgs:[]],ErrorLevel:ERROR,MessageName:E1000002,MessageArgs:[]
	// Error:
	// 	Column=8, Line=2
	//	Message=CSVError Line:2,Column:8,Cause:[CSVError Line:0,Column:0,Cause:[],ErrorLevel:ERROR,MessageName:E0000017,MessageArgs:[]],ErrorLevel:ERROR,MessageName:E1000002,MessageArgs:[SINGLE_SIGN_ON]
	// {"user_id":"U001","email":"user001@example.com","display_name":"取込 太郎","password":"","age":0,"weight":0,"tags":["default001","default002"],"enable_single_sign_on":false}
	// ---
	// Error:
	// 	Column=5, Line=3
	// 	Message=CSVError Line:3,Column:5,Cause:[CSVError Line:0,Column:0,Cause:[strconv.ParseInt: parsing "２０": invalid syntax],ErrorLevel:ERROR,MessageName:E0000004,MessageArgs:[]],ErrorLevel:ERROR,MessageName:E1000002,MessageArgs:[AGE]
	// {"user_id":"U002","email":"user002@example.com","display_name":"取込 次郎","password":"","age":0,"weight":0.1,"tags":["a","b"],"enable_single_sign_on":true}
	// ---
	// Error:
	// 	Column=5, Line=5
	// 	Message=CSVError Line:5,Column:5,Cause:[CSVError Line:0,Column:0,Cause:[strconv.ParseInt: parsing "error\n30": invalid syntax],ErrorLevel:ERROR,MessageName:E0000004,MessageArgs:[]],ErrorLevel:ERROR,MessageName:E1000002,MessageArgs:[AGE]
	// Error:
	// 	Column=8, Line=6
	// 	Message=CSVError Line:6,Column:8,Cause:[CSVError Line:0,Column:0,Cause:[],ErrorLevel:ERROR,MessageName:E0000017,MessageArgs:[]],ErrorLevel:ERROR,MessageName:E1000002,MessageArgs:[SINGLE_SIGN_ON]
	// {"user_id":"U003","email":"user003@example.com","display_name":"取込 \n三郎","password":"","age":0,"weight":0.01,"tags":["b","c"],"enable_single_sign_on":false}
}

func Example_writeAll() {
	file, err := os.CreateTemp("", "")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := encoding.NewWriter(file, encoding.ShiftJIS)

	writer := csv.NewWriter(io.MultiWriter(w, os.Stdout))
	writer.Comma = ','
	writer.UseCRLF = false

	records := []*pb.SampleRecord{
		{
			UserId:             "U001",
			Email:              "user001@example.com",
			DisplayName:        `取込 "クォート`,
			Password:           "XXX",
			Age:                10,
			EnableSingleSignOn: false,
		},
		{
			UserId:             "U002",
			Email:              "user002@example.com",
			DisplayName:        "取込\n改行",
			Password:           "XXX",
			Age:                20,
			Weight:             0.1,
			Tags:               []string{"A", "B"},
			EnableSingleSignOn: true,
		},
	}
	if err := writer.WriteAll(records); err != nil {
		panic(err)
	}

	// Output:
	// user001@example.com,U001,"取込 ""クォート",,10,0,,0
	// user002@example.com,U002,"取込
	// 改行",,20,0.1,"A,B",1
}

func Example_writeAll_forceQuote() {
	file, err := os.CreateTemp("", "")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := encoding.NewWriter(file, encoding.ShiftJIS)

	writer := csv.NewWriter(io.MultiWriter(w, os.Stdout))
	writer.Comma = ','
	writer.UseCRLF = false
	writer.ForceQuote = true

	records := []*pb.SampleRecord{
		{
			UserId:      "U001",
			Email:       "user001@example.com",
			DisplayName: "取込 太郎",
			Password:    "XXX",
			Age:         10,
		},
		{
			UserId:             "U002",
			Email:              "user002@example.com",
			DisplayName:        "取込 次郎",
			Password:           "XXX",
			Age:                20,
			Weight:             0.1,
			Tags:               []string{"A", "B"},
			EnableSingleSignOn: bool(true),
		},
		{
			UserId:      "U003",
			Email:       "user003@example.com",
			DisplayName: "取込 三郎",
			Password:    "XXX",
			Age:         30,
			Weight:      0.01,
			Tags:        []string{"B", "C"},
		},
	}
	if err := writer.WriteAll(records); err != nil {
		panic(err)
	}

	// Output:
	// "user001@example.com","U001","取込 太郎","","10","0","","0"
	// "user002@example.com","U002","取込 次郎","","20","0.1","A,B","1"
	// "user003@example.com","U003","取込 三郎","","30","0.01","B,C","0"
}

func jsonc(v proto.Message) string {
	var opts protojson.MarshalOptions
	opts.EmitUnpopulated = true
	opts.UseProtoNames = true
	data := json.RawMessage(opts.Format(v))
	var buffer bytes.Buffer
	json.Compact(&buffer, data)
	return buffer.String()
}

package csv_test

import (
	pb "mtechnavi/sharelib/encoding/csv/testdata/protobuf"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalUnmarshalCSV(t *testing.T) {
	t.Run("SampleRecord", func(t *testing.T) {
		v := &pb.SampleRecord{
			Email:              "c1",
			UserId:             "c2",
			DisplayName:        "c3",
			Password:           "ignore",
			Age:                5,
			Weight:             7.7,
			Tags:               []string{"c7.0", "c7.1"},
			EnableSingleSignOn: true,
		}

		record, err := v.MarshalCSV()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, []string{
			"c1",
			"c2",
			"c3",
			"",
			"5",
			"7.7",
			"c7.0,c7.1",
			"1",
		}, record)

		var v2 pb.SampleRecord
		err = v2.UnmarshalCSV(record)
		if !assert.NoError(t, err) {
			return
		}
		v.Password = ""
		assert.JSONEq(t, jsonc(v), jsonc(&v2))
	})
	t.Run("SampleRecordWithError", func(t *testing.T) {
		v := &pb.SampleRecordWithError{
			Base: &pb.SampleRecord{
				Email:              "c1",
				UserId:             "c2",
				DisplayName:        "c3",
				Password:           "ignore",
				Age:                5,
				Weight:             7.7,
				Tags:               []string{"c7.0", "c7.1"},
				EnableSingleSignOn: false,
			},
			ErrorContent: "c9",
		}

		record, err := v.MarshalCSV()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, []string{
			"c1",
			"c2",
			"c3",
			"",
			"5",
			"7.7",
			"c7.0,c7.1",
			"0",
			"c9",
		}, record)

		var v2 pb.SampleRecordWithError
		err = v2.UnmarshalCSV(record)
		if !assert.NoError(t, err) {
			return
		}
		v.Base.Password = ""
		assert.JSONEq(t, jsonc(v), jsonc(&v2))
	})
	t.Run("SampleRecordWithError", func(t *testing.T) {
		v := &pb.SampleRecordWithError{}

		record, err := v.MarshalCSV()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, []string{
			"",
			"",
			"",
			"",
			"0",
			"0",
			"",
			"0",
			"",
		}, record)

		var v2 pb.SampleRecordWithError
		err = v2.UnmarshalCSV(record)
		if !assert.NoError(t, err) {
			return
		}
		var base pb.SampleRecord
		base.Tags = []string{"default001", "default002"}
		v.Base = &base
		assert.JSONEq(t, jsonc(v), jsonc(&v2))
	})
	t.Run("SampleRecordWithError", func(t *testing.T) {
		v := &pb.SampleRecordWithError{
			ErrorContent: "c8",
		}

		record, err := v.MarshalCSV()
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, []string{
			"",
			"",
			"",
			"",
			"0",
			"0",
			"",
			"0",
			"c8",
		}, record)

		var v2 pb.SampleRecordWithError
		err = v2.UnmarshalCSV(record)
		if !assert.NoError(t, err) {
			return
		}
		var base pb.SampleRecord
		base.Tags = []string{"default001", "default002"}
		v.Base = &base
		assert.JSONEq(t, jsonc(v), jsonc(&v2))
	})
}

func TestDefaultCSV(t *testing.T) {
	t.Run("SampleRecord_EmptyDefault", func(t *testing.T) {
		record := []string{
			"c1",  // Email:          {column: 2}
			"c2",  // UserId:         {column: 1}
			"c3",  // DisplayName:    {column: 3}
			"",    // empty:          {column: 4}
			"",    // Age:            {column: 5, default: "0"}
			"7.7", // weight:         {column: 6}
			"",    // tag:            {column: 7, default: "default001,default002"}
			"",    // SINGLE_SIGN_ON: {column: 8, default: "1"}
		}
		var v2 pb.SampleRecord
		err := v2.UnmarshalCSV(record)
		if !assert.NoError(t, err) {
			return
		}
		expect := pb.SampleRecord{
			Email:       "c1",
			UserId:      "c2",
			DisplayName: "c3",
			// empty col
			Age:                0,
			Weight:             7.7,
			Tags:               []string{"default001", "default002"},
			EnableSingleSignOn: true,
		}
		assert.JSONEq(t, jsonc(&expect), jsonc(&v2))
	})
	t.Run("SampleRecord_FilledDefault", func(t *testing.T) {
		record := []string{
			"",         // Email:          {column: 2}
			"UU001",    // UserId:         {column: 1}
			"",         // DisplayName:    {column: 3}
			"",         // empty:          {column: 4}
			"5",        // Age:            {column: 5, default: "0"}
			"",         // weight:         {column: 6}
			"6,66,666", // tag:            {column: 7, default: "default001,default002"}
			"0",        // SINGLE_SIGN_ON: {column: 8, default: "1"}
		}
		var v2 pb.SampleRecord
		err := v2.UnmarshalCSV(record)
		if !assert.NoError(t, err) {
			return
		}
		expect := pb.SampleRecord{
			Email:       "",
			UserId:      "UU001",
			DisplayName: "",
			// empty col
			Age:                5,
			Weight:             0,
			Tags:               []string{"6", "66", "666"},
			EnableSingleSignOn: false,
		}
		assert.JSONEq(t, jsonc(&expect), jsonc(&v2))
	})
}

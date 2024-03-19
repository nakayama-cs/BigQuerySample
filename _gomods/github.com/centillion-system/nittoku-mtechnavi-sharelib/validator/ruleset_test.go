package validator_test

import (
	"context"
	"mtechnavi/sharelib/constants"
	pb "mtechnavi/sharelib/protobuf"
	"mtechnavi/sharelib/validator"
	testpb "mtechnavi/sharelib/validator/testdata/protobuf"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuleSet_Eval(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		ctx := context.Background()

		rs := validator.NewRuleSet()
		src := &testpb.TestRecord{
			String_: "",
		}
		rules := []validator.Rule{
			validator.NotEmptyf("string_", pb.ErrorLevel_ERROR, constants.MessageName_E0000003, "XXX"),
			validator.GreaterThanOrEqual("string_", 0),
		}
		expect := struct {
			wantError bool
			messages  []constants.MessageName
		}{
			wantError: true,
			messages:  []constants.MessageName{constants.MessageName_E0000003},
		}
		rs.Add(rules...)
		err := rs.Eval(ctx, src)
		if expect.wantError {
			assert.NotNil(t, err)
			e, ok := err.(*validator.Error)
			assert.True(t, ok)
			if !assert.Len(t, e.Results, len(expect.messages)) {
				return
			}
			for i, r := range e.Results {
				m := expect.messages[i]
				assert.Equal(t, m, r.MessageName)
			}
		} else {
			assert.Nil(t, err)
		}
	})
	t.Run("conditioned", func(t *testing.T) {
		ctx := context.Background()

		rs := validator.NewRuleSet()
		src := &testpb.TestRecord{
			String_: "a",
		}
		rules := []validator.Rule{
			rs.EvalIf(validator.NotEmptyField("string_"),
				validator.NotEmptyf("uint64_", pb.ErrorLevel_ERROR, constants.MessageName_E0000003, "UINT64_"),
				validator.GreaterThanOrEqual("int64_", 1),
			),
		}
		expect := struct {
			wantError   bool
			errMessages []constants.MessageName
		}{
			wantError: true,
			errMessages: []constants.MessageName{
				constants.MessageName_E0000003,
				constants.MessageName_E0000013,
			},
		}
		rs.Add(rules...)
		err := rs.Eval(ctx, src)
		if expect.wantError {
			assert.NotNil(t, err)
			e, ok := err.(*validator.Error)
			assert.True(t, ok)
			if !assert.Len(t, e.Results, len(expect.errMessages)) {
				return
			}
			for i, r := range e.Results {
				m := expect.errMessages[i]
				assert.Equal(t, m, r.MessageName)
			}
		} else {
			assert.Nil(t, err)
		}
	})
}

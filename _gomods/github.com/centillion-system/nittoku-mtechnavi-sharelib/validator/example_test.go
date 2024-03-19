package validator_test

import (
	"context"
	"fmt"
	"mtechnavi/sharelib/constants"
	pb "mtechnavi/sharelib/protobuf"
	"mtechnavi/sharelib/validator"
	testpb "mtechnavi/sharelib/validator/testdata/protobuf"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func testServerMethod(ctx context.Context, msgs []proto.Message) error {
	rules := validator.NewRuleSet()
	rules.Add(
		validator.NotEmptyf("x", pb.ErrorLevel_ERROR, constants.MessageName_E1000002, "STRING_X"),
		rules.Unique("x", pb.ErrorLevel_ERROR, constants.MessageName_E0000026),
		customNotEmptyf("x", constants.MessageName_E1000003, "X"),
	)
	for _, m := range msgs {
		if err := rules.Eval(ctx, m); err != nil {
			outputValidateErr(err)
		}
	}

	return status.Error(codes.OK, "example OK")
}

func Example() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	fmt.Println("- 1 -")
	_ = testServerMethod(ctx, []proto.Message{
		&testpb.TestingRequest{
			X: "",
		},
	})

	fmt.Println("- 2 -")
	_ = testServerMethod(ctx, []proto.Message{
		&testpb.TestingRequest{
			X: "win",
		},
		&testpb.TestingRequest{
			X: "winwin",
		},
		&testpb.TestingRequest{
			X: "",
		},
		&testpb.TestingRequest{
			X: "win",
		},
	})

	// Output:
	// - 1 -
	// grpcStatus: Unimplemented
	// message   : RuleResult Field:x,ErrorLevel:ERROR,MessageName:E1000002,MessageArgs:[STRING_X],RuleResult Field:x,ErrorLevel:ERROR,MessageName:E0000003,MessageArgs:[X]
	// - 2 -
	// grpcStatus: Unimplemented
	// message   : RuleResult Field:x,ErrorLevel:ERROR,MessageName:E1000002,MessageArgs:[STRING_X],RuleResult Field:x,ErrorLevel:ERROR,MessageName:E0000003,MessageArgs:[X]
	// grpcStatus: Unimplemented
	// message   : RuleResult Field:x,ErrorLevel:ERROR,MessageName:E0000026,MessageArgs:[]
}

func outputValidateErr(err error) {
	if st, ok := status.FromError(err); ok {
		fmt.Printf("grpcStatus: %v\n", st.Code())
	}
	if ve, ok := err.(*validator.Error); ok {
		fmt.Printf("message   : %v\n", ve.Error())
	}
}

func customNotEmptyf(field string, messageName constants.MessageName, args ...any) validator.Rule {
	return validator.RuleFunc(func(ctx context.Context, m *validator.MessageWrapper) []*validator.RuleResult {
		if res := validator.NotEmptyf(field, pb.ErrorLevel_ERROR, messageName, args...).Eval(ctx, m); res != nil {
			for _, r := range res {
				r.MessageName = constants.MessageName_E0000003
				r.ErrorLevel = pb.ErrorLevel_ERROR
			}
			return res
		}
		return nil
	})
}

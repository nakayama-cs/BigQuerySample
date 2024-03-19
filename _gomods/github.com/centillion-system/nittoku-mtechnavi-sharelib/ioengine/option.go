package ioengine

import (
	"mtechnavi/sharelib/validator"
)

type ImportParams struct {
	RuleSet *validator.RuleSet
}

func GetImportParamsFromOptions(ops ...ImportOption) (*ImportParams, error) {
	var params ImportParams
	for _, fn := range ops {
		if err := fn(&params); err != nil {
			return nil, err
		}
	}
	return &params, nil
}

type ImportOption func(params *ImportParams) error

func WithRuleSet(rs *validator.RuleSet) ImportOption {
	return func(params *ImportParams) error {
		params.RuleSet = rs
		return nil
	}
}

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: survey_io.proto

package protobuf

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ExportSurveyResultFilesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ExportSurveyResultFilesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExportSurveyResultFilesRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ExportSurveyResultFilesRequestMultiError, or nil if none found.
func (m *ExportSurveyResultFilesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ExportSurveyResultFilesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetTargets() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ExportSurveyResultFilesRequestValidationError{
						field:  fmt.Sprintf("Targets[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ExportSurveyResultFilesRequestValidationError{
						field:  fmt.Sprintf("Targets[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExportSurveyResultFilesRequestValidationError{
					field:  fmt.Sprintf("Targets[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for IncludeExported

	if len(errors) > 0 {
		return ExportSurveyResultFilesRequestMultiError(errors)
	}

	return nil
}

// ExportSurveyResultFilesRequestMultiError is an error wrapping multiple
// validation errors returned by ExportSurveyResultFilesRequest.ValidateAll()
// if the designated constraints aren't met.
type ExportSurveyResultFilesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExportSurveyResultFilesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExportSurveyResultFilesRequestMultiError) AllErrors() []error { return m }

// ExportSurveyResultFilesRequestValidationError is the validation error
// returned by ExportSurveyResultFilesRequest.Validate if the designated
// constraints aren't met.
type ExportSurveyResultFilesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExportSurveyResultFilesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExportSurveyResultFilesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExportSurveyResultFilesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExportSurveyResultFilesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExportSurveyResultFilesRequestValidationError) ErrorName() string {
	return "ExportSurveyResultFilesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ExportSurveyResultFilesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExportSurveyResultFilesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExportSurveyResultFilesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExportSurveyResultFilesRequestValidationError{}

// Validate checks the field values on ExportSurveyResultFilesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ExportSurveyResultFilesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExportSurveyResultFilesResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ExportSurveyResultFilesResponseMultiError, or nil if none found.
func (m *ExportSurveyResultFilesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ExportSurveyResultFilesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for AssetId

	if len(errors) > 0 {
		return ExportSurveyResultFilesResponseMultiError(errors)
	}

	return nil
}

// ExportSurveyResultFilesResponseMultiError is an error wrapping multiple
// validation errors returned by ExportSurveyResultFilesResponse.ValidateAll()
// if the designated constraints aren't met.
type ExportSurveyResultFilesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExportSurveyResultFilesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExportSurveyResultFilesResponseMultiError) AllErrors() []error { return m }

// ExportSurveyResultFilesResponseValidationError is the validation error
// returned by ExportSurveyResultFilesResponse.Validate if the designated
// constraints aren't met.
type ExportSurveyResultFilesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExportSurveyResultFilesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExportSurveyResultFilesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExportSurveyResultFilesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExportSurveyResultFilesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExportSurveyResultFilesResponseValidationError) ErrorName() string {
	return "ExportSurveyResultFilesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ExportSurveyResultFilesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExportSurveyResultFilesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExportSurveyResultFilesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExportSurveyResultFilesResponseValidationError{}

// Validate checks the field values on ExportSurveyResultFilesRequest_Target
// with the rules defined in the proto definition for this message. If any
// rules are violated, the first error encountered is returned, or nil if
// there are no violations.
func (m *ExportSurveyResultFilesRequest_Target) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExportSurveyResultFilesRequest_Target
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// ExportSurveyResultFilesRequest_TargetMultiError, or nil if none found.
func (m *ExportSurveyResultFilesRequest_Target) ValidateAll() error {
	return m.validate(true)
}

func (m *ExportSurveyResultFilesRequest_Target) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SurveyBaseRequestId

	// no validation rules for StatusCode

	if len(errors) > 0 {
		return ExportSurveyResultFilesRequest_TargetMultiError(errors)
	}

	return nil
}

// ExportSurveyResultFilesRequest_TargetMultiError is an error wrapping
// multiple validation errors returned by
// ExportSurveyResultFilesRequest_Target.ValidateAll() if the designated
// constraints aren't met.
type ExportSurveyResultFilesRequest_TargetMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExportSurveyResultFilesRequest_TargetMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExportSurveyResultFilesRequest_TargetMultiError) AllErrors() []error { return m }

// ExportSurveyResultFilesRequest_TargetValidationError is the validation error
// returned by ExportSurveyResultFilesRequest_Target.Validate if the
// designated constraints aren't met.
type ExportSurveyResultFilesRequest_TargetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExportSurveyResultFilesRequest_TargetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExportSurveyResultFilesRequest_TargetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExportSurveyResultFilesRequest_TargetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExportSurveyResultFilesRequest_TargetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExportSurveyResultFilesRequest_TargetValidationError) ErrorName() string {
	return "ExportSurveyResultFilesRequest_TargetValidationError"
}

// Error satisfies the builtin error interface
func (e ExportSurveyResultFilesRequest_TargetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExportSurveyResultFilesRequest_Target.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExportSurveyResultFilesRequest_TargetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExportSurveyResultFilesRequest_TargetValidationError{}

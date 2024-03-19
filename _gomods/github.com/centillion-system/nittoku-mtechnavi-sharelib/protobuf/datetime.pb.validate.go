// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: datetime.proto

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

// Validate checks the field values on Datetime with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Datetime) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Datetime with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DatetimeMultiError, or nil
// if none found.
func (m *Datetime) ValidateAll() error {
	return m.validate(true)
}

func (m *Datetime) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TimezoneName

	// no validation rules for Timestamp

	// no validation rules for Accuracy

	if len(errors) > 0 {
		return DatetimeMultiError(errors)
	}

	return nil
}

// DatetimeMultiError is an error wrapping multiple validation errors returned
// by Datetime.ValidateAll() if the designated constraints aren't met.
type DatetimeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DatetimeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DatetimeMultiError) AllErrors() []error { return m }

// DatetimeValidationError is the validation error returned by
// Datetime.Validate if the designated constraints aren't met.
type DatetimeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DatetimeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DatetimeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DatetimeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DatetimeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DatetimeValidationError) ErrorName() string { return "DatetimeValidationError" }

// Error satisfies the builtin error interface
func (e DatetimeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDatetime.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DatetimeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DatetimeValidationError{}

// Validate checks the field values on Date with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Date) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Date with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DateMultiError, or nil if none found.
func (m *Date) ValidateAll() error {
	return m.validate(true)
}

func (m *Date) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Year

	// no validation rules for Month

	// no validation rules for Day

	if len(errors) > 0 {
		return DateMultiError(errors)
	}

	return nil
}

// DateMultiError is an error wrapping multiple validation errors returned by
// Date.ValidateAll() if the designated constraints aren't met.
type DateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DateMultiError) AllErrors() []error { return m }

// DateValidationError is the validation error returned by Date.Validate if the
// designated constraints aren't met.
type DateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DateValidationError) ErrorName() string { return "DateValidationError" }

// Error satisfies the builtin error interface
func (e DateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DateValidationError{}

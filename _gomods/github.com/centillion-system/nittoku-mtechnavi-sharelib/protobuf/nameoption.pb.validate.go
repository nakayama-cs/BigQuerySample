// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: nameoption.proto

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

// Validate checks the field values on NameOption with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NameOption) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NameOption with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NameOptionMultiError, or
// nil if none found.
func (m *NameOption) ValidateAll() error {
	return m.validate(true)
}

func (m *NameOption) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CategoryName

	// no validation rules for Code

	// no validation rules for DisplayNameLang

	// no validation rules for SystemName

	if len(errors) > 0 {
		return NameOptionMultiError(errors)
	}

	return nil
}

// NameOptionMultiError is an error wrapping multiple validation errors
// returned by NameOption.ValidateAll() if the designated constraints aren't met.
type NameOptionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NameOptionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NameOptionMultiError) AllErrors() []error { return m }

// NameOptionValidationError is the validation error returned by
// NameOption.Validate if the designated constraints aren't met.
type NameOptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NameOptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NameOptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NameOptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NameOptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NameOptionValidationError) ErrorName() string { return "NameOptionValidationError" }

// Error satisfies the builtin error interface
func (e NameOptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNameOption.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NameOptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NameOptionValidationError{}

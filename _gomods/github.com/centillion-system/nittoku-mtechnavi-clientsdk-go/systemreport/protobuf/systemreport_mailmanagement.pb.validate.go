// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: systemreport_mailmanagement.proto

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

// Validate checks the field values on DailyMailManagement with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DailyMailManagement) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DailyMailManagement with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DailyMailManagementMultiError, or nil if none found.
func (m *DailyMailManagement) ValidateAll() error {
	return m.validate(true)
}

func (m *DailyMailManagement) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DailyMailManagementId

	// no validation rules for GroupId

	// no validation rules for UserId

	for idx, item := range m.GetNotificationSettings() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DailyMailManagementValidationError{
						field:  fmt.Sprintf("NotificationSettings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DailyMailManagementValidationError{
						field:  fmt.Sprintf("NotificationSettings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DailyMailManagementValidationError{
					field:  fmt.Sprintf("NotificationSettings[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	// no validation rules for DeletedAt

	if len(errors) > 0 {
		return DailyMailManagementMultiError(errors)
	}

	return nil
}

// DailyMailManagementMultiError is an error wrapping multiple validation
// errors returned by DailyMailManagement.ValidateAll() if the designated
// constraints aren't met.
type DailyMailManagementMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DailyMailManagementMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DailyMailManagementMultiError) AllErrors() []error { return m }

// DailyMailManagementValidationError is the validation error returned by
// DailyMailManagement.Validate if the designated constraints aren't met.
type DailyMailManagementValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DailyMailManagementValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DailyMailManagementValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DailyMailManagementValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DailyMailManagementValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DailyMailManagementValidationError) ErrorName() string {
	return "DailyMailManagementValidationError"
}

// Error satisfies the builtin error interface
func (e DailyMailManagementValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDailyMailManagement.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DailyMailManagementValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DailyMailManagementValidationError{}

// Validate checks the field values on OneOffMailManagement with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OneOffMailManagement) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OneOffMailManagement with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OneOffMailManagementMultiError, or nil if none found.
func (m *OneOffMailManagement) ValidateAll() error {
	return m.validate(true)
}

func (m *OneOffMailManagement) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OneOffMailManagementId

	// no validation rules for UserId

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	// no validation rules for DeletedAt

	if len(errors) > 0 {
		return OneOffMailManagementMultiError(errors)
	}

	return nil
}

// OneOffMailManagementMultiError is an error wrapping multiple validation
// errors returned by OneOffMailManagement.ValidateAll() if the designated
// constraints aren't met.
type OneOffMailManagementMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OneOffMailManagementMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OneOffMailManagementMultiError) AllErrors() []error { return m }

// OneOffMailManagementValidationError is the validation error returned by
// OneOffMailManagement.Validate if the designated constraints aren't met.
type OneOffMailManagementValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OneOffMailManagementValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OneOffMailManagementValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OneOffMailManagementValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OneOffMailManagementValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OneOffMailManagementValidationError) ErrorName() string {
	return "OneOffMailManagementValidationError"
}

// Error satisfies the builtin error interface
func (e OneOffMailManagementValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOneOffMailManagement.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OneOffMailManagementValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OneOffMailManagementValidationError{}

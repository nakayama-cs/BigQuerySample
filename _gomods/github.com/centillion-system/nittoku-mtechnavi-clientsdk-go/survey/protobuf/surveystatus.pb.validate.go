// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: surveystatus.proto

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

// Validate checks the field values on Complete with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Complete) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Complete with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CompleteMultiError, or nil
// if none found.
func (m *Complete) ValidateAll() error {
	return m.validate(true)
}

func (m *Complete) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CompleteId

	// no validation rules for Urn

	// no validation rules for TypeName

	// no validation rules for RecordId

	// no validation rules for CompletedAt

	if all {
		switch v := interface{}(m.GetSharedProperties()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CompleteValidationError{
					field:  "SharedProperties",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CompleteValidationError{
					field:  "SharedProperties",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSharedProperties()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CompleteValidationError{
				field:  "SharedProperties",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	// no validation rules for DeletedAt

	if len(errors) > 0 {
		return CompleteMultiError(errors)
	}

	return nil
}

// CompleteMultiError is an error wrapping multiple validation errors returned
// by Complete.ValidateAll() if the designated constraints aren't met.
type CompleteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CompleteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CompleteMultiError) AllErrors() []error { return m }

// CompleteValidationError is the validation error returned by
// Complete.Validate if the designated constraints aren't met.
type CompleteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CompleteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CompleteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CompleteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CompleteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CompleteValidationError) ErrorName() string { return "CompleteValidationError" }

// Error satisfies the builtin error interface
func (e CompleteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sComplete.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CompleteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CompleteValidationError{}

// Validate checks the field values on Discard with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Discard) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Discard with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in DiscardMultiError, or nil if none found.
func (m *Discard) ValidateAll() error {
	return m.validate(true)
}

func (m *Discard) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DiscardId

	// no validation rules for Urn

	// no validation rules for TypeName

	// no validation rules for RecordId

	// no validation rules for Comment

	// no validation rules for DiscardedAt

	if all {
		switch v := interface{}(m.GetSharedProperties()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DiscardValidationError{
					field:  "SharedProperties",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DiscardValidationError{
					field:  "SharedProperties",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSharedProperties()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DiscardValidationError{
				field:  "SharedProperties",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	// no validation rules for DeletedAt

	if len(errors) > 0 {
		return DiscardMultiError(errors)
	}

	return nil
}

// DiscardMultiError is an error wrapping multiple validation errors returned
// by Discard.ValidateAll() if the designated constraints aren't met.
type DiscardMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DiscardMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DiscardMultiError) AllErrors() []error { return m }

// DiscardValidationError is the validation error returned by Discard.Validate
// if the designated constraints aren't met.
type DiscardValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DiscardValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DiscardValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DiscardValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DiscardValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DiscardValidationError) ErrorName() string { return "DiscardValidationError" }

// Error satisfies the builtin error interface
func (e DiscardValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDiscard.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DiscardValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DiscardValidationError{}

// Validate checks the field values on Opened with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Opened) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Opened with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in OpenedMultiError, or nil if none found.
func (m *Opened) ValidateAll() error {
	return m.validate(true)
}

func (m *Opened) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OpenedId

	// no validation rules for Urn

	// no validation rules for TypeName

	// no validation rules for RecordId

	if all {
		switch v := interface{}(m.GetOpenUserProcess()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OpenedValidationError{
					field:  "OpenUserProcess",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OpenedValidationError{
					field:  "OpenUserProcess",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOpenUserProcess()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenedValidationError{
				field:  "OpenUserProcess",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetSharedProperties()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OpenedValidationError{
					field:  "SharedProperties",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OpenedValidationError{
					field:  "SharedProperties",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSharedProperties()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenedValidationError{
				field:  "SharedProperties",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	// no validation rules for DeletedAt

	if len(errors) > 0 {
		return OpenedMultiError(errors)
	}

	return nil
}

// OpenedMultiError is an error wrapping multiple validation errors returned by
// Opened.ValidateAll() if the designated constraints aren't met.
type OpenedMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OpenedMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OpenedMultiError) AllErrors() []error { return m }

// OpenedValidationError is the validation error returned by Opened.Validate if
// the designated constraints aren't met.
type OpenedValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpenedValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpenedValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpenedValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpenedValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpenedValidationError) ErrorName() string { return "OpenedValidationError" }

// Error satisfies the builtin error interface
func (e OpenedValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpened.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpenedValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpenedValidationError{}

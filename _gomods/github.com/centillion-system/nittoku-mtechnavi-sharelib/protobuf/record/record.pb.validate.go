// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: record/record.proto

package record

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

// Validate checks the field values on Record with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Record) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Record with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RecordMultiError, or nil if none found.
func (m *Record) ValidateAll() error {
	return m.validate(true)
}

func (m *Record) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RecordId

	if utf8.RuneCountInString(m.GetTypeName()) < 1 {
		err := RecordValidationError{
			field:  "TypeName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	{
		sorted_keys := make([]int32, len(m.GetFields()))
		i := 0
		for key := range m.GetFields() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetFields()[key]
			_ = val

			// no validation rules for Fields[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, RecordValidationError{
							field:  fmt.Sprintf("Fields[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, RecordValidationError{
							field:  fmt.Sprintf("Fields[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return RecordValidationError{
						field:  fmt.Sprintf("Fields[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	for idx, item := range m.GetIndexs() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RecordValidationError{
						field:  fmt.Sprintf("Indexs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RecordValidationError{
						field:  fmt.Sprintf("Indexs[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RecordValidationError{
					field:  fmt.Sprintf("Indexs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetSharedProperties()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RecordValidationError{
					field:  "SharedProperties",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RecordValidationError{
					field:  "SharedProperties",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSharedProperties()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RecordValidationError{
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
		return RecordMultiError(errors)
	}

	return nil
}

// RecordMultiError is an error wrapping multiple validation errors returned by
// Record.ValidateAll() if the designated constraints aren't met.
type RecordMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RecordMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RecordMultiError) AllErrors() []error { return m }

// RecordValidationError is the validation error returned by Record.Validate if
// the designated constraints aren't met.
type RecordValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RecordValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RecordValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RecordValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RecordValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RecordValidationError) ErrorName() string { return "RecordValidationError" }

// Error satisfies the builtin error interface
func (e RecordValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRecord.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RecordValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RecordValidationError{}

// Validate checks the field values on Field with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Field) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Field with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in FieldMultiError, or nil if none found.
func (m *Field) ValidateAll() error {
	return m.validate(true)
}

func (m *Field) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Value

	if all {
		switch v := interface{}(m.GetVisibility()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FieldValidationError{
					field:  "Visibility",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FieldValidationError{
					field:  "Visibility",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetVisibility()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FieldValidationError{
				field:  "Visibility",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return FieldMultiError(errors)
	}

	return nil
}

// FieldMultiError is an error wrapping multiple validation errors returned by
// Field.ValidateAll() if the designated constraints aren't met.
type FieldMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FieldMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FieldMultiError) AllErrors() []error { return m }

// FieldValidationError is the validation error returned by Field.Validate if
// the designated constraints aren't met.
type FieldValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FieldValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FieldValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FieldValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FieldValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FieldValidationError) ErrorName() string { return "FieldValidationError" }

// Error satisfies the builtin error interface
func (e FieldValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sField.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FieldValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FieldValidationError{}

// Validate checks the field values on Index with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Index) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Index with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in IndexMultiError, or nil if none found.
func (m *Index) ValidateAll() error {
	return m.validate(true)
}

func (m *Index) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Value

	if len(errors) > 0 {
		return IndexMultiError(errors)
	}

	return nil
}

// IndexMultiError is an error wrapping multiple validation errors returned by
// Index.ValidateAll() if the designated constraints aren't met.
type IndexMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndexMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndexMultiError) AllErrors() []error { return m }

// IndexValidationError is the validation error returned by Index.Validate if
// the designated constraints aren't met.
type IndexValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndexValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndexValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndexValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndexValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndexValidationError) ErrorName() string { return "IndexValidationError" }

// Error satisfies the builtin error interface
func (e IndexValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndex.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndexValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndexValidationError{}

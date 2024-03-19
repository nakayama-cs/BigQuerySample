// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: tenantprovision.proto

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

// Validate checks the field values on GetTenantRequestRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTenantRequestRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTenantRequestRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTenantRequestRequestMultiError, or nil if none found.
func (m *GetTenantRequestRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTenantRequestRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TenantRequestId

	if len(errors) > 0 {
		return GetTenantRequestRequestMultiError(errors)
	}

	return nil
}

// GetTenantRequestRequestMultiError is an error wrapping multiple validation
// errors returned by GetTenantRequestRequest.ValidateAll() if the designated
// constraints aren't met.
type GetTenantRequestRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTenantRequestRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTenantRequestRequestMultiError) AllErrors() []error { return m }

// GetTenantRequestRequestValidationError is the validation error returned by
// GetTenantRequestRequest.Validate if the designated constraints aren't met.
type GetTenantRequestRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTenantRequestRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTenantRequestRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTenantRequestRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTenantRequestRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTenantRequestRequestValidationError) ErrorName() string {
	return "GetTenantRequestRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTenantRequestRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTenantRequestRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTenantRequestRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTenantRequestRequestValidationError{}

// Validate checks the field values on GetTenantRequestResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetTenantRequestResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTenantRequestResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTenantRequestResponseMultiError, or nil if none found.
func (m *GetTenantRequestResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTenantRequestResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTenantRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetTenantRequestResponseValidationError{
					field:  "TenantRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetTenantRequestResponseValidationError{
					field:  "TenantRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTenantRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTenantRequestResponseValidationError{
				field:  "TenantRequest",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetTenantRequestResponseMultiError(errors)
	}

	return nil
}

// GetTenantRequestResponseMultiError is an error wrapping multiple validation
// errors returned by GetTenantRequestResponse.ValidateAll() if the designated
// constraints aren't met.
type GetTenantRequestResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTenantRequestResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTenantRequestResponseMultiError) AllErrors() []error { return m }

// GetTenantRequestResponseValidationError is the validation error returned by
// GetTenantRequestResponse.Validate if the designated constraints aren't met.
type GetTenantRequestResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTenantRequestResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTenantRequestResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTenantRequestResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTenantRequestResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTenantRequestResponseValidationError) ErrorName() string {
	return "GetTenantRequestResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTenantRequestResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTenantRequestResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTenantRequestResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTenantRequestResponseValidationError{}

// Validate checks the field values on VerifyTenantRequestRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyTenantRequestRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyTenantRequestRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyTenantRequestRequestMultiError, or nil if none found.
func (m *VerifyTenantRequestRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyTenantRequestRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTenantRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, VerifyTenantRequestRequestValidationError{
					field:  "TenantRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, VerifyTenantRequestRequestValidationError{
					field:  "TenantRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTenantRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return VerifyTenantRequestRequestValidationError{
				field:  "TenantRequest",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return VerifyTenantRequestRequestMultiError(errors)
	}

	return nil
}

// VerifyTenantRequestRequestMultiError is an error wrapping multiple
// validation errors returned by VerifyTenantRequestRequest.ValidateAll() if
// the designated constraints aren't met.
type VerifyTenantRequestRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyTenantRequestRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyTenantRequestRequestMultiError) AllErrors() []error { return m }

// VerifyTenantRequestRequestValidationError is the validation error returned
// by VerifyTenantRequestRequest.Validate if the designated constraints aren't met.
type VerifyTenantRequestRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyTenantRequestRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyTenantRequestRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyTenantRequestRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyTenantRequestRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyTenantRequestRequestValidationError) ErrorName() string {
	return "VerifyTenantRequestRequestValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyTenantRequestRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyTenantRequestRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyTenantRequestRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyTenantRequestRequestValidationError{}

// Validate checks the field values on VerifyTenantRequestResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *VerifyTenantRequestResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on VerifyTenantRequestResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// VerifyTenantRequestResponseMultiError, or nil if none found.
func (m *VerifyTenantRequestResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *VerifyTenantRequestResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return VerifyTenantRequestResponseMultiError(errors)
	}

	return nil
}

// VerifyTenantRequestResponseMultiError is an error wrapping multiple
// validation errors returned by VerifyTenantRequestResponse.ValidateAll() if
// the designated constraints aren't met.
type VerifyTenantRequestResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VerifyTenantRequestResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VerifyTenantRequestResponseMultiError) AllErrors() []error { return m }

// VerifyTenantRequestResponseValidationError is the validation error returned
// by VerifyTenantRequestResponse.Validate if the designated constraints
// aren't met.
type VerifyTenantRequestResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VerifyTenantRequestResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VerifyTenantRequestResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VerifyTenantRequestResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VerifyTenantRequestResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VerifyTenantRequestResponseValidationError) ErrorName() string {
	return "VerifyTenantRequestResponseValidationError"
}

// Error satisfies the builtin error interface
func (e VerifyTenantRequestResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVerifyTenantRequestResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VerifyTenantRequestResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VerifyTenantRequestResponseValidationError{}

// Validate checks the field values on ApplyTenantRequestRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ApplyTenantRequestRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ApplyTenantRequestRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ApplyTenantRequestRequestMultiError, or nil if none found.
func (m *ApplyTenantRequestRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ApplyTenantRequestRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TenantRequestId

	// no validation rules for VerifyTenantRequestId

	if len(errors) > 0 {
		return ApplyTenantRequestRequestMultiError(errors)
	}

	return nil
}

// ApplyTenantRequestRequestMultiError is an error wrapping multiple validation
// errors returned by ApplyTenantRequestRequest.ValidateAll() if the
// designated constraints aren't met.
type ApplyTenantRequestRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApplyTenantRequestRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApplyTenantRequestRequestMultiError) AllErrors() []error { return m }

// ApplyTenantRequestRequestValidationError is the validation error returned by
// ApplyTenantRequestRequest.Validate if the designated constraints aren't met.
type ApplyTenantRequestRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplyTenantRequestRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplyTenantRequestRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplyTenantRequestRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplyTenantRequestRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplyTenantRequestRequestValidationError) ErrorName() string {
	return "ApplyTenantRequestRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ApplyTenantRequestRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplyTenantRequestRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplyTenantRequestRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplyTenantRequestRequestValidationError{}

// Validate checks the field values on ApplyTenantRequestResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ApplyTenantRequestResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ApplyTenantRequestResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ApplyTenantRequestResponseMultiError, or nil if none found.
func (m *ApplyTenantRequestResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ApplyTenantRequestResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ApplyTenantRequestResponseMultiError(errors)
	}

	return nil
}

// ApplyTenantRequestResponseMultiError is an error wrapping multiple
// validation errors returned by ApplyTenantRequestResponse.ValidateAll() if
// the designated constraints aren't met.
type ApplyTenantRequestResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ApplyTenantRequestResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ApplyTenantRequestResponseMultiError) AllErrors() []error { return m }

// ApplyTenantRequestResponseValidationError is the validation error returned
// by ApplyTenantRequestResponse.Validate if the designated constraints aren't met.
type ApplyTenantRequestResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ApplyTenantRequestResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ApplyTenantRequestResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ApplyTenantRequestResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ApplyTenantRequestResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ApplyTenantRequestResponseValidationError) ErrorName() string {
	return "ApplyTenantRequestResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ApplyTenantRequestResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sApplyTenantRequestResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ApplyTenantRequestResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ApplyTenantRequestResponseValidationError{}

// Validate checks the field values on TenantRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TenantRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TenantRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TenantRequestMultiError, or
// nil if none found.
func (m *TenantRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TenantRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TenantRequestId

	if all {
		switch v := interface{}(m.GetTenant()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TenantRequestValidationError{
					field:  "Tenant",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TenantRequestValidationError{
					field:  "Tenant",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTenant()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TenantRequestValidationError{
				field:  "Tenant",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTenantAdmin()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TenantRequestValidationError{
					field:  "TenantAdmin",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TenantRequestValidationError{
					field:  "TenantAdmin",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTenantAdmin()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TenantRequestValidationError{
				field:  "TenantAdmin",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for InvitedAt

	// no validation rules for RequestVerifiedAt

	// no validation rules for InviteBy

	// no validation rules for ReleateBusinessUnitManagementId

	if len(errors) > 0 {
		return TenantRequestMultiError(errors)
	}

	return nil
}

// TenantRequestMultiError is an error wrapping multiple validation errors
// returned by TenantRequest.ValidateAll() if the designated constraints
// aren't met.
type TenantRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TenantRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TenantRequestMultiError) AllErrors() []error { return m }

// TenantRequestValidationError is the validation error returned by
// TenantRequest.Validate if the designated constraints aren't met.
type TenantRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TenantRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TenantRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TenantRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TenantRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TenantRequestValidationError) ErrorName() string { return "TenantRequestValidationError" }

// Error satisfies the builtin error interface
func (e TenantRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTenantRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TenantRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TenantRequestValidationError{}

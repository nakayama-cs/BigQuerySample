// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: worktaskcatalog.proto

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

// Validate checks the field values on WorkTaskCatalog with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *WorkTaskCatalog) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on WorkTaskCatalog with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// WorkTaskCatalogMultiError, or nil if none found.
func (m *WorkTaskCatalog) ValidateAll() error {
	return m.validate(true)
}

func (m *WorkTaskCatalog) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for WorkTaskCatalogId

	// no validation rules for Code

	// no validation rules for DisplayName

	if all {
		switch v := interface{}(m.GetCategory()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WorkTaskCatalogValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WorkTaskCatalogValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCategory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkTaskCatalogValidationError{
				field:  "Category",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Contents

	// no validation rules for Approval

	if all {
		switch v := interface{}(m.GetManagementOrganization()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WorkTaskCatalogValidationError{
					field:  "ManagementOrganization",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WorkTaskCatalogValidationError{
					field:  "ManagementOrganization",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetManagementOrganization()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkTaskCatalogValidationError{
				field:  "ManagementOrganization",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetWorkOrganization()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WorkTaskCatalogValidationError{
					field:  "WorkOrganization",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WorkTaskCatalogValidationError{
					field:  "WorkOrganization",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWorkOrganization()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkTaskCatalogValidationError{
				field:  "WorkOrganization",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetCommonAttachments() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, WorkTaskCatalogValidationError{
						field:  fmt.Sprintf("CommonAttachments[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, WorkTaskCatalogValidationError{
						field:  fmt.Sprintf("CommonAttachments[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return WorkTaskCatalogValidationError{
					field:  fmt.Sprintf("CommonAttachments[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for WorkTaskCatalogCatalogUpdatedAt

	if all {
		switch v := interface{}(m.GetWorkTaskCatalogCatalogUpdatedBy()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, WorkTaskCatalogValidationError{
					field:  "WorkTaskCatalogCatalogUpdatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, WorkTaskCatalogValidationError{
					field:  "WorkTaskCatalogCatalogUpdatedBy",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetWorkTaskCatalogCatalogUpdatedBy()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WorkTaskCatalogValidationError{
				field:  "WorkTaskCatalogCatalogUpdatedBy",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for CreatedAt

	// no validation rules for UpdatedAt

	// no validation rules for DeletedAt

	if len(errors) > 0 {
		return WorkTaskCatalogMultiError(errors)
	}

	return nil
}

// WorkTaskCatalogMultiError is an error wrapping multiple validation errors
// returned by WorkTaskCatalog.ValidateAll() if the designated constraints
// aren't met.
type WorkTaskCatalogMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m WorkTaskCatalogMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m WorkTaskCatalogMultiError) AllErrors() []error { return m }

// WorkTaskCatalogValidationError is the validation error returned by
// WorkTaskCatalog.Validate if the designated constraints aren't met.
type WorkTaskCatalogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WorkTaskCatalogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WorkTaskCatalogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WorkTaskCatalogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WorkTaskCatalogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WorkTaskCatalogValidationError) ErrorName() string { return "WorkTaskCatalogValidationError" }

// Error satisfies the builtin error interface
func (e WorkTaskCatalogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWorkTaskCatalog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WorkTaskCatalogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WorkTaskCatalogValidationError{}

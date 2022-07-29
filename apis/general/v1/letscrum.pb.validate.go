// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/general/v1/letscrum.proto

package v1

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

// Validate checks the field values on Version with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Version) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Version with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in VersionMultiError, or nil if none found.
func (m *Version) ValidateAll() error {
	return m.validate(true)
}

func (m *Version) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Version

	// no validation rules for GitCommit

	// no validation rules for BuildDate

	// no validation rules for GoVersion

	if len(errors) > 0 {
		return VersionMultiError(errors)
	}

	return nil
}

// VersionMultiError is an error wrapping multiple validation errors returned
// by Version.ValidateAll() if the designated constraints aren't met.
type VersionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m VersionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m VersionMultiError) AllErrors() []error { return m }

// VersionValidationError is the validation error returned by Version.Validate
// if the designated constraints aren't met.
type VersionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e VersionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e VersionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e VersionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e VersionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e VersionValidationError) ErrorName() string { return "VersionValidationError" }

// Error satisfies the builtin error interface
func (e VersionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sVersion.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = VersionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = VersionValidationError{}

// Validate checks the field values on GetVersionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetVersionResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetVersionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetVersionResponseMultiError, or nil if none found.
func (m *GetVersionResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetVersionResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetVersion()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetVersionResponseValidationError{
					field:  "Version",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetVersionResponseValidationError{
					field:  "Version",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetVersion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetVersionResponseValidationError{
				field:  "Version",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetVersionResponseMultiError(errors)
	}

	return nil
}

// GetVersionResponseMultiError is an error wrapping multiple validation errors
// returned by GetVersionResponse.ValidateAll() if the designated constraints
// aren't met.
type GetVersionResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetVersionResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetVersionResponseMultiError) AllErrors() []error { return m }

// GetVersionResponseValidationError is the validation error returned by
// GetVersionResponse.Validate if the designated constraints aren't met.
type GetVersionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetVersionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetVersionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetVersionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetVersionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetVersionResponseValidationError) ErrorName() string {
	return "GetVersionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetVersionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetVersionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetVersionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetVersionResponseValidationError{}

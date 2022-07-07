// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/quotes/v1/quotes.proto

package v1Quotes

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

	v1Screener "github.com/arktos-venture/buf/proto/screener/v1"
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

	_ = v1Screener.TSDB(0)
)

// Validate checks the field values on QuoteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuoteRequestMultiError, or
// nil if none found.
func (m *QuoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QuoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := v1Screener.TSDB_name[int32(m.GetTsdb())]; !ok {
		err := QuoteRequestValidationError{
			field:  "Tsdb",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := v1Screener.Interval_name[int32(m.GetInterval())]; !ok {
		err := QuoteRequestValidationError{
			field:  "Interval",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := len(m.GetFilters()); l < 1 || l > 20 {
		err := QuoteRequestValidationError{
			field:  "Filters",
			reason: "value must contain between 1 and 20 items, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetFilters() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuoteRequestValidationError{
						field:  fmt.Sprintf("Filters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuoteRequestValidationError{
						field:  fmt.Sprintf("Filters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuoteRequestValidationError{
					field:  fmt.Sprintf("Filters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QuoteRequestMultiError(errors)
	}

	return nil
}

// QuoteRequestMultiError is an error wrapping multiple validation errors
// returned by QuoteRequest.ValidateAll() if the designated constraints aren't met.
type QuoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteRequestMultiError) AllErrors() []error { return m }

// QuoteRequestValidationError is the validation error returned by
// QuoteRequest.Validate if the designated constraints aren't met.
type QuoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteRequestValidationError) ErrorName() string { return "QuoteRequestValidationError" }

// Error satisfies the builtin error interface
func (e QuoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteRequestValidationError{}

// Validate checks the field values on QuoteLastRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *QuoteLastRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuoteLastRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuoteLastRequestMultiError, or nil if none found.
func (m *QuoteLastRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QuoteLastRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := v1Screener.TSDB_name[int32(m.GetTsdb())]; !ok {
		err := QuoteLastRequestValidationError{
			field:  "Tsdb",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Exchange

	// no validation rules for Ticker

	if len(errors) > 0 {
		return QuoteLastRequestMultiError(errors)
	}

	return nil
}

// QuoteLastRequestMultiError is an error wrapping multiple validation errors
// returned by QuoteLastRequest.ValidateAll() if the designated constraints
// aren't met.
type QuoteLastRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteLastRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteLastRequestMultiError) AllErrors() []error { return m }

// QuoteLastRequestValidationError is the validation error returned by
// QuoteLastRequest.Validate if the designated constraints aren't met.
type QuoteLastRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteLastRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteLastRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteLastRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteLastRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteLastRequestValidationError) ErrorName() string { return "QuoteLastRequestValidationError" }

// Error satisfies the builtin error interface
func (e QuoteLastRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuoteLastRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteLastRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteLastRequestValidationError{}

// Validate checks the field values on QuoteDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QuoteDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuoteDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuoteDeleteRequestMultiError, or nil if none found.
func (m *QuoteDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QuoteDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := v1Screener.TSDB_name[int32(m.GetTsdb())]; !ok {
		err := QuoteDeleteRequestValidationError{
			field:  "Tsdb",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return QuoteDeleteRequestMultiError(errors)
	}

	return nil
}

// QuoteDeleteRequestMultiError is an error wrapping multiple validation errors
// returned by QuoteDeleteRequest.ValidateAll() if the designated constraints
// aren't met.
type QuoteDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteDeleteRequestMultiError) AllErrors() []error { return m }

// QuoteDeleteRequestValidationError is the validation error returned by
// QuoteDeleteRequest.Validate if the designated constraints aren't met.
type QuoteDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteDeleteRequestValidationError) ErrorName() string {
	return "QuoteDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e QuoteDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuoteDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteDeleteRequestValidationError{}

// Validate checks the field values on QuoteReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuoteReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuoteReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuoteReplyMultiError, or
// nil if none found.
func (m *QuoteReply) ValidateAll() error {
	return m.validate(true)
}

func (m *QuoteReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCreatedAt() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuoteReplyValidationError{
						field:  fmt.Sprintf("CreatedAt[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuoteReplyValidationError{
						field:  fmt.Sprintf("CreatedAt[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuoteReplyValidationError{
					field:  fmt.Sprintf("CreatedAt[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QuoteReplyMultiError(errors)
	}

	return nil
}

// QuoteReplyMultiError is an error wrapping multiple validation errors
// returned by QuoteReply.ValidateAll() if the designated constraints aren't met.
type QuoteReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteReplyMultiError) AllErrors() []error { return m }

// QuoteReplyValidationError is the validation error returned by
// QuoteReply.Validate if the designated constraints aren't met.
type QuoteReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteReplyValidationError) ErrorName() string { return "QuoteReplyValidationError" }

// Error satisfies the builtin error interface
func (e QuoteReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuoteReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteReplyValidationError{}

// Validate checks the field values on QuoteLastReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuoteLastReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuoteLastReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuoteLastReplyMultiError,
// or nil if none found.
func (m *QuoteLastReply) ValidateAll() error {
	return m.validate(true)
}

func (m *QuoteLastReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Open

	// no validation rules for Close

	// no validation rules for Low

	// no validation rules for High

	// no validation rules for Volume

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuoteLastReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuoteLastReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuoteLastReplyValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return QuoteLastReplyMultiError(errors)
	}

	return nil
}

// QuoteLastReplyMultiError is an error wrapping multiple validation errors
// returned by QuoteLastReply.ValidateAll() if the designated constraints
// aren't met.
type QuoteLastReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteLastReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteLastReplyMultiError) AllErrors() []error { return m }

// QuoteLastReplyValidationError is the validation error returned by
// QuoteLastReply.Validate if the designated constraints aren't met.
type QuoteLastReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteLastReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteLastReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteLastReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteLastReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteLastReplyValidationError) ErrorName() string { return "QuoteLastReplyValidationError" }

// Error satisfies the builtin error interface
func (e QuoteLastReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuoteLastReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteLastReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteLastReplyValidationError{}

// Validate checks the field values on QuoteDelete with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuoteDelete) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuoteDelete with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuoteDeleteMultiError, or
// nil if none found.
func (m *QuoteDelete) ValidateAll() error {
	return m.validate(true)
}

func (m *QuoteDelete) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	if len(errors) > 0 {
		return QuoteDeleteMultiError(errors)
	}

	return nil
}

// QuoteDeleteMultiError is an error wrapping multiple validation errors
// returned by QuoteDelete.ValidateAll() if the designated constraints aren't met.
type QuoteDeleteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteDeleteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteDeleteMultiError) AllErrors() []error { return m }

// QuoteDeleteValidationError is the validation error returned by
// QuoteDelete.Validate if the designated constraints aren't met.
type QuoteDeleteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteDeleteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteDeleteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteDeleteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteDeleteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteDeleteValidationError) ErrorName() string { return "QuoteDeleteValidationError" }

// Error satisfies the builtin error interface
func (e QuoteDeleteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuoteDelete.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteDeleteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteDeleteValidationError{}

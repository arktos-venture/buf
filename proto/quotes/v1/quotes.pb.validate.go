// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/quotes/v1/quotes.proto

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

// Validate checks the field values on Filter with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Filter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Filter with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in FilterMultiError, or nil if none found.
func (m *Filter) ValidateAll() error {
	return m.validate(true)
}

func (m *Filter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := Operator_name[int32(m.GetOperator())]; !ok {
		err := FilterValidationError{
			field:  "Operator",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := Argument_name[int32(m.GetArgument())]; !ok {
		err := FilterValidationError{
			field:  "Argument",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Value

	if len(errors) > 0 {
		return FilterMultiError(errors)
	}

	return nil
}

// FilterMultiError is an error wrapping multiple validation errors returned by
// Filter.ValidateAll() if the designated constraints aren't met.
type FilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FilterMultiError) AllErrors() []error { return m }

// FilterValidationError is the validation error returned by Filter.Validate if
// the designated constraints aren't met.
type FilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterValidationError) ErrorName() string { return "FilterValidationError" }

// Error satisfies the builtin error interface
func (e FilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterValidationError{}

// Validate checks the field values on QuotesRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuotesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuotesRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuotesRequestMultiError, or
// nil if none found.
func (m *QuotesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QuotesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := TSDB_name[int32(m.GetTsdb())]; !ok {
		err := QuotesRequestValidationError{
			field:  "Tsdb",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := Interval_name[int32(m.GetInterval())]; !ok {
		err := QuotesRequestValidationError{
			field:  "Interval",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := len(m.GetFilters()); l < 1 || l > 20 {
		err := QuotesRequestValidationError{
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
					errors = append(errors, QuotesRequestValidationError{
						field:  fmt.Sprintf("Filters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuotesRequestValidationError{
						field:  fmt.Sprintf("Filters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuotesRequestValidationError{
					field:  fmt.Sprintf("Filters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QuotesRequestMultiError(errors)
	}

	return nil
}

// QuotesRequestMultiError is an error wrapping multiple validation errors
// returned by QuotesRequest.ValidateAll() if the designated constraints
// aren't met.
type QuotesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuotesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuotesRequestMultiError) AllErrors() []error { return m }

// QuotesRequestValidationError is the validation error returned by
// QuotesRequest.Validate if the designated constraints aren't met.
type QuotesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuotesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuotesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuotesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuotesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuotesRequestValidationError) ErrorName() string { return "QuotesRequestValidationError" }

// Error satisfies the builtin error interface
func (e QuotesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuotesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuotesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuotesRequestValidationError{}

// Validate checks the field values on QuotesReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuotesReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuotesReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuotesReplyMultiError, or
// nil if none found.
func (m *QuotesReply) ValidateAll() error {
	return m.validate(true)
}

func (m *QuotesReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetDate() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuotesReplyValidationError{
						field:  fmt.Sprintf("Date[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuotesReplyValidationError{
						field:  fmt.Sprintf("Date[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuotesReplyValidationError{
					field:  fmt.Sprintf("Date[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QuotesReplyMultiError(errors)
	}

	return nil
}

// QuotesReplyMultiError is an error wrapping multiple validation errors
// returned by QuotesReply.ValidateAll() if the designated constraints aren't met.
type QuotesReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuotesReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuotesReplyMultiError) AllErrors() []error { return m }

// QuotesReplyValidationError is the validation error returned by
// QuotesReply.Validate if the designated constraints aren't met.
type QuotesReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuotesReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuotesReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuotesReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuotesReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuotesReplyValidationError) ErrorName() string { return "QuotesReplyValidationError" }

// Error satisfies the builtin error interface
func (e QuotesReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuotesReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuotesReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuotesReplyValidationError{}

// Validate checks the field values on QuotesRequest_Sort with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QuotesRequest_Sort) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuotesRequest_Sort with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuotesRequest_SortMultiError, or nil if none found.
func (m *QuotesRequest_Sort) ValidateAll() error {
	return m.validate(true)
}

func (m *QuotesRequest_Sort) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := Argument_name[int32(m.GetArgument())]; !ok {
		err := QuotesRequest_SortValidationError{
			field:  "Argument",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := QuotesRequest_Sort_Orientation_name[int32(m.GetOrientation())]; !ok {
		err := QuotesRequest_SortValidationError{
			field:  "Orientation",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return QuotesRequest_SortMultiError(errors)
	}

	return nil
}

// QuotesRequest_SortMultiError is an error wrapping multiple validation errors
// returned by QuotesRequest_Sort.ValidateAll() if the designated constraints
// aren't met.
type QuotesRequest_SortMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuotesRequest_SortMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuotesRequest_SortMultiError) AllErrors() []error { return m }

// QuotesRequest_SortValidationError is the validation error returned by
// QuotesRequest_Sort.Validate if the designated constraints aren't met.
type QuotesRequest_SortValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuotesRequest_SortValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuotesRequest_SortValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuotesRequest_SortValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuotesRequest_SortValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuotesRequest_SortValidationError) ErrorName() string {
	return "QuotesRequest_SortValidationError"
}

// Error satisfies the builtin error interface
func (e QuotesRequest_SortValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuotesRequest_Sort.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuotesRequest_SortValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuotesRequest_SortValidationError{}

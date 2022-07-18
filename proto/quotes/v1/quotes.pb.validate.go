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

	_ = v1Screener.Interval(0)
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

	if l := utf8.RuneCountInString(m.GetTicker()); l < 1 || l > 16 {
		err := QuoteRequestValidationError{
			field:  "Ticker",
			reason: "value length must be between 1 and 16 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetExchange()); l < 1 || l > 16 {
		err := QuoteRequestValidationError{
			field:  "Exchange",
			reason: "value length must be between 1 and 16 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
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

// Validate checks the field values on QuoteSearchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *QuoteSearchRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuoteSearchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// QuoteSearchRequestMultiError, or nil if none found.
func (m *QuoteSearchRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QuoteSearchRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := v1Screener.Interval_name[int32(m.GetInterval())]; !ok {
		err := QuoteSearchRequestValidationError{
			field:  "Interval",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := len(m.GetFilters()); l < 1 || l > 20 {
		err := QuoteSearchRequestValidationError{
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
					errors = append(errors, QuoteSearchRequestValidationError{
						field:  fmt.Sprintf("Filters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuoteSearchRequestValidationError{
						field:  fmt.Sprintf("Filters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuoteSearchRequestValidationError{
					field:  fmt.Sprintf("Filters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QuoteSearchRequestMultiError(errors)
	}

	return nil
}

// QuoteSearchRequestMultiError is an error wrapping multiple validation errors
// returned by QuoteSearchRequest.ValidateAll() if the designated constraints
// aren't met.
type QuoteSearchRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteSearchRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteSearchRequestMultiError) AllErrors() []error { return m }

// QuoteSearchRequestValidationError is the validation error returned by
// QuoteSearchRequest.Validate if the designated constraints aren't met.
type QuoteSearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteSearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteSearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteSearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteSearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteSearchRequestValidationError) ErrorName() string {
	return "QuoteSearchRequestValidationError"
}

// Error satisfies the builtin error interface
func (e QuoteSearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuoteSearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteSearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteSearchRequestValidationError{}

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

	// no validation rules for Open

	// no validation rules for Close

	// no validation rules for Low

	// no validation rules for High

	// no validation rules for Volume

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuoteReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuoteReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuoteReplyValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
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

// Validate checks the field values on QuoteReplies with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuoteReplies) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuoteReplies with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuoteRepliesMultiError, or
// nil if none found.
func (m *QuoteReplies) ValidateAll() error {
	return m.validate(true)
}

func (m *QuoteReplies) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetResults() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuoteRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuoteRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuoteRepliesValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return QuoteRepliesMultiError(errors)
	}

	return nil
}

// QuoteRepliesMultiError is an error wrapping multiple validation errors
// returned by QuoteReplies.ValidateAll() if the designated constraints aren't met.
type QuoteRepliesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteRepliesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteRepliesMultiError) AllErrors() []error { return m }

// QuoteRepliesValidationError is the validation error returned by
// QuoteReplies.Validate if the designated constraints aren't met.
type QuoteRepliesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteRepliesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteRepliesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteRepliesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteRepliesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteRepliesValidationError) ErrorName() string { return "QuoteRepliesValidationError" }

// Error satisfies the builtin error interface
func (e QuoteRepliesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuoteReplies.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteRepliesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteRepliesValidationError{}

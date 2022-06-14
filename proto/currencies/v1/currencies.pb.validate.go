// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/currencies/v1/currencies.proto

package currencies_v1

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

// Validate checks the field values on CurrencyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CurrencyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CurrencyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CurrencyRequestMultiError, or nil if none found.
func (m *CurrencyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CurrencyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTicker()) != 3 {
		err := CurrencyRequestValidationError{
			field:  "Ticker",
			reason: "value length must be 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return CurrencyRequestMultiError(errors)
	}

	return nil
}

// CurrencyRequestMultiError is an error wrapping multiple validation errors
// returned by CurrencyRequest.ValidateAll() if the designated constraints
// aren't met.
type CurrencyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CurrencyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CurrencyRequestMultiError) AllErrors() []error { return m }

// CurrencyRequestValidationError is the validation error returned by
// CurrencyRequest.Validate if the designated constraints aren't met.
type CurrencyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CurrencyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CurrencyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CurrencyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CurrencyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CurrencyRequestValidationError) ErrorName() string { return "CurrencyRequestValidationError" }

// Error satisfies the builtin error interface
func (e CurrencyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCurrencyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CurrencyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CurrencyRequestValidationError{}

// Validate checks the field values on CurrencyDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CurrencyDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CurrencyDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CurrencyDeleteRequestMultiError, or nil if none found.
func (m *CurrencyDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CurrencyDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetTickers()) < 1 {
		err := CurrencyDeleteRequestValidationError{
			field:  "Tickers",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	_CurrencyDeleteRequest_Tickers_Unique := make(map[string]struct{}, len(m.GetTickers()))

	for idx, item := range m.GetTickers() {
		_, _ = idx, item

		if _, exists := _CurrencyDeleteRequest_Tickers_Unique[item]; exists {
			err := CurrencyDeleteRequestValidationError{
				field:  fmt.Sprintf("Tickers[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_CurrencyDeleteRequest_Tickers_Unique[item] = struct{}{}
		}

		// no validation rules for Tickers[idx]
	}

	if len(errors) > 0 {
		return CurrencyDeleteRequestMultiError(errors)
	}

	return nil
}

// CurrencyDeleteRequestMultiError is an error wrapping multiple validation
// errors returned by CurrencyDeleteRequest.ValidateAll() if the designated
// constraints aren't met.
type CurrencyDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CurrencyDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CurrencyDeleteRequestMultiError) AllErrors() []error { return m }

// CurrencyDeleteRequestValidationError is the validation error returned by
// CurrencyDeleteRequest.Validate if the designated constraints aren't met.
type CurrencyDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CurrencyDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CurrencyDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CurrencyDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CurrencyDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CurrencyDeleteRequestValidationError) ErrorName() string {
	return "CurrencyDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CurrencyDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCurrencyDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CurrencyDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CurrencyDeleteRequestValidationError{}

// Validate checks the field values on CurrencyReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CurrencyReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CurrencyReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CurrencyReplyMultiError, or
// nil if none found.
func (m *CurrencyReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CurrencyReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ticker

	if len(errors) > 0 {
		return CurrencyReplyMultiError(errors)
	}

	return nil
}

// CurrencyReplyMultiError is an error wrapping multiple validation errors
// returned by CurrencyReply.ValidateAll() if the designated constraints
// aren't met.
type CurrencyReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CurrencyReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CurrencyReplyMultiError) AllErrors() []error { return m }

// CurrencyReplyValidationError is the validation error returned by
// CurrencyReply.Validate if the designated constraints aren't met.
type CurrencyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CurrencyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CurrencyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CurrencyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CurrencyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CurrencyReplyValidationError) ErrorName() string { return "CurrencyReplyValidationError" }

// Error satisfies the builtin error interface
func (e CurrencyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCurrencyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CurrencyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CurrencyReplyValidationError{}

// Validate checks the field values on CurrencyReplies with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CurrencyReplies) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CurrencyReplies with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CurrencyRepliesMultiError, or nil if none found.
func (m *CurrencyReplies) ValidateAll() error {
	return m.validate(true)
}

func (m *CurrencyReplies) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	if len(errors) > 0 {
		return CurrencyRepliesMultiError(errors)
	}

	return nil
}

// CurrencyRepliesMultiError is an error wrapping multiple validation errors
// returned by CurrencyReplies.ValidateAll() if the designated constraints
// aren't met.
type CurrencyRepliesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CurrencyRepliesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CurrencyRepliesMultiError) AllErrors() []error { return m }

// CurrencyRepliesValidationError is the validation error returned by
// CurrencyReplies.Validate if the designated constraints aren't met.
type CurrencyRepliesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CurrencyRepliesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CurrencyRepliesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CurrencyRepliesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CurrencyRepliesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CurrencyRepliesValidationError) ErrorName() string { return "CurrencyRepliesValidationError" }

// Error satisfies the builtin error interface
func (e CurrencyRepliesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCurrencyReplies.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CurrencyRepliesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CurrencyRepliesValidationError{}

// Validate checks the field values on CurrencyDelete with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CurrencyDelete) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CurrencyDelete with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CurrencyDeleteMultiError,
// or nil if none found.
func (m *CurrencyDelete) ValidateAll() error {
	return m.validate(true)
}

func (m *CurrencyDelete) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	if len(errors) > 0 {
		return CurrencyDeleteMultiError(errors)
	}

	return nil
}

// CurrencyDeleteMultiError is an error wrapping multiple validation errors
// returned by CurrencyDelete.ValidateAll() if the designated constraints
// aren't met.
type CurrencyDeleteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CurrencyDeleteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CurrencyDeleteMultiError) AllErrors() []error { return m }

// CurrencyDeleteValidationError is the validation error returned by
// CurrencyDelete.Validate if the designated constraints aren't met.
type CurrencyDeleteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CurrencyDeleteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CurrencyDeleteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CurrencyDeleteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CurrencyDeleteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CurrencyDeleteValidationError) ErrorName() string { return "CurrencyDeleteValidationError" }

// Error satisfies the builtin error interface
func (e CurrencyDeleteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCurrencyDelete.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CurrencyDeleteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CurrencyDeleteValidationError{}

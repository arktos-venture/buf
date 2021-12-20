// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/strategy/v1/strategy.proto

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

// Validate checks the field values on StrategyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StrategyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrategyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StrategyRequestMultiError, or nil if none found.
func (m *StrategyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StrategyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTicker()); l < 1 || l > 8 {
		err := StrategyRequestValidationError{
			field:  "Ticker",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetExchange()); l < 2 || l > 8 {
		err := StrategyRequestValidationError{
			field:  "Exchange",
			reason: "value length must be between 2 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCurrency()) != 3 {
		err := StrategyRequestValidationError{
			field:  "Currency",
			reason: "value length must be 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if _, ok := StrategyRequest_Interval_name[int32(m.GetInterval())]; !ok {
		err := StrategyRequestValidationError{
			field:  "Interval",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _StrategyRequest_Period_InLookup[m.GetPeriod()]; !ok {
		err := StrategyRequestValidationError{
			field:  "Period",
			reason: "value must be in list [last 3d 1w 2w 1m 2m 3m 6m 1y 2y 3y 5y 10y 20y 30y]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := StrategyRequest_StrategyType_name[int32(m.GetStrategy())]; !ok {
		err := StrategyRequestValidationError{
			field:  "Strategy",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StrategyRequestMultiError(errors)
	}
	return nil
}

// StrategyRequestMultiError is an error wrapping multiple validation errors
// returned by StrategyRequest.ValidateAll() if the designated constraints
// aren't met.
type StrategyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyRequestMultiError) AllErrors() []error { return m }

// StrategyRequestValidationError is the validation error returned by
// StrategyRequest.Validate if the designated constraints aren't met.
type StrategyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyRequestValidationError) ErrorName() string { return "StrategyRequestValidationError" }

// Error satisfies the builtin error interface
func (e StrategyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyRequestValidationError{}

var _StrategyRequest_Period_InLookup = map[string]struct{}{
	"last": {},
	"3d":   {},
	"1w":   {},
	"2w":   {},
	"1m":   {},
	"2m":   {},
	"3m":   {},
	"6m":   {},
	"1y":   {},
	"2y":   {},
	"3y":   {},
	"5y":   {},
	"10y":  {},
	"20y":  {},
	"30y":  {},
}

// Validate checks the field values on StrategyReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StrategyReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrategyReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StrategyReplyMultiError, or
// nil if none found.
func (m *StrategyReply) ValidateAll() error {
	return m.validate(true)
}

func (m *StrategyReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	if len(errors) > 0 {
		return StrategyReplyMultiError(errors)
	}
	return nil
}

// StrategyReplyMultiError is an error wrapping multiple validation errors
// returned by StrategyReply.ValidateAll() if the designated constraints
// aren't met.
type StrategyReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyReplyMultiError) AllErrors() []error { return m }

// StrategyReplyValidationError is the validation error returned by
// StrategyReply.Validate if the designated constraints aren't met.
type StrategyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyReplyValidationError) ErrorName() string { return "StrategyReplyValidationError" }

// Error satisfies the builtin error interface
func (e StrategyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyReplyValidationError{}

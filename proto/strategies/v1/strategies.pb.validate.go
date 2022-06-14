// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/strategies/v1/strategies.proto

package strategies_v1

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

// Validate checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Request) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RequestMultiError, or nil if none found.
func (m *Request) ValidateAll() error {
	return m.validate(true)
}

func (m *Request) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Parameters

	// no validation rules for Period

	if len(errors) > 0 {
		return RequestMultiError(errors)
	}

	return nil
}

// RequestMultiError is an error wrapping multiple validation errors returned
// by Request.ValidateAll() if the designated constraints aren't met.
type RequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RequestMultiError) AllErrors() []error { return m }

// RequestValidationError is the validation error returned by Request.Validate
// if the designated constraints aren't met.
type RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestValidationError) ErrorName() string { return "RequestValidationError" }

// Error satisfies the builtin error interface
func (e RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestValidationError{}

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

	// no validation rules for Ticker

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

// Validate checks the field values on StrategyUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StrategyUpdateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrategyUpdateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StrategyUpdateRequestMultiError, or nil if none found.
func (m *StrategyUpdateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StrategyUpdateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ticker

	// no validation rules for Account

	// no validation rules for Name

	// no validation rules for Description

	if all {
		switch v := interface{}(m.GetParameters()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StrategyUpdateRequestValidationError{
					field:  "Parameters",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StrategyUpdateRequestValidationError{
					field:  "Parameters",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetParameters()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StrategyUpdateRequestValidationError{
				field:  "Parameters",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StrategyUpdateRequestMultiError(errors)
	}

	return nil
}

// StrategyUpdateRequestMultiError is an error wrapping multiple validation
// errors returned by StrategyUpdateRequest.ValidateAll() if the designated
// constraints aren't met.
type StrategyUpdateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyUpdateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyUpdateRequestMultiError) AllErrors() []error { return m }

// StrategyUpdateRequestValidationError is the validation error returned by
// StrategyUpdateRequest.Validate if the designated constraints aren't met.
type StrategyUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyUpdateRequestValidationError) ErrorName() string {
	return "StrategyUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StrategyUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategyUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyUpdateRequestValidationError{}

// Validate checks the field values on StrategyDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StrategyDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrategyDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StrategyDeleteRequestMultiError, or nil if none found.
func (m *StrategyDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StrategyDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetTickers()) < 1 {
		err := StrategyDeleteRequestValidationError{
			field:  "Tickers",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	_StrategyDeleteRequest_Tickers_Unique := make(map[string]struct{}, len(m.GetTickers()))

	for idx, item := range m.GetTickers() {
		_, _ = idx, item

		if _, exists := _StrategyDeleteRequest_Tickers_Unique[item]; exists {
			err := StrategyDeleteRequestValidationError{
				field:  fmt.Sprintf("Tickers[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_StrategyDeleteRequest_Tickers_Unique[item] = struct{}{}
		}

		// no validation rules for Tickers[idx]
	}

	if len(errors) > 0 {
		return StrategyDeleteRequestMultiError(errors)
	}

	return nil
}

// StrategyDeleteRequestMultiError is an error wrapping multiple validation
// errors returned by StrategyDeleteRequest.ValidateAll() if the designated
// constraints aren't met.
type StrategyDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyDeleteRequestMultiError) AllErrors() []error { return m }

// StrategyDeleteRequestValidationError is the validation error returned by
// StrategyDeleteRequest.Validate if the designated constraints aren't met.
type StrategyDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyDeleteRequestValidationError) ErrorName() string {
	return "StrategyDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StrategyDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategyDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyDeleteRequestValidationError{}

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

	// no validation rules for Ticker

	// no validation rules for Description

	// no validation rules for Account

	if all {
		switch v := interface{}(m.GetRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StrategyReplyValidationError{
					field:  "Request",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StrategyReplyValidationError{
					field:  "Request",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StrategyReplyValidationError{
				field:  "Request",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StrategyReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StrategyReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StrategyReplyValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StrategyReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StrategyReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StrategyReplyValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

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

// Validate checks the field values on StrategyReplies with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StrategyReplies) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrategyReplies with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StrategyRepliesMultiError, or nil if none found.
func (m *StrategyReplies) ValidateAll() error {
	return m.validate(true)
}

func (m *StrategyReplies) validate(all bool) error {
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
					errors = append(errors, StrategyRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, StrategyRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return StrategyRepliesValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return StrategyRepliesMultiError(errors)
	}

	return nil
}

// StrategyRepliesMultiError is an error wrapping multiple validation errors
// returned by StrategyReplies.ValidateAll() if the designated constraints
// aren't met.
type StrategyRepliesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyRepliesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyRepliesMultiError) AllErrors() []error { return m }

// StrategyRepliesValidationError is the validation error returned by
// StrategyReplies.Validate if the designated constraints aren't met.
type StrategyRepliesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyRepliesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyRepliesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyRepliesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyRepliesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyRepliesValidationError) ErrorName() string { return "StrategyRepliesValidationError" }

// Error satisfies the builtin error interface
func (e StrategyRepliesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategyReplies.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyRepliesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyRepliesValidationError{}

// Validate checks the field values on StrategyDelete with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StrategyDelete) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrategyDelete with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StrategyDeleteMultiError,
// or nil if none found.
func (m *StrategyDelete) ValidateAll() error {
	return m.validate(true)
}

func (m *StrategyDelete) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	if len(errors) > 0 {
		return StrategyDeleteMultiError(errors)
	}

	return nil
}

// StrategyDeleteMultiError is an error wrapping multiple validation errors
// returned by StrategyDelete.ValidateAll() if the designated constraints
// aren't met.
type StrategyDeleteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyDeleteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyDeleteMultiError) AllErrors() []error { return m }

// StrategyDeleteValidationError is the validation error returned by
// StrategyDelete.Validate if the designated constraints aren't met.
type StrategyDeleteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyDeleteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyDeleteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyDeleteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyDeleteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyDeleteValidationError) ErrorName() string { return "StrategyDeleteValidationError" }

// Error satisfies the builtin error interface
func (e StrategyDeleteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategyDelete.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyDeleteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyDeleteValidationError{}

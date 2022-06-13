// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/positions/v1/positions.proto

package positions_v1

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

// Validate checks the field values on PositionRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PositionRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PositionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PositionRequestMultiError, or nil if none found.
func (m *PositionRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PositionRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetAccount()); l < 3 || l > 36 {
		err := PositionRequestValidationError{
			field:  "Account",
			reason: "value length must be between 3 and 36 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PositionRequestMultiError(errors)
	}

	return nil
}

// PositionRequestMultiError is an error wrapping multiple validation errors
// returned by PositionRequest.ValidateAll() if the designated constraints
// aren't met.
type PositionRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PositionRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PositionRequestMultiError) AllErrors() []error { return m }

// PositionRequestValidationError is the validation error returned by
// PositionRequest.Validate if the designated constraints aren't met.
type PositionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PositionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PositionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PositionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PositionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PositionRequestValidationError) ErrorName() string { return "PositionRequestValidationError" }

// Error satisfies the builtin error interface
func (e PositionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPositionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PositionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PositionRequestValidationError{}

// Validate checks the field values on PositionDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PositionDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PositionDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PositionDeleteRequestMultiError, or nil if none found.
func (m *PositionDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PositionDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetAccount()); l < 3 || l > 36 {
		err := PositionDeleteRequestValidationError{
			field:  "Account",
			reason: "value length must be between 3 and 36 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PositionDeleteRequestMultiError(errors)
	}

	return nil
}

// PositionDeleteRequestMultiError is an error wrapping multiple validation
// errors returned by PositionDeleteRequest.ValidateAll() if the designated
// constraints aren't met.
type PositionDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PositionDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PositionDeleteRequestMultiError) AllErrors() []error { return m }

// PositionDeleteRequestValidationError is the validation error returned by
// PositionDeleteRequest.Validate if the designated constraints aren't met.
type PositionDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PositionDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PositionDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PositionDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PositionDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PositionDeleteRequestValidationError) ErrorName() string {
	return "PositionDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PositionDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPositionDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PositionDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PositionDeleteRequestValidationError{}

// Validate checks the field values on PositionReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PositionReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PositionReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PositionReplyMultiError, or
// nil if none found.
func (m *PositionReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PositionReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ticker

	// no validation rules for Exchange

	// no validation rules for Size

	if all {
		switch v := interface{}(m.GetCost()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PositionReplyValidationError{
					field:  "Cost",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PositionReplyValidationError{
					field:  "Cost",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PositionReplyValidationError{
				field:  "Cost",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPerformance()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PositionReplyValidationError{
					field:  "Performance",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PositionReplyValidationError{
					field:  "Performance",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPerformance()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PositionReplyValidationError{
				field:  "Performance",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PositionReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PositionReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PositionReplyValidationError{
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
				errors = append(errors, PositionReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PositionReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PositionReplyValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PositionReplyMultiError(errors)
	}

	return nil
}

// PositionReplyMultiError is an error wrapping multiple validation errors
// returned by PositionReply.ValidateAll() if the designated constraints
// aren't met.
type PositionReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PositionReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PositionReplyMultiError) AllErrors() []error { return m }

// PositionReplyValidationError is the validation error returned by
// PositionReply.Validate if the designated constraints aren't met.
type PositionReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PositionReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PositionReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PositionReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PositionReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PositionReplyValidationError) ErrorName() string { return "PositionReplyValidationError" }

// Error satisfies the builtin error interface
func (e PositionReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPositionReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PositionReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PositionReplyValidationError{}

// Validate checks the field values on PositionReplies with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PositionReplies) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PositionReplies with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PositionRepliesMultiError, or nil if none found.
func (m *PositionReplies) ValidateAll() error {
	return m.validate(true)
}

func (m *PositionReplies) validate(all bool) error {
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
					errors = append(errors, PositionRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PositionRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PositionRepliesValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return PositionRepliesMultiError(errors)
	}

	return nil
}

// PositionRepliesMultiError is an error wrapping multiple validation errors
// returned by PositionReplies.ValidateAll() if the designated constraints
// aren't met.
type PositionRepliesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PositionRepliesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PositionRepliesMultiError) AllErrors() []error { return m }

// PositionRepliesValidationError is the validation error returned by
// PositionReplies.Validate if the designated constraints aren't met.
type PositionRepliesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PositionRepliesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PositionRepliesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PositionRepliesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PositionRepliesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PositionRepliesValidationError) ErrorName() string { return "PositionRepliesValidationError" }

// Error satisfies the builtin error interface
func (e PositionRepliesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPositionReplies.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PositionRepliesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PositionRepliesValidationError{}

// Validate checks the field values on PositionDelete with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PositionDelete) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PositionDelete with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PositionDeleteMultiError,
// or nil if none found.
func (m *PositionDelete) ValidateAll() error {
	return m.validate(true)
}

func (m *PositionDelete) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	if len(errors) > 0 {
		return PositionDeleteMultiError(errors)
	}

	return nil
}

// PositionDeleteMultiError is an error wrapping multiple validation errors
// returned by PositionDelete.ValidateAll() if the designated constraints
// aren't met.
type PositionDeleteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PositionDeleteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PositionDeleteMultiError) AllErrors() []error { return m }

// PositionDeleteValidationError is the validation error returned by
// PositionDelete.Validate if the designated constraints aren't met.
type PositionDeleteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PositionDeleteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PositionDeleteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PositionDeleteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PositionDeleteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PositionDeleteValidationError) ErrorName() string { return "PositionDeleteValidationError" }

// Error satisfies the builtin error interface
func (e PositionDeleteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPositionDelete.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PositionDeleteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PositionDeleteValidationError{}

// Validate checks the field values on PositionReply_Performance with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PositionReply_Performance) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PositionReply_Performance with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PositionReply_PerformanceMultiError, or nil if none found.
func (m *PositionReply_Performance) ValidateAll() error {
	return m.validate(true)
}

func (m *PositionReply_Performance) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Percent

	// no validation rules for Money

	if len(errors) > 0 {
		return PositionReply_PerformanceMultiError(errors)
	}

	return nil
}

// PositionReply_PerformanceMultiError is an error wrapping multiple validation
// errors returned by PositionReply_Performance.ValidateAll() if the
// designated constraints aren't met.
type PositionReply_PerformanceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PositionReply_PerformanceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PositionReply_PerformanceMultiError) AllErrors() []error { return m }

// PositionReply_PerformanceValidationError is the validation error returned by
// PositionReply_Performance.Validate if the designated constraints aren't met.
type PositionReply_PerformanceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PositionReply_PerformanceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PositionReply_PerformanceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PositionReply_PerformanceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PositionReply_PerformanceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PositionReply_PerformanceValidationError) ErrorName() string {
	return "PositionReply_PerformanceValidationError"
}

// Error satisfies the builtin error interface
func (e PositionReply_PerformanceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPositionReply_Performance.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PositionReply_PerformanceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PositionReply_PerformanceValidationError{}

// Validate checks the field values on PositionReply_Cost with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PositionReply_Cost) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PositionReply_Cost with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PositionReply_CostMultiError, or nil if none found.
func (m *PositionReply_Cost) ValidateAll() error {
	return m.validate(true)
}

func (m *PositionReply_Cost) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Unit

	// no validation rules for Total

	if len(errors) > 0 {
		return PositionReply_CostMultiError(errors)
	}

	return nil
}

// PositionReply_CostMultiError is an error wrapping multiple validation errors
// returned by PositionReply_Cost.ValidateAll() if the designated constraints
// aren't met.
type PositionReply_CostMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PositionReply_CostMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PositionReply_CostMultiError) AllErrors() []error { return m }

// PositionReply_CostValidationError is the validation error returned by
// PositionReply_Cost.Validate if the designated constraints aren't met.
type PositionReply_CostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PositionReply_CostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PositionReply_CostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PositionReply_CostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PositionReply_CostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PositionReply_CostValidationError) ErrorName() string {
	return "PositionReply_CostValidationError"
}

// Error satisfies the builtin error interface
func (e PositionReply_CostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPositionReply_Cost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PositionReply_CostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PositionReply_CostValidationError{}

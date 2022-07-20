// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/forexes/v1/forexes.proto

package v1Forexes

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

// Validate checks the field values on ForexRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ForexRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ForexRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ForexRequestMultiError, or
// nil if none found.
func (m *ForexRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ForexRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTicker()) != 6 {
		err := ForexRequestValidationError{
			field:  "Ticker",
			reason: "value length must be 6 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return ForexRequestMultiError(errors)
	}

	return nil
}

// ForexRequestMultiError is an error wrapping multiple validation errors
// returned by ForexRequest.ValidateAll() if the designated constraints aren't met.
type ForexRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ForexRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ForexRequestMultiError) AllErrors() []error { return m }

// ForexRequestValidationError is the validation error returned by
// ForexRequest.Validate if the designated constraints aren't met.
type ForexRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForexRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForexRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForexRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForexRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForexRequestValidationError) ErrorName() string { return "ForexRequestValidationError" }

// Error satisfies the builtin error interface
func (e ForexRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForexRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForexRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForexRequestValidationError{}

// Validate checks the field values on ForexStrategiesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ForexStrategiesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ForexStrategiesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ForexStrategiesRequestMultiError, or nil if none found.
func (m *ForexStrategiesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ForexStrategiesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetAccount()); l < 3 || l > 36 {
		err := ForexStrategiesRequestValidationError{
			field:  "Account",
			reason: "value length must be between 3 and 36 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetTicker()) != 6 {
		err := ForexStrategiesRequestValidationError{
			field:  "Ticker",
			reason: "value length must be 6 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return ForexStrategiesRequestMultiError(errors)
	}

	return nil
}

// ForexStrategiesRequestMultiError is an error wrapping multiple validation
// errors returned by ForexStrategiesRequest.ValidateAll() if the designated
// constraints aren't met.
type ForexStrategiesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ForexStrategiesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ForexStrategiesRequestMultiError) AllErrors() []error { return m }

// ForexStrategiesRequestValidationError is the validation error returned by
// ForexStrategiesRequest.Validate if the designated constraints aren't met.
type ForexStrategiesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForexStrategiesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForexStrategiesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForexStrategiesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForexStrategiesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForexStrategiesRequestValidationError) ErrorName() string {
	return "ForexStrategiesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ForexStrategiesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForexStrategiesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForexStrategiesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForexStrategiesRequestValidationError{}

// Validate checks the field values on ForexListRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ForexListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ForexListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ForexListRequestMultiError, or nil if none found.
func (m *ForexListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ForexListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetCurrency()) != 3 {
		err := ForexListRequestValidationError{
			field:  "Currency",
			reason: "value length must be 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return ForexListRequestMultiError(errors)
	}

	return nil
}

// ForexListRequestMultiError is an error wrapping multiple validation errors
// returned by ForexListRequest.ValidateAll() if the designated constraints
// aren't met.
type ForexListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ForexListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ForexListRequestMultiError) AllErrors() []error { return m }

// ForexListRequestValidationError is the validation error returned by
// ForexListRequest.Validate if the designated constraints aren't met.
type ForexListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForexListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForexListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForexListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForexListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForexListRequestValidationError) ErrorName() string { return "ForexListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ForexListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForexListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForexListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForexListRequestValidationError{}

// Validate checks the field values on ForexReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ForexReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ForexReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ForexReplyMultiError, or
// nil if none found.
func (m *ForexReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ForexReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ticker

	if all {
		switch v := interface{}(m.GetQuote()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ForexReplyValidationError{
					field:  "Quote",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ForexReplyValidationError{
					field:  "Quote",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuote()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ForexReplyValidationError{
				field:  "Quote",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetStats()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ForexReplyValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ForexReplyValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStats()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ForexReplyValidationError{
				field:  "Stats",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFrom()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ForexReplyValidationError{
					field:  "From",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ForexReplyValidationError{
					field:  "From",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFrom()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ForexReplyValidationError{
				field:  "From",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetTo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ForexReplyValidationError{
					field:  "To",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ForexReplyValidationError{
					field:  "To",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ForexReplyValidationError{
				field:  "To",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ForexReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ForexReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ForexReplyValidationError{
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
				errors = append(errors, ForexReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ForexReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ForexReplyValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ForexReplyMultiError(errors)
	}

	return nil
}

// ForexReplyMultiError is an error wrapping multiple validation errors
// returned by ForexReply.ValidateAll() if the designated constraints aren't met.
type ForexReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ForexReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ForexReplyMultiError) AllErrors() []error { return m }

// ForexReplyValidationError is the validation error returned by
// ForexReply.Validate if the designated constraints aren't met.
type ForexReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForexReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForexReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForexReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForexReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForexReplyValidationError) ErrorName() string { return "ForexReplyValidationError" }

// Error satisfies the builtin error interface
func (e ForexReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForexReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForexReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForexReplyValidationError{}

// Validate checks the field values on ForexStatsReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ForexStatsReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ForexStatsReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ForexStatsReplyMultiError, or nil if none found.
func (m *ForexStatsReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ForexStatsReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPrice()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ForexStatsReplyValidationError{
					field:  "Price",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ForexStatsReplyValidationError{
					field:  "Price",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPrice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ForexStatsReplyValidationError{
				field:  "Price",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetVolume()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ForexStatsReplyValidationError{
					field:  "Volume",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ForexStatsReplyValidationError{
					field:  "Volume",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetVolume()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ForexStatsReplyValidationError{
				field:  "Volume",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ForexStatsReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ForexStatsReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ForexStatsReplyValidationError{
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
				errors = append(errors, ForexStatsReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ForexStatsReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ForexStatsReplyValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ForexStatsReplyMultiError(errors)
	}

	return nil
}

// ForexStatsReplyMultiError is an error wrapping multiple validation errors
// returned by ForexStatsReply.ValidateAll() if the designated constraints
// aren't met.
type ForexStatsReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ForexStatsReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ForexStatsReplyMultiError) AllErrors() []error { return m }

// ForexStatsReplyValidationError is the validation error returned by
// ForexStatsReply.Validate if the designated constraints aren't met.
type ForexStatsReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForexStatsReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForexStatsReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForexStatsReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForexStatsReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForexStatsReplyValidationError) ErrorName() string { return "ForexStatsReplyValidationError" }

// Error satisfies the builtin error interface
func (e ForexStatsReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForexStatsReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForexStatsReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForexStatsReplyValidationError{}

// Validate checks the field values on ForexReplies with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ForexReplies) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ForexReplies with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ForexRepliesMultiError, or
// nil if none found.
func (m *ForexReplies) ValidateAll() error {
	return m.validate(true)
}

func (m *ForexReplies) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	if len(errors) > 0 {
		return ForexRepliesMultiError(errors)
	}

	return nil
}

// ForexRepliesMultiError is an error wrapping multiple validation errors
// returned by ForexReplies.ValidateAll() if the designated constraints aren't met.
type ForexRepliesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ForexRepliesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ForexRepliesMultiError) AllErrors() []error { return m }

// ForexRepliesValidationError is the validation error returned by
// ForexReplies.Validate if the designated constraints aren't met.
type ForexRepliesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForexRepliesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForexRepliesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForexRepliesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForexRepliesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForexRepliesValidationError) ErrorName() string { return "ForexRepliesValidationError" }

// Error satisfies the builtin error interface
func (e ForexRepliesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForexReplies.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForexRepliesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForexRepliesValidationError{}

// Validate checks the field values on ForexStatsReply_Price with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ForexStatsReply_Price) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ForexStatsReply_Price with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ForexStatsReply_PriceMultiError, or nil if none found.
func (m *ForexStatsReply_Price) ValidateAll() error {
	return m.validate(true)
}

func (m *ForexStatsReply_Price) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MaxAnnual

	// no validation rules for MinAnnual

	// no validation rules for ReturnYear

	// no validation rules for Beta5Y

	// no validation rules for Start

	if len(errors) > 0 {
		return ForexStatsReply_PriceMultiError(errors)
	}

	return nil
}

// ForexStatsReply_PriceMultiError is an error wrapping multiple validation
// errors returned by ForexStatsReply_Price.ValidateAll() if the designated
// constraints aren't met.
type ForexStatsReply_PriceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ForexStatsReply_PriceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ForexStatsReply_PriceMultiError) AllErrors() []error { return m }

// ForexStatsReply_PriceValidationError is the validation error returned by
// ForexStatsReply_Price.Validate if the designated constraints aren't met.
type ForexStatsReply_PriceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForexStatsReply_PriceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForexStatsReply_PriceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForexStatsReply_PriceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForexStatsReply_PriceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForexStatsReply_PriceValidationError) ErrorName() string {
	return "ForexStatsReply_PriceValidationError"
}

// Error satisfies the builtin error interface
func (e ForexStatsReply_PriceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForexStatsReply_Price.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForexStatsReply_PriceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForexStatsReply_PriceValidationError{}

// Validate checks the field values on ForexStatsReply_Volume with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ForexStatsReply_Volume) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ForexStatsReply_Volume with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ForexStatsReply_VolumeMultiError, or nil if none found.
func (m *ForexStatsReply_Volume) ValidateAll() error {
	return m.validate(true)
}

func (m *ForexStatsReply_Volume) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Avg_10D

	// no validation rules for Avg_30D

	// no validation rules for Avg_90D

	if len(errors) > 0 {
		return ForexStatsReply_VolumeMultiError(errors)
	}

	return nil
}

// ForexStatsReply_VolumeMultiError is an error wrapping multiple validation
// errors returned by ForexStatsReply_Volume.ValidateAll() if the designated
// constraints aren't met.
type ForexStatsReply_VolumeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ForexStatsReply_VolumeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ForexStatsReply_VolumeMultiError) AllErrors() []error { return m }

// ForexStatsReply_VolumeValidationError is the validation error returned by
// ForexStatsReply_Volume.Validate if the designated constraints aren't met.
type ForexStatsReply_VolumeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ForexStatsReply_VolumeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ForexStatsReply_VolumeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ForexStatsReply_VolumeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ForexStatsReply_VolumeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ForexStatsReply_VolumeValidationError) ErrorName() string {
	return "ForexStatsReply_VolumeValidationError"
}

// Error satisfies the builtin error interface
func (e ForexStatsReply_VolumeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sForexStatsReply_Volume.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ForexStatsReply_VolumeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ForexStatsReply_VolumeValidationError{}

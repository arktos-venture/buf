// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/indices/v1/indices.proto

package v1Indices

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

	_ = v1Screener.Asset(0)
)

// Validate checks the field values on IndiceRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IndiceRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndiceRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IndiceRequestMultiError, or
// nil if none found.
func (m *IndiceRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IndiceRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTicker()); l < 2 || l > 16 {
		err := IndiceRequestValidationError{
			field:  "Ticker",
			reason: "value length must be between 2 and 16 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IndiceRequestMultiError(errors)
	}

	return nil
}

// IndiceRequestMultiError is an error wrapping multiple validation errors
// returned by IndiceRequest.ValidateAll() if the designated constraints
// aren't met.
type IndiceRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndiceRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndiceRequestMultiError) AllErrors() []error { return m }

// IndiceRequestValidationError is the validation error returned by
// IndiceRequest.Validate if the designated constraints aren't met.
type IndiceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndiceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndiceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndiceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndiceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndiceRequestValidationError) ErrorName() string { return "IndiceRequestValidationError" }

// Error satisfies the builtin error interface
func (e IndiceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndiceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndiceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndiceRequestValidationError{}

// Validate checks the field values on IndiceStrategiesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IndiceStrategiesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndiceStrategiesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndiceStrategiesRequestMultiError, or nil if none found.
func (m *IndiceStrategiesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IndiceStrategiesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetAccount()); l < 3 || l > 36 {
		err := IndiceStrategiesRequestValidationError{
			field:  "Account",
			reason: "value length must be between 3 and 36 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetTicker()); l < 1 || l > 16 {
		err := IndiceStrategiesRequestValidationError{
			field:  "Ticker",
			reason: "value length must be between 1 and 16 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IndiceStrategiesRequestMultiError(errors)
	}

	return nil
}

// IndiceStrategiesRequestMultiError is an error wrapping multiple validation
// errors returned by IndiceStrategiesRequest.ValidateAll() if the designated
// constraints aren't met.
type IndiceStrategiesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndiceStrategiesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndiceStrategiesRequestMultiError) AllErrors() []error { return m }

// IndiceStrategiesRequestValidationError is the validation error returned by
// IndiceStrategiesRequest.Validate if the designated constraints aren't met.
type IndiceStrategiesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndiceStrategiesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndiceStrategiesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndiceStrategiesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndiceStrategiesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndiceStrategiesRequestValidationError) ErrorName() string {
	return "IndiceStrategiesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e IndiceStrategiesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndiceStrategiesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndiceStrategiesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndiceStrategiesRequestValidationError{}

// Validate checks the field values on IndiceSearchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IndiceSearchRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndiceSearchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndiceSearchRequestMultiError, or nil if none found.
func (m *IndiceSearchRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IndiceSearchRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetExchanges()) < 1 {
		err := IndiceSearchRequestValidationError{
			field:  "Exchanges",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	_IndiceSearchRequest_Exchanges_Unique := make(map[string]struct{}, len(m.GetExchanges()))

	for idx, item := range m.GetExchanges() {
		_, _ = idx, item

		if _, exists := _IndiceSearchRequest_Exchanges_Unique[item]; exists {
			err := IndiceSearchRequestValidationError{
				field:  fmt.Sprintf("Exchanges[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_IndiceSearchRequest_Exchanges_Unique[item] = struct{}{}
		}

		// no validation rules for Exchanges[idx]
	}

	if len(errors) > 0 {
		return IndiceSearchRequestMultiError(errors)
	}

	return nil
}

// IndiceSearchRequestMultiError is an error wrapping multiple validation
// errors returned by IndiceSearchRequest.ValidateAll() if the designated
// constraints aren't met.
type IndiceSearchRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndiceSearchRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndiceSearchRequestMultiError) AllErrors() []error { return m }

// IndiceSearchRequestValidationError is the validation error returned by
// IndiceSearchRequest.Validate if the designated constraints aren't met.
type IndiceSearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndiceSearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndiceSearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndiceSearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndiceSearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndiceSearchRequestValidationError) ErrorName() string {
	return "IndiceSearchRequestValidationError"
}

// Error satisfies the builtin error interface
func (e IndiceSearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndiceSearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndiceSearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndiceSearchRequestValidationError{}

// Validate checks the field values on IndiceReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IndiceReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndiceReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IndiceReplyMultiError, or
// nil if none found.
func (m *IndiceReply) ValidateAll() error {
	return m.validate(true)
}

func (m *IndiceReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ticker

	// no validation rules for Description

	// no validation rules for Exchange

	// no validation rules for Asset

	if all {
		switch v := interface{}(m.GetQuote()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IndiceReplyValidationError{
					field:  "Quote",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IndiceReplyValidationError{
					field:  "Quote",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuote()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IndiceReplyValidationError{
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
				errors = append(errors, IndiceReplyValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IndiceReplyValidationError{
					field:  "Stats",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetStats()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IndiceReplyValidationError{
				field:  "Stats",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFilters()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IndiceReplyValidationError{
					field:  "Filters",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IndiceReplyValidationError{
					field:  "Filters",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilters()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IndiceReplyValidationError{
				field:  "Filters",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IndiceReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IndiceReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IndiceReplyValidationError{
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
				errors = append(errors, IndiceReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IndiceReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IndiceReplyValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return IndiceReplyMultiError(errors)
	}

	return nil
}

// IndiceReplyMultiError is an error wrapping multiple validation errors
// returned by IndiceReply.ValidateAll() if the designated constraints aren't met.
type IndiceReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndiceReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndiceReplyMultiError) AllErrors() []error { return m }

// IndiceReplyValidationError is the validation error returned by
// IndiceReply.Validate if the designated constraints aren't met.
type IndiceReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndiceReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndiceReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndiceReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndiceReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndiceReplyValidationError) ErrorName() string { return "IndiceReplyValidationError" }

// Error satisfies the builtin error interface
func (e IndiceReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndiceReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndiceReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndiceReplyValidationError{}

// Validate checks the field values on IndiceStatsReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IndiceStatsReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndiceStatsReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndiceStatsReplyMultiError, or nil if none found.
func (m *IndiceStatsReply) ValidateAll() error {
	return m.validate(true)
}

func (m *IndiceStatsReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPrice()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IndiceStatsReplyValidationError{
					field:  "Price",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IndiceStatsReplyValidationError{
					field:  "Price",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPrice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IndiceStatsReplyValidationError{
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
				errors = append(errors, IndiceStatsReplyValidationError{
					field:  "Volume",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IndiceStatsReplyValidationError{
					field:  "Volume",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetVolume()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IndiceStatsReplyValidationError{
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
				errors = append(errors, IndiceStatsReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IndiceStatsReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IndiceStatsReplyValidationError{
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
				errors = append(errors, IndiceStatsReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IndiceStatsReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IndiceStatsReplyValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return IndiceStatsReplyMultiError(errors)
	}

	return nil
}

// IndiceStatsReplyMultiError is an error wrapping multiple validation errors
// returned by IndiceStatsReply.ValidateAll() if the designated constraints
// aren't met.
type IndiceStatsReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndiceStatsReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndiceStatsReplyMultiError) AllErrors() []error { return m }

// IndiceStatsReplyValidationError is the validation error returned by
// IndiceStatsReply.Validate if the designated constraints aren't met.
type IndiceStatsReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndiceStatsReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndiceStatsReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndiceStatsReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndiceStatsReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndiceStatsReplyValidationError) ErrorName() string { return "IndiceStatsReplyValidationError" }

// Error satisfies the builtin error interface
func (e IndiceStatsReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndiceStatsReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndiceStatsReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndiceStatsReplyValidationError{}

// Validate checks the field values on IndiceReplies with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IndiceReplies) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndiceReplies with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IndiceRepliesMultiError, or
// nil if none found.
func (m *IndiceReplies) ValidateAll() error {
	return m.validate(true)
}

func (m *IndiceReplies) validate(all bool) error {
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
					errors = append(errors, IndiceRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IndiceRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IndiceRepliesValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return IndiceRepliesMultiError(errors)
	}

	return nil
}

// IndiceRepliesMultiError is an error wrapping multiple validation errors
// returned by IndiceReplies.ValidateAll() if the designated constraints
// aren't met.
type IndiceRepliesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndiceRepliesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndiceRepliesMultiError) AllErrors() []error { return m }

// IndiceRepliesValidationError is the validation error returned by
// IndiceReplies.Validate if the designated constraints aren't met.
type IndiceRepliesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndiceRepliesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndiceRepliesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndiceRepliesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndiceRepliesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndiceRepliesValidationError) ErrorName() string { return "IndiceRepliesValidationError" }

// Error satisfies the builtin error interface
func (e IndiceRepliesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndiceReplies.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndiceRepliesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndiceRepliesValidationError{}

// Validate checks the field values on IndiceStatsReply_Price with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IndiceStatsReply_Price) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndiceStatsReply_Price with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndiceStatsReply_PriceMultiError, or nil if none found.
func (m *IndiceStatsReply_Price) ValidateAll() error {
	return m.validate(true)
}

func (m *IndiceStatsReply_Price) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for MaxAnnual

	// no validation rules for MinAnnual

	// no validation rules for ReturnYear

	// no validation rules for MarketCapUsd

	// no validation rules for Beta5Y

	// no validation rules for Start

	if len(errors) > 0 {
		return IndiceStatsReply_PriceMultiError(errors)
	}

	return nil
}

// IndiceStatsReply_PriceMultiError is an error wrapping multiple validation
// errors returned by IndiceStatsReply_Price.ValidateAll() if the designated
// constraints aren't met.
type IndiceStatsReply_PriceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndiceStatsReply_PriceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndiceStatsReply_PriceMultiError) AllErrors() []error { return m }

// IndiceStatsReply_PriceValidationError is the validation error returned by
// IndiceStatsReply_Price.Validate if the designated constraints aren't met.
type IndiceStatsReply_PriceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndiceStatsReply_PriceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndiceStatsReply_PriceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndiceStatsReply_PriceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndiceStatsReply_PriceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndiceStatsReply_PriceValidationError) ErrorName() string {
	return "IndiceStatsReply_PriceValidationError"
}

// Error satisfies the builtin error interface
func (e IndiceStatsReply_PriceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndiceStatsReply_Price.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndiceStatsReply_PriceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndiceStatsReply_PriceValidationError{}

// Validate checks the field values on IndiceStatsReply_Volume with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IndiceStatsReply_Volume) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndiceStatsReply_Volume with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndiceStatsReply_VolumeMultiError, or nil if none found.
func (m *IndiceStatsReply_Volume) ValidateAll() error {
	return m.validate(true)
}

func (m *IndiceStatsReply_Volume) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Avg_10D

	// no validation rules for Avg_30D

	// no validation rules for Avg_90D

	if len(errors) > 0 {
		return IndiceStatsReply_VolumeMultiError(errors)
	}

	return nil
}

// IndiceStatsReply_VolumeMultiError is an error wrapping multiple validation
// errors returned by IndiceStatsReply_Volume.ValidateAll() if the designated
// constraints aren't met.
type IndiceStatsReply_VolumeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndiceStatsReply_VolumeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndiceStatsReply_VolumeMultiError) AllErrors() []error { return m }

// IndiceStatsReply_VolumeValidationError is the validation error returned by
// IndiceStatsReply_Volume.Validate if the designated constraints aren't met.
type IndiceStatsReply_VolumeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndiceStatsReply_VolumeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndiceStatsReply_VolumeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndiceStatsReply_VolumeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndiceStatsReply_VolumeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndiceStatsReply_VolumeValidationError) ErrorName() string {
	return "IndiceStatsReply_VolumeValidationError"
}

// Error satisfies the builtin error interface
func (e IndiceStatsReply_VolumeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndiceStatsReply_Volume.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndiceStatsReply_VolumeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndiceStatsReply_VolumeValidationError{}

// Validate checks the field values on IndiceReplies_Result with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IndiceReplies_Result) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndiceReplies_Result with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndiceReplies_ResultMultiError, or nil if none found.
func (m *IndiceReplies_Result) ValidateAll() error {
	return m.validate(true)
}

func (m *IndiceReplies_Result) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ticker

	// no validation rules for Description

	// no validation rules for Asset

	if len(errors) > 0 {
		return IndiceReplies_ResultMultiError(errors)
	}

	return nil
}

// IndiceReplies_ResultMultiError is an error wrapping multiple validation
// errors returned by IndiceReplies_Result.ValidateAll() if the designated
// constraints aren't met.
type IndiceReplies_ResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndiceReplies_ResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndiceReplies_ResultMultiError) AllErrors() []error { return m }

// IndiceReplies_ResultValidationError is the validation error returned by
// IndiceReplies_Result.Validate if the designated constraints aren't met.
type IndiceReplies_ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndiceReplies_ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndiceReplies_ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndiceReplies_ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndiceReplies_ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndiceReplies_ResultValidationError) ErrorName() string {
	return "IndiceReplies_ResultValidationError"
}

// Error satisfies the builtin error interface
func (e IndiceReplies_ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndiceReplies_Result.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndiceReplies_ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndiceReplies_ResultValidationError{}

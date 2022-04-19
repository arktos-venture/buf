// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/splits/v1/splits.proto

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

// Validate checks the field values on DatePeriod with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DatePeriod) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DatePeriod with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DatePeriodMultiError, or
// nil if none found.
func (m *DatePeriod) ValidateAll() error {
	return m.validate(true)
}

func (m *DatePeriod) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := Interval_name[int32(m.GetInterval())]; !ok {
		err := DatePeriodValidationError{
			field:  "Interval",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _DatePeriod_Period_InLookup[m.GetPeriod()]; !ok {
		err := DatePeriodValidationError{
			field:  "Period",
			reason: "value must be in list [last 3d 1w 2w 1m 2m 3m 6m 1y 2y 3y 5y 10y 20y 30y]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DatePeriodMultiError(errors)
	}

	return nil
}

// DatePeriodMultiError is an error wrapping multiple validation errors
// returned by DatePeriod.ValidateAll() if the designated constraints aren't met.
type DatePeriodMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DatePeriodMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DatePeriodMultiError) AllErrors() []error { return m }

// DatePeriodValidationError is the validation error returned by
// DatePeriod.Validate if the designated constraints aren't met.
type DatePeriodValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DatePeriodValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DatePeriodValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DatePeriodValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DatePeriodValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DatePeriodValidationError) ErrorName() string { return "DatePeriodValidationError" }

// Error satisfies the builtin error interface
func (e DatePeriodValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDatePeriod.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DatePeriodValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DatePeriodValidationError{}

var _DatePeriod_Period_InLookup = map[string]struct{}{
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

// Validate checks the field values on DateTimestamp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DateTimestamp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DateTimestamp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DateTimestampMultiError, or
// nil if none found.
func (m *DateTimestamp) ValidateAll() error {
	return m.validate(true)
}

func (m *DateTimestamp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := Interval_name[int32(m.GetInterval())]; !ok {
		err := DateTimestampValidationError{
			field:  "Interval",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DateTimestampValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DateTimestampValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DateTimestampValidationError{
				field:  "Date",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DateTimestampMultiError(errors)
	}

	return nil
}

// DateTimestampMultiError is an error wrapping multiple validation errors
// returned by DateTimestamp.ValidateAll() if the designated constraints
// aren't met.
type DateTimestampMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DateTimestampMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DateTimestampMultiError) AllErrors() []error { return m }

// DateTimestampValidationError is the validation error returned by
// DateTimestamp.Validate if the designated constraints aren't met.
type DateTimestampValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DateTimestampValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DateTimestampValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DateTimestampValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DateTimestampValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DateTimestampValidationError) ErrorName() string { return "DateTimestampValidationError" }

// Error satisfies the builtin error interface
func (e DateTimestampValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDateTimestamp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DateTimestampValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DateTimestampValidationError{}

// Validate checks the field values on SplitPeriodRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SplitPeriodRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SplitPeriodRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SplitPeriodRequestMultiError, or nil if none found.
func (m *SplitPeriodRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SplitPeriodRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTicker()); l < 1 || l > 8 {
		err := SplitPeriodRequestValidationError{
			field:  "Ticker",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetExchange()); l < 1 || l > 8 {
		err := SplitPeriodRequestValidationError{
			field:  "Exchange",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetDate() == nil {
		err := SplitPeriodRequestValidationError{
			field:  "Date",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SplitPeriodRequestValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SplitPeriodRequestValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SplitPeriodRequestValidationError{
				field:  "Date",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SplitPeriodRequestMultiError(errors)
	}

	return nil
}

// SplitPeriodRequestMultiError is an error wrapping multiple validation errors
// returned by SplitPeriodRequest.ValidateAll() if the designated constraints
// aren't met.
type SplitPeriodRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SplitPeriodRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SplitPeriodRequestMultiError) AllErrors() []error { return m }

// SplitPeriodRequestValidationError is the validation error returned by
// SplitPeriodRequest.Validate if the designated constraints aren't met.
type SplitPeriodRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SplitPeriodRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SplitPeriodRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SplitPeriodRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SplitPeriodRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SplitPeriodRequestValidationError) ErrorName() string {
	return "SplitPeriodRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SplitPeriodRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSplitPeriodRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SplitPeriodRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SplitPeriodRequestValidationError{}

// Validate checks the field values on SplitTimestampRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SplitTimestampRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SplitTimestampRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SplitTimestampRequestMultiError, or nil if none found.
func (m *SplitTimestampRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SplitTimestampRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTicker()); l < 1 || l > 8 {
		err := SplitTimestampRequestValidationError{
			field:  "Ticker",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetExchange()); l < 1 || l > 8 {
		err := SplitTimestampRequestValidationError{
			field:  "Exchange",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetDate() == nil {
		err := SplitTimestampRequestValidationError{
			field:  "Date",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SplitTimestampRequestValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SplitTimestampRequestValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SplitTimestampRequestValidationError{
				field:  "Date",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SplitTimestampRequestMultiError(errors)
	}

	return nil
}

// SplitTimestampRequestMultiError is an error wrapping multiple validation
// errors returned by SplitTimestampRequest.ValidateAll() if the designated
// constraints aren't met.
type SplitTimestampRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SplitTimestampRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SplitTimestampRequestMultiError) AllErrors() []error { return m }

// SplitTimestampRequestValidationError is the validation error returned by
// SplitTimestampRequest.Validate if the designated constraints aren't met.
type SplitTimestampRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SplitTimestampRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SplitTimestampRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SplitTimestampRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SplitTimestampRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SplitTimestampRequestValidationError) ErrorName() string {
	return "SplitTimestampRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SplitTimestampRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSplitTimestampRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SplitTimestampRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SplitTimestampRequestValidationError{}

// Validate checks the field values on SplitBulkRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SplitBulkRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SplitBulkRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SplitBulkRequestMultiError, or nil if none found.
func (m *SplitBulkRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SplitBulkRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetExchange()); l < 1 || l > 8 {
		err := SplitBulkRequestValidationError{
			field:  "Exchange",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetDate() == nil {
		err := SplitBulkRequestValidationError{
			field:  "Date",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SplitBulkRequestValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SplitBulkRequestValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SplitBulkRequestValidationError{
				field:  "Date",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SplitBulkRequestMultiError(errors)
	}

	return nil
}

// SplitBulkRequestMultiError is an error wrapping multiple validation errors
// returned by SplitBulkRequest.ValidateAll() if the designated constraints
// aren't met.
type SplitBulkRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SplitBulkRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SplitBulkRequestMultiError) AllErrors() []error { return m }

// SplitBulkRequestValidationError is the validation error returned by
// SplitBulkRequest.Validate if the designated constraints aren't met.
type SplitBulkRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SplitBulkRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SplitBulkRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SplitBulkRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SplitBulkRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SplitBulkRequestValidationError) ErrorName() string { return "SplitBulkRequestValidationError" }

// Error satisfies the builtin error interface
func (e SplitBulkRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSplitBulkRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SplitBulkRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SplitBulkRequestValidationError{}

// Validate checks the field values on SplitsReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SplitsReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SplitsReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SplitsReplyMultiError, or
// nil if none found.
func (m *SplitsReply) ValidateAll() error {
	return m.validate(true)
}

func (m *SplitsReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for From

	// no validation rules for To

	if all {
		switch v := interface{}(m.GetDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SplitsReplyValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SplitsReplyValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SplitsReplyValidationError{
				field:  "Date",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SplitsReplyMultiError(errors)
	}

	return nil
}

// SplitsReplyMultiError is an error wrapping multiple validation errors
// returned by SplitsReply.ValidateAll() if the designated constraints aren't met.
type SplitsReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SplitsReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SplitsReplyMultiError) AllErrors() []error { return m }

// SplitsReplyValidationError is the validation error returned by
// SplitsReply.Validate if the designated constraints aren't met.
type SplitsReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SplitsReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SplitsReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SplitsReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SplitsReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SplitsReplyValidationError) ErrorName() string { return "SplitsReplyValidationError" }

// Error satisfies the builtin error interface
func (e SplitsReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSplitsReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SplitsReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SplitsReplyValidationError{}

// Validate checks the field values on SplitReplies with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SplitReplies) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SplitReplies with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SplitRepliesMultiError, or
// nil if none found.
func (m *SplitReplies) ValidateAll() error {
	return m.validate(true)
}

func (m *SplitReplies) validate(all bool) error {
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
					errors = append(errors, SplitRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SplitRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SplitRepliesValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return SplitRepliesMultiError(errors)
	}

	return nil
}

// SplitRepliesMultiError is an error wrapping multiple validation errors
// returned by SplitReplies.ValidateAll() if the designated constraints aren't met.
type SplitRepliesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SplitRepliesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SplitRepliesMultiError) AllErrors() []error { return m }

// SplitRepliesValidationError is the validation error returned by
// SplitReplies.Validate if the designated constraints aren't met.
type SplitRepliesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SplitRepliesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SplitRepliesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SplitRepliesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SplitRepliesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SplitRepliesValidationError) ErrorName() string { return "SplitRepliesValidationError" }

// Error satisfies the builtin error interface
func (e SplitRepliesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSplitReplies.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SplitRepliesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SplitRepliesValidationError{}

// Validate checks the field values on SplitBulkReplies with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SplitBulkReplies) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SplitBulkReplies with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SplitBulkRepliesMultiError, or nil if none found.
func (m *SplitBulkReplies) ValidateAll() error {
	return m.validate(true)
}

func (m *SplitBulkReplies) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCompanies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SplitBulkRepliesValidationError{
						field:  fmt.Sprintf("Companies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SplitBulkRepliesValidationError{
						field:  fmt.Sprintf("Companies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SplitBulkRepliesValidationError{
					field:  fmt.Sprintf("Companies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return SplitBulkRepliesMultiError(errors)
	}

	return nil
}

// SplitBulkRepliesMultiError is an error wrapping multiple validation errors
// returned by SplitBulkReplies.ValidateAll() if the designated constraints
// aren't met.
type SplitBulkRepliesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SplitBulkRepliesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SplitBulkRepliesMultiError) AllErrors() []error { return m }

// SplitBulkRepliesValidationError is the validation error returned by
// SplitBulkReplies.Validate if the designated constraints aren't met.
type SplitBulkRepliesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SplitBulkRepliesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SplitBulkRepliesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SplitBulkRepliesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SplitBulkRepliesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SplitBulkRepliesValidationError) ErrorName() string { return "SplitBulkRepliesValidationError" }

// Error satisfies the builtin error interface
func (e SplitBulkRepliesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSplitBulkReplies.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SplitBulkRepliesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SplitBulkRepliesValidationError{}

// Validate checks the field values on SplitBulkReplies_Company with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SplitBulkReplies_Company) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SplitBulkReplies_Company with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SplitBulkReplies_CompanyMultiError, or nil if none found.
func (m *SplitBulkReplies_Company) ValidateAll() error {
	return m.validate(true)
}

func (m *SplitBulkReplies_Company) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ticker

	// no validation rules for Exchange

	if all {
		switch v := interface{}(m.GetResults()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SplitBulkReplies_CompanyValidationError{
					field:  "Results",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SplitBulkReplies_CompanyValidationError{
					field:  "Results",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResults()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SplitBulkReplies_CompanyValidationError{
				field:  "Results",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SplitBulkReplies_CompanyMultiError(errors)
	}

	return nil
}

// SplitBulkReplies_CompanyMultiError is an error wrapping multiple validation
// errors returned by SplitBulkReplies_Company.ValidateAll() if the designated
// constraints aren't met.
type SplitBulkReplies_CompanyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SplitBulkReplies_CompanyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SplitBulkReplies_CompanyMultiError) AllErrors() []error { return m }

// SplitBulkReplies_CompanyValidationError is the validation error returned by
// SplitBulkReplies_Company.Validate if the designated constraints aren't met.
type SplitBulkReplies_CompanyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SplitBulkReplies_CompanyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SplitBulkReplies_CompanyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SplitBulkReplies_CompanyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SplitBulkReplies_CompanyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SplitBulkReplies_CompanyValidationError) ErrorName() string {
	return "SplitBulkReplies_CompanyValidationError"
}

// Error satisfies the builtin error interface
func (e SplitBulkReplies_CompanyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSplitBulkReplies_Company.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SplitBulkReplies_CompanyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SplitBulkReplies_CompanyValidationError{}

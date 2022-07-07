// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/industries/v1/industries.proto

package v1Industries

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

// Validate checks the field values on Ref with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Ref) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Ref with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RefMultiError, or nil if none found.
func (m *Ref) ValidateAll() error {
	return m.validate(true)
}

func (m *Ref) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Classication

	// no validation rules for Name

	{
		sorted_keys := make([]int64, len(m.GetNext()))
		i := 0
		for key := range m.GetNext() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetNext()[key]
			_ = val

			// no validation rules for Next[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, RefValidationError{
							field:  fmt.Sprintf("Next[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, RefValidationError{
							field:  fmt.Sprintf("Next[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return RefValidationError{
						field:  fmt.Sprintf("Next[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return RefMultiError(errors)
	}

	return nil
}

// RefMultiError is an error wrapping multiple validation errors returned by
// Ref.ValidateAll() if the designated constraints aren't met.
type RefMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RefMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RefMultiError) AllErrors() []error { return m }

// RefValidationError is the validation error returned by Ref.Validate if the
// designated constraints aren't met.
type RefValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RefValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RefValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RefValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RefValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RefValidationError) ErrorName() string { return "RefValidationError" }

// Error satisfies the builtin error interface
func (e RefValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRef.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RefValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RefValidationError{}

// Validate checks the field values on IndustrySearchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IndustrySearchRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndustrySearchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndustrySearchRequestMultiError, or nil if none found.
func (m *IndustrySearchRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IndustrySearchRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ticker

	if len(errors) > 0 {
		return IndustrySearchRequestMultiError(errors)
	}

	return nil
}

// IndustrySearchRequestMultiError is an error wrapping multiple validation
// errors returned by IndustrySearchRequest.ValidateAll() if the designated
// constraints aren't met.
type IndustrySearchRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndustrySearchRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndustrySearchRequestMultiError) AllErrors() []error { return m }

// IndustrySearchRequestValidationError is the validation error returned by
// IndustrySearchRequest.Validate if the designated constraints aren't met.
type IndustrySearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndustrySearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndustrySearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndustrySearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndustrySearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndustrySearchRequestValidationError) ErrorName() string {
	return "IndustrySearchRequestValidationError"
}

// Error satisfies the builtin error interface
func (e IndustrySearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndustrySearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndustrySearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndustrySearchRequestValidationError{}

// Validate checks the field values on IndustryActivitiesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IndustryActivitiesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndustryActivitiesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndustryActivitiesRequestMultiError, or nil if none found.
func (m *IndustryActivitiesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IndustryActivitiesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	_IndustryActivitiesRequest_Tickers_Unique := make(map[int64]struct{}, len(m.GetTickers()))

	for idx, item := range m.GetTickers() {
		_, _ = idx, item

		if _, exists := _IndustryActivitiesRequest_Tickers_Unique[item]; exists {
			err := IndustryActivitiesRequestValidationError{
				field:  fmt.Sprintf("Tickers[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_IndustryActivitiesRequest_Tickers_Unique[item] = struct{}{}
		}

		// no validation rules for Tickers[idx]
	}

	if len(errors) > 0 {
		return IndustryActivitiesRequestMultiError(errors)
	}

	return nil
}

// IndustryActivitiesRequestMultiError is an error wrapping multiple validation
// errors returned by IndustryActivitiesRequest.ValidateAll() if the
// designated constraints aren't met.
type IndustryActivitiesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndustryActivitiesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndustryActivitiesRequestMultiError) AllErrors() []error { return m }

// IndustryActivitiesRequestValidationError is the validation error returned by
// IndustryActivitiesRequest.Validate if the designated constraints aren't met.
type IndustryActivitiesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndustryActivitiesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndustryActivitiesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndustryActivitiesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndustryActivitiesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndustryActivitiesRequestValidationError) ErrorName() string {
	return "IndustryActivitiesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e IndustryActivitiesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndustryActivitiesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndustryActivitiesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndustryActivitiesRequestValidationError{}

// Validate checks the field values on IndustryReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IndustryReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndustryReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IndustryReplyMultiError, or
// nil if none found.
func (m *IndustryReply) ValidateAll() error {
	return m.validate(true)
}

func (m *IndustryReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]int64, len(m.GetResults()))
		i := 0
		for key := range m.GetResults() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetResults()[key]
			_ = val

			// no validation rules for Results[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, IndustryReplyValidationError{
							field:  fmt.Sprintf("Results[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, IndustryReplyValidationError{
							field:  fmt.Sprintf("Results[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return IndustryReplyValidationError{
						field:  fmt.Sprintf("Results[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return IndustryReplyMultiError(errors)
	}

	return nil
}

// IndustryReplyMultiError is an error wrapping multiple validation errors
// returned by IndustryReply.ValidateAll() if the designated constraints
// aren't met.
type IndustryReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndustryReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndustryReplyMultiError) AllErrors() []error { return m }

// IndustryReplyValidationError is the validation error returned by
// IndustryReply.Validate if the designated constraints aren't met.
type IndustryReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndustryReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndustryReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndustryReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndustryReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndustryReplyValidationError) ErrorName() string { return "IndustryReplyValidationError" }

// Error satisfies the builtin error interface
func (e IndustryReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndustryReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndustryReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndustryReplyValidationError{}

// Validate checks the field values on IndustrySearchReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IndustrySearchReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndustrySearchReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndustrySearchReplyMultiError, or nil if none found.
func (m *IndustrySearchReply) ValidateAll() error {
	return m.validate(true)
}

func (m *IndustrySearchReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Results

	if len(errors) > 0 {
		return IndustrySearchReplyMultiError(errors)
	}

	return nil
}

// IndustrySearchReplyMultiError is an error wrapping multiple validation
// errors returned by IndustrySearchReply.ValidateAll() if the designated
// constraints aren't met.
type IndustrySearchReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndustrySearchReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndustrySearchReplyMultiError) AllErrors() []error { return m }

// IndustrySearchReplyValidationError is the validation error returned by
// IndustrySearchReply.Validate if the designated constraints aren't met.
type IndustrySearchReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndustrySearchReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndustrySearchReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndustrySearchReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndustrySearchReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndustrySearchReplyValidationError) ErrorName() string {
	return "IndustrySearchReplyValidationError"
}

// Error satisfies the builtin error interface
func (e IndustrySearchReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndustrySearchReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndustrySearchReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndustrySearchReplyValidationError{}

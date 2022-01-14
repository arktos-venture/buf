// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/industry/v1/industry.proto

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

// Validate checks the field values on IndustryRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IndustryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndustryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndustryRequestMultiError, or nil if none found.
func (m *IndustryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IndustryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetRef(); val < 5010101010 || val > 6310301010 {
		err := IndustryRequestValidationError{
			field:  "Ref",
			reason: "value must be inside range [5010101010, 6310301010]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IndustryRequestMultiError(errors)
	}

	return nil
}

// IndustryRequestMultiError is an error wrapping multiple validation errors
// returned by IndustryRequest.ValidateAll() if the designated constraints
// aren't met.
type IndustryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndustryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndustryRequestMultiError) AllErrors() []error { return m }

// IndustryRequestValidationError is the validation error returned by
// IndustryRequest.Validate if the designated constraints aren't met.
type IndustryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndustryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndustryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndustryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndustryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndustryRequestValidationError) ErrorName() string { return "IndustryRequestValidationError" }

// Error satisfies the builtin error interface
func (e IndustryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndustryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndustryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndustryRequestValidationError{}

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

// Validate checks the field values on IndustryReply_Ref with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IndustryReply_Ref) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndustryReply_Ref with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndustryReply_RefMultiError, or nil if none found.
func (m *IndustryReply_Ref) ValidateAll() error {
	return m.validate(true)
}

func (m *IndustryReply_Ref) validate(all bool) error {
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
						errors = append(errors, IndustryReply_RefValidationError{
							field:  fmt.Sprintf("Next[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, IndustryReply_RefValidationError{
							field:  fmt.Sprintf("Next[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return IndustryReply_RefValidationError{
						field:  fmt.Sprintf("Next[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return IndustryReply_RefMultiError(errors)
	}

	return nil
}

// IndustryReply_RefMultiError is an error wrapping multiple validation errors
// returned by IndustryReply_Ref.ValidateAll() if the designated constraints
// aren't met.
type IndustryReply_RefMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndustryReply_RefMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndustryReply_RefMultiError) AllErrors() []error { return m }

// IndustryReply_RefValidationError is the validation error returned by
// IndustryReply_Ref.Validate if the designated constraints aren't met.
type IndustryReply_RefValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndustryReply_RefValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndustryReply_RefValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndustryReply_RefValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndustryReply_RefValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndustryReply_RefValidationError) ErrorName() string {
	return "IndustryReply_RefValidationError"
}

// Error satisfies the builtin error interface
func (e IndustryReply_RefValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndustryReply_Ref.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndustryReply_RefValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndustryReply_RefValidationError{}

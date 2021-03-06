// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/dividends/v1/dividends.proto

package v1Dividends

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

// Validate checks the field values on DividendRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DividendRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DividendRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DividendRequestMultiError, or nil if none found.
func (m *DividendRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DividendRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := v1Screener.Interval_name[int32(m.GetInterval())]; !ok {
		err := DividendRequestValidationError{
			field:  "Interval",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := len(m.GetFilters()); l < 1 || l > 20 {
		err := DividendRequestValidationError{
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
					errors = append(errors, DividendRequestValidationError{
						field:  fmt.Sprintf("Filters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DividendRequestValidationError{
						field:  fmt.Sprintf("Filters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DividendRequestValidationError{
					field:  fmt.Sprintf("Filters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DividendRequestMultiError(errors)
	}

	return nil
}

// DividendRequestMultiError is an error wrapping multiple validation errors
// returned by DividendRequest.ValidateAll() if the designated constraints
// aren't met.
type DividendRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DividendRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DividendRequestMultiError) AllErrors() []error { return m }

// DividendRequestValidationError is the validation error returned by
// DividendRequest.Validate if the designated constraints aren't met.
type DividendRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DividendRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DividendRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DividendRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DividendRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DividendRequestValidationError) ErrorName() string { return "DividendRequestValidationError" }

// Error satisfies the builtin error interface
func (e DividendRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDividendRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DividendRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DividendRequestValidationError{}

// Validate checks the field values on DividendLastRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DividendLastRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DividendLastRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DividendLastRequestMultiError, or nil if none found.
func (m *DividendLastRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DividendLastRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetInstrumentId() >= 0 {
		err := DividendLastRequestValidationError{
			field:  "InstrumentId",
			reason: "value must be less than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DividendLastRequestMultiError(errors)
	}

	return nil
}

// DividendLastRequestMultiError is an error wrapping multiple validation
// errors returned by DividendLastRequest.ValidateAll() if the designated
// constraints aren't met.
type DividendLastRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DividendLastRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DividendLastRequestMultiError) AllErrors() []error { return m }

// DividendLastRequestValidationError is the validation error returned by
// DividendLastRequest.Validate if the designated constraints aren't met.
type DividendLastRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DividendLastRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DividendLastRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DividendLastRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DividendLastRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DividendLastRequestValidationError) ErrorName() string {
	return "DividendLastRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DividendLastRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDividendLastRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DividendLastRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DividendLastRequestValidationError{}

// Validate checks the field values on DividendLastReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DividendLastReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DividendLastReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DividendLastReplyMultiError, or nil if none found.
func (m *DividendLastReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DividendLastReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Values

	// no validation rules for Yield

	if all {
		switch v := interface{}(m.GetDeclarationDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DividendLastReplyValidationError{
					field:  "DeclarationDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DividendLastReplyValidationError{
					field:  "DeclarationDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDeclarationDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DividendLastReplyValidationError{
				field:  "DeclarationDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRecordDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DividendLastReplyValidationError{
					field:  "RecordDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DividendLastReplyValidationError{
					field:  "RecordDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRecordDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DividendLastReplyValidationError{
				field:  "RecordDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPaymentDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DividendLastReplyValidationError{
					field:  "PaymentDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DividendLastReplyValidationError{
					field:  "PaymentDate",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaymentDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DividendLastReplyValidationError{
				field:  "PaymentDate",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DividendLastReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DividendLastReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DividendLastReplyValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DividendLastReplyMultiError(errors)
	}

	return nil
}

// DividendLastReplyMultiError is an error wrapping multiple validation errors
// returned by DividendLastReply.ValidateAll() if the designated constraints
// aren't met.
type DividendLastReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DividendLastReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DividendLastReplyMultiError) AllErrors() []error { return m }

// DividendLastReplyValidationError is the validation error returned by
// DividendLastReply.Validate if the designated constraints aren't met.
type DividendLastReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DividendLastReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DividendLastReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DividendLastReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DividendLastReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DividendLastReplyValidationError) ErrorName() string {
	return "DividendLastReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DividendLastReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDividendLastReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DividendLastReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DividendLastReplyValidationError{}

// Validate checks the field values on DividendReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DividendReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DividendReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DividendReplyMultiError, or
// nil if none found.
func (m *DividendReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DividendReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetDeclarationDate() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DividendReplyValidationError{
						field:  fmt.Sprintf("DeclarationDate[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DividendReplyValidationError{
						field:  fmt.Sprintf("DeclarationDate[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DividendReplyValidationError{
					field:  fmt.Sprintf("DeclarationDate[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetRecordDate() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DividendReplyValidationError{
						field:  fmt.Sprintf("RecordDate[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DividendReplyValidationError{
						field:  fmt.Sprintf("RecordDate[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DividendReplyValidationError{
					field:  fmt.Sprintf("RecordDate[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetPaymentDate() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DividendReplyValidationError{
						field:  fmt.Sprintf("PaymentDate[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DividendReplyValidationError{
						field:  fmt.Sprintf("PaymentDate[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DividendReplyValidationError{
					field:  fmt.Sprintf("PaymentDate[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetCreatedAt() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DividendReplyValidationError{
						field:  fmt.Sprintf("CreatedAt[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DividendReplyValidationError{
						field:  fmt.Sprintf("CreatedAt[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DividendReplyValidationError{
					field:  fmt.Sprintf("CreatedAt[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return DividendReplyMultiError(errors)
	}

	return nil
}

// DividendReplyMultiError is an error wrapping multiple validation errors
// returned by DividendReply.ValidateAll() if the designated constraints
// aren't met.
type DividendReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DividendReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DividendReplyMultiError) AllErrors() []error { return m }

// DividendReplyValidationError is the validation error returned by
// DividendReply.Validate if the designated constraints aren't met.
type DividendReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DividendReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DividendReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DividendReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DividendReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DividendReplyValidationError) ErrorName() string { return "DividendReplyValidationError" }

// Error satisfies the builtin error interface
func (e DividendReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDividendReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DividendReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DividendReplyValidationError{}

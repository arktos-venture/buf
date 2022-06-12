// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/dividends/v1/dividends.proto

package dividends_v1

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

	screener_v1 "github.com/arktos-venture/buf/proto/screener/v1"
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

	_ = screener_v1.Operator(0)
)

// Validate checks the field values on Filter with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Filter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Filter with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in FilterMultiError, or nil if none found.
func (m *Filter) ValidateAll() error {
	return m.validate(true)
}

func (m *Filter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := screener_v1.Operator_name[int32(m.GetOperator())]; !ok {
		err := FilterValidationError{
			field:  "Operator",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := screener_v1.Argument_name[int32(m.GetArgument())]; !ok {
		err := FilterValidationError{
			field:  "Argument",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Value

	if len(errors) > 0 {
		return FilterMultiError(errors)
	}

	return nil
}

// FilterMultiError is an error wrapping multiple validation errors returned by
// Filter.ValidateAll() if the designated constraints aren't met.
type FilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FilterMultiError) AllErrors() []error { return m }

// FilterValidationError is the validation error returned by Filter.Validate if
// the designated constraints aren't met.
type FilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterValidationError) ErrorName() string { return "FilterValidationError" }

// Error satisfies the builtin error interface
func (e FilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterValidationError{}

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

	if _, ok := Interval_name[int32(m.GetInterval())]; !ok {
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

	// no validation rules for Exchange

	// no validation rules for Ticker

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

// Validate checks the field values on DividendDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DividendDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DividendDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DividendDeleteRequestMultiError, or nil if none found.
func (m *DividendDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DividendDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	_DividendDeleteRequest_Tickers_Unique := make(map[string]struct{}, len(m.GetTickers()))

	for idx, item := range m.GetTickers() {
		_, _ = idx, item

		if _, exists := _DividendDeleteRequest_Tickers_Unique[item]; exists {
			err := DividendDeleteRequestValidationError{
				field:  fmt.Sprintf("Tickers[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_DividendDeleteRequest_Tickers_Unique[item] = struct{}{}
		}

		// no validation rules for Tickers[idx]
	}

	_DividendDeleteRequest_Exchanges_Unique := make(map[string]struct{}, len(m.GetExchanges()))

	for idx, item := range m.GetExchanges() {
		_, _ = idx, item

		if _, exists := _DividendDeleteRequest_Exchanges_Unique[item]; exists {
			err := DividendDeleteRequestValidationError{
				field:  fmt.Sprintf("Exchanges[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_DividendDeleteRequest_Exchanges_Unique[item] = struct{}{}
		}

		// no validation rules for Exchanges[idx]
	}

	if len(errors) > 0 {
		return DividendDeleteRequestMultiError(errors)
	}

	return nil
}

// DividendDeleteRequestMultiError is an error wrapping multiple validation
// errors returned by DividendDeleteRequest.ValidateAll() if the designated
// constraints aren't met.
type DividendDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DividendDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DividendDeleteRequestMultiError) AllErrors() []error { return m }

// DividendDeleteRequestValidationError is the validation error returned by
// DividendDeleteRequest.Validate if the designated constraints aren't met.
type DividendDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DividendDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DividendDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DividendDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DividendDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DividendDeleteRequestValidationError) ErrorName() string {
	return "DividendDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DividendDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDividendDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DividendDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DividendDeleteRequestValidationError{}

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

// Validate checks the field values on DividendDeleteReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DividendDeleteReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DividendDeleteReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DividendDeleteReplyMultiError, or nil if none found.
func (m *DividendDeleteReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DividendDeleteReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	if len(errors) > 0 {
		return DividendDeleteReplyMultiError(errors)
	}

	return nil
}

// DividendDeleteReplyMultiError is an error wrapping multiple validation
// errors returned by DividendDeleteReply.ValidateAll() if the designated
// constraints aren't met.
type DividendDeleteReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DividendDeleteReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DividendDeleteReplyMultiError) AllErrors() []error { return m }

// DividendDeleteReplyValidationError is the validation error returned by
// DividendDeleteReply.Validate if the designated constraints aren't met.
type DividendDeleteReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DividendDeleteReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DividendDeleteReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DividendDeleteReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DividendDeleteReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DividendDeleteReplyValidationError) ErrorName() string {
	return "DividendDeleteReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DividendDeleteReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDividendDeleteReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DividendDeleteReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DividendDeleteReplyValidationError{}

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/order/v1/order.proto

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

// Validate checks the field values on OrderRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderRequestMultiError, or
// nil if none found.
func (m *OrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetAccount()); l < 3 || l > 15 {
		err := OrderRequestValidationError{
			field:  "Account",
			reason: "value length must be between 3 and 15 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OrderRequestMultiError(errors)
	}

	return nil
}

// OrderRequestMultiError is an error wrapping multiple validation errors
// returned by OrderRequest.ValidateAll() if the designated constraints aren't met.
type OrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderRequestMultiError) AllErrors() []error { return m }

// OrderRequestValidationError is the validation error returned by
// OrderRequest.Validate if the designated constraints aren't met.
type OrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderRequestValidationError) ErrorName() string { return "OrderRequestValidationError" }

// Error satisfies the builtin error interface
func (e OrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderRequestValidationError{}

// Validate checks the field values on OrderSearchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OrderSearchRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderSearchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderSearchRequestMultiError, or nil if none found.
func (m *OrderSearchRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderSearchRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetAccount()); l < 3 || l > 15 {
		err := OrderSearchRequestValidationError{
			field:  "Account",
			reason: "value length must be between 3 and 15 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := Action_name[int32(m.GetAction())]; !ok {
		err := OrderSearchRequestValidationError{
			field:  "Action",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := Duration_name[int32(m.GetDuration())]; !ok {
		err := OrderSearchRequestValidationError{
			field:  "Duration",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := OrderType_name[int32(m.GetOrderType())]; !ok {
		err := OrderSearchRequestValidationError{
			field:  "OrderType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _OrderSearchRequest_Period_InLookup[m.GetPeriod()]; !ok {
		err := OrderSearchRequestValidationError{
			field:  "Period",
			reason: "value must be in list [last 3d 1w 2w 1m 2m 3m 6m 1y]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OrderSearchRequestMultiError(errors)
	}

	return nil
}

// OrderSearchRequestMultiError is an error wrapping multiple validation errors
// returned by OrderSearchRequest.ValidateAll() if the designated constraints
// aren't met.
type OrderSearchRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderSearchRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderSearchRequestMultiError) AllErrors() []error { return m }

// OrderSearchRequestValidationError is the validation error returned by
// OrderSearchRequest.Validate if the designated constraints aren't met.
type OrderSearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderSearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderSearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderSearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderSearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderSearchRequestValidationError) ErrorName() string {
	return "OrderSearchRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OrderSearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderSearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderSearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderSearchRequestValidationError{}

var _OrderSearchRequest_Period_InLookup = map[string]struct{}{
	"last": {},
	"3d":   {},
	"1w":   {},
	"2w":   {},
	"1m":   {},
	"2m":   {},
	"3m":   {},
	"6m":   {},
	"1y":   {},
}

// Validate checks the field values on OrderCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OrderCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderCreateRequestMultiError, or nil if none found.
func (m *OrderCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetAccount()); l < 3 || l > 15 {
		err := OrderCreateRequestValidationError{
			field:  "Account",
			reason: "value length must be between 3 and 15 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetTicker()); l < 1 || l > 8 {
		err := OrderCreateRequestValidationError{
			field:  "Ticker",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetExchange()); l < 1 || l > 8 {
		err := OrderCreateRequestValidationError{
			field:  "Exchange",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCurrency()) != 3 {
		err := OrderCreateRequestValidationError{
			field:  "Currency",
			reason: "value length must be 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if _, ok := Action_name[int32(m.GetAction())]; !ok {
		err := OrderCreateRequestValidationError{
			field:  "Action",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := OrderType_name[int32(m.GetOrderType())]; !ok {
		err := OrderCreateRequestValidationError{
			field:  "OrderType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetSize(); val <= 0.01 || val >= 1e+08 {
		err := OrderCreateRequestValidationError{
			field:  "Size",
			reason: "value must be inside range (0.01, 1e+08)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := Duration_name[int32(m.GetDuration())]; !ok {
		err := OrderCreateRequestValidationError{
			field:  "Duration",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPrice(); val <= 1e-06 || val >= 1e+08 {
		err := OrderCreateRequestValidationError{
			field:  "Price",
			reason: "value must be inside range (1e-06, 1e+08)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OrderCreateRequestMultiError(errors)
	}

	return nil
}

// OrderCreateRequestMultiError is an error wrapping multiple validation errors
// returned by OrderCreateRequest.ValidateAll() if the designated constraints
// aren't met.
type OrderCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderCreateRequestMultiError) AllErrors() []error { return m }

// OrderCreateRequestValidationError is the validation error returned by
// OrderCreateRequest.Validate if the designated constraints aren't met.
type OrderCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderCreateRequestValidationError) ErrorName() string {
	return "OrderCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e OrderCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderCreateRequestValidationError{}

// Validate checks the field values on OrderReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderReplyMultiError, or
// nil if none found.
func (m *OrderReply) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ticker

	// no validation rules for Exchange

	// no validation rules for Currency

	// no validation rules for Action

	// no validation rules for OrderType

	// no validation rules for Routing

	// no validation rules for Duration

	if all {
		switch v := interface{}(m.GetSize()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderReplyValidationError{
					field:  "Size",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderReplyValidationError{
					field:  "Size",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSize()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderReplyValidationError{
				field:  "Size",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPrice()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrderReplyValidationError{
					field:  "Price",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrderReplyValidationError{
					field:  "Price",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPrice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrderReplyValidationError{
				field:  "Price",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Status

	// no validation rules for Date

	if len(errors) > 0 {
		return OrderReplyMultiError(errors)
	}

	return nil
}

// OrderReplyMultiError is an error wrapping multiple validation errors
// returned by OrderReply.ValidateAll() if the designated constraints aren't met.
type OrderReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderReplyMultiError) AllErrors() []error { return m }

// OrderReplyValidationError is the validation error returned by
// OrderReply.Validate if the designated constraints aren't met.
type OrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderReplyValidationError) ErrorName() string { return "OrderReplyValidationError" }

// Error satisfies the builtin error interface
func (e OrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderReplyValidationError{}

// Validate checks the field values on OrderReplies with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderReplies) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderReplies with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderRepliesMultiError, or
// nil if none found.
func (m *OrderReplies) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderReplies) validate(all bool) error {
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
					errors = append(errors, OrderRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OrderRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OrderRepliesValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return OrderRepliesMultiError(errors)
	}

	return nil
}

// OrderRepliesMultiError is an error wrapping multiple validation errors
// returned by OrderReplies.ValidateAll() if the designated constraints aren't met.
type OrderRepliesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderRepliesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderRepliesMultiError) AllErrors() []error { return m }

// OrderRepliesValidationError is the validation error returned by
// OrderReplies.Validate if the designated constraints aren't met.
type OrderRepliesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderRepliesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderRepliesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderRepliesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderRepliesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderRepliesValidationError) ErrorName() string { return "OrderRepliesValidationError" }

// Error satisfies the builtin error interface
func (e OrderRepliesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderReplies.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderRepliesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderRepliesValidationError{}

// Validate checks the field values on OrderReply_Size with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OrderReply_Size) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderReply_Size with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderReply_SizeMultiError, or nil if none found.
func (m *OrderReply_Size) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderReply_Size) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ask

	// no validation rules for Buy

	if len(errors) > 0 {
		return OrderReply_SizeMultiError(errors)
	}

	return nil
}

// OrderReply_SizeMultiError is an error wrapping multiple validation errors
// returned by OrderReply_Size.ValidateAll() if the designated constraints
// aren't met.
type OrderReply_SizeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderReply_SizeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderReply_SizeMultiError) AllErrors() []error { return m }

// OrderReply_SizeValidationError is the validation error returned by
// OrderReply_Size.Validate if the designated constraints aren't met.
type OrderReply_SizeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderReply_SizeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderReply_SizeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderReply_SizeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderReply_SizeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderReply_SizeValidationError) ErrorName() string { return "OrderReply_SizeValidationError" }

// Error satisfies the builtin error interface
func (e OrderReply_SizeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderReply_Size.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderReply_SizeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderReply_SizeValidationError{}

// Validate checks the field values on OrderReply_Price with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OrderReply_Price) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderReply_Price with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrderReply_PriceMultiError, or nil if none found.
func (m *OrderReply_Price) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderReply_Price) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ask

	// no validation rules for Buy

	// no validation rules for Commission

	if len(errors) > 0 {
		return OrderReply_PriceMultiError(errors)
	}

	return nil
}

// OrderReply_PriceMultiError is an error wrapping multiple validation errors
// returned by OrderReply_Price.ValidateAll() if the designated constraints
// aren't met.
type OrderReply_PriceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderReply_PriceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderReply_PriceMultiError) AllErrors() []error { return m }

// OrderReply_PriceValidationError is the validation error returned by
// OrderReply_Price.Validate if the designated constraints aren't met.
type OrderReply_PriceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderReply_PriceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderReply_PriceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderReply_PriceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderReply_PriceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderReply_PriceValidationError) ErrorName() string { return "OrderReply_PriceValidationError" }

// Error satisfies the builtin error interface
func (e OrderReply_PriceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderReply_Price.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderReply_PriceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderReply_PriceValidationError{}

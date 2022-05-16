// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/notifications/v1/notifications.proto

package notifications_v1

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

// Validate checks the field values on Page with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Page) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Page with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PageMultiError, or nil if none found.
func (m *Page) ValidateAll() error {
	return m.validate(true)
}

func (m *Page) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetNumber(); val <= 0 || val > 10000 {
		err := PageValidationError{
			field:  "Number",
			reason: "value must be inside range (0, 10000]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetLimit(); val <= 0 || val > 150 {
		err := PageValidationError{
			field:  "Limit",
			reason: "value must be inside range (0, 150]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PageMultiError(errors)
	}

	return nil
}

// PageMultiError is an error wrapping multiple validation errors returned by
// Page.ValidateAll() if the designated constraints aren't met.
type PageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageMultiError) AllErrors() []error { return m }

// PageValidationError is the validation error returned by Page.Validate if the
// designated constraints aren't met.
type PageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageValidationError) ErrorName() string { return "PageValidationError" }

// Error satisfies the builtin error interface
func (e PageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageValidationError{}

// Validate checks the field values on NotificationCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *NotificationCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NotificationCreateRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NotificationCreateRequestMultiError, or nil if none found.
func (m *NotificationCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *NotificationCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Account

	if _, ok := Level_name[int32(m.GetLevel())]; !ok {
		err := NotificationCreateRequestValidationError{
			field:  "Level",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Subject

	// no validation rules for Title

	// no validation rules for Description

	if len(errors) > 0 {
		return NotificationCreateRequestMultiError(errors)
	}

	return nil
}

// NotificationCreateRequestMultiError is an error wrapping multiple validation
// errors returned by NotificationCreateRequest.ValidateAll() if the
// designated constraints aren't met.
type NotificationCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NotificationCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NotificationCreateRequestMultiError) AllErrors() []error { return m }

// NotificationCreateRequestValidationError is the validation error returned by
// NotificationCreateRequest.Validate if the designated constraints aren't met.
type NotificationCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotificationCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotificationCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotificationCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotificationCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotificationCreateRequestValidationError) ErrorName() string {
	return "NotificationCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e NotificationCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotificationCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotificationCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotificationCreateRequestValidationError{}

// Validate checks the field values on NotificationSearchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *NotificationSearchRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NotificationSearchRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NotificationSearchRequestMultiError, or nil if none found.
func (m *NotificationSearchRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *NotificationSearchRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Account

	if _, ok := Level_name[int32(m.GetLevel())]; !ok {
		err := NotificationSearchRequestValidationError{
			field:  "Level",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _NotificationSearchRequest_Period_InLookup[m.GetPeriod()]; !ok {
		err := NotificationSearchRequestValidationError{
			field:  "Period",
			reason: "value must be in list [last 3d 1w 2w 1m]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPage() == nil {
		err := NotificationSearchRequestValidationError{
			field:  "Page",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetPage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NotificationSearchRequestValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NotificationSearchRequestValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NotificationSearchRequestValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return NotificationSearchRequestMultiError(errors)
	}

	return nil
}

// NotificationSearchRequestMultiError is an error wrapping multiple validation
// errors returned by NotificationSearchRequest.ValidateAll() if the
// designated constraints aren't met.
type NotificationSearchRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NotificationSearchRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NotificationSearchRequestMultiError) AllErrors() []error { return m }

// NotificationSearchRequestValidationError is the validation error returned by
// NotificationSearchRequest.Validate if the designated constraints aren't met.
type NotificationSearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotificationSearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotificationSearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotificationSearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotificationSearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotificationSearchRequestValidationError) ErrorName() string {
	return "NotificationSearchRequestValidationError"
}

// Error satisfies the builtin error interface
func (e NotificationSearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotificationSearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotificationSearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotificationSearchRequestValidationError{}

var _NotificationSearchRequest_Period_InLookup = map[string]struct{}{
	"last": {},
	"3d":   {},
	"1w":   {},
	"2w":   {},
	"1m":   {},
}

// Validate checks the field values on NotificationReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *NotificationReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NotificationReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NotificationReplyMultiError, or nil if none found.
func (m *NotificationReply) ValidateAll() error {
	return m.validate(true)
}

func (m *NotificationReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Subject

	// no validation rules for Level

	// no validation rules for Title

	// no validation rules for Description

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NotificationReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NotificationReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NotificationReplyValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return NotificationReplyMultiError(errors)
	}

	return nil
}

// NotificationReplyMultiError is an error wrapping multiple validation errors
// returned by NotificationReply.ValidateAll() if the designated constraints
// aren't met.
type NotificationReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NotificationReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NotificationReplyMultiError) AllErrors() []error { return m }

// NotificationReplyValidationError is the validation error returned by
// NotificationReply.Validate if the designated constraints aren't met.
type NotificationReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotificationReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotificationReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotificationReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotificationReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotificationReplyValidationError) ErrorName() string {
	return "NotificationReplyValidationError"
}

// Error satisfies the builtin error interface
func (e NotificationReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotificationReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotificationReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotificationReplyValidationError{}

// Validate checks the field values on NotificationReplies with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *NotificationReplies) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NotificationReplies with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// NotificationRepliesMultiError, or nil if none found.
func (m *NotificationReplies) ValidateAll() error {
	return m.validate(true)
}

func (m *NotificationReplies) validate(all bool) error {
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
					errors = append(errors, NotificationRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NotificationRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NotificationRepliesValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return NotificationRepliesMultiError(errors)
	}

	return nil
}

// NotificationRepliesMultiError is an error wrapping multiple validation
// errors returned by NotificationReplies.ValidateAll() if the designated
// constraints aren't met.
type NotificationRepliesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NotificationRepliesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NotificationRepliesMultiError) AllErrors() []error { return m }

// NotificationRepliesValidationError is the validation error returned by
// NotificationReplies.Validate if the designated constraints aren't met.
type NotificationRepliesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NotificationRepliesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NotificationRepliesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NotificationRepliesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NotificationRepliesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NotificationRepliesValidationError) ErrorName() string {
	return "NotificationRepliesValidationError"
}

// Error satisfies the builtin error interface
func (e NotificationRepliesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNotificationReplies.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NotificationRepliesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NotificationRepliesValidationError{}

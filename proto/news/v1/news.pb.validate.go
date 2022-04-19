// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/news/v1/news.proto

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

// Validate checks the field values on NewsRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NewsRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NewsRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NewsRequestMultiError, or
// nil if none found.
func (m *NewsRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *NewsRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTicker()); l < 1 || l > 8 {
		err := NewsRequestValidationError{
			field:  "Ticker",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetExchange()); l < 1 || l > 8 {
		err := NewsRequestValidationError{
			field:  "Exchange",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _NewsRequest_Period_InLookup[m.GetPeriod()]; !ok {
		err := NewsRequestValidationError{
			field:  "Period",
			reason: "value must be in list [last 3d 1w]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return NewsRequestMultiError(errors)
	}

	return nil
}

// NewsRequestMultiError is an error wrapping multiple validation errors
// returned by NewsRequest.ValidateAll() if the designated constraints aren't met.
type NewsRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NewsRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NewsRequestMultiError) AllErrors() []error { return m }

// NewsRequestValidationError is the validation error returned by
// NewsRequest.Validate if the designated constraints aren't met.
type NewsRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewsRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewsRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewsRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewsRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewsRequestValidationError) ErrorName() string { return "NewsRequestValidationError" }

// Error satisfies the builtin error interface
func (e NewsRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewsRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewsRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewsRequestValidationError{}

var _NewsRequest_Period_InLookup = map[string]struct{}{
	"last": {},
	"3d":   {},
	"1w":   {},
}

// Validate checks the field values on NewsReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NewsReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NewsReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NewsReplyMultiError, or nil
// if none found.
func (m *NewsReply) ValidateAll() error {
	return m.validate(true)
}

func (m *NewsReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Title

	// no validation rules for Link

	if all {
		switch v := interface{}(m.GetDate()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, NewsReplyValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, NewsReplyValidationError{
					field:  "Date",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDate()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NewsReplyValidationError{
				field:  "Date",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return NewsReplyMultiError(errors)
	}

	return nil
}

// NewsReplyMultiError is an error wrapping multiple validation errors returned
// by NewsReply.ValidateAll() if the designated constraints aren't met.
type NewsReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NewsReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NewsReplyMultiError) AllErrors() []error { return m }

// NewsReplyValidationError is the validation error returned by
// NewsReply.Validate if the designated constraints aren't met.
type NewsReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewsReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewsReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewsReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewsReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewsReplyValidationError) ErrorName() string { return "NewsReplyValidationError" }

// Error satisfies the builtin error interface
func (e NewsReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewsReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewsReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewsReplyValidationError{}

// Validate checks the field values on NewsReplies with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *NewsReplies) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on NewsReplies with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in NewsRepliesMultiError, or
// nil if none found.
func (m *NewsReplies) ValidateAll() error {
	return m.validate(true)
}

func (m *NewsReplies) validate(all bool) error {
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
					errors = append(errors, NewsRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, NewsRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NewsRepliesValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return NewsRepliesMultiError(errors)
	}

	return nil
}

// NewsRepliesMultiError is an error wrapping multiple validation errors
// returned by NewsReplies.ValidateAll() if the designated constraints aren't met.
type NewsRepliesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m NewsRepliesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m NewsRepliesMultiError) AllErrors() []error { return m }

// NewsRepliesValidationError is the validation error returned by
// NewsReplies.Validate if the designated constraints aren't met.
type NewsRepliesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewsRepliesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewsRepliesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewsRepliesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewsRepliesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewsRepliesValidationError) ErrorName() string { return "NewsRepliesValidationError" }

// Error satisfies the builtin error interface
func (e NewsRepliesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewsReplies.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewsRepliesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewsRepliesValidationError{}

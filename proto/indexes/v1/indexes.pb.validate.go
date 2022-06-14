// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/indexes/v1/indexes.proto

package indexes_v1

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

// Validate checks the field values on IndexRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IndexRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndexRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IndexRequestMultiError, or
// nil if none found.
func (m *IndexRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IndexRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTicker()); l < 2 || l > 36 {
		err := IndexRequestValidationError{
			field:  "Ticker",
			reason: "value length must be between 2 and 36 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IndexRequestMultiError(errors)
	}

	return nil
}

// IndexRequestMultiError is an error wrapping multiple validation errors
// returned by IndexRequest.ValidateAll() if the designated constraints aren't met.
type IndexRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndexRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndexRequestMultiError) AllErrors() []error { return m }

// IndexRequestValidationError is the validation error returned by
// IndexRequest.Validate if the designated constraints aren't met.
type IndexRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndexRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndexRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndexRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndexRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndexRequestValidationError) ErrorName() string { return "IndexRequestValidationError" }

// Error satisfies the builtin error interface
func (e IndexRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndexRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndexRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndexRequestValidationError{}

// Validate checks the field values on IndexSearchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IndexSearchRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndexSearchRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndexSearchRequestMultiError, or nil if none found.
func (m *IndexSearchRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IndexSearchRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetCurrency()) != 3 {
		err := IndexSearchRequestValidationError{
			field:  "Currency",
			reason: "value length must be 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if m.GetPage() == nil {
		err := IndexSearchRequestValidationError{
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
				errors = append(errors, IndexSearchRequestValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IndexSearchRequestValidationError{
					field:  "Page",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IndexSearchRequestValidationError{
				field:  "Page",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return IndexSearchRequestMultiError(errors)
	}

	return nil
}

// IndexSearchRequestMultiError is an error wrapping multiple validation errors
// returned by IndexSearchRequest.ValidateAll() if the designated constraints
// aren't met.
type IndexSearchRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndexSearchRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndexSearchRequestMultiError) AllErrors() []error { return m }

// IndexSearchRequestValidationError is the validation error returned by
// IndexSearchRequest.Validate if the designated constraints aren't met.
type IndexSearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndexSearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndexSearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndexSearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndexSearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndexSearchRequestValidationError) ErrorName() string {
	return "IndexSearchRequestValidationError"
}

// Error satisfies the builtin error interface
func (e IndexSearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndexSearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndexSearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndexSearchRequestValidationError{}

// Validate checks the field values on IndexCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IndexCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndexCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndexCreateRequestMultiError, or nil if none found.
func (m *IndexCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IndexCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTicker()); l < 2 || l > 36 {
		err := IndexCreateRequestValidationError{
			field:  "Ticker",
			reason: "value length must be between 2 and 36 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetRequest()); l < 8 || l > 8196 {
		err := IndexCreateRequestValidationError{
			field:  "Request",
			reason: "value length must be between 8 and 8196 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Public

	if len(errors) > 0 {
		return IndexCreateRequestMultiError(errors)
	}

	return nil
}

// IndexCreateRequestMultiError is an error wrapping multiple validation errors
// returned by IndexCreateRequest.ValidateAll() if the designated constraints
// aren't met.
type IndexCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndexCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndexCreateRequestMultiError) AllErrors() []error { return m }

// IndexCreateRequestValidationError is the validation error returned by
// IndexCreateRequest.Validate if the designated constraints aren't met.
type IndexCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndexCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndexCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndexCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndexCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndexCreateRequestValidationError) ErrorName() string {
	return "IndexCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e IndexCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndexCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndexCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndexCreateRequestValidationError{}

// Validate checks the field values on IndexDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IndexDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndexDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndexDeleteRequestMultiError, or nil if none found.
func (m *IndexDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *IndexDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetTickers()) < 1 {
		err := IndexDeleteRequestValidationError{
			field:  "Tickers",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	_IndexDeleteRequest_Tickers_Unique := make(map[string]struct{}, len(m.GetTickers()))

	for idx, item := range m.GetTickers() {
		_, _ = idx, item

		if _, exists := _IndexDeleteRequest_Tickers_Unique[item]; exists {
			err := IndexDeleteRequestValidationError{
				field:  fmt.Sprintf("Tickers[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_IndexDeleteRequest_Tickers_Unique[item] = struct{}{}
		}

		// no validation rules for Tickers[idx]
	}

	if len(errors) > 0 {
		return IndexDeleteRequestMultiError(errors)
	}

	return nil
}

// IndexDeleteRequestMultiError is an error wrapping multiple validation errors
// returned by IndexDeleteRequest.ValidateAll() if the designated constraints
// aren't met.
type IndexDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndexDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndexDeleteRequestMultiError) AllErrors() []error { return m }

// IndexDeleteRequestValidationError is the validation error returned by
// IndexDeleteRequest.Validate if the designated constraints aren't met.
type IndexDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndexDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndexDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndexDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndexDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndexDeleteRequestValidationError) ErrorName() string {
	return "IndexDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e IndexDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndexDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndexDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndexDeleteRequestValidationError{}

// Validate checks the field values on IndexReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IndexReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndexReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IndexReplyMultiError, or
// nil if none found.
func (m *IndexReply) ValidateAll() error {
	return m.validate(true)
}

func (m *IndexReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ticker

	// no validation rules for Public

	for idx, item := range m.GetCompanies() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, IndexReplyValidationError{
						field:  fmt.Sprintf("Companies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IndexReplyValidationError{
						field:  fmt.Sprintf("Companies[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IndexReplyValidationError{
					field:  fmt.Sprintf("Companies[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IndexReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IndexReplyValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IndexReplyValidationError{
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
				errors = append(errors, IndexReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IndexReplyValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IndexReplyValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return IndexReplyMultiError(errors)
	}

	return nil
}

// IndexReplyMultiError is an error wrapping multiple validation errors
// returned by IndexReply.ValidateAll() if the designated constraints aren't met.
type IndexReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndexReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndexReplyMultiError) AllErrors() []error { return m }

// IndexReplyValidationError is the validation error returned by
// IndexReply.Validate if the designated constraints aren't met.
type IndexReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndexReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndexReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndexReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndexReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndexReplyValidationError) ErrorName() string { return "IndexReplyValidationError" }

// Error satisfies the builtin error interface
func (e IndexReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndexReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndexReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndexReplyValidationError{}

// Validate checks the field values on IndexReplies with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IndexReplies) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndexReplies with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IndexRepliesMultiError, or
// nil if none found.
func (m *IndexReplies) ValidateAll() error {
	return m.validate(true)
}

func (m *IndexReplies) validate(all bool) error {
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
					errors = append(errors, IndexRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IndexRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IndexRepliesValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return IndexRepliesMultiError(errors)
	}

	return nil
}

// IndexRepliesMultiError is an error wrapping multiple validation errors
// returned by IndexReplies.ValidateAll() if the designated constraints aren't met.
type IndexRepliesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndexRepliesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndexRepliesMultiError) AllErrors() []error { return m }

// IndexRepliesValidationError is the validation error returned by
// IndexReplies.Validate if the designated constraints aren't met.
type IndexRepliesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndexRepliesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndexRepliesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndexRepliesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndexRepliesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndexRepliesValidationError) ErrorName() string { return "IndexRepliesValidationError" }

// Error satisfies the builtin error interface
func (e IndexRepliesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndexReplies.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndexRepliesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndexRepliesValidationError{}

// Validate checks the field values on IndexSearchReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IndexSearchReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndexSearchReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndexSearchReplyMultiError, or nil if none found.
func (m *IndexSearchReply) ValidateAll() error {
	return m.validate(true)
}

func (m *IndexSearchReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ticker

	if len(errors) > 0 {
		return IndexSearchReplyMultiError(errors)
	}

	return nil
}

// IndexSearchReplyMultiError is an error wrapping multiple validation errors
// returned by IndexSearchReply.ValidateAll() if the designated constraints
// aren't met.
type IndexSearchReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndexSearchReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndexSearchReplyMultiError) AllErrors() []error { return m }

// IndexSearchReplyValidationError is the validation error returned by
// IndexSearchReply.Validate if the designated constraints aren't met.
type IndexSearchReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndexSearchReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndexSearchReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndexSearchReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndexSearchReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndexSearchReplyValidationError) ErrorName() string { return "IndexSearchReplyValidationError" }

// Error satisfies the builtin error interface
func (e IndexSearchReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndexSearchReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndexSearchReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndexSearchReplyValidationError{}

// Validate checks the field values on IndexSearchReplies with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IndexSearchReplies) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndexSearchReplies with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndexSearchRepliesMultiError, or nil if none found.
func (m *IndexSearchReplies) ValidateAll() error {
	return m.validate(true)
}

func (m *IndexSearchReplies) validate(all bool) error {
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
					errors = append(errors, IndexSearchRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, IndexSearchRepliesValidationError{
						field:  fmt.Sprintf("Results[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return IndexSearchRepliesValidationError{
					field:  fmt.Sprintf("Results[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return IndexSearchRepliesMultiError(errors)
	}

	return nil
}

// IndexSearchRepliesMultiError is an error wrapping multiple validation errors
// returned by IndexSearchReplies.ValidateAll() if the designated constraints
// aren't met.
type IndexSearchRepliesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndexSearchRepliesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndexSearchRepliesMultiError) AllErrors() []error { return m }

// IndexSearchRepliesValidationError is the validation error returned by
// IndexSearchReplies.Validate if the designated constraints aren't met.
type IndexSearchRepliesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndexSearchRepliesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndexSearchRepliesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndexSearchRepliesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndexSearchRepliesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndexSearchRepliesValidationError) ErrorName() string {
	return "IndexSearchRepliesValidationError"
}

// Error satisfies the builtin error interface
func (e IndexSearchRepliesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndexSearchReplies.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndexSearchRepliesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndexSearchRepliesValidationError{}

// Validate checks the field values on IndexDelete with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IndexDelete) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndexDelete with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IndexDeleteMultiError, or
// nil if none found.
func (m *IndexDelete) ValidateAll() error {
	return m.validate(true)
}

func (m *IndexDelete) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	if len(errors) > 0 {
		return IndexDeleteMultiError(errors)
	}

	return nil
}

// IndexDeleteMultiError is an error wrapping multiple validation errors
// returned by IndexDelete.ValidateAll() if the designated constraints aren't met.
type IndexDeleteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndexDeleteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndexDeleteMultiError) AllErrors() []error { return m }

// IndexDeleteValidationError is the validation error returned by
// IndexDelete.Validate if the designated constraints aren't met.
type IndexDeleteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndexDeleteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndexDeleteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndexDeleteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndexDeleteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndexDeleteValidationError) ErrorName() string { return "IndexDeleteValidationError" }

// Error satisfies the builtin error interface
func (e IndexDeleteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndexDelete.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndexDeleteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndexDeleteValidationError{}

// Validate checks the field values on IndexSearchRequest_Page with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IndexSearchRequest_Page) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IndexSearchRequest_Page with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IndexSearchRequest_PageMultiError, or nil if none found.
func (m *IndexSearchRequest_Page) ValidateAll() error {
	return m.validate(true)
}

func (m *IndexSearchRequest_Page) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetNumber(); val <= 0 || val > 10000 {
		err := IndexSearchRequest_PageValidationError{
			field:  "Number",
			reason: "value must be inside range (0, 10000]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetLimit(); val <= 0 || val > 150 {
		err := IndexSearchRequest_PageValidationError{
			field:  "Limit",
			reason: "value must be inside range (0, 150]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IndexSearchRequest_PageMultiError(errors)
	}

	return nil
}

// IndexSearchRequest_PageMultiError is an error wrapping multiple validation
// errors returned by IndexSearchRequest_Page.ValidateAll() if the designated
// constraints aren't met.
type IndexSearchRequest_PageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IndexSearchRequest_PageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IndexSearchRequest_PageMultiError) AllErrors() []error { return m }

// IndexSearchRequest_PageValidationError is the validation error returned by
// IndexSearchRequest_Page.Validate if the designated constraints aren't met.
type IndexSearchRequest_PageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IndexSearchRequest_PageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IndexSearchRequest_PageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IndexSearchRequest_PageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IndexSearchRequest_PageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IndexSearchRequest_PageValidationError) ErrorName() string {
	return "IndexSearchRequest_PageValidationError"
}

// Error satisfies the builtin error interface
func (e IndexSearchRequest_PageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIndexSearchRequest_Page.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IndexSearchRequest_PageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IndexSearchRequest_PageValidationError{}

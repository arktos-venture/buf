// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/broker/v1/broker.proto

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

// Validate checks the field values on BrokerRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BrokerRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BrokerRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BrokerRequestMultiError, or
// nil if none found.
func (m *BrokerRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BrokerRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetBroker()); l < 3 || l > 15 {
		err := BrokerRequestValidationError{
			field:  "Broker",
			reason: "value length must be between 3 and 15 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetUser()); l < 3 || l > 15 {
		err := BrokerRequestValidationError{
			field:  "User",
			reason: "value length must be between 3 and 15 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return BrokerRequestMultiError(errors)
	}
	return nil
}

// BrokerRequestMultiError is an error wrapping multiple validation errors
// returned by BrokerRequest.ValidateAll() if the designated constraints
// aren't met.
type BrokerRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BrokerRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BrokerRequestMultiError) AllErrors() []error { return m }

// BrokerRequestValidationError is the validation error returned by
// BrokerRequest.Validate if the designated constraints aren't met.
type BrokerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BrokerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BrokerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BrokerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BrokerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BrokerRequestValidationError) ErrorName() string { return "BrokerRequestValidationError" }

// Error satisfies the builtin error interface
func (e BrokerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBrokerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BrokerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BrokerRequestValidationError{}

// Validate checks the field values on BrokerCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *BrokerCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BrokerCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BrokerCreateRequestMultiError, or nil if none found.
func (m *BrokerCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BrokerCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetBroker()); l < 3 || l > 15 {
		err := BrokerCreateRequestValidationError{
			field:  "Broker",
			reason: "value length must be between 3 and 15 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 5 || l > 30 {
		err := BrokerCreateRequestValidationError{
			field:  "Name",
			reason: "value length must be between 5 and 30 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetDescription()); l < 15 || l > 1024 {
		err := BrokerCreateRequestValidationError{
			field:  "Description",
			reason: "value length must be between 15 and 1024 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if uri, err := url.Parse(m.GetWebsite()); err != nil {
		err = BrokerCreateRequestValidationError{
			field:  "Website",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	} else if !uri.IsAbs() {
		err := BrokerCreateRequestValidationError{
			field:  "Website",
			reason: "value must be absolute",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = BrokerCreateRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetExchanges()) < 1 {
		err := BrokerCreateRequestValidationError{
			field:  "Exchanges",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	_BrokerCreateRequest_Exchanges_Unique := make(map[string]struct{}, len(m.GetExchanges()))

	for idx, item := range m.GetExchanges() {
		_, _ = idx, item

		if _, exists := _BrokerCreateRequest_Exchanges_Unique[item]; exists {
			err := BrokerCreateRequestValidationError{
				field:  fmt.Sprintf("Exchanges[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_BrokerCreateRequest_Exchanges_Unique[item] = struct{}{}
		}

		// no validation rules for Exchanges[idx]
	}

	if len(errors) > 0 {
		return BrokerCreateRequestMultiError(errors)
	}
	return nil
}

func (m *BrokerCreateRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *BrokerCreateRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// BrokerCreateRequestMultiError is an error wrapping multiple validation
// errors returned by BrokerCreateRequest.ValidateAll() if the designated
// constraints aren't met.
type BrokerCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BrokerCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BrokerCreateRequestMultiError) AllErrors() []error { return m }

// BrokerCreateRequestValidationError is the validation error returned by
// BrokerCreateRequest.Validate if the designated constraints aren't met.
type BrokerCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BrokerCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BrokerCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BrokerCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BrokerCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BrokerCreateRequestValidationError) ErrorName() string {
	return "BrokerCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e BrokerCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBrokerCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BrokerCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BrokerCreateRequestValidationError{}

// Validate checks the field values on BrokerReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BrokerReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BrokerReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BrokerReplyMultiError, or
// nil if none found.
func (m *BrokerReply) ValidateAll() error {
	return m.validate(true)
}

func (m *BrokerReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ref

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Website

	// no validation rules for Active

	// no validation rules for Date

	if len(errors) > 0 {
		return BrokerReplyMultiError(errors)
	}
	return nil
}

// BrokerReplyMultiError is an error wrapping multiple validation errors
// returned by BrokerReply.ValidateAll() if the designated constraints aren't met.
type BrokerReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BrokerReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BrokerReplyMultiError) AllErrors() []error { return m }

// BrokerReplyValidationError is the validation error returned by
// BrokerReply.Validate if the designated constraints aren't met.
type BrokerReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BrokerReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BrokerReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BrokerReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BrokerReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BrokerReplyValidationError) ErrorName() string { return "BrokerReplyValidationError" }

// Error satisfies the builtin error interface
func (e BrokerReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBrokerReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BrokerReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BrokerReplyValidationError{}

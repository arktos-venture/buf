// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/strategy-results/v1/strategy-results.proto

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

// Validate checks the field values on StrategyCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StrategyCompanyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrategyCompanyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StrategyCompanyRequestMultiError, or nil if none found.
func (m *StrategyCompanyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StrategyCompanyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTicker()); l < 1 || l > 8 {
		err := StrategyCompanyRequestValidationError{
			field:  "Ticker",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetExchange()); l < 1 || l > 8 {
		err := StrategyCompanyRequestValidationError{
			field:  "Exchange",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCurrency()) != 3 {
		err := StrategyCompanyRequestValidationError{
			field:  "Currency",
			reason: "value length must be 3 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if _, ok := Interval_name[int32(m.GetInterval())]; !ok {
		err := StrategyCompanyRequestValidationError{
			field:  "Interval",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := StrategyType_name[int32(m.GetStrategy())]; !ok {
		err := StrategyCompanyRequestValidationError{
			field:  "Strategy",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StrategyCompanyRequestMultiError(errors)
	}

	return nil
}

// StrategyCompanyRequestMultiError is an error wrapping multiple validation
// errors returned by StrategyCompanyRequest.ValidateAll() if the designated
// constraints aren't met.
type StrategyCompanyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyCompanyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyCompanyRequestMultiError) AllErrors() []error { return m }

// StrategyCompanyRequestValidationError is the validation error returned by
// StrategyCompanyRequest.Validate if the designated constraints aren't met.
type StrategyCompanyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyCompanyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyCompanyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyCompanyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyCompanyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyCompanyRequestValidationError) ErrorName() string {
	return "StrategyCompanyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StrategyCompanyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategyCompanyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyCompanyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyCompanyRequestValidationError{}

// Validate checks the field values on StrategyCurrencyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StrategyCurrencyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrategyCurrencyRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StrategyCurrencyRequestMultiError, or nil if none found.
func (m *StrategyCurrencyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StrategyCurrencyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetTicker()); l < 1 || l > 8 {
		err := StrategyCurrencyRequestValidationError{
			field:  "Ticker",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := Interval_name[int32(m.GetInterval())]; !ok {
		err := StrategyCurrencyRequestValidationError{
			field:  "Interval",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := StrategyType_name[int32(m.GetStrategy())]; !ok {
		err := StrategyCurrencyRequestValidationError{
			field:  "Strategy",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StrategyCurrencyRequestMultiError(errors)
	}

	return nil
}

// StrategyCurrencyRequestMultiError is an error wrapping multiple validation
// errors returned by StrategyCurrencyRequest.ValidateAll() if the designated
// constraints aren't met.
type StrategyCurrencyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyCurrencyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyCurrencyRequestMultiError) AllErrors() []error { return m }

// StrategyCurrencyRequestValidationError is the validation error returned by
// StrategyCurrencyRequest.Validate if the designated constraints aren't met.
type StrategyCurrencyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyCurrencyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyCurrencyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyCurrencyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyCurrencyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyCurrencyRequestValidationError) ErrorName() string {
	return "StrategyCurrencyRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StrategyCurrencyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategyCurrencyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyCurrencyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyCurrencyRequestValidationError{}

// Validate checks the field values on StrategyIndustryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StrategyIndustryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrategyIndustryRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StrategyIndustryRequestMultiError, or nil if none found.
func (m *StrategyIndustryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StrategyIndustryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetIndustry(); val < 5010101010 || val > 6310301010 {
		err := StrategyIndustryRequestValidationError{
			field:  "Industry",
			reason: "value must be inside range [5010101010, 6310301010]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetExchange()); l < 1 || l > 8 {
		err := StrategyIndustryRequestValidationError{
			field:  "Exchange",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := Interval_name[int32(m.GetInterval())]; !ok {
		err := StrategyIndustryRequestValidationError{
			field:  "Interval",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := StrategyType_name[int32(m.GetStrategy())]; !ok {
		err := StrategyIndustryRequestValidationError{
			field:  "Strategy",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StrategyIndustryRequestMultiError(errors)
	}

	return nil
}

// StrategyIndustryRequestMultiError is an error wrapping multiple validation
// errors returned by StrategyIndustryRequest.ValidateAll() if the designated
// constraints aren't met.
type StrategyIndustryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyIndustryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyIndustryRequestMultiError) AllErrors() []error { return m }

// StrategyIndustryRequestValidationError is the validation error returned by
// StrategyIndustryRequest.Validate if the designated constraints aren't met.
type StrategyIndustryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyIndustryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyIndustryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyIndustryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyIndustryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyIndustryRequestValidationError) ErrorName() string {
	return "StrategyIndustryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StrategyIndustryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategyIndustryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyIndustryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyIndustryRequestValidationError{}

// Validate checks the field values on StrategyExchangeRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StrategyExchangeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrategyExchangeRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StrategyExchangeRequestMultiError, or nil if none found.
func (m *StrategyExchangeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StrategyExchangeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetExchange()); l < 1 || l > 8 {
		err := StrategyExchangeRequestValidationError{
			field:  "Exchange",
			reason: "value length must be between 1 and 8 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCountry()) != 2 {
		err := StrategyExchangeRequestValidationError{
			field:  "Country",
			reason: "value length must be 2 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if _, ok := Interval_name[int32(m.GetInterval())]; !ok {
		err := StrategyExchangeRequestValidationError{
			field:  "Interval",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := StrategyType_name[int32(m.GetStrategy())]; !ok {
		err := StrategyExchangeRequestValidationError{
			field:  "Strategy",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StrategyExchangeRequestMultiError(errors)
	}

	return nil
}

// StrategyExchangeRequestMultiError is an error wrapping multiple validation
// errors returned by StrategyExchangeRequest.ValidateAll() if the designated
// constraints aren't met.
type StrategyExchangeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyExchangeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyExchangeRequestMultiError) AllErrors() []error { return m }

// StrategyExchangeRequestValidationError is the validation error returned by
// StrategyExchangeRequest.Validate if the designated constraints aren't met.
type StrategyExchangeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyExchangeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyExchangeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyExchangeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyExchangeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyExchangeRequestValidationError) ErrorName() string {
	return "StrategyExchangeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StrategyExchangeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategyExchangeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyExchangeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyExchangeRequestValidationError{}

// Validate checks the field values on StrategyCountryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StrategyCountryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrategyCountryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StrategyCountryRequestMultiError, or nil if none found.
func (m *StrategyCountryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StrategyCountryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetCountry()) != 2 {
		err := StrategyCountryRequestValidationError{
			field:  "Country",
			reason: "value length must be 2 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if _, ok := Interval_name[int32(m.GetInterval())]; !ok {
		err := StrategyCountryRequestValidationError{
			field:  "Interval",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := StrategyType_name[int32(m.GetStrategy())]; !ok {
		err := StrategyCountryRequestValidationError{
			field:  "Strategy",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StrategyCountryRequestMultiError(errors)
	}

	return nil
}

// StrategyCountryRequestMultiError is an error wrapping multiple validation
// errors returned by StrategyCountryRequest.ValidateAll() if the designated
// constraints aren't met.
type StrategyCountryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyCountryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyCountryRequestMultiError) AllErrors() []error { return m }

// StrategyCountryRequestValidationError is the validation error returned by
// StrategyCountryRequest.Validate if the designated constraints aren't met.
type StrategyCountryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyCountryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyCountryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyCountryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyCountryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyCountryRequestValidationError) ErrorName() string {
	return "StrategyCountryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StrategyCountryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategyCountryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyCountryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyCountryRequestValidationError{}

// Validate checks the field values on StrategyIndexRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StrategyIndexRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrategyIndexRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StrategyIndexRequestMultiError, or nil if none found.
func (m *StrategyIndexRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StrategyIndexRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetIndex()); l < 3 || l > 32 {
		err := StrategyIndexRequestValidationError{
			field:  "Index",
			reason: "value length must be between 3 and 32 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := Interval_name[int32(m.GetInterval())]; !ok {
		err := StrategyIndexRequestValidationError{
			field:  "Interval",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := StrategyType_name[int32(m.GetStrategy())]; !ok {
		err := StrategyIndexRequestValidationError{
			field:  "Strategy",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StrategyIndexRequestMultiError(errors)
	}

	return nil
}

// StrategyIndexRequestMultiError is an error wrapping multiple validation
// errors returned by StrategyIndexRequest.ValidateAll() if the designated
// constraints aren't met.
type StrategyIndexRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyIndexRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyIndexRequestMultiError) AllErrors() []error { return m }

// StrategyIndexRequestValidationError is the validation error returned by
// StrategyIndexRequest.Validate if the designated constraints aren't met.
type StrategyIndexRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyIndexRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyIndexRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyIndexRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyIndexRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyIndexRequestValidationError) ErrorName() string {
	return "StrategyIndexRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StrategyIndexRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategyIndexRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyIndexRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyIndexRequestValidationError{}

// Validate checks the field values on StrategyAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *StrategyAccountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrategyAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StrategyAccountRequestMultiError, or nil if none found.
func (m *StrategyAccountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StrategyAccountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetAccount()); l < 3 || l > 32 {
		err := StrategyAccountRequestValidationError{
			field:  "Account",
			reason: "value length must be between 3 and 32 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := Interval_name[int32(m.GetInterval())]; !ok {
		err := StrategyAccountRequestValidationError{
			field:  "Interval",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := StrategyType_name[int32(m.GetStrategy())]; !ok {
		err := StrategyAccountRequestValidationError{
			field:  "Strategy",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StrategyAccountRequestMultiError(errors)
	}

	return nil
}

// StrategyAccountRequestMultiError is an error wrapping multiple validation
// errors returned by StrategyAccountRequest.ValidateAll() if the designated
// constraints aren't met.
type StrategyAccountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyAccountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyAccountRequestMultiError) AllErrors() []error { return m }

// StrategyAccountRequestValidationError is the validation error returned by
// StrategyAccountRequest.Validate if the designated constraints aren't met.
type StrategyAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyAccountRequestValidationError) ErrorName() string {
	return "StrategyAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e StrategyAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategyAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyAccountRequestValidationError{}

// Validate checks the field values on StrategyReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StrategyReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StrategyReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StrategyReplyMultiError, or
// nil if none found.
func (m *StrategyReply) ValidateAll() error {
	return m.validate(true)
}

func (m *StrategyReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Strategy

	// no validation rules for Interval

	// no validation rules for Status

	// no validation rules for StatusAgo

	// no validation rules for Date

	if len(errors) > 0 {
		return StrategyReplyMultiError(errors)
	}

	return nil
}

// StrategyReplyMultiError is an error wrapping multiple validation errors
// returned by StrategyReply.ValidateAll() if the designated constraints
// aren't met.
type StrategyReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StrategyReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StrategyReplyMultiError) AllErrors() []error { return m }

// StrategyReplyValidationError is the validation error returned by
// StrategyReply.Validate if the designated constraints aren't met.
type StrategyReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StrategyReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StrategyReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StrategyReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StrategyReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StrategyReplyValidationError) ErrorName() string { return "StrategyReplyValidationError" }

// Error satisfies the builtin error interface
func (e StrategyReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStrategyReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StrategyReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StrategyReplyValidationError{}

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: shorted/collector/v1/collector_service.proto

package collector

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _collector_service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetSourceRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetSourceRequest) Validate() error {
	if m == nil {
		return nil
	}

	if _, err := url.Parse(m.GetUrl()); err != nil {
		return GetSourceRequestValidationError{
			field:  "Url",
			reason: "value must be a valid URI",
			cause:  err,
		}
	}

	// no validation rules for Format

	// no validation rules for Parser

	return nil
}

// GetSourceRequestValidationError is the validation error returned by
// GetSourceRequest.Validate if the designated constraints aren't met.
type GetSourceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSourceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSourceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSourceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSourceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSourceRequestValidationError) ErrorName() string { return "GetSourceRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetSourceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSourceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSourceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSourceRequestValidationError{}

// Validate checks the field values on GetSourceResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetSourceResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetSourceResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetSourceResponseValidationError is the validation error returned by
// GetSourceResponse.Validate if the designated constraints aren't met.
type GetSourceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetSourceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetSourceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetSourceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetSourceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetSourceResponseValidationError) ErrorName() string {
	return "GetSourceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetSourceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetSourceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetSourceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetSourceResponseValidationError{}

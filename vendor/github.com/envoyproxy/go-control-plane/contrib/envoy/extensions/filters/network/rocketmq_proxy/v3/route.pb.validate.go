// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: contrib/envoy/extensions/filters/network/rocketmq_proxy/v3/route.proto

package envoy_extensions_filters_network_rocketmq_proxy_v3

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

// Validate checks the field values on RouteConfiguration with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RouteConfiguration) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RouteConfiguration with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RouteConfigurationMultiError, or nil if none found.
func (m *RouteConfiguration) ValidateAll() error {
	return m.validate(true)
}

func (m *RouteConfiguration) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	for idx, item := range m.GetRoutes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RouteConfigurationValidationError{
						field:  fmt.Sprintf("Routes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RouteConfigurationValidationError{
						field:  fmt.Sprintf("Routes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteConfigurationValidationError{
					field:  fmt.Sprintf("Routes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RouteConfigurationMultiError(errors)
	}
	return nil
}

// RouteConfigurationMultiError is an error wrapping multiple validation errors
// returned by RouteConfiguration.ValidateAll() if the designated constraints
// aren't met.
type RouteConfigurationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RouteConfigurationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RouteConfigurationMultiError) AllErrors() []error { return m }

// RouteConfigurationValidationError is the validation error returned by
// RouteConfiguration.Validate if the designated constraints aren't met.
type RouteConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteConfigurationValidationError) ErrorName() string {
	return "RouteConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e RouteConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteConfigurationValidationError{}

// Validate checks the field values on Route with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Route) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Route with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RouteMultiError, or nil if none found.
func (m *Route) ValidateAll() error {
	return m.validate(true)
}

func (m *Route) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetMatch() == nil {
		err := RouteValidationError{
			field:  "Match",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetMatch()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RouteValidationError{
					field:  "Match",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RouteValidationError{
					field:  "Match",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteValidationError{
				field:  "Match",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetRoute() == nil {
		err := RouteValidationError{
			field:  "Route",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRoute()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RouteValidationError{
					field:  "Route",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RouteValidationError{
					field:  "Route",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRoute()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteValidationError{
				field:  "Route",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RouteMultiError(errors)
	}
	return nil
}

// RouteMultiError is an error wrapping multiple validation errors returned by
// Route.ValidateAll() if the designated constraints aren't met.
type RouteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RouteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RouteMultiError) AllErrors() []error { return m }

// RouteValidationError is the validation error returned by Route.Validate if
// the designated constraints aren't met.
type RouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteValidationError) ErrorName() string { return "RouteValidationError" }

// Error satisfies the builtin error interface
func (e RouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteValidationError{}

// Validate checks the field values on RouteMatch with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RouteMatch) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RouteMatch with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RouteMatchMultiError, or
// nil if none found.
func (m *RouteMatch) ValidateAll() error {
	return m.validate(true)
}

func (m *RouteMatch) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTopic() == nil {
		err := RouteMatchValidationError{
			field:  "Topic",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTopic()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RouteMatchValidationError{
					field:  "Topic",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RouteMatchValidationError{
					field:  "Topic",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTopic()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteMatchValidationError{
				field:  "Topic",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RouteMatchValidationError{
						field:  fmt.Sprintf("Headers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RouteMatchValidationError{
						field:  fmt.Sprintf("Headers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RouteMatchValidationError{
					field:  fmt.Sprintf("Headers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RouteMatchMultiError(errors)
	}
	return nil
}

// RouteMatchMultiError is an error wrapping multiple validation errors
// returned by RouteMatch.ValidateAll() if the designated constraints aren't met.
type RouteMatchMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RouteMatchMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RouteMatchMultiError) AllErrors() []error { return m }

// RouteMatchValidationError is the validation error returned by
// RouteMatch.Validate if the designated constraints aren't met.
type RouteMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteMatchValidationError) ErrorName() string { return "RouteMatchValidationError" }

// Error satisfies the builtin error interface
func (e RouteMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteMatchValidationError{}

// Validate checks the field values on RouteAction with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RouteAction) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RouteAction with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RouteActionMultiError, or
// nil if none found.
func (m *RouteAction) ValidateAll() error {
	return m.validate(true)
}

func (m *RouteAction) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetCluster()) < 1 {
		err := RouteActionValidationError{
			field:  "Cluster",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetMetadataMatch()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, RouteActionValidationError{
					field:  "MetadataMatch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RouteActionValidationError{
					field:  "MetadataMatch",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadataMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RouteActionValidationError{
				field:  "MetadataMatch",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return RouteActionMultiError(errors)
	}
	return nil
}

// RouteActionMultiError is an error wrapping multiple validation errors
// returned by RouteAction.ValidateAll() if the designated constraints aren't met.
type RouteActionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RouteActionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RouteActionMultiError) AllErrors() []error { return m }

// RouteActionValidationError is the validation error returned by
// RouteAction.Validate if the designated constraints aren't met.
type RouteActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteActionValidationError) ErrorName() string { return "RouteActionValidationError" }

// Error satisfies the builtin error interface
func (e RouteActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteActionValidationError{}
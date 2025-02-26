// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/ext_authz/v3/ext_authz.proto

package envoy_extensions_filters_http_ext_authz_v3

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

	v3 "github.com/datawire/ambassador/v2/pkg/api/envoy/config/core/v3"
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

	_ = v3.ApiVersion(0)
)

// Validate checks the field values on ExtAuthz with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *ExtAuthz) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := v3.ApiVersion_name[int32(m.GetTransportApiVersion())]; !ok {
		return ExtAuthzValidationError{
			field:  "TransportApiVersion",
			reason: "value must be one of the defined enum values",
		}
	}

	// no validation rules for FailureModeAllow

	if v, ok := interface{}(m.GetWithRequestBody()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtAuthzValidationError{
				field:  "WithRequestBody",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ClearRouteCache

	if v, ok := interface{}(m.GetStatusOnError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtAuthzValidationError{
				field:  "StatusOnError",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFilterEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtAuthzValidationError{
				field:  "FilterEnabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFilterEnabledMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtAuthzValidationError{
				field:  "FilterEnabledMetadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDenyAtDisable()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtAuthzValidationError{
				field:  "DenyAtDisable",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IncludePeerCertificate

	// no validation rules for StatPrefix

	// no validation rules for HiddenEnvoyDeprecatedUseAlpha

	switch m.Services.(type) {

	case *ExtAuthz_GrpcService:

		if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExtAuthzValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ExtAuthz_HttpService:

		if v, ok := interface{}(m.GetHttpService()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExtAuthzValidationError{
					field:  "HttpService",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ExtAuthzValidationError is the validation error returned by
// ExtAuthz.Validate if the designated constraints aren't met.
type ExtAuthzValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtAuthzValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtAuthzValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtAuthzValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtAuthzValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtAuthzValidationError) ErrorName() string { return "ExtAuthzValidationError" }

// Error satisfies the builtin error interface
func (e ExtAuthzValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtAuthz.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtAuthzValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtAuthzValidationError{}

// Validate checks the field values on BufferSettings with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *BufferSettings) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetMaxRequestBytes() <= 0 {
		return BufferSettingsValidationError{
			field:  "MaxRequestBytes",
			reason: "value must be greater than 0",
		}
	}

	// no validation rules for AllowPartialMessage

	// no validation rules for PackAsBytes

	return nil
}

// BufferSettingsValidationError is the validation error returned by
// BufferSettings.Validate if the designated constraints aren't met.
type BufferSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BufferSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BufferSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BufferSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BufferSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BufferSettingsValidationError) ErrorName() string { return "BufferSettingsValidationError" }

// Error satisfies the builtin error interface
func (e BufferSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBufferSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BufferSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BufferSettingsValidationError{}

// Validate checks the field values on HttpService with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HttpService) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetServerUri()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpServiceValidationError{
				field:  "ServerUri",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for PathPrefix

	if v, ok := interface{}(m.GetAuthorizationRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpServiceValidationError{
				field:  "AuthorizationRequest",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAuthorizationResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HttpServiceValidationError{
				field:  "AuthorizationResponse",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HttpServiceValidationError is the validation error returned by
// HttpService.Validate if the designated constraints aren't met.
type HttpServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpServiceValidationError) ErrorName() string { return "HttpServiceValidationError" }

// Error satisfies the builtin error interface
func (e HttpServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpService.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpServiceValidationError{}

// Validate checks the field values on AuthorizationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AuthorizationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetAllowedHeaders()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuthorizationRequestValidationError{
				field:  "AllowedHeaders",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetHeadersToAdd() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AuthorizationRequestValidationError{
					field:  fmt.Sprintf("HeadersToAdd[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// AuthorizationRequestValidationError is the validation error returned by
// AuthorizationRequest.Validate if the designated constraints aren't met.
type AuthorizationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthorizationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthorizationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthorizationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthorizationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthorizationRequestValidationError) ErrorName() string {
	return "AuthorizationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AuthorizationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthorizationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthorizationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthorizationRequestValidationError{}

// Validate checks the field values on AuthorizationResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AuthorizationResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetAllowedUpstreamHeaders()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuthorizationResponseValidationError{
				field:  "AllowedUpstreamHeaders",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAllowedUpstreamHeadersToAppend()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuthorizationResponseValidationError{
				field:  "AllowedUpstreamHeadersToAppend",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAllowedClientHeaders()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AuthorizationResponseValidationError{
				field:  "AllowedClientHeaders",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// AuthorizationResponseValidationError is the validation error returned by
// AuthorizationResponse.Validate if the designated constraints aren't met.
type AuthorizationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthorizationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthorizationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthorizationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthorizationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthorizationResponseValidationError) ErrorName() string {
	return "AuthorizationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthorizationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthorizationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthorizationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthorizationResponseValidationError{}

// Validate checks the field values on ExtAuthzPerRoute with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ExtAuthzPerRoute) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Override.(type) {

	case *ExtAuthzPerRoute_Disabled:

		if m.GetDisabled() != true {
			return ExtAuthzPerRouteValidationError{
				field:  "Disabled",
				reason: "value must equal true",
			}
		}

	case *ExtAuthzPerRoute_CheckSettings:

		if m.GetCheckSettings() == nil {
			return ExtAuthzPerRouteValidationError{
				field:  "CheckSettings",
				reason: "value is required",
			}
		}

		if v, ok := interface{}(m.GetCheckSettings()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ExtAuthzPerRouteValidationError{
					field:  "CheckSettings",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return ExtAuthzPerRouteValidationError{
			field:  "Override",
			reason: "value is required",
		}

	}

	return nil
}

// ExtAuthzPerRouteValidationError is the validation error returned by
// ExtAuthzPerRoute.Validate if the designated constraints aren't met.
type ExtAuthzPerRouteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtAuthzPerRouteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtAuthzPerRouteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtAuthzPerRouteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtAuthzPerRouteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtAuthzPerRouteValidationError) ErrorName() string { return "ExtAuthzPerRouteValidationError" }

// Error satisfies the builtin error interface
func (e ExtAuthzPerRouteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtAuthzPerRoute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtAuthzPerRouteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtAuthzPerRouteValidationError{}

// Validate checks the field values on CheckSettings with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *CheckSettings) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ContextExtensions

	// no validation rules for DisableRequestBodyBuffering

	return nil
}

// CheckSettingsValidationError is the validation error returned by
// CheckSettings.Validate if the designated constraints aren't met.
type CheckSettingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckSettingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckSettingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckSettingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckSettingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckSettingsValidationError) ErrorName() string { return "CheckSettingsValidationError" }

// Error satisfies the builtin error interface
func (e CheckSettingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckSettings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckSettingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckSettingsValidationError{}

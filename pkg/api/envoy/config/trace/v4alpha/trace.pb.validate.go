// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/trace/v4alpha/trace.proto

package envoy_config_trace_v4alpha

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _trace_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Tracing with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Tracing) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetHttp()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return TracingValidationError{
					field:  "Http",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// TracingValidationError is the validation error returned by Tracing.Validate
// if the designated constraints aren't met.
type TracingValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TracingValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TracingValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TracingValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TracingValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TracingValidationError) ErrorName() string { return "TracingValidationError" }

// Error satisfies the builtin error interface
func (e TracingValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTracing.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TracingValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TracingValidationError{}

// Validate checks the field values on LightstepConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *LightstepConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetCollectorCluster()) < 1 {
		return LightstepConfigValidationError{
			field:  "CollectorCluster",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetAccessTokenFile()) < 1 {
		return LightstepConfigValidationError{
			field:  "AccessTokenFile",
			reason: "value length must be at least 1 bytes",
		}
	}

	for idx, item := range m.GetPropagationModes() {
		_, _ = idx, item

		if _, ok := LightstepConfig_PropagationMode_name[int32(item)]; !ok {
			return LightstepConfigValidationError{
				field:  fmt.Sprintf("PropagationModes[%v]", idx),
				reason: "value must be one of the defined enum values",
			}
		}

	}

	return nil
}

// LightstepConfigValidationError is the validation error returned by
// LightstepConfig.Validate if the designated constraints aren't met.
type LightstepConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LightstepConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LightstepConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LightstepConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LightstepConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LightstepConfigValidationError) ErrorName() string { return "LightstepConfigValidationError" }

// Error satisfies the builtin error interface
func (e LightstepConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLightstepConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LightstepConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LightstepConfigValidationError{}

// Validate checks the field values on ZipkinConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ZipkinConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetCollectorCluster()) < 1 {
		return ZipkinConfigValidationError{
			field:  "CollectorCluster",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetCollectorEndpoint()) < 1 {
		return ZipkinConfigValidationError{
			field:  "CollectorEndpoint",
			reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for TraceId_128Bit

	{
		tmp := m.GetSharedSpanContext()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ZipkinConfigValidationError{
					field:  "SharedSpanContext",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for CollectorEndpointVersion

	return nil
}

// ZipkinConfigValidationError is the validation error returned by
// ZipkinConfig.Validate if the designated constraints aren't met.
type ZipkinConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ZipkinConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ZipkinConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ZipkinConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ZipkinConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ZipkinConfigValidationError) ErrorName() string { return "ZipkinConfigValidationError" }

// Error satisfies the builtin error interface
func (e ZipkinConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sZipkinConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ZipkinConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ZipkinConfigValidationError{}

// Validate checks the field values on DynamicOtConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DynamicOtConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetLibrary()) < 1 {
		return DynamicOtConfigValidationError{
			field:  "Library",
			reason: "value length must be at least 1 bytes",
		}
	}

	{
		tmp := m.GetConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return DynamicOtConfigValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// DynamicOtConfigValidationError is the validation error returned by
// DynamicOtConfig.Validate if the designated constraints aren't met.
type DynamicOtConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DynamicOtConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DynamicOtConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DynamicOtConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DynamicOtConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DynamicOtConfigValidationError) ErrorName() string { return "DynamicOtConfigValidationError" }

// Error satisfies the builtin error interface
func (e DynamicOtConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDynamicOtConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DynamicOtConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DynamicOtConfigValidationError{}

// Validate checks the field values on DatadogConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DatadogConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetCollectorCluster()) < 1 {
		return DatadogConfigValidationError{
			field:  "CollectorCluster",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetServiceName()) < 1 {
		return DatadogConfigValidationError{
			field:  "ServiceName",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// DatadogConfigValidationError is the validation error returned by
// DatadogConfig.Validate if the designated constraints aren't met.
type DatadogConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DatadogConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DatadogConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DatadogConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DatadogConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DatadogConfigValidationError) ErrorName() string { return "DatadogConfigValidationError" }

// Error satisfies the builtin error interface
func (e DatadogConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDatadogConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DatadogConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DatadogConfigValidationError{}

// Validate checks the field values on OpenCensusConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OpenCensusConfig) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetTraceConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return OpenCensusConfigValidationError{
					field:  "TraceConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for StdoutExporterEnabled

	// no validation rules for StackdriverExporterEnabled

	// no validation rules for StackdriverProjectId

	// no validation rules for StackdriverAddress

	{
		tmp := m.GetStackdriverGrpcService()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return OpenCensusConfigValidationError{
					field:  "StackdriverGrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for ZipkinExporterEnabled

	// no validation rules for ZipkinUrl

	// no validation rules for OcagentExporterEnabled

	// no validation rules for OcagentAddress

	{
		tmp := m.GetOcagentGrpcService()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return OpenCensusConfigValidationError{
					field:  "OcagentGrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// OpenCensusConfigValidationError is the validation error returned by
// OpenCensusConfig.Validate if the designated constraints aren't met.
type OpenCensusConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OpenCensusConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OpenCensusConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OpenCensusConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OpenCensusConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OpenCensusConfigValidationError) ErrorName() string { return "OpenCensusConfigValidationError" }

// Error satisfies the builtin error interface
func (e OpenCensusConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOpenCensusConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OpenCensusConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OpenCensusConfigValidationError{}

// Validate checks the field values on TraceServiceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *TraceServiceConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetGrpcService() == nil {
		return TraceServiceConfigValidationError{
			field:  "GrpcService",
			reason: "value is required",
		}
	}

	{
		tmp := m.GetGrpcService()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return TraceServiceConfigValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// TraceServiceConfigValidationError is the validation error returned by
// TraceServiceConfig.Validate if the designated constraints aren't met.
type TraceServiceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TraceServiceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TraceServiceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TraceServiceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TraceServiceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TraceServiceConfigValidationError) ErrorName() string {
	return "TraceServiceConfigValidationError"
}

// Error satisfies the builtin error interface
func (e TraceServiceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTraceServiceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TraceServiceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TraceServiceConfigValidationError{}

// Validate checks the field values on Tracing_Http with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Tracing_Http) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return Tracing_HttpValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	switch m.ConfigType.(type) {

	case *Tracing_Http_TypedConfig:

		{
			tmp := m.GetTypedConfig()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return Tracing_HttpValidationError{
						field:  "TypedConfig",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// Tracing_HttpValidationError is the validation error returned by
// Tracing_Http.Validate if the designated constraints aren't met.
type Tracing_HttpValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Tracing_HttpValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Tracing_HttpValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Tracing_HttpValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Tracing_HttpValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Tracing_HttpValidationError) ErrorName() string { return "Tracing_HttpValidationError" }

// Error satisfies the builtin error interface
func (e Tracing_HttpValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTracing_Http.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Tracing_HttpValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Tracing_HttpValidationError{}

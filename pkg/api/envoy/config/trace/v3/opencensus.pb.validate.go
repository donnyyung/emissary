// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/trace/v3/opencensus.proto

package envoy_config_trace_v3

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
var _opencensus_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on OpenCensusConfig with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *OpenCensusConfig) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetTraceConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenCensusConfigValidationError{
				field:  "TraceConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for StdoutExporterEnabled

	// no validation rules for StackdriverExporterEnabled

	// no validation rules for StackdriverProjectId

	// no validation rules for StackdriverAddress

	if v, ok := interface{}(m.GetStackdriverGrpcService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenCensusConfigValidationError{
				field:  "StackdriverGrpcService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ZipkinExporterEnabled

	// no validation rules for ZipkinUrl

	// no validation rules for OcagentExporterEnabled

	// no validation rules for OcagentAddress

	if v, ok := interface{}(m.GetOcagentGrpcService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OpenCensusConfigValidationError{
				field:  "OcagentGrpcService",
				reason: "embedded message failed validation",
				cause:  err,
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

// Copyright 2020 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: bufbuild/buf/image/v1/image.proto

package imagev1

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
var _image_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Image with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Image) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetFile() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ImageValidationError{
					field:  fmt.Sprintf("File[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetBufbuildImageExtension()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImageValidationError{
				field:  "BufbuildImageExtension",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ImageValidationError is the validation error returned by Image.Validate if
// the designated constraints aren't met.
type ImageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageValidationError) ErrorName() string { return "ImageValidationError" }

// Error satisfies the builtin error interface
func (e ImageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageValidationError{}

// Validate checks the field values on ImageExtension with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ImageExtension) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetImageImportRefs() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ImageExtensionValidationError{
					field:  fmt.Sprintf("ImageImportRefs[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ImageExtensionValidationError is the validation error returned by
// ImageExtension.Validate if the designated constraints aren't met.
type ImageExtensionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageExtensionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageExtensionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageExtensionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageExtensionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageExtensionValidationError) ErrorName() string { return "ImageExtensionValidationError" }

// Error satisfies the builtin error interface
func (e ImageExtensionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImageExtension.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageExtensionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageExtensionValidationError{}

// Validate checks the field values on ImageImportRef with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ImageImportRef) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for FileIndex

	return nil
}

// ImageImportRefValidationError is the validation error returned by
// ImageImportRef.Validate if the designated constraints aren't met.
type ImageImportRefValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImageImportRefValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImageImportRefValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImageImportRefValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImageImportRefValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImageImportRefValidationError) ErrorName() string { return "ImageImportRefValidationError" }

// Error satisfies the builtin error interface
func (e ImageImportRefValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImageImportRef.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImageImportRefValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImageImportRefValidationError{}

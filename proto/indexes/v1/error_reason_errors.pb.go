// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsContentMissing(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CONTENT_MISSING.String() && e.Code == 400
}

func ErrorContentMissing(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_CONTENT_MISSING.String(), fmt.Sprintf(format, args...))
}

func IsFORBIDDEN(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_FORBIDDEN.String() && e.Code == 403
}

func ErrorFORBIDDEN(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ErrorReason_FORBIDDEN.String(), fmt.Sprintf(format, args...))
}

func IsNotFound(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NOT_FOUND.String() && e.Code == 404
}

func ErrorNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsCONFLICT(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CONFLICT.String() && e.Code == 409
}

func ErrorCONFLICT(format string, args ...interface{}) *errors.Error {
	return errors.New(409, ErrorReason_CONFLICT.String(), fmt.Sprintf(format, args...))
}

func IsNotImplemented(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NOT_IMPLEMENTED.String() && e.Code == 501
}

func ErrorNotImplemented(format string, args ...interface{}) *errors.Error {
	return errors.New(501, ErrorReason_NOT_IMPLEMENTED.String(), fmt.Sprintf(format, args...))
}

func IsUNAVAILABLE(err error) bool {
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UNAVAILABLE.String() && e.Code == 503
}

func ErrorUNAVAILABLE(format string, args ...interface{}) *errors.Error {
	return errors.New(503, ErrorReason_UNAVAILABLE.String(), fmt.Sprintf(format, args...))
}
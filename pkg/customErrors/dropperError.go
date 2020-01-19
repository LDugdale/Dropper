package customErrors

import (
	"fmt"

	"github.com/pkg/errors"
)

// ErrorType is the type of an error
type ErrorType uint

const (
	OK                 ErrorType = 0
	Canceled           ErrorType = 1
	Unknown            ErrorType = 2
	InvalidArgument    ErrorType = 3
	DeadlineExceeded   ErrorType = 4
	NotFound           ErrorType = 5
	AlreadyExists      ErrorType = 6
	PermissionDenied   ErrorType = 7
	ResourceExhausted  ErrorType = 8
	FailedPrecondition ErrorType = 9
	Aborted            ErrorType = 10
	OutOfRange         ErrorType = 11
	Unimplemented      ErrorType = 12
	Internal           ErrorType = 13
	Unavailable        ErrorType = 14
	DataLoss           ErrorType = 15
	Unauthenticated    ErrorType = 16
)

type DropperError struct {
	errorType     ErrorType
	originalError error
	context       errorContext
	details       errorDetails
}

type errorDetails struct {
	Message     string
	Description string
}

type errorContext struct {
	Field   string
	Message string
}

// New creates a new dropperError
func (errorType ErrorType) New(msg string) error {
	return DropperError{errorType: errorType, originalError: errors.New(msg)}
}

func (errorType ErrorType) NewWithDetails(msg string, description string) error {
	details := errorDetails{Message: msg, Description: description}
	return DropperError{errorType: errorType, originalError: errors.New(msg), details: details}
}

// New creates a new dropperError with formatted message
func (errorType ErrorType) Newf(msg string, args ...interface{}) error {
	return DropperError{errorType: errorType, originalError: fmt.Errorf(msg, args...)}
}

// Wrap creates a new wrapped error
func (errorType ErrorType) Wrap(err error, msg string) error {
	return errorType.Wrapf(err, msg)
}

func (errorType ErrorType) WrapWithDetails(err error, description string, msg string) error {
	return errorType.WrapfWithDetails(err, description, msg)
}

// Wrap creates a new wrapped error with formatted message
func (errorType ErrorType) Wrapf(err error, msg string, args ...interface{}) error {
	return DropperError{errorType: errorType, originalError: errors.Wrapf(err, msg, args...)}
}

func (errorType ErrorType) WrapfWithDetails(err error, description string, msg string, args ...interface{}) error {
	details := errorDetails{Message: msg, Description: description}
	return DropperError{errorType: errorType, details: details, originalError: errors.Wrapf(err, msg, args...)}

}

// Error returns the mssage of a dropperError
func (error DropperError) Error() string {
	return error.originalError.Error()
}

// New creates a no type error
func New(msg string) error {
	return DropperError{errorType: Unknown, originalError: errors.New(msg)}
}

// Newf creates a no type error with formatted message
func Newf(msg string, args ...interface{}) error {
	return DropperError{errorType: Unknown, originalError: errors.New(fmt.Sprintf(msg, args...))}
}

// Wrap an error with a string
func Wrap(err error, msg string) error {
	return Wrapf(err, msg)
}

// Cause gives the original error
func Cause(err error) error {
	return errors.Cause(err)
}

// Wrapf an error with format string
func Wrapf(err error, msg string, args ...interface{}) error {
	wrappedError := errors.Wrapf(err, msg, args...)
	if customErr, ok := err.(DropperError); ok {
		return DropperError{
			errorType:     customErr.errorType,
			originalError: wrappedError,
			context:       customErr.context,
		}
	}

	return DropperError{errorType: Unknown, originalError: wrappedError}
}

// AddErrorContext adds a context to an error
func AddErrorContext(err error, field, message string) error {
	context := errorContext{Field: field, Message: message}
	if customErr, ok := err.(DropperError); ok {
		return DropperError{errorType: customErr.errorType, originalError: customErr.originalError, context: context}
	}

	return DropperError{errorType: Unknown, originalError: err, context: context}
}

// GetErrorContext returns the error context
func GetErrorContext(err error) map[string]string {
	emptyContext := errorContext{}
	if customErr, ok := err.(DropperError); ok || customErr.context != emptyContext {

		return map[string]string{"field": customErr.context.Field, "message": customErr.context.Message}
	}

	return nil
}

// GetErrorContext returns the error details
func GetErrorDetails(err error) *errorDetails {
	emptyDetails := errorDetails{}
	if customErr, ok := err.(DropperError); ok || customErr.details != emptyDetails {
		return &customErr.details
	}

	return nil
}

// GetType returns the error type
func GetType(err error) ErrorType {
	if customErr, ok := err.(DropperError); ok {
		return customErr.errorType
	}

	return Unknown
}

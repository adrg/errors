package errors

import (
	"fmt"
	"reflect"
)

type wrapper interface {
	Unwrap() error
}

type withCode interface {
	Code() int
}

// New returns a new error with the specified message.
func New(message string) error {
	return newPrimitive(nil, message)
}

// Errorf returns a new error formatted according to the format specifier.
func Errorf(format string, args ...interface{}) error {
	return newPrimitive(nil, fmt.Sprintf(format, args...))
}

// Annotate annotates the provided error with the specified message.
// Returns nil if the err argument is nil.
func Annotate(err error, message string) error {
	if err == nil {
		return nil
	}

	return newPrimitive(err, message)
}

// Annotatef annotates the provided error according to the format specifier.
// Returns nil if the err argument is nil.
func Annotatef(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return newPrimitive(err, fmt.Sprintf(format, args...))
}

// Unwrap returns the next error in the error chain.
// If there is no next error, Unwrap returns nil.
func Unwrap(err error) error {
	if err != nil {
		if p, ok := err.(wrapper); ok {
			return p.Unwrap()
		}
	}

	return nil
}

// Is reports whether err or any of the errors in its chain is equal to target.
func Is(err, target error) bool {
	if target == nil {
		return err == target
	}

	for {
		if err == target {
			return true
		}
		if err = Unwrap(err); err == nil {
			return false
		}
	}
}

// As checks whether err or any of the errors in its chain is a value of the
// same type as target. If so, it sets target with the found value and returns
// true. Otherwise, target is left unchanged and false is returned.
func As(err error, target interface{}) bool {
	if err == nil || target == nil {
		return false
	}

	targetType := reflect.TypeOf(target)
	if targetType.Kind() != reflect.Ptr {
		return false
	}

	targetType = targetType.Elem()
	for {
		if reflect.TypeOf(err) == targetType {
			reflect.ValueOf(target).Elem().Set(reflect.ValueOf(err))
			return true
		}
		if err = Unwrap(err); err == nil {
			return false
		}
	}
}

// Code returns the code of the provided error.
// Returns 0 if the error has no error code.
func Code(err error) int {
	if err != nil {
		if h, ok := err.(withCode); ok {
			return h.Code()
		}
	}

	return 0
}

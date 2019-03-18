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

func New(message string) error {
	return newPrimitive(nil, message)
}

func Errorf(format string, args ...interface{}) error {
	return newPrimitive(nil, fmt.Sprintf(format, args...))
}

func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}

	return newPrimitive(err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}

	return newPrimitive(err, fmt.Sprintf(format, args...))
}

func Unwrap(err error) error {
	if err != nil {
		if p, ok := err.(wrapper); ok {
			return p.Unwrap()
		}
	}

	return nil
}

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

func Code(err error) int {
	if err != nil {
		if h, ok := err.(withCode); ok {
			return h.Code()
		}
	}

	return 0
}

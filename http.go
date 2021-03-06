package errors

import "fmt"

// HTTP represents an HTTP error.
type HTTP interface {
	error

	// Code returns the HTTP status code of the error.
	Code() int
}

type http struct {
	*primitive
	code int
}

func (h http) Code() int {
	return h.code
}

// NewHTTP returns a new HTTP error which annotates err with the specified
// message and has the provided status code.
func NewHTTP(err error, code int, message string) error {
	return &http{
		primitive: newPrimitive(err, message),
		code:      code,
	}
}

// HTTPf returns a new HTTP error which annotates err according to the format
// specifier and has the provided status code.
func HTTPf(err error, code int, format string, args ...interface{}) error {
	return &http{
		primitive: newPrimitive(err, fmt.Sprintf(format, args...)),
		code:      code,
	}
}

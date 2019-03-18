package errors

import "fmt"

type HTTP struct {
	*primitive
	code int
}

func (h HTTP) Code() int {
	return h.code
}

func NewHTTP(cause error, code int, message string) error {
	return &HTTP{
		primitive: newPrimitive(cause, message),
		code:      code,
	}
}

func HTTPf(cause error, code int, format string, args ...interface{}) error {
	return &HTTP{
		primitive: newPrimitive(cause, fmt.Sprintf(format, args...)),
		code:      code,
	}
}

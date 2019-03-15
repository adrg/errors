package errors

import "fmt"

type http struct {
	*primitive
	code int
}

func (h http) Code() int {
	return h.code
}

func HTTP(cause error, code int, message string) error {
	return &http{
		primitive: newPrimitive(cause, message),
		code:      code,
	}
}

func HTTPf(cause error, code int, format string, args ...interface{}) error {
	return &http{
		primitive: newPrimitive(cause, fmt.Sprintf(format, args...)),
		code:      code,
	}
}

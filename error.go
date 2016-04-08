package errors

import (
	"fmt"
	"runtime"
)

type (
	Fields  map[string]interface{}
	FmtFunc func(err *Err) string
	ErrType uint
)

const (
	Generic ErrType = iota
	NotFound
	Unauthorized
	NotImplemented
	AlreadyExists
	NotSupported
	NotValid
	NotProvisioned
	NotAssigned
	BadRequest
	MethodNotAllowed
)

type Err struct {
	cause   error
	message string
	fields  Fields
	errType ErrType

	line    int
	file    string
	fmtFunc FmtFunc
}

func NewErr(cause error, fields Fields, fmtFunc FmtFunc,
	format string, args ...interface{}) *Err {
	return newErr(cause, fields, defaultFmtFunc, format, args...)
}

func newErr(cause error, fields Fields, fmtFunc FmtFunc,
	format string, args ...interface{}) *Err {
	if fmtFunc == nil {
		fmtFunc = defaultFmtFunc
	}

	err := &Err{
		cause:   cause,
		message: fmt.Sprintf(format, args...),
		fields:  fields,
		fmtFunc: fmtFunc,
	}

	err.SetLocation(2)
	return err
}

func (e *Err) Error() string {
	return e.fmtFunc(e)
}

func (e *Err) Message() string {
	return e.message
}

func (e *Err) Cause() error {
	return e.cause
}

func (e *Err) Fields() Fields {
	return e.fields
}

func (e *Err) Type() ErrType {
	return e.errType
}

func (e *Err) Location() (string, int) {
	return e.file, e.line
}

func (e *Err) WithMessage(format string, args ...interface{}) *Err {
	e.message = fmt.Sprintf(format, args...)
	return e
}

func (e *Err) WithCause(err error) *Err {
	e.cause = err
	return e
}

func (e *Err) WithFields(fields Fields) *Err {
	e.fields = fields
	return e
}

func (e *Err) WithType(errType ErrType) *Err {
	e.errType = errType
	return e
}

func (e *Err) WithFormat(fmtFunc FmtFunc) *Err {
	e.fmtFunc = fmtFunc
	return e
}

func (e *Err) SetLocation(depth int) {
	_, e.file, e.line, _ = runtime.Caller(depth + 1)
}

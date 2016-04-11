package errors

func New(format string, args ...interface{}) *Err {
	return newErr(nil, nil, defaultFmtFunc, format, args...)
}

func Wrap(err error) *Err {
	return newErr(err, nil, defaultFmtFunc, "")
}

func Cause(err error) error {
	if e, ok := err.(*Err); ok {
		return e.Cause()
	}

	return err
}

func Message(err error) string {
	if e, ok := err.(*Err); ok {
		return e.Message()
	}

	return err.Error()
}

func Type(err error) ErrType {
	if e, ok := err.(*Err); ok {
		return e.Type()
	}

	return Generic
}

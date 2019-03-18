package errors

import (
	"fmt"
	"io"
	"strings"
)

type primitive struct {
	cause   error
	message string
	frame   *frame
}

func newPrimitive(cause error, message string) *primitive {
	return &primitive{
		cause:   cause,
		message: message,
		frame:   callerFrame(2),
	}
}

func (p *primitive) Error() string {
	return p.message
}

func (p *primitive) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		io.WriteString(s, p.format(s.Flag('+')))
	case 's':
		io.WriteString(s, p.Error())
	case 'q':
		io.WriteString(s, fmt.Sprintf("%q", p.Error()))
	}
}

func (p *primitive) Unwrap() error {
	return p.cause
}

func (p *primitive) format(detailed bool) string {
	var b strings.Builder

	var cause error = p
	for {
		if cause == nil {
			break
		}

		if b.Len() > 0 {
			b.WriteString("\n--- ")
		}

		if err, ok := cause.(*primitive); ok {
			b.WriteString(err.message)
			printFrame := detailed && err.frame != nil
			cause = err.cause

			if cause != nil || printFrame {
				b.WriteString(":")
				if printFrame {
					b.WriteString("\n    ")
					b.WriteString(err.frame.String())
				}
			}
		} else {
			b.WriteString(cause.Error())
			if w, ok := cause.(wrapper); ok {
				if cause = w.Unwrap(); cause != nil {
					b.WriteString(":")
					continue
				}
			}

			break
		}
	}

	return b.String()
}

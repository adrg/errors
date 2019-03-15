package errors

import (
	"runtime"
	"strconv"
)

type frame struct {
	Function string
	File     string
	Line     int
}

func callerFrame(skip int) *frame {
	pcs := make([]uintptr, 3)
	if n := runtime.Callers(skip+1, pcs[:]); n == 0 {
		return nil
	}

	callerFrames := runtime.CallersFrames(pcs[:])
	if _, ok := callerFrames.Next(); !ok {
		return nil
	}
	callerFrame, ok := callerFrames.Next()
	if !ok {
		return nil
	}

	return &frame{
		Function: callerFrame.Function,
		File:     callerFrame.File,
		Line:     callerFrame.Line,
	}
}

func (f *frame) String() string {
	return f.File + ":" + strconv.Itoa(f.Line) + " " + f.Function
}

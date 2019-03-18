package errors

import (
	"runtime"
	"strconv"
	"strings"
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

	cfs := runtime.CallersFrames(pcs[:])
	if _, ok := cfs.Next(); !ok {
		return nil
	}
	cf, ok := cfs.Next()
	if !ok {
		return nil
	}

	return &frame{
		Function: cf.Function[strings.LastIndex(cf.Function, "/")+1:],
		File:     cf.File,
		Line:     cf.Line,
	}
}

func (f *frame) String() string {
	return f.File + ":" + strconv.Itoa(f.Line) + " " + f.Function
}

package errors

import (
	"fmt"
	"strings"
)

func defaultFmtFunc(err *Err) string {
	var errFmt string
	var params []interface{}

	if file, line := err.Location(); file != "" {
		errFmt += " %s:%d"
		params = append(params, file, line)
	}

	if message := err.Message(); message != "" {
		errFmt += " %s"
		params = append(params, message)
	}

	if cause := err.Cause(); cause != nil {
		errFmt += " %v"
		params = append(params, cause)
	}

	if fields := err.Fields(); fields != nil {
		for key, value := range fields {
			errFmt += " %s=%v"
			params = append(params, key, value)
		}
	}

	return fmt.Sprintf(strings.TrimSpace(errFmt), params...)
}

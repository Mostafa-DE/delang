package object

import "fmt"

func throwError(format string, a ...interface{}) *Error {
	return &Error{Msg: fmt.Sprintf(format, a...)}
}

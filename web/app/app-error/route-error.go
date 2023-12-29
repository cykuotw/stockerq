package apperror

import (
	"fmt"
	"runtime"
)

type RoutingError struct {
	Layer string
	Err   appError
}

func (e *RoutingError) Error() string {
	return fmt.Sprintf("[%s]\n%s", e.Layer, e.Err.Error())
}

func (e *RoutingError) Unwrap() error {
	return e.Err.Unwrap()
}

func NewRoutingError(op string, err error, route string) (e *RoutingError) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return &RoutingError{
		Layer: "Routing",
		Err: appError{
			Err:            err,
			CallerFile:     frame.File,
			CallerLine:     frame.Line,
			CallerFunction: frame.Function,
		},
	}
}

package apperror

import (
	"fmt"
	"runtime"
)

type RoutingError struct {
	Layer string
	appError
}

func (e *RoutingError) Error() string {
	return fmt.Sprintf("[%s]\n%s", e.Layer, e.Err.Error())
}

func (e *RoutingError) Unwrap() error {
	return e.Err
}

func NewRoutingError(err error) (e *RoutingError) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return &RoutingError{
		Layer: "Routing",
		appError: appError{
			Err:            err,
			CallerFile:     frame.File,
			CallerLine:     frame.Line,
			CallerFunction: frame.Function,
		},
	}
}

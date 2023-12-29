package apperror

import (
	"fmt"
	"runtime"
)

type ControllerError struct {
	Layer string
	Err   appError
}

func (e *ControllerError) Error() string {
	return fmt.Sprintf("[%s]\n%s", e.Layer, e.Err.Error())
}

func (e *ControllerError) Unwrap() error {
	return e.Err.Unwrap()
}

func NewControllerError(op string, err error, handler string) (e *ControllerError) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return &ControllerError{
		Layer: "Controller",
		Err: appError{
			Err:            err,
			CallerFile:     frame.File,
			CallerLine:     frame.Line,
			CallerFunction: frame.Function,
		},
	}
}

package apperror

import (
	"fmt"
	"runtime"
)

type ModelError struct {
	Layer string
	appError
}

func (e *ModelError) Error() string {
	return fmt.Sprintf("[%s]\n%s", e.Layer, e.Err.Error())
}

func (e *ModelError) Unwrap() error {
	return e.Err
}

func NewModelError(err error) (e *ModelError) {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()

	return &ModelError{
		Layer: "Model",
		appError: appError{
			Err:            err,
			CallerFile:     frame.File,
			CallerLine:     frame.Line,
			CallerFunction: frame.Function,
		},
	}
}

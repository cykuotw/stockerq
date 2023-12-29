package apperror

import (
	"errors"
	"fmt"
)

var (
	// [Routing Layer]

	// [Controller Layer]

	// [Model Layer]
	// db connection
	ErrDbConnectFail    = errors.New("db connection fail")
	ErrDbTestFail       = errors.New("db connection fail")
	ErrDbDisconnectFail = errors.New("db disconnection fail")

	// insert stock price

	// update stock price
	ErrInputPriceNotValid = errors.New("input price is not valid")
	ErrIdNotExist         = errors.New("record with this id not exist")

	// get stock price
	ErrZeroDate    = errors.New("startDate and endDate should be not zero")
	ErrReverseDate = errors.New("startDate much be ealier than endDate")
)

type appError struct {
	Err            error
	CallerFile     string
	CallerLine     int
	CallerFunction string
}

func (e *appError) Error() string {
	return fmt.Sprintf("%s:%d %s\nerror: %s",
		e.CallerFile, e.CallerLine, e.CallerFunction, e.Err)
}

func (e *appError) Unwrap() error {
	return e.Err
}

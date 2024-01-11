package apperror

import (
	"errors"
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

type Error interface {
	Error() string
	Unwrap() error
	GetFile() string
	GetLine() int
	GetFunction() string
}

type appError struct {
	Err            error
	CallerFile     string
	CallerLine     int
	CallerFunction string
}

func (e appError) GetFile() string {
	return e.CallerFile
}

func (e appError) GetLine() int {
	return e.CallerLine
}

func (e appError) GetFunction() string {
	return e.CallerFunction
}

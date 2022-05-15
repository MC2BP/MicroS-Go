package errorlib

import (
	"errors"
	"fmt"
	"runtime/debug"
)

var (
	serviceID int
)

type ErrorStruct struct {
	err        error
	stackTrace string
	code       int
	data       interface{}
}

func Err(text string) error {
	return NewErr(errors.New(text))
}

func Errf(format string, v ...interface{}) error {
	return NewErr(errors.New(fmt.Sprintf(format, v...)))
}

func NewErr(err error) error {
	return ErrorStruct{
		err:        err,
		stackTrace: string(debug.Stack()),
	}
}

func WithoutStackTrace(err error) error {
	if errorStruct, ok := err.(ErrorStruct); ok {
		errorStruct.stackTrace = ""
		return errorStruct
	}
	return ErrorStruct{
		err: err,
	}
}

func WithErrorCode(err error, code int) error {
	if errorStruct, ok := err.(ErrorStruct); ok {
		errorStruct.code = code
		return errorStruct
	}
	return ErrorStruct{
		err:        err,
		stackTrace: string(debug.Stack()),
		code:       code,
	}
}

func WithData(err error, data interface{}) error {
	if errorStruct, ok := err.(ErrorStruct); ok {
		errorStruct.data = data
		return errorStruct
	}
	return ErrorStruct{
		err:        err,
		stackTrace: string(debug.Stack()),
		data:       data,
	}
}

func GetError(err error) error {
	if errorStruct, ok := err.(ErrorStruct); ok {
		return errorStruct.err
	}
	return err
}

func GetErrorCode(err error) int {
	if errorStruct, ok := err.(ErrorStruct); ok {
		return errorStruct.code
	}
	return 0
}

func GetData(err error) interface{} {
	if errorStruct, ok := err.(ErrorStruct); ok {
		return errorStruct.data
	}
	return nil
}

func (err ErrorStruct) Error() string {
	if err.stackTrace != "" {
		return fmt.Sprint(err.err.Error(), "\n", err.stackTrace)
	}
	return err.err.Error()
}

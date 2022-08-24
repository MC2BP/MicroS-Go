package errorlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime/debug"
)

var (
	serviceID int
)

type ErrorStruct struct {
	Err        error
	stackTrace string
	Code       int
	Data       interface{}
}

// Err creates an error consisting of a simple text
func Err(text string) ErrorStruct {
	return WrapError(errors.New(text))
}

// Errf creates an error consisting of a text, formatted like fmt.Sprintf
func Errf(format string, v ...interface{}) ErrorStruct {
	return WrapError(errors.New(fmt.Sprintf(format, v...)))
}

// WrapError wraps the error into an ErrorStruct with Stacktrace
func WrapError(err error) ErrorStruct {
	return ErrorStruct{
		Err:        err,
		stackTrace: string(debug.Stack()),
	}
}

// WithoutStackTrace removes the stacktrace from the error
func WithoutStackTrace(err error) error {
	if errorStruct, ok := err.(ErrorStruct); ok {
		errorStruct.stackTrace = ""
		return errorStruct
	}
	return err
}

// WithErrorCode sets the errorcode
func (e ErrorStruct) WithErrorCode(code int) ErrorStruct {
	e.Code = code
	return e
}

// WithData sets the data of the error
func (e ErrorStruct) WithData(data interface{}) ErrorStruct {
	e.Data = data
	return e
}

// Prints out the error
func (e ErrorStruct) Error() string {
	stackTrace := ""
	if e.stackTrace != "" {
		stackTrace = "\n" + e.stackTrace
	}
	data := ""
	if e.Data == nil {
		rawData, err := json.Marshal(e.Data)
		if err != nil {
			data = "\n" + err.Error()
		} else {
			data = "\n" + string(rawData)
		}
	}
	return fmt.Sprint(e.Err.Error(), data, stackTrace)
}

// return the errorcode of the error
func GetErrorCode(err error) int {
	if errorStruct, ok := err.(ErrorStruct); ok {
		return errorStruct.Code
	}
	return 0
}

// returns the underlying error if available
func GetError(err error) error {
	
	if errorStruct, ok := err.(ErrorStruct); ok {
		return errorStruct.Err
	}
	return err
}


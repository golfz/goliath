package goliath

import (
	"runtime/debug"
	"time"
)

type Error interface {
	// SetStatus sets `int`
	SetStatus(status int)

	// SetMessage sets `string`
	SetMessage(msg string)

	// SetLogID sets `string`
	SetLogID(logID string)

	// SetErrorCode sets `string`
	SetErrorCode(errorCode string)

	// SetErrorArgs sets `map[string]interface{}`
	SetErrorArgs(args map[string]interface{})

	// SetErrorArg sets a pair of `(string, interface{})`
	SetErrorArg(key string, v interface{})

	// SetError sets `error`
	SetError(err error)
}

type goliathError struct {
	Status    int                    `json:"status"`
	Message   string                 `json:"message"`
	Time      time.Time              `json:"time"`
	LogID     string                 `json:"log_id"`
	ErrorCode string                 `json:"error_code"`
	ErrorArgs map[string]interface{} `json:"error_args"`
	ErrorDev  errorDev               `json:"error_dev"`
	err       error
}

type errorDev struct {
	Error      string `json:"error"`
	Stacktrace string `json:"stacktrace"`
}

// NewError create a new goliathError
func NewError(status int, errCode string, errArgs map[string]interface{}, err error, logID string, optionalMsg string) *goliathError {
	return &goliathError{
		Status:    status,
		Message:   optionalMsg,
		Time:      time.Now(),
		LogID:     logID,
		ErrorCode: errCode,
		ErrorArgs: errArgs,
		ErrorDev: errorDev{
			Error:      err.Error(),
			Stacktrace: string(debug.Stack()),
		},
		err: err,
	}
}

// Error implements `error interface` form `builtin` package
func (e *goliathError) Error() string {
	return e.Message
}

// Unwrap implements `Wrapper interface` from `xerrors` package
func (e *goliathError) Unwrap() error {
	return e.err
}

func (e *goliathError) SetStatus(status int) {
	e.Status = status
}

func (e *goliathError) SetMessage(msg string) {
	e.Message = msg
}

func (e *goliathError) SetLogID(logID string) {
	e.LogID = logID
}

func (e *goliathError) SetErrorCode(errorCode string) {
	e.ErrorCode = errorCode
}

func (e *goliathError) SetErrorArgs(args map[string]interface{}) {
	e.ErrorArgs = args
}

func (e *goliathError) SetErrorArg(key string, v interface{}) {
	e.ErrorArgs[key] = v
}

func (e *goliathError) SetError(err error) {
	e.err = err
	e.ErrorDev.Error = err.Error()
}

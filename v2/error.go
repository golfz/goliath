package goliath

import (
	"time"
)

type Error interface {
	// Data returns `interface{}`
	Data() interface{}

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

	// SetStacktrace sets `string`
	SetStacktrace(s string)
}

type errorData struct {
	Status    int                    `json:"status"`
	Message   string                 `json:"message"`
	Time      time.Time              `json:"time"`
	LogID     string                 `json:"log_id"`
	ErrorCode string                 `json:"error_code"`
	ErrorArgs map[string]interface{} `json:"error_args"`
	ErrorDev  errorDev               `json:"error_dev"`
}

type errorDev struct {
	Stacktrace string `json:"stacktrace"`
}

// NewError create a new errorData
func NewError() *errorData {
	return &errorData{}
}

func (e *errorData) Error() string {
	return e.Message
}

func (e *errorData) Data() interface{} {
	return e
}

func (e *errorData) SetStatus(status int) {
	e.Status = status
}

func (e *errorData) SetMessage(msg string) {
	e.Message = msg
}

func (e *errorData) SetLogID(logID string) {
	e.LogID = logID
}

func (e *errorData) SetErrorCode(errorCode string) {
	e.ErrorCode = errorCode
}

func (e *errorData) SetErrorArgs(args map[string]interface{}) {
	e.ErrorArgs = args
}

func (e *errorData) SetErrorArg(key string, v interface{}) {
	e.ErrorArgs[key] = v
}

func (e *errorData) SetStacktrace(s string) {
	e.ErrorDev.Stacktrace = s
}

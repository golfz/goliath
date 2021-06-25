package output

import (
	"github.com/golfz/goliath/utils/contexts"
	"runtime/debug"
	"time"
)

type GoliathError interface {
	Errors() Error
	SetStatus(s int)
	GetStatus() int
}

type Error struct {
	// if field is used in V2 or V3 api only, must omitempty it.
	Status    int         `json:"status"`               // V2, V3
	Message   string      `json:"message"`              // V2, V3
	Time      time.Time   `json:"time"`                 // V2, V3
	LogId     string      `json:"log_id,omitempty"`     // V3
	ErrorCode string      `json:"error_code,omitempty"` // V3
	ErrorArgs interface{} `json:"error_args,omitempty"` // V3
	Type      string      `json:"type,omitempty"`       // V2
	Code      string      `json:"code,omitempty"`       // V2
	Error     string      `json:"error,omitempty"`      // V2
	ErrorDev  ErrorDev    `json:"error_dev"`            // V2 V3
}

type ErrorDev struct {
	Error      string `json:"error"`
	Stacktrace string `json:"stacktrace"`
}

func (e *Error) Errors() Error {
	return *e
}

func (e *Error) SetStatus(s int) {
	e.Status = s
}

func (e *Error) GetStatus() int {
	return e.Status
}

func SmallError(status int, code string, errType string, msg string) GoliathError {
	return &Error{
		Status:  status,
		Time:    time.Now(),
		Type:    errType,
		Code:    code,
		Error:   msg,
		Message: msg,
		ErrorDev: ErrorDev{
			Error:      msg,
			Stacktrace: string(debug.Stack()),
		},
	}
}

func ErrorV3(status int, errCode string, errArgs interface{}, err error, ctx contexts.GoliathContextor, optionalMsg string) GoliathError {
	return &Error{
		Status:    status,
		Message:   optionalMsg,
		Time:      time.Now(),
		LogId:     ctx.GetLogId(),
		ErrorCode: errCode,
		ErrorArgs: errArgs,
		ErrorDev: ErrorDev{
			Error:      err.Error(),
			Stacktrace: string(debug.Stack()),
		},
	}
}

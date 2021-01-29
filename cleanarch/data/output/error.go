package output

import (
	"runtime/debug"
	"time"
)

type GoliathError interface {
	Errors() Error
	SetStatus(s int)
}

type Error struct {
	Status   int
	Time     time.Time
	Type     string
	Code     string
	Error    string
	Message  string
	ErrorDev ErrorDev
}

func (e *Error) Errors() Error {
	return *e
}

func (e *Error) SetStatus(s int) {
	e.Status = s
}

type ErrorDev struct {
	Error      string
	Stacktrace string
}

func SmallError(status int, code string, errType string, msg string) GoliathError {
	return &Error{
		Status:   status,
		Time:     time.Now(),
		Type:     errType,
		Code:     code,
		Error:    msg,
		Message:  msg,
		ErrorDev: ErrorDev{
			Error:      msg,
			Stacktrace: string(debug.Stack()),
		},
	}
}

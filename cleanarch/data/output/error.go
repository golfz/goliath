package output

import (
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

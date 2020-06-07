package output

import (
	"time"
)

type GoliathError interface {
	Errors() Error
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

type ErrorDev struct {
	Error      string
	Stacktrace string
}

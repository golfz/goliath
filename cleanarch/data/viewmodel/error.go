package viewmodel

import "time"

type Error struct {
	Status    int      `json:"status"`
	Timestamp string   `json:"timestamp"`
	Type      string   `json:"type"`
	Code      string   `json:"code"`
	Error     string   `json:"error"`
	Message   string   `json:"message"`
	ErrorDev  ErrorDev `json:"error_dev"`
}

func (e *Error) Time(t time.Time) {
	e.Timestamp = t.Format(time.RFC3339)
}

type ErrorDev struct {
	Error      string `json:"error"`
	Stacktrace string `json:"stacktrace"`
}

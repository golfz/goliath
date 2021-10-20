package goliath

import (
	"database/sql"
	"net/http"
)

type Goliath interface {
	// Request returns `*http.Request`.
	Request() *http.Request

	// SetRequest sets `*http.Request`.
	SetRequest(r *http.Request)

	// SetResponse sets `http.ResponseWriter`.
	SetResponse(w http.ResponseWriter)

	// Response returns `http.ResponseWriter`.
	Response() http.ResponseWriter

	// SetDBConnection sets `*sql.DB`
	SetDBConnection(db *sql.DB)

	// DB returns `*sql.DB`
	DB() *sql.DB

	// SetLogID sets `string`.
	SetLogID(logID string)

	// LogID returns `string`.
	LogID() string
}

type goliath struct {
	writer  http.ResponseWriter
	request *http.Request
	db      *sql.DB
	logID   string
}

func New() *goliath {
	return &goliath{}
}

func (g *goliath) Request() *http.Request {
	panic("implement me")
}

func (g *goliath) SetRequest(r *http.Request) {
	panic("implement me")
}

func (g *goliath) SetResponse(w http.ResponseWriter) {
	panic("implement me")
}

func (g *goliath) Response() http.ResponseWriter {
	panic("implement me")
}

func (g *goliath) SetDBConnection(db *sql.DB) {
	panic("implement me")
}

func (g *goliath) DB() *sql.DB {
	panic("implement me")
}

func (g *goliath) SetLogID(logID string) {

}

func (g *goliath) LogID() string {
	return g.logID
}

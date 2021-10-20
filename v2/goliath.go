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

	// SetResponseWriter sets `http.ResponseWriter`.
	SetResponseWriter(w http.ResponseWriter)

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

func New() Goliath {
	return &goliath{}
}

func (g *goliath) Request() *http.Request {
	return g.request
}

func (g *goliath) SetRequest(r *http.Request) {
	g.request = r
	logID, ok := g.request.Context().Value(ContextLogIdKey).(string)
	if ok {
		g.logID = logID
	}
}

func (g *goliath) SetResponseWriter(w http.ResponseWriter) {
	g.writer = w
}

func (g *goliath) Response() http.ResponseWriter {
	return g.writer
}

func (g *goliath) SetDBConnection(db *sql.DB) {
	g.db = db
}

func (g *goliath) DB() *sql.DB {
	return g.db
}

func (g *goliath) SetLogID(logID string) {
	g.logID = logID
}

func (g *goliath) LogID() string {
	return g.logID
}

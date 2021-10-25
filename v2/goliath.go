package goliath

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Goliath interface {
	// SetRequest sets `*http.Request`.
	SetRequest(r *http.Request)

	// Request returns `*http.Request`.
	Request() *http.Request

	// SetResponseWriter sets `http.ResponseWriter`.
	SetResponseWriter(w http.ResponseWriter)

	// Response returns `http.ResponseWriter`.
	Response() http.ResponseWriter

	// SetSqlDB sets map of `string -> *sql.DB`
	SetSqlDB(id string, db *sql.DB)

	// DB returns `*sql.DB`
	sqlDB(id string) (*sql.DB, error)

	// SetLogID sets `string`.
	SetLogID(logID string)

	// LogID returns `string`.
	LogID() string

	// NewInternalError return `*goliathError`
	NewInternalError(errCode string, errArgs map[string]interface{}, err error, optionalMsg string) *goliathError
}

type goliath struct {
	writer   http.ResponseWriter
	request  *http.Request
	sqlDBMap map[string]*sql.DB
	logID    string
}

func New() Goliath {
	return &goliath{}
}

func (g *goliath) SetRequest(r *http.Request) {
	g.request = r
	logID, ok := g.request.Context().Value(ContextLogIdKey).(string)
	if ok {
		g.logID = logID
	}
}

func (g *goliath) Request() *http.Request {
	return g.request
}

func (g *goliath) SetResponseWriter(w http.ResponseWriter) {
	g.writer = w
}

func (g *goliath) Response() http.ResponseWriter {
	return g.writer
}

func (g *goliath) SetSqlDB(id string, db *sql.DB) {
	g.sqlDBMap[id] = db
}

func (g *goliath) sqlDB(id string) (*sql.DB, error) {
	db, ok := g.sqlDBMap[id]
	if !ok {
		return nil, fmt.Errorf("no sqlDB with id: %s", id)
	}
	return db, nil
}

func (g *goliath) SetLogID(logID string) {
	g.logID = logID
}

func (g *goliath) LogID() string {
	return g.logID
}

func (g *goliath) NewInternalError(errCode string, errArgs map[string]interface{}, err error, optionalMsg string) *goliathError {
	errStatus := http.StatusInternalServerError
	return NewError(errStatus, errCode, errArgs, err, g.logID, optionalMsg)
}

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

	// SqlDB returns `*sql.DB`
	SqlDB(id string) (*sql.DB, error)

	// SetLogID sets `string`.
	SetLogID(logID string)

	// LogIDKey returns `string`.
	LogIDKey() string

	// SetLogIDKey sets `string`.
	SetLogIDKey(key string)

	// LogID returns `string`.
	LogID() string

	// NewInternalError return `*goliathError`
	NewInternalError(errCode string, errArgs map[string]interface{}, err error, optionalMsg string) *goliathError

	// SetValue sets `interface{}`.
	SetValue(key string, value interface{})

	// GetValue returns `(interface{}, bool)`.
	GetValue(key string) (interface{}, bool)
}

type goliath struct {
	writer   http.ResponseWriter
	request  *http.Request
	sqlDBMap map[string]*sql.DB
	logID    string
	logIdKey string
	valueMap map[string]interface{}
}

func New() Goliath {
	return &goliath{
		sqlDBMap: make(map[string]*sql.DB),
		valueMap: make(map[string]interface{}),
	}
}

func (g *goliath) SetRequest(r *http.Request) {
	g.request = r

	key := ContextLogIDKey
	if g.logIdKey != "" {
		key = g.logIdKey
	}

	logID, ok := g.request.Context().Value(key).(string)
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

func (g *goliath) SqlDB(id string) (*sql.DB, error) {
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

func (g *goliath) LogIDKey() string {
	return g.logIdKey
}

func (g *goliath) SetLogIDKey(key string) {
	g.logIdKey = key
}

func (g *goliath) NewInternalError(errCode string, errArgs map[string]interface{}, err error, optionalMsg string) *goliathError {
	errStatus := http.StatusInternalServerError
	return NewError(errStatus, errCode, errArgs, err, g.logID, optionalMsg)
}

func (g *goliath) SetValue(key string, value interface{}) {
	g.valueMap[key] = value
}

func (g *goliath) GetValue(key string) (interface{}, bool) {
	v, ok := g.valueMap[key]
	return v, ok
}

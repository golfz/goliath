package contexts

import (
	"database/sql"
	"net/http"
)

type GoliathContext struct {
	Writer      http.ResponseWriter
	Request     *http.Request
	Db          *sql.DB
	authContext *authContext
}

type GoliathContextor interface {
	GetResponseWriter() http.ResponseWriter
	GetRequest() *http.Request
	GetAuthContext() *authContext
}

func (ctx *GoliathContext) GetResponseWriter() http.ResponseWriter {
	return ctx.Writer
}

func (ctx *GoliathContext) GetRequest() *http.Request {
	return ctx.Request
}

func (ctx *GoliathContext) GetAuthContext() *authContext {
	if ctx.authContext == nil {
		ctx.authContext = &authContext{request: ctx.Request}
	}
	return ctx.authContext
}

package contexts

import (
	"net/http"
)

type goliathContext struct {
	writer      http.ResponseWriter
	request     *http.Request
	authContext *authContext
}

func NewGoliathContext(w http.ResponseWriter, r *http.Request) *goliathContext {
	return &goliathContext{writer: w, request: r}
}

type GoliathContextor interface {
	GetResponseWriter() http.ResponseWriter
	GetRequest() *http.Request
	AuthContext() *authContext
}

func (ctx *goliathContext) GetResponseWriter() http.ResponseWriter {
	return ctx.writer
}

func (ctx *goliathContext) GetRequest() *http.Request {
	return ctx.request
}

func (ctx *goliathContext) AuthContext() *authContext {
	if ctx.authContext == nil {
		ctx.authContext = &authContext{request: ctx.request}
	}
	return ctx.authContext
}

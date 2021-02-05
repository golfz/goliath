package contexts

import (
	"net/http"
)

type goliathContext struct {
	request     *http.Request
	authContext *authContext
}

func NewGoliathContext(r *http.Request) *goliathContext {
	return &goliathContext{request: r}
}

type GoliathContextor interface {
	Auth() *authContext
}

func (ctx *goliathContext) Auth() *authContext {
	if ctx.authContext == nil {
		ctx.authContext = &authContext{request: ctx.request}
	}
	return ctx.authContext
}

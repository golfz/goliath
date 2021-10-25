package goliath

import (
	"encoding/json"
)

type Viewer interface {
	Write(status int, data interface{})
}

type JSONRestfulView struct {
	Ctx Goliath
}

func (v *JSONRestfulView) Write(status int, body interface{}) {
	v.Ctx.Response().Header().Set("Content-Type", "application/json")
	v.Ctx.Response().WriteHeader(status)
	if body != nil {
		json.NewEncoder(v.Ctx.Response()).Encode(body)
	}
}

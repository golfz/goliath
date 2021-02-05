package json

import (
	"encoding/json"
	"github.com/golfz/goliath/cleanarch/data/viewmodel"
	"github.com/golfz/goliath/utils/contexts"
)

type JsonRestfulView struct {
	Ctx contexts.GoliathContextor
}

func (v *JsonRestfulView) Write(header viewmodel.HttpHeader, body interface{}) {
	v.Ctx.GetResponseWriter().Header().Set("Content-Type", "application/json")
	v.Ctx.GetResponseWriter().WriteHeader(header.StatusCode)
	if body != nil {
		json.NewEncoder(v.Ctx.GetResponseWriter()).Encode(body)
	}
}

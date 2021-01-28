package json

import (
	"encoding/json"
	"github.com/golfz/goliath/cleanarch/data/viewmodel"
	"net/http"
)

type JsonRestfulView struct {
	Writer http.ResponseWriter
}

func (v *JsonRestfulView) Write(header viewmodel.HttpHeader, body interface{}) {
	v.Writer.Header().Set("Content-Type", "application/json")
	v.Writer.WriteHeader(header.StatusCode)
	if body != nil {
		json.NewEncoder(v.Writer).Encode(body)
	}
}

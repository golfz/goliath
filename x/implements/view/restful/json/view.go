package json

import (
	"encoding/json"
	"github.com/golfz/goliath/x/data/viewmodel"
	"net/http"
)

type jsonRestfulView struct {
	writer http.ResponseWriter
}

func (v *jsonRestfulView) Write(header viewmodel.HttpHeader, body interface{}) {
	v.writer.Header().Set("Content-Type", "application/json")
	v.writer.WriteHeader(header.StatusCode)
	if body != nil {
		json.NewEncoder(v.writer).Encode(body)
	}
}

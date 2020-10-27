package xutils

import (
	"encoding/json"
	"github.com/golfz/goliath/x/data/output"
	"io"
)

func JsonDecode(r io.Reader, v interface{}) *output.Error {

	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		//return output.Error{
		//	ErrorCode: http.StatusBadRequest,
		//	Message:   "",
		//}
	}
	return nil
}

package qrystr

import (
	"net/http"
	"strconv"
)

func hasKey(r *http.Request, key string) bool {
	v := r.URL.Query().Get(key)
	hasKey := v != ""
	return hasKey
}

func keyToIntPtr(r *http.Request, key string) *int {
	q := r.URL.Query()

	val := 0
	var err error

	if val, err = strconv.Atoi(q.Get(key)); err != nil {
		return nil
	}

	return &val
}

func hasStringInSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

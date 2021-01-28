package qrystr

import (
	"net/http"
)

const (
	QUERY_STRING_FIELDNAME_LIMIT = "limit"
	QUERY_STRING_FIELDNAME_PAGE  = "page"
)

func getLimit(r *http.Request) *int {
	limitIntPtr := keyToIntPtr(r, QUERY_STRING_FIELDNAME_LIMIT)

	if limitIntPtr == nil {
		return nil
	}

	if *limitIntPtr < 1 {
		return nil
	}

	return limitIntPtr
}

func getPage(r *http.Request) *int {
	if getLimit(r) == nil {
		return nil
	}

	pageIntPtr := keyToIntPtr(r, QUERY_STRING_FIELDNAME_PAGE)

	if pageIntPtr == nil {
		return nil
	}

	if *pageIntPtr < 1 {
		return nil
	}

	return pageIntPtr
}

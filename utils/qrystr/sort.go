package qrystr

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	QUERY_STRING_FIELDNAME_SORT_BY = "sort_by"
)

type SortBy struct {
	Field   string `json:"field"`
	OrderBy string `json:"order_by"`
}

const (
	ORDER_BY_ASC  string = "asc"
	ORDER_BY_DESC string = "desc"
)

func getSort(r *http.Request) []SortBy {
	var out []SortBy

	if !hasKey(r, QUERY_STRING_FIELDNAME_SORT_BY) {
		return out
	}

	sortQryValue := strings.TrimSpace(r.URL.Query().Get(QUERY_STRING_FIELDNAME_SORT_BY))
	if sortQryValue == "" {
		return out
	}

	sortList := strings.Split(sortQryValue, ",")
	for _, sortItem := range sortList {
		sortItem = strings.TrimSpace(sortItem)
		if sortItem == "" {
			continue
		}

		colonIndex := strings.Index(sortItem, ":")
		if colonIndex == -1 {
			continue
		}

		sortElements := strings.Split(sortItem, ":")
		if len(sortElements) != 2 {
			continue
		}

		sortField := strings.TrimSpace(sortElements[0])
		sortOrderByStr := strings.ToLower(strings.TrimSpace(sortElements[1]))

		if sortField == "" || sortOrderByStr == "" {
			continue
		}

		var sortOrderBy string
		switch sortOrderByStr {
		case "asc":
			sortOrderBy = ORDER_BY_ASC
		case "desc":
			sortOrderBy = ORDER_BY_DESC
		default:
			return out
		}

		aSort := SortBy{
			Field:   sortField,
			OrderBy: sortOrderBy,
		}

		out = append(out, aSort)
	}

	return out
}

func (q *QueryString) GetSortSql(isNeedSortKeyword bool) string {
	if len(q.Sort) == 0 {
		return ""
	}

	out := " "

	if isNeedSortKeyword {
		out += " ORDER BY "
	}

	for i, sort := range q.Sort {
		if i > 0 {
			out += " , "
		}
		out += fmt.Sprintf(" `%s` %s ", sort.Field, strings.ToUpper(sort.OrderBy))
	}

	out = strings.ReplaceAll(out, "  ", " ")

	return out
}

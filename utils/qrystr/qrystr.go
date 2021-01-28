package qrystr

import (
	"errors"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strings"
)

// spec
//
//ops.Tag()
//
//ops.Tag("eq").Datetime(0, "2015-07-01T15:03:04Z")
//ops.Tag("eq").Int(0)
//ops.Tag("eq").String(0)
//ops.Tag("eq").Float(0)
//ops.Tag("eq").Double(0)
//ops.Tag("eq").Value(0)
//
//ops.Tag("eq").Int()
//ops.Tag("eq").String()
//ops.Tag("eq").Float()
//ops.Tag("eq").Double()
//ops.Tag("eq").Value()
//
//filter.Field("user_id").Tag("eq").Int(0)

/*
qry.Sort() -> []SortBy
qry.Sort("name") -> string("desc")

qry.FilterByField("country").Eq().Int(n ..int)


*/

// use this instead of const []string
func getStandardLimitAndSortFieldList() []string {
	return []string{
		QUERY_STRING_FIELDNAME_SORT_BY,
		QUERY_STRING_FIELDNAME_LIMIT,
		QUERY_STRING_FIELDNAME_PAGE,
	}
}

// **********************************************
// Query String
// **********************************************
type QueryString struct {
	RequestPath string
	Filter      []FilterField
	Sort        []SortBy
	Limit       *int
	Page        *int
	ResultSize  int
	TotalSize   int
	ResultPage  int
	TotalPage   int
}

func New(r *http.Request) QueryString {
	return QueryString{
		RequestPath: r.URL.Path,
		Filter:      getFilter(r),
		Sort:        getSort(r),
		Limit:       getLimit(r),
		Page:        getPage(r),
		ResultSize:  0,
		TotalSize:   0,
		ResultPage:  0,
		TotalPage:   0,
	}
}

type CollectionResponse struct {
	Size      int           `json:"size"`
	TotalSize int           `json:"total_size"`
	Page      int           `json:"page"`
	TotalPage int           `json:"total_page"`
	Filter    []FilterField `json:"filter"`
	SortBy    []SortBy      `json:"sort_by"`
	Limit     Limit         `json:"limit"`
	Link      Link          `json:"link"`
	Data      interface{}   `json:"data"`
}

type Limit struct {
	Limit *int `json:"limit"`
	Page  *int `json:"page"`
}

type Link struct {
	Self     string  `json:"self"`
	First    string  `json:"first"`
	Last     string  `json:"last"`
	Previous *string `json:"prev"`
	Next     *string `json:"next"`
}

type QueryStringer interface {
	GetCollectionResponse(data []interface{}) (CollectionResponse, error)

	GetSortSql(isNeedSortKeyword bool) string

	GetFilterSql(isNeedWhereKeyword bool) string
}

func (q *QueryString) GetCollectionResponse(data []interface{}) (CollectionResponse, error) {
	q.TotalSize = len(data)

	q.TotalPage, _ = calculateTotalPage(q.TotalSize, q.Limit)

	q.ResultPage, _ = calculateResultPage(q.TotalPage, q.Page)

	q.Page, _ = reCalculateRequestPage(q.TotalPage, q.Limit, q.Page)

	data, _ = applyLimit(data, q.Limit, q.Page)

	q.ResultSize = len(data)

	return CollectionResponse{
		Size:      q.ResultSize,
		TotalSize: q.TotalSize,
		Page:      q.ResultPage,
		TotalPage: q.TotalPage,
		Filter:    q.Filter,
		SortBy:    q.Sort,
		Limit: Limit{
			Limit: q.Limit,
			Page:  q.Page,
		},
		Link: getLink(*q),
		Data: data,
	}, nil
}

func calculateTotalPage(totalSize int, limit *int) (int, error) {
	if totalSize < 0 {
		return 0, errors.New("totalPage can not least than 0")
	} else if totalSize == 0 {
		return 1, nil
	}

	if limit == nil {
		return 1, nil
	} else if *limit < 1 {
		return 0, errors.New("limit can not least than 1")
	}

	return int(math.Ceil(float64(totalSize) / float64(*limit))), nil
}

func calculateResultPage(totalPage int, page *int) (int, error) {
	if totalPage < 1 {
		return 0, errors.New("totalPage can not least than 1")
	}

	if page == nil {
		return 1, nil
	} else if *page < 1 {
		return 0, errors.New("page can not least than 1")
	} else if *page > totalPage {
		return totalPage, nil
	}

	return *page, nil
}

func reCalculateRequestPage(totalPage int, limit *int, page *int) (*int, error) {
	if totalPage < 1 {
		return nil, errors.New("totalPage can not least than 1")
	}

	if limit == nil {
		return nil, nil
	}

	if page == nil {
		return nil, nil
	} else if *page > totalPage {
		return &totalPage, nil
	}

	return page, nil
}

func applyLimit(data []interface{}, limit *int, page *int) ([]interface{}, error) {
	if limit == nil {
		return data, nil
	} else if *limit < 1 {
		return nil, errors.New("limit can not least than 1")
	}

	if page != nil {
		if *page < 1 {
			return nil, errors.New("page can not least than 1")
		}
	}

	startIndex := 0

	offset := getOffset(page, limit)

	if offset > 0 {
		startIndex = offset
	}
	if offset > len(data) {
		startIndex = len(data)
	}

	endIndex := len(data)
	if (startIndex + *limit) < endIndex {
		endIndex = startIndex + *limit
	}

	return data[startIndex:endIndex], nil
}

func getOffset(page *int, limit *int) int {
	if limit == nil {
		return 0
	} else if *limit < 1 {
		return 0
	}

	if page == nil {
		return 0
	} else if *page < 1 {
		return 0
	}
	return (*page - 1) * (*limit)
}

func getLink(q QueryString) Link {
	qrtStr := ""

	filter := getFilterQryString(q)
	if len(filter) > 0 {
		qrtStr = filter
	}

	sort := getSortQryString(q)
	if len(sort) > 0 {
		if len(qrtStr) == 0 {
			qrtStr = sort
		} else {
			qrtStr += fmt.Sprintf("&%s", sort)
		}
	}

	fmt.Println("query string:", qrtStr)

	// link self
	self := getLinkSelfQryString(q)
	self = *combineQryString(q.RequestPath, qrtStr, &self)
	fmt.Println("Self:", self)

	// link first
	first := getLinkFirstQryString(q)
	first = *combineQryString(q.RequestPath, qrtStr, &first)
	fmt.Println("First:", first)

	// link last
	last := getLinkLastQryString(q)
	last = *combineQryString(q.RequestPath, qrtStr, &last)
	fmt.Println("Last:", last)

	// link previous
	previous := getLinkPreviousQryString(q)
	previous = combineQryString(q.RequestPath, qrtStr, previous)
	if previous == nil {
		fmt.Println("Previous:", previous)
	} else {
		fmt.Println("Previous:", *previous)
	}

	// link next
	next := getLinkNextQryString(q)
	next = combineQryString(q.RequestPath, qrtStr, next)
	if next == nil {
		fmt.Println("Previous:", next)
	} else {
		fmt.Println("Previous:", *next)
	}

	return Link{
		Self:     self,
		First:    first,
		Last:     last,
		Previous: previous,
		Next:     next,
	}
}

func getFilterQryString(q QueryString) string {
	outQueryString := ""

	// Filter
	for _, field := range q.Filter {
		outField := ""
		if len(outQueryString) > 0 {
			outField += "&"
		}
		outField += fmt.Sprintf("%s=", field.Field)

		for iTag, tag := range field.Condition {
			outTag := fmt.Sprintf("%s:", tag.Operator)

			if tag.Operator == "in" {
				// in tag
				inVal := ""
				for _, v := range tag.Values {
					if len(inVal) > 0 {
						inVal += fmt.Sprintf("+%s", v)
					} else {
						inVal = v
					}
				}
				outTag += url.QueryEscape(inVal)

			} else {
				// all tag (other tags except "in" tag)
				outTag += url.QueryEscape(tag.Values[0])
			}

			if iTag > 0 {
				outField += ","
			}
			// assign tag to field
			outField += outTag
		}

		outQueryString += outField
	}

	return outQueryString
}

func getSortQryString(q QueryString) string {
	outQueryString := ""

	// Filter
	for _, sortBy := range q.Sort {
		if len(outQueryString) == 0 {
			outQueryString = fmt.Sprintf("sort_by=%s:%s", sortBy.Field, strings.ToLower(sortBy.OrderBy))
		} else {
			outQueryString += fmt.Sprintf(",%s:%s", sortBy.Field, strings.ToLower(sortBy.OrderBy))
		}
	}

	return outQueryString
}

func getLinkSelfQryString(q QueryString) string {
	outQueryString := ""

	if q.Limit != nil {
		outQueryString = fmt.Sprintf("page=%d&limit=%d", q.ResultPage, *q.Limit)
	}

	return outQueryString
}

func getLinkFirstQryString(q QueryString) string {
	outQueryString := ""

	if q.Limit != nil {
		outQueryString = fmt.Sprintf("page=1&limit=%d", *q.Limit)
	}

	return outQueryString
}

func getLinkLastQryString(q QueryString) string {
	outQueryString := ""

	if q.Limit != nil {
		outQueryString = fmt.Sprintf("page=%d&limit=%d", q.TotalPage, *q.Limit)
	}

	return outQueryString
}

func getLinkPreviousQryString(q QueryString) *string {
	outQueryString := ""

	if q.Limit == nil {
		return nil
	} else if q.ResultPage == 1 {
		return nil
	} else {
		lastPage := q.ResultPage - 1
		outQueryString = fmt.Sprintf("page=%d&limit=%d", lastPage, *q.Limit)
	}

	return &outQueryString
}

func getLinkNextQryString(q QueryString) *string {
	outQueryString := ""

	if q.Limit == nil {
		return nil
	} else if q.ResultPage == q.TotalPage {
		return nil
	} else {
		nextPage := q.ResultPage + 1
		outQueryString = fmt.Sprintf("page=%d&limit=%d", nextPage, *q.Limit)
	}

	return &outQueryString
}

func combineQryString(requestPath string, filterAndSort string, linkToPage *string) *string {
	if linkToPage == nil {
		return nil
	}

	pageLink := *linkToPage

	if len(pageLink) == 0 {
		pageLink = fmt.Sprintf("?%s", filterAndSort)

	} else if len(pageLink) > 0 {
		if len(filterAndSort) == 0 {
			pageLink = fmt.Sprintf("?%s", pageLink)
		} else {
			pageLink = fmt.Sprintf("?%s&%s", filterAndSort, pageLink)
		}
	}

	out := fmt.Sprintf("%s%s", requestPath, pageLink)

	return &out
}

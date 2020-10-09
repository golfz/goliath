package common

import "time"

type CollectionDescription struct {
	Size        int
	TotalSize   int
	Page        int
	TotalPage   int
	LimitOffset int
	LimitCount  int
	OrderBy     []CollectionOrderBy
}

type CollectionOrderBy struct {
	Field  string
	IsDesc bool
}

type CollectionFilter struct {
	Page          *int
	LimitOffset   *int
	LimitCount    *int
	OrderBy       *[]CollectionOrderBy
	DatetimeStart *time.Time
	DatetimeEnd   *time.Time
}

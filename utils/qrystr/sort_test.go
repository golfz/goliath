package qrystr

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGetSort_NoSortBy(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	assert.Empty(t, getSort(r))

	r = httptest.NewRequest("GET", "/other_field_name=name:asc,age:desc", nil)
	assert.Empty(t, getSort(r))
}

func TestGetSort_SortByWithoutValue(t *testing.T) {
	r := httptest.NewRequest("GET", "/?sort_by", nil)
	assert.Empty(t, getSort(r))

	r = httptest.NewRequest("GET", "/?sort_by=", nil)
	assert.Empty(t, getSort(r))
}

func TestGetSort_SortByWithWhitespaceValue(t *testing.T) {
	r := httptest.NewRequest("GET", "/?sort_by=%20", nil)
	assert.Empty(t, getSort(r))

	r = httptest.NewRequest("GET", "/?sort_by=%20%20", nil)
	assert.Empty(t, getSort(r))
}

func TestGetSort_InvalidValue(t *testing.T) {
	r := httptest.NewRequest("GET", "/?sort_by=,,,", nil)
	assert.Empty(t, getSort(r))

	r = httptest.NewRequest("GET", "/?sort_by=,,%20,,", nil)
	assert.Empty(t, getSort(r))

	r = httptest.NewRequest("GET", "/?sort_by=name,,,,", nil)
	assert.Empty(t, getSort(r))

	r = httptest.NewRequest("GET", "/?sort_by=,,name,,", nil)
	assert.Empty(t, getSort(r))
}

func TestGetSort_WithoutOrderBy(t *testing.T) {
	r := httptest.NewRequest("GET", "/?sort_by=name:", nil)
	assert.Empty(t, getSort(r))

	r = httptest.NewRequest("GET", "/?sort_by=name:,,,,", nil)
	assert.Empty(t, getSort(r))

	r = httptest.NewRequest("GET", "/?sort_by=,,name:,,", nil)
	assert.Empty(t, getSort(r))
}

func TestGetSort_WrongOrderBy(t *testing.T) {
	r := httptest.NewRequest("GET", "/?sort_by=name:abc", nil)
	assert.Empty(t, getSort(r))

	r = httptest.NewRequest("GET", "/?sort_by=name:abc,,,,", nil)
	assert.Empty(t, getSort(r))

	r = httptest.NewRequest("GET", "/?sort_by=,,name:abc,,", nil)
	assert.Empty(t, getSort(r))
}

func TestGetSort_With2Colon(t *testing.T) {
	r := httptest.NewRequest("GET", "/?sort_by=name:asc:", nil)
	assert.Empty(t, getSort(r))

	r = httptest.NewRequest("GET", "/?sort_by=name:asc:,,,,", nil)
	assert.Empty(t, getSort(r))

	r = httptest.NewRequest("GET", "/?sort_by=,,name:asc:,,", nil)
	assert.Empty(t, getSort(r))
}

func TestGetSort_SuccessWithOneField(t *testing.T) {
	r := httptest.NewRequest("GET", "/?sort_by=name:desc", nil)
	sortBy := getSort(r)
	assert.NotEmpty(t, sortBy)
	assert.Len(t, sortBy, 1)
	assert.Equal(t, "name", sortBy[0].Field)
	assert.Equal(t, ORDER_BY_DESC, sortBy[0].OrderBy)

	r = httptest.NewRequest("GET", "/?sort_by=name:desc,,,,", nil)
	sortBy = getSort(r)
	assert.NotEmpty(t, sortBy)
	assert.Len(t, sortBy, 1)
	assert.Equal(t, "name", sortBy[0].Field)
	assert.Equal(t, ORDER_BY_DESC, sortBy[0].OrderBy)

	r = httptest.NewRequest("GET", "/?sort_by=,,name:desc,,", nil)
	sortBy = getSort(r)
	assert.NotEmpty(t, sortBy)
	assert.Len(t, sortBy, 1)
	assert.Equal(t, "name", sortBy[0].Field)
	assert.Equal(t, ORDER_BY_DESC, sortBy[0].OrderBy)
}

func TestGetSort_SuccessWith2Field(t *testing.T) {
	r := httptest.NewRequest("GET", "/?sort_by=name:desc,age:asc", nil)
	sortBy := getSort(r)
	assert.NotEmpty(t, sortBy)
	assert.Len(t, sortBy, 2)
	assert.Equal(t, "name", sortBy[0].Field)
	assert.Equal(t, ORDER_BY_DESC, sortBy[0].OrderBy)
	assert.Equal(t, "age", sortBy[1].Field)
	assert.Equal(t, ORDER_BY_ASC, sortBy[1].OrderBy)

	r = httptest.NewRequest("GET", "/?sort_by=name:desc,age:asc,,,,", nil)
	sortBy = getSort(r)
	assert.NotEmpty(t, sortBy)
	assert.Len(t, sortBy, 2)
	assert.Equal(t, "name", sortBy[0].Field)
	assert.Equal(t, ORDER_BY_DESC, sortBy[0].OrderBy)
	assert.Equal(t, "age", sortBy[1].Field)
	assert.Equal(t, ORDER_BY_ASC, sortBy[1].OrderBy)

	r = httptest.NewRequest("GET", "/?sort_by=,,name:desc,age:asc,,", nil)
	sortBy = getSort(r)
	assert.NotEmpty(t, sortBy)
	assert.Len(t, sortBy, 2)
	assert.Equal(t, "name", sortBy[0].Field)
	assert.Equal(t, ORDER_BY_DESC, sortBy[0].OrderBy)
	assert.Equal(t, "age", sortBy[1].Field)
	assert.Equal(t, ORDER_BY_ASC, sortBy[1].OrderBy)

	r = httptest.NewRequest("GET", "/?sort_by=,,name:desc,,,,age:asc,,", nil)
	sortBy = getSort(r)
	assert.NotEmpty(t, sortBy)
	assert.Len(t, sortBy, 2)
	assert.Equal(t, "name", sortBy[0].Field)
	assert.Equal(t, ORDER_BY_DESC, sortBy[0].OrderBy)
	assert.Equal(t, "age", sortBy[1].Field)
	assert.Equal(t, ORDER_BY_ASC, sortBy[1].OrderBy)
}

func TestQueryString_GetSortSql(t *testing.T) {
	type args struct {
		uriTarget         string
		isNeedSortKeyword bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no sort",
			args: args{
				uriTarget:         "/user?name=eq:micheal",
				isNeedSortKeyword: true,
			},
			want: "",
		},
		{
			name: "sort 1 field, need ORDER BY",
			args: args{
				uriTarget:         "/user?sort_by=name:asc",
				isNeedSortKeyword: true,
			},
			want: " ORDER BY `name` ASC ",
		},
		{
			name: "sort 1 field, dont need ORDER BY",
			args: args{
				uriTarget:         "/user?sort_by=name:asc",
				isNeedSortKeyword: false,
			},
			want: " `name` ASC ",
		},
		{
			name: "sort 3 field, need ORDER BY",
			args: args{
				uriTarget:         "/user?sort_by=name:asc,age:desc,created:asc",
				isNeedSortKeyword: true,
			},
			want: " ORDER BY `name` ASC , `age` DESC , `created` ASC ",
		},
		{
			name: "sort 3 field, dont need ORDER BY",
			args: args{
				uriTarget:         "/user?sort_by=name:asc,age:desc,created:asc",
				isNeedSortKeyword: false,
			},
			want: " `name` ASC , `age` DESC , `created` ASC ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New(httptest.NewRequest("GET", tt.args.uriTarget, nil))
			got := q.GetSortSql(tt.args.isNeedSortKeyword)
			assert.Equal(t, tt.want, got)
		})
	}
}

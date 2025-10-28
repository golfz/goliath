package qrystr

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// **********************************************
// getFilter
// **********************************************
func TestGetFilter_NoFilter(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	filter := getFilter(r)
	assert.NotNil(t, filter)
	assert.Len(t, filter, 0)

	r = httptest.NewRequest("GET", "/?", nil)
	filter = getFilter(r)
	assert.NotNil(t, filter)
	assert.Len(t, filter, 0)
}

func TestGetFilter_SortAndLimit_NoFilter(t *testing.T) {
	r := httptest.NewRequest("GET", "/?sort_by=firstname:asc&limit=50&page=3", nil)
	filter := getFilter(r)

	assert.NotNil(t, filter)
	assert.Len(t, filter, 0)
}

func TestGetFilter_OneFilter(t *testing.T) {
	expected := []FilterField{
		{
			Field: "age",
			Condition: []FilterFieldCondition{
				{
					Operator: "gte",
					Values:   []string{"30"},
				},
			},
		},
	}

	r := httptest.NewRequest("GET", "/?age=gte:30", nil)
	filter := getFilter(r)

	assert.NotNil(t, filter)
	assert.Equal(t, expected, filter)
}

func TestGetFilter_MultiFilter(t *testing.T) {
	expected := []FilterField{
		{
			Field: "name",
			Condition: []FilterFieldCondition{
				{Operator: "like", Values: []string{"mi"}},
			},
		},
		{
			Field: "age",
			Condition: []FilterFieldCondition{
				{Operator: "gte", Values: []string{"30"}},
				{Operator: "lt", Values: []string{"40"}},
			},
		},
		{
			Field: "country",
			Condition: []FilterFieldCondition{
				{Operator: "in", Values: []string{"thailand", "usa"}},
			},
		},
	}

	r := httptest.NewRequest("GET", "/?name=like:mi&age=gte:30,lt:40&country=in:thailand%2Busa", nil)
	filter := getFilter(r)

	assert.NotNil(t, filter)
	assert.Len(t, filter, 3)
	assert.Contains(t, filter, expected[0])
	assert.Contains(t, filter, expected[1])
	assert.Contains(t, filter, expected[2])
}

func TestGetFilter_OrderFieldName_MultiFilter(t *testing.T) {
	expected := []FilterField{
		{
			Field: "age",
			Condition: []FilterFieldCondition{
				{Operator: "gte", Values: []string{"30"}},
				{Operator: "lt", Values: []string{"40"}},
			},
		},
		{
			Field: "country",
			Condition: []FilterFieldCondition{
				{Operator: "in", Values: []string{"thailand", "usa"}},
			},
		},
		{
			Field: "name",
			Condition: []FilterFieldCondition{
				{Operator: "like", Values: []string{"mi"}},
			},
		},
	}

	r := httptest.NewRequest("GET", "/?name=like:mi&age=gte:30,lt:40&country=in:thailand%2Busa", nil)
	got := getFilter(r)

	assert.Equal(t, expected, got)
}

func TestGetFilter_MultiFilter_SomeTagValueInvalid(t *testing.T) {
	expected := []FilterField{
		{
			Field: "name",
			Condition: []FilterFieldCondition{
				{Operator: "like", Values: []string{"mi"}},
			},
		},
		{
			Field: "age",
			Condition: []FilterFieldCondition{
				{Operator: "gte", Values: []string{"30"}},
				{Operator: "lt", Values: []string{"40"}},
			},
		},
	}

	r := httptest.NewRequest("GET", "/?name=like:mi&age=gte:30,lt:40,eq:", nil)
	filter := getFilter(r)

	assert.NotNil(t, filter)
	assert.Len(t, filter, 2)
	assert.Contains(t, filter, expected[0])
	assert.Contains(t, filter, expected[1])
}

func TestGetFilter_MultiFilter_SomeFieldValueInvalid(t *testing.T) {
	expected := []FilterField{
		{
			Field: "name",
			Condition: []FilterFieldCondition{
				{Operator: "like", Values: []string{"mi"}},
			},
		},
	}

	r := httptest.NewRequest("GET", "/?name=like:mi&age=gte:,:::,eq:", nil)
	filter := getFilter(r)

	assert.NotNil(t, filter)
	assert.Equal(t, expected, filter)
}

// **********************************************
// isFilterField
// **********************************************
func TestIsFilterField_SortOrLimit(t *testing.T) {
	assert.False(t, isFilterField("sort_by"))
	assert.False(t, isFilterField("limit"))
	assert.False(t, isFilterField("page"))
}

func TestIsFilterField_NoSortOrLimit_IsFilter(t *testing.T) {
	assert.True(t, isFilterField("some_field"))
	assert.True(t, isFilterField("user_id"))
}

// **********************************************
// getFilterFieldCondition
// **********************************************
func TestGetFilterFieldCondition_BlankString(t *testing.T) {
	_, err := getFilterFieldCondition("")
	assert.Error(t, err)
}

func TestGetFilterFieldCondition_WhiteSpace_AsEq(t *testing.T) {
	cond, err := getFilterFieldCondition(" ")
	assert.Nil(t, err)
	assert.Equal(t, "eq", cond.Operator)
	assert.Len(t, cond.Values, 1)
	assert.Equal(t, " ", cond.Values[0])
}

func TestGetFilterFieldCondition_EqValueWithComma_MustError(t *testing.T) {
	_, err := getFilterFieldCondition("eq:100,200")
	assert.Error(t, err)
}

func TestGetFilterFieldCondition_OnlyColon_MustError(t *testing.T) {
	_, err := getFilterFieldCondition(":")
	assert.Error(t, err)
}

func TestGetFilterFieldCondition_2Colon_MustError(t *testing.T) {
	_, err := getFilterFieldCondition("::")
	assert.Error(t, err)
}

func TestGetFilterFieldCondition_EqWithoutValue_MustError(t *testing.T) {
	_, err := getFilterFieldCondition("eq:")
	assert.Error(t, err)
}

func TestGetFilterFieldCondition_WithoutFilterTag_MustError(t *testing.T) {
	_, err := getFilterFieldCondition(":100")
	assert.Error(t, err)
}

func TestGetFilterFieldCondition_UnknownFilterTag_MustError(t *testing.T) {
	_, err := getFilterFieldCondition("unknown:100")
	assert.Error(t, err)
}

func TestGetFilterFieldCondition_OnlyValueIsEq(t *testing.T) {
	expected := FilterFieldCondition{
		Operator: "eq",
		Values:   []string{"100"},
	}
	cond, err := getFilterFieldCondition("100")
	assert.Nil(t, err)
	assert.Equal(t, expected, cond)
}

func TestGetFilterFieldCondition_In(t *testing.T) {
	expected := FilterFieldCondition{
		Operator: "in",
		Values:   []string{"100", "200", "300"},
	}
	cond, err := getFilterFieldCondition("in:100+200+300")
	assert.Nil(t, err)
	assert.Equal(t, expected, cond)
}

func TestGetFilterFieldCondition_Is_WithUnknown_MustError(t *testing.T) {
	_, err := getFilterFieldCondition("is:unknown")
	assert.Error(t, err)
}

func TestGetFilterFieldCondition_Is_WithBlank_MustError(t *testing.T) {
	_, err := getFilterFieldCondition("is:")
	assert.Error(t, err)
}

func TestGetFilterFieldCondition_Is_NullOrNotnull(t *testing.T) {
	expected := FilterFieldCondition{
		Operator: "is",
		Values:   []string{"null"},
	}
	cond, err := getFilterFieldCondition("is:null")
	assert.Nil(t, err)
	assert.Equal(t, expected, cond)

	expected = FilterFieldCondition{
		Operator: "is",
		Values:   []string{"notnull"},
	}
	cond, err = getFilterFieldCondition("is:notnull")
	assert.Nil(t, err)
	assert.Equal(t, expected, cond)
}

func TestGetFilterFieldCondition_Eq(t *testing.T) {
	expected := FilterFieldCondition{
		Operator: "eq",
		Values:   []string{"100"},
	}
	cond, err := getFilterFieldCondition("eq:100")
	assert.Nil(t, err)
	assert.Equal(t, expected, cond)
}

func TestGetFilterFieldCondition_Not(t *testing.T) {
	expected := FilterFieldCondition{
		Operator: "not",
		Values:   []string{"100"},
	}
	cond, err := getFilterFieldCondition("not:100")
	assert.Nil(t, err)
	assert.Equal(t, expected, cond)
}

func TestQueryString_GetFilterSql(t *testing.T) {
	type args struct {
		uriTarget          string
		isNeedWhereKeyword bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no filter",
			args: args{
				uriTarget:          "/user?sort_by=name:asc",
				isNeedWhereKeyword: true,
			},
			want: "",
		},
		{
			name: "eq without WHERE",
			args: args{
				uriTarget:          "/user?age=eq:25",
				isNeedWhereKeyword: false,
			},
			want: " `age` = '25' ",
		},
		{
			name: "eq with WHERE",
			args: args{
				uriTarget:          "/user?age=eq:25",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `age` = '25' ",
		},
		{
			name: "not with WHERE",
			args: args{
				uriTarget:          "/user?age=not:25",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `age` <> '25' ",
		},
		{
			name: "gt with WHERE",
			args: args{
				uriTarget:          "/user?age=gt:25",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `age` > '25' ",
		},
		{
			name: "gte with WHERE",
			args: args{
				uriTarget:          "/user?age=gte:25",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `age` >= '25' ",
		},
		{
			name: "lt with WHERE",
			args: args{
				uriTarget:          "/user?age=lt:25",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `age` < '25' ",
		},
		{
			name: "lte with WHERE",
			args: args{
				uriTarget:          "/user?age=lte:25",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `age` <= '25' ",
		},
		{
			name: "like no wildcard",
			args: args{
				uriTarget:          "/user?name=like:tom",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `name` LIKE 'tom' ",
		},
		{
			name: "like with wildcard",
			args: args{
				uriTarget:          "/user?name=like:~tom_s~",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `name` LIKE '%tom_s%' ",
		},
		{
			name: "is null",
			args: args{
				uriTarget:          "/user?name=is:null",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `name` IS NULL ",
		},
		{
			name: "is notnull",
			args: args{
				uriTarget:          "/user?name=is:notnull",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `name` IS NOT NULL ",
		},
		{
			name: "in with 1 element",
			args: args{
				uriTarget:          "/user?country=in:thailand",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `country` IN ('thailand') ",
		},
		{
			name: "in with 2 element",
			args: args{
				uriTarget:          "/user?country=in:thailand%2Busa",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `country` IN ('thailand', 'usa') ",
		},
		{
			name: "in with 3 element",
			args: args{
				uriTarget:          "/user?country=in:thailand%2Busa%2Bcanada",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `country` IN ('thailand', 'usa', 'canada') ",
		},
		{
			name: "eq + in with 3 element",
			args: args{
				uriTarget:          "/user?age=eq:25&country=in:thailand%2Busa%2Bcanada",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `age` = '25' AND `country` IN ('thailand', 'usa', 'canada') ",
		},
		{
			name: "eq + is + like + in with 3 element",
			args: args{
				uriTarget:          "/user?age=gte:20,lt:30&org=is:null&com=is:notnull&name=like:~tom_s~&country=in:thailand%2Busa%2Bcanada",
				isNeedWhereKeyword: true,
			},
			want: " WHERE `age` >= '20' AND `age` < '30' AND `com` IS NOT NULL AND `country` IN ('thailand', 'usa', 'canada') AND `name` LIKE '%tom_s%' AND `org` IS NULL ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := New(httptest.NewRequest("GET", tt.args.uriTarget, nil))
			got := q.GetFilterSql(tt.args.isNeedWhereKeyword)
			assert.Equal(t, tt.want, got)
		})
	}
}

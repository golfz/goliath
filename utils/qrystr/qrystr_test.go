package qrystr

import (
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
	"mastertime-service/utils/sliceutils"
	"net/http/httptest"
	"testing"
)

func Test_GetStandardLimitAndSortFieldList(t *testing.T) {
	list := getStandardLimitAndSortFieldList()

	assert.Contains(t, list, QUERY_STRING_FIELDNAME_LIMIT)
	assert.Contains(t, list, QUERY_STRING_FIELDNAME_PAGE)
	assert.Contains(t, list, QUERY_STRING_FIELDNAME_SORT_BY)
	//assert.Contains(t, list, QUERY_STRING_FIELDNAME_MARK)
}

func Test_New_PageWithoutLimit_ExpectedAllLimitFieldNil(t *testing.T) {
	target := "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=3"
	r := httptest.NewRequest("GET", target, nil)
	qryCondition := New(r)

	expected := QueryString{
		RequestPath: "/users",
		Filter: []FilterField{
			{
				Field: "country",
				Condition: []FilterFieldCondition{
					{
						Operator: "in",
						Values:   []string{"thailand", "usa"},
					},
				},
			},
		},
		Sort: []SortBy{
			{
				Field:   "firstname",
				OrderBy: "asc",
			},
			{
				Field:   "age",
				OrderBy: "desc",
			},
		},
		Limit:      nil,
		Page:       nil,
		ResultSize: 0,
		TotalSize:  0,
		ResultPage: 0,
		TotalPage:  0,
	}

	assert.Equal(t, expected, qryCondition)
}

func Test_New_LimitWithoutPage(t *testing.T) {
	target := "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&limit=50"
	r := httptest.NewRequest("GET", target, nil)
	qryCondition := New(r)

	limit := 50
	expected := QueryString{
		RequestPath: "/users",
		Filter: []FilterField{
			{
				Field: "country",
				Condition: []FilterFieldCondition{
					{
						Operator: "in",
						Values:   []string{"thailand", "usa"},
					},
				},
			},
		},
		Sort: []SortBy{
			{
				Field:   "firstname",
				OrderBy: "asc",
			},
			{
				Field:   "age",
				OrderBy: "desc",
			},
		},
		Limit:      &limit,
		Page:       nil,
		ResultSize: 0,
		TotalSize:  0,
		ResultPage: 0,
		TotalPage:  0,
	}

	assert.Equal(t, expected, qryCondition)
}

func Test_New_LimitWithPage(t *testing.T) {
	target := "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=3&limit=50"
	r := httptest.NewRequest("GET", target, nil)
	qryCondition := New(r)

	page := 3
	limit := 50
	expected := QueryString{
		RequestPath: "/users",
		Filter: []FilterField{
			{
				Field: "country",
				Condition: []FilterFieldCondition{
					{
						Operator: "in",
						Values:   []string{"thailand", "usa"},
					},
				},
			},
		},
		Sort: []SortBy{
			{
				Field:   "firstname",
				OrderBy: "asc",
			},
			{
				Field:   "age",
				OrderBy: "desc",
			},
		},
		Limit:      &limit,
		Page:       &page,
		ResultSize: 0,
		TotalSize:  0,
		ResultPage: 0,
		TotalPage:  0,
	}

	assert.Equal(t, expected, qryCondition)
}

// **********************************************
// GetCollectionResponse
// **********************************************

type sample struct {
	Id        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	TelNumber *string `json:"tel_number"`
}

func Test_GetCollectionResponse_EveryField(t *testing.T) {
	target := "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=3&limit=10"
	r := httptest.NewRequest("GET", target, nil)
	q := New(r)

	getOnlyInterface := func(s []interface{}, ok bool) []interface{} {
		return s
	}

	data := getOnlyInterface(sliceutils.TakeSliceArg(make([]sample, 100)))
	for i, _ := range data {
		data[i] = sample{
			Id:        i + 1,
			FirstName: "",
			LastName:  "",
			TelNumber: nil,
		}
	}
	colRes, err := q.GetCollectionResponse(data)

	assert.Nil(t, err)
	assert.NotNil(t, colRes)

	var (
		limit      = 10
		page       = 3
		outputData = data[20:30]
	)

	expected := CollectionResponse{
		Size:      10,
		TotalSize: 100,
		Page:      3,
		TotalPage: 10,
		Filter: []FilterField{
			{
				Field: "country",
				Condition: []FilterFieldCondition{
					{
						Operator: "in",
						Values:   []string{"thailand", "usa"},
					},
				},
			},
		},
		SortBy: []SortBy{
			{
				Field:   "firstname",
				OrderBy: "asc",
			},
			{
				Field:   "age",
				OrderBy: "desc",
			},
		},
		Limit: Limit{
			Limit: &limit,
			Page:  &page,
		},
		Link: Link{
			Self:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=3&limit=10",
			First:    "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=1&limit=10",
			Last:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=10&limit=10",
			Previous: pointy.String("/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=2&limit=10"),
			Next:     pointy.String("/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=4&limit=10"),
		},
		Data: outputData,
	}

	assert.Equal(t, expected, colRes)
}

func Test_GetCollectionResponse_EveryField_1stPage_ExpectNoPrevious(t *testing.T) {
	target := "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=1&limit=10"
	r := httptest.NewRequest("GET", target, nil)
	q := New(r)

	getOnlyInterface := func(s []interface{}, ok bool) []interface{} {
		return s
	}

	data := getOnlyInterface(sliceutils.TakeSliceArg(make([]sample, 100)))
	for i, _ := range data {
		data[i] = sample{
			Id:        i + 1,
			FirstName: "",
			LastName:  "",
			TelNumber: nil,
		}
	}
	colRes, err := q.GetCollectionResponse(data)

	assert.Nil(t, err)
	assert.NotNil(t, colRes)

	var (
		limit      = 10
		page       = 1
		outputData = data[0:10]
	)

	expected := CollectionResponse{
		Size:      10,
		TotalSize: 100,
		Page:      1,
		TotalPage: 10,
		Filter: []FilterField{
			{
				Field: "country",
				Condition: []FilterFieldCondition{
					{
						Operator: "in",
						Values:   []string{"thailand", "usa"},
					},
				},
			},
		},
		SortBy: []SortBy{
			{
				Field:   "firstname",
				OrderBy: "asc",
			},
			{
				Field:   "age",
				OrderBy: "desc",
			},
		},
		Limit: Limit{
			Limit: &limit,
			Page:  &page,
		},
		Link: Link{
			Self:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=1&limit=10",
			First:    "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=1&limit=10",
			Last:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=10&limit=10",
			Previous: nil,
			Next:     pointy.String("/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=2&limit=10"),
		},
		Data: outputData,
	}

	assert.Equal(t, expected, colRes)
}

func Test_GetCollectionResponse_EveryField_LatestPage_ExpectNoPrevious(t *testing.T) {
	target := "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=10&limit=10"
	r := httptest.NewRequest("GET", target, nil)
	q := New(r)

	getOnlyInterface := func(s []interface{}, ok bool) []interface{} {
		return s
	}

	data := getOnlyInterface(sliceutils.TakeSliceArg(make([]sample, 100)))
	for i, _ := range data {
		data[i] = sample{
			Id:        i + 1,
			FirstName: "",
			LastName:  "",
			TelNumber: nil,
		}
	}
	colRes, err := q.GetCollectionResponse(data)

	assert.Nil(t, err)
	assert.NotNil(t, colRes)

	var (
		limit      = 10
		page       = 10
		outputData = data[90:100]
	)

	expected := CollectionResponse{
		Size:      10,
		TotalSize: 100,
		Page:      10,
		TotalPage: 10,
		Filter: []FilterField{
			{
				Field: "country",
				Condition: []FilterFieldCondition{
					{
						Operator: "in",
						Values:   []string{"thailand", "usa"},
					},
				},
			},
		},
		SortBy: []SortBy{
			{
				Field:   "firstname",
				OrderBy: "asc",
			},
			{
				Field:   "age",
				OrderBy: "desc",
			},
		},
		Limit: Limit{
			Limit: &limit,
			Page:  &page,
		},
		Link: Link{
			Self:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=10&limit=10",
			First:    "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=1&limit=10",
			Last:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=10&limit=10",
			Previous: pointy.String("/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=9&limit=10"),
			Next:     nil,
		},
		Data: outputData,
	}

	assert.Equal(t, expected, colRes)
}

func Test_GetCollectionResponse_EveryField_LimitWithoutPage_Expected1stPage(t *testing.T) {
	target := "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&limit=10"
	r := httptest.NewRequest("GET", target, nil)
	q := New(r)

	getOnlyInterface := func(s []interface{}, ok bool) []interface{} {
		return s
	}

	data := getOnlyInterface(sliceutils.TakeSliceArg(make([]sample, 100)))
	for i, _ := range data {
		data[i] = sample{
			Id:        i + 1,
			FirstName: "",
			LastName:  "",
			TelNumber: nil,
		}
	}
	colRes, err := q.GetCollectionResponse(data)

	assert.Nil(t, err)
	assert.NotNil(t, colRes)

	var (
		limit      = 10
		outputData = data[0:10]
	)

	expected := CollectionResponse{
		Size:      10,
		TotalSize: 100,
		Page:      1,
		TotalPage: 10,
		Filter: []FilterField{
			{
				Field: "country",
				Condition: []FilterFieldCondition{
					{
						Operator: "in",
						Values:   []string{"thailand", "usa"},
					},
				},
			},
		},
		SortBy: []SortBy{
			{
				Field:   "firstname",
				OrderBy: "asc",
			},
			{
				Field:   "age",
				OrderBy: "desc",
			},
		},
		Limit: Limit{
			Limit: &limit,
			Page:  nil,
		},
		Link: Link{
			Self:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=1&limit=10",
			First:    "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=1&limit=10",
			Last:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=10&limit=10",
			Previous: nil,
			Next:     pointy.String("/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=2&limit=10"),
		},
		Data: outputData,
	}

	assert.Equal(t, expected, colRes)
}

func Test_GetCollectionResponse_EveryField_PageWithoutLimit_IgnorePage_GetAll(t *testing.T) {
	target := "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=3"
	r := httptest.NewRequest("GET", target, nil)
	q := New(r)

	getOnlyInterface := func(s []interface{}, ok bool) []interface{} {
		return s
	}

	data := getOnlyInterface(sliceutils.TakeSliceArg(make([]sample, 100)))
	for i, _ := range data {
		data[i] = sample{
			Id:        i + 1,
			FirstName: "",
			LastName:  "",
			TelNumber: nil,
		}
	}
	colRes, err := q.GetCollectionResponse(data)

	assert.Nil(t, err)
	assert.NotNil(t, colRes)

	var (
		outputData = data
	)

	expected := CollectionResponse{
		Size:      100,
		TotalSize: 100,
		Page:      1,
		TotalPage: 1,
		Filter: []FilterField{
			{
				Field: "country",
				Condition: []FilterFieldCondition{
					{
						Operator: "in",
						Values:   []string{"thailand", "usa"},
					},
				},
			},
		},
		SortBy: []SortBy{
			{
				Field:   "firstname",
				OrderBy: "asc",
			},
			{
				Field:   "age",
				OrderBy: "desc",
			},
		},
		Limit: Limit{
			Limit: nil,
			Page:  nil,
		},
		Link: Link{
			Self:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc",
			First:    "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc",
			Last:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc",
			Previous: nil,
			Next:     nil,
		},
		Data: outputData,
	}

	assert.Equal(t, expected, colRes)
}

func Test_calculateTotalPage(t *testing.T) {
	type args struct {
		totalSize int
		limit     *int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "totalPage < 0, should error",
			args: args{
				totalSize: -1,
				limit:     nil,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "totalPage = 0, totalPage should be 1",
			args: args{
				totalSize: 0,
				limit:     pointy.Int(100),
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "limit = nil, totalPage should be 1",
			args: args{
				totalSize: 10,
				limit:     nil,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "limit = 0, should error",
			args: args{
				totalSize: 10,
				limit:     pointy.Int(0),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "limit < 0, should error",
			args: args{
				totalSize: 10,
				limit:     pointy.Int(-5),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "total=1010, limit=10, expect=101",
			args: args{
				totalSize: 1010,
				limit:     pointy.Int(10),
			},
			want:    101,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateTotalPage(tt.args.totalSize, tt.args.limit)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_calculateResultPage(t *testing.T) {
	type args struct {
		totalPage int
		page      *int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "totalPage < 0, should error",
			args: args{
				totalPage: -10,
				page:      pointy.Int(3),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "totalPage = 0, should error",
			args: args{
				totalPage: 0,
				page:      pointy.Int(3),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "page = nil, expected 1",
			args: args{
				totalPage: 1,
				page:      nil,
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "page < 1, expected error",
			args: args{
				totalPage: 1,
				page:      pointy.Int(0),
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "page > totalPage, expected page = totalPage",
			args: args{
				totalPage: 10,
				page:      pointy.Int(100),
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "page == totalPage, expected page = page",
			args: args{
				totalPage: 100,
				page:      pointy.Int(100),
			},
			want:    100,
			wantErr: false,
		},
		{
			name: "page < totalPage, expected page = page",
			args: args{
				totalPage: 100,
				page:      pointy.Int(10),
			},
			want:    10,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateResultPage(tt.args.totalPage, tt.args.page)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_reCalculateRequestPage(t *testing.T) {
	type args struct {
		totalPage int
		limit     *int
		page      *int
	}
	tests := []struct {
		name    string
		args    args
		want    *int
		wantErr bool
	}{
		{
			name: "totalPage < 1, expected error",
			args: args{
				totalPage: 0,
				limit:     nil,
				page:      nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "limit = nil, expected nil",
			args: args{
				totalPage: 10,
				limit:     nil,
				page:      nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "page = nil, expected nil",
			args: args{
				totalPage: 10,
				limit:     pointy.Int(50),
				page:      nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "page > totalPage, expected page = totalPage",
			args: args{
				totalPage: 10,
				limit:     pointy.Int(50),
				page:      pointy.Int(12),
			},
			want:    pointy.Int(10),
			wantErr: false,
		},
		{
			name: "page <= totalPage, expected page",
			args: args{
				totalPage: 10,
				limit:     pointy.Int(50),
				page:      pointy.Int(3),
			},
			want:    pointy.Int(3),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reCalculateRequestPage(tt.args.totalPage, tt.args.limit, tt.args.page)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})

	}
}

func TestApplyLimit(t *testing.T) {
	getOnlySlice := func(s []interface{}, ok bool) []interface{} {
		return s
	}
	type args struct {
		data  []interface{}
		limit *int
		page  *int
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr bool
	}{
		{
			name: "limit = nil, should get whole data",
			args: args{
				data:  getOnlySlice(sliceutils.TakeSliceArg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})),
				limit: nil,
				page:  pointy.Int(4),
			},
			want:    getOnlySlice(sliceutils.TakeSliceArg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})),
			wantErr: false,
		},
		{
			name: "limit = 0, should error",
			args: args{
				data:  getOnlySlice(sliceutils.TakeSliceArg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})),
				limit: pointy.Int(0),
				page:  pointy.Int(4),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "limit < 0, should error",
			args: args{
				data:  getOnlySlice(sliceutils.TakeSliceArg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})),
				limit: pointy.Int(-5),
				page:  pointy.Int(4),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "page == nil, expected first page (accord to limit)",
			args: args{
				data:  getOnlySlice(sliceutils.TakeSliceArg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})),
				limit: pointy.Int(4),
				page:  pointy.Int(-3),
			},
			want:    getOnlySlice(sliceutils.TakeSliceArg([]int{1, 2, 3, 4})),
			wantErr: false,
		},
		{
			name: "page < 1, expected error",
			args: args{
				data:  getOnlySlice(sliceutils.TakeSliceArg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})),
				limit: pointy.Int(4),
				page:  pointy.Int(0),
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "page = 1, expected trim from No. 1",
			args: args{
				data:  getOnlySlice(sliceutils.TakeSliceArg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})),
				limit: pointy.Int(4),
				page:  pointy.Int(1),
			},
			want:    getOnlySlice(sliceutils.TakeSliceArg([]int{1, 2, 3, 4})),
			wantErr: false,
		},
		{
			name: "page = 3, expected trim from No. 9",
			args: args{
				data:  getOnlySlice(sliceutils.TakeSliceArg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})),
				limit: pointy.Int(4),
				page:  pointy.Int(3),
			},
			want:    getOnlySlice(sliceutils.TakeSliceArg([]int{9, 10, 11, 12})),
			wantErr: false,
		},
		{
			name: "page > totalPage(3), expected empty",
			args: args{
				data:  getOnlySlice(sliceutils.TakeSliceArg([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})),
				limit: pointy.Int(4),
				page:  pointy.Int(7),
			},
			want:    getOnlySlice(sliceutils.TakeSliceArg([]int{})),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := applyLimit(tt.args.data, tt.args.limit, tt.args.page)
			if (err != nil) != tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getOffset(t *testing.T) {
	type args struct {
		page  *int
		limit *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "limit == nil, expected 0",
			args: args{
				page:  pointy.Int(3),
				limit: nil,
			},
			want: 0,
		},
		{
			name: "limit < 1, expected 0",
			args: args{
				page:  pointy.Int(3),
				limit: pointy.Int(0),
			},
			want: 0,
		},
		{
			name: "page == nil, expected 0",
			args: args{
				page:  nil,
				limit: pointy.Int(50),
			},
			want: 0,
		},
		{
			name: "page < 1, expected 0",
			args: args{
				page:  pointy.Int(0),
				limit: pointy.Int(50),
			},
			want: 0,
		},
		{
			name: "page = 1, limit 50, expected 0",
			args: args{
				page:  pointy.Int(1),
				limit: pointy.Int(50),
			},
			want: 0,
		},
		{
			name: "page = 3, limit 50, expected 100",
			args: args{
				page:  pointy.Int(3),
				limit: pointy.Int(50),
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getOffset(tt.args.page, tt.args.limit)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getLink(t *testing.T) {
	type args struct {
		target     string
		totalSize  int
		resultSize int
		totalPage  int
		resultPage int
	}
	tests := []struct {
		name string
		args args
		want Link
	}{
		{
			name: "all field, page 3, expected has all link",
			args: args{
				target:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=3&limit=10",
				totalSize:  100,
				resultSize: 10,
				totalPage:  10,
				resultPage: 3,
			},
			want: Link{
				Self:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=3&limit=10",
				First:    "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=1&limit=10",
				Last:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=10&limit=10",
				Previous: pointy.String("/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=2&limit=10"),
				Next:     pointy.String("/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=4&limit=10"),
			},
		},
		{
			name: "has filter, no sort, page 3, expected has all link",
			args: args{
				target:     "/users?country=in:thailand%2Busa&page=3&limit=10",
				totalSize:  100,
				resultSize: 10,
				totalPage:  10,
				resultPage: 3,
			},
			want: Link{
				Self:     "/users?country=in:thailand%2Busa&page=3&limit=10",
				First:    "/users?country=in:thailand%2Busa&page=1&limit=10",
				Last:     "/users?country=in:thailand%2Busa&page=10&limit=10",
				Previous: pointy.String("/users?country=in:thailand%2Busa&page=2&limit=10"),
				Next:     pointy.String("/users?country=in:thailand%2Busa&page=4&limit=10"),
			},
		},
		{
			name: "no filter, has sort, page 3, expected has all link",
			args: args{
				target:     "/users?sort_by=firstname:asc,age:desc&page=3&limit=10",
				totalSize:  100,
				resultSize: 10,
				totalPage:  10,
				resultPage: 3,
			},
			want: Link{
				Self:     "/users?sort_by=firstname:asc,age:desc&page=3&limit=10",
				First:    "/users?sort_by=firstname:asc,age:desc&page=1&limit=10",
				Last:     "/users?sort_by=firstname:asc,age:desc&page=10&limit=10",
				Previous: pointy.String("/users?sort_by=firstname:asc,age:desc&page=2&limit=10"),
				Next:     pointy.String("/users?sort_by=firstname:asc,age:desc&page=4&limit=10"),
			},
		},
		{
			name: "no filter, no sort, page 3, expected has all link",
			args: args{
				target:     "/users?page=3&limit=10",
				totalSize:  100,
				resultSize: 10,
				totalPage:  10,
				resultPage: 3,
			},
			want: Link{
				Self:     "/users?page=3&limit=10",
				First:    "/users?page=1&limit=10",
				Last:     "/users?page=10&limit=10",
				Previous: pointy.String("/users?page=2&limit=10"),
				Next:     pointy.String("/users?page=4&limit=10"),
			},
		},
		{
			name: "all field, on 1st page, expected previous=nil",
			args: args{
				target:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=1&limit=10",
				totalSize:  100,
				resultSize: 10,
				totalPage:  10,
				resultPage: 1,
			},
			want: Link{
				Self:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=1&limit=10",
				First:    "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=1&limit=10",
				Last:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=10&limit=10",
				Previous: nil,
				Next:     pointy.String("/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=2&limit=10"),
			},
		},
		{
			name: "all field, on latest page, expected next=nil",
			args: args{
				target:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=10&limit=10",
				totalSize:  100,
				resultSize: 10,
				totalPage:  10,
				resultPage: 10,
			},
			want: Link{
				Self:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=10&limit=10",
				First:    "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=1&limit=10",
				Last:     "/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=10&limit=10",
				Previous: pointy.String("/users?country=in:thailand%2Busa&sort_by=firstname:asc,age:desc&page=9&limit=10"),
				Next:     nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qObj := New(httptest.NewRequest("GET", tt.args.target, nil))
			qObj.TotalSize = tt.args.totalSize
			qObj.TotalPage = tt.args.totalPage
			qObj.ResultSize = tt.args.resultSize
			qObj.ResultPage = tt.args.resultPage
			got := getLink(qObj)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getFilterQryString(t *testing.T) {
	type args struct {
		q QueryString
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				q: QueryString{
					Filter: []FilterField{
						{
							Field: "org",
							Condition: []FilterFieldCondition{
								{
									Operator: "eq",
									Values:   []string{"r&d"},
								},
							},
						},
						{
							Field: "color",
							Condition: []FilterFieldCondition{
								{
									Operator: "not",
									Values:   []string{"red"},
								},
								{
									Operator: "not",
									Values:   []string{"blue"},
								},
							},
						},
						{
							Field: "country",
							Condition: []FilterFieldCondition{
								{
									Operator: "in",
									Values:   []string{"thailand", "usa", "canada"},
								},
							},
						},
						{
							Field: "name",
							Condition: []FilterFieldCondition{
								{
									Operator: "like",
									Values:   []string{"~a_d~"},
								},
							},
						},
					},
				},
			},
			want: "org=eq:r%26d&color=not:red,not:blue&country=in:thailand%2Busa%2Bcanada&name=like:~a_d~",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getFilterQryString(tt.args.q)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getSortQryString(t *testing.T) {
	type args struct {
		q QueryString
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{q: QueryString{
				Sort: []SortBy{
					{
						Field:   "name",
						OrderBy: "asc",
					},
					{
						Field:   "age",
						OrderBy: "desc",
					},
				},
			}},
			want: "sort_by=name:asc,age:desc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSortQryString(tt.args.q)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getLinkSelfQryString(t *testing.T) {
	type args struct {
		q QueryString
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "limit=nil, page=3, expect empty",
			args: args{q: QueryString{
				Limit:      nil,
				ResultPage: 3,
			}},
			want: "",
		},
		{
			name: "limit=50, page=3, expect page=3&limit=50",
			args: args{q: QueryString{
				Limit:      pointy.Int(50),
				ResultPage: 3,
			}},
			want: "page=3&limit=50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLinkSelfQryString(tt.args.q)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getLinkFirstQryString(t *testing.T) {
	type args struct {
		q QueryString
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "limit=nil, expect empty",
			args: args{
				q: QueryString{
					Limit: nil,
				},
			},
			want: "",
		},
		{
			name: "limit=50, expect page=1&limit=50",
			args: args{
				q: QueryString{
					Limit: pointy.Int(50),
				},
			},
			want: "page=1&limit=50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLinkFirstQryString(tt.args.q)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getLinkLastQryString(t *testing.T) {
	type args struct {
		q QueryString
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "limit=nil, expect empty",
			args: args{
				q: QueryString{
					Limit:     nil,
					TotalPage: 10,
				},
			},
			want: "",
		},
		{
			name: "limit=50, expect page=10&limit=50",
			args: args{
				q: QueryString{
					Limit:     pointy.Int(50),
					TotalPage: 10,
				},
			},
			want: "page=10&limit=50",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLinkLastQryString(tt.args.q)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getLinkPreviousQryString(t *testing.T) {
	type args struct {
		q QueryString
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "limit=nil, expect nil",
			args: args{
				q: QueryString{
					Limit: nil,
				},
			},
			want: nil,
		},
		{
			name: "page=1, expect nil",
			args: args{
				q: QueryString{
					ResultPage: 1,
					Limit:      pointy.Int(50),
				},
			},
			want: nil,
		},
		{
			name: "page=3, expect page=2&limit=50",
			args: args{
				q: QueryString{
					ResultPage: 3,
					Limit:      pointy.Int(50),
				},
			},
			want: pointy.String("page=2&limit=50"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLinkPreviousQryString(tt.args.q)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getLinkNextQryString(t *testing.T) {
	type args struct {
		q QueryString
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "limit=nil, expect nil",
			args: args{
				q: QueryString{
					Limit: nil,
				},
			},
			want: nil,
		},
		{
			name: "page == totalPage, expect nil",
			args: args{
				q: QueryString{
					ResultPage: 10,
					TotalPage:  10,
					Limit:      pointy.Int(50),
				},
			},
			want: nil,
		},
		{
			name: "page=3, totalPage=10 expect page=4&limit=50",
			args: args{
				q: QueryString{
					ResultPage: 3,
					TotalPage:  10,
					Limit:      pointy.Int(50),
				},
			},
			want: pointy.String("page=4&limit=50"),
		},
		{
			name: "page=9, totalPage=10 expect page=10&limit=50",
			args: args{
				q: QueryString{
					ResultPage: 9,
					TotalPage:  10,
					Limit:      pointy.Int(50),
				},
			},
			want: pointy.String("page=10&limit=50"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLinkNextQryString(tt.args.q)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_combineQryString(t *testing.T) {
	type args struct {
		requestPath   string
		filterAndSort string
		linkToPage    *string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "linkToPage = nil, expected nil",
			args: args{
				requestPath:   "",
				filterAndSort: "",
				linkToPage:    nil,
			},
			want: nil,
		},
		{
			name: "linkToPage = empty string, expected /requestPath?filterAndSort",
			args: args{
				requestPath:   "/requestPath",
				filterAndSort: "filterAndSort",
				linkToPage:    pointy.String(""),
			},
			want: pointy.String("/requestPath?filterAndSort"),
		},
		{
			name: "linkToPage=linkToPage, filterAndSort=empty string , expected /requestPath?linkToPage",
			args: args{
				requestPath:   "/requestPath",
				filterAndSort: "",
				linkToPage:    pointy.String("linkToPage"),
			},
			want: pointy.String("/requestPath?linkToPage"),
		},
		{
			name: "linkToPage=linkToPage, filterAndSort=filterAndSort , expected /requestPath?filterAndSort&linkToPage",
			args: args{
				requestPath:   "/requestPath",
				filterAndSort: "filterAndSort",
				linkToPage:    pointy.String("linkToPage"),
			},
			want: pointy.String("/requestPath?filterAndSort&linkToPage"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combineQryString(tt.args.requestPath, tt.args.filterAndSort, tt.args.linkToPage)
			assert.Equal(t, tt.want, got)
		})
	}
}

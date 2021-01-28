package qrystr

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

// **********************************************
// Limit
// **********************************************
func TestGetLimit_NoLimit(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	assert.Nil(t, getLimit(r))

	r = httptest.NewRequest("GET", "/?", nil)
	assert.Nil(t, getLimit(r))

	r = httptest.NewRequest("GET", "/?other=10", nil)
	assert.Nil(t, getLimit(r))
}

func TestGetLimit_NoLimitValue(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit", nil)
	assert.Nil(t, getLimit(r))

	r = httptest.NewRequest("GET", "/?limit=", nil)
	assert.Nil(t, getLimit(r))
}

func TestGetLimit_WhitespaceValue(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=%20", nil)
	assert.Nil(t, getLimit(r))

	r = httptest.NewRequest("GET", "/?limit=%20%20", nil)
	assert.Nil(t, getLimit(r))
}

func TestGetLimit_StringValue(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=abc", nil)
	assert.Nil(t, getLimit(r))
}

func TestGetLimit_FloatValue(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=12.5", nil)
	assert.Nil(t, getLimit(r))
}

func TestGetLimit_NegativeInt(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=-10", nil)
	assert.Nil(t, getLimit(r))

	r = httptest.NewRequest("GET", "/?limit=-12.5", nil)
	assert.Nil(t, getLimit(r))
}

func TestGetLimit_Zero(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=0", nil)
	assert.Nil(t, getLimit(r))
}

func TestGetLimit_MultiLimit(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=100&limit=10", nil)
	limit := getLimit(r)
	assert.NotNil(t, limit)
	assert.Equal(t, 100, *limit)
}

func TestGetLimit_Success(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=100", nil)
	limit := getLimit(r)
	assert.NotNil(t, limit)
	assert.Equal(t, 100, *limit)
}

// **********************************************
// Page
// **********************************************
func TestGetPage_NoPage(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	assert.Nil(t, getPage(r))

	r = httptest.NewRequest("GET", "/?", nil)
	assert.Nil(t, getPage(r))

	r = httptest.NewRequest("GET", "/?other=101", nil)
	assert.Nil(t, getPage(r))
}

func TestGetPage_NoPageValue(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=100&page", nil)
	assert.Nil(t, getPage(r))

	r = httptest.NewRequest("GET", "/?limit=100&page=", nil)
	assert.Nil(t, getPage(r))
}

func TestGetPage_NoLimit_IgnorePage(t *testing.T) {
	r := httptest.NewRequest("GET", "/?page=3", nil)
	assert.Nil(t, getPage(r))
}

func TestGetPage_ZeroLimit_IgnorePage(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=0&page=3", nil)
	assert.Nil(t, getPage(r))
}

func TestGetPage_NegativeLimit_IgnorePage(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=-50&page=3", nil)
	assert.Nil(t, getPage(r))
}

func TestGetPage_WhitespaceValue(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=100&page=%20", nil)
	assert.Nil(t, getPage(r))

	r = httptest.NewRequest("GET", "/?limit=100&page=%20%20", nil)
	assert.Nil(t, getPage(r))
}

func TestGetPage_StringValue(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=100&page=abc", nil)
	assert.Nil(t, getPage(r))
}

func TestGetPage_FloatValue(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=100&page=12.5", nil)
	assert.Nil(t, getPage(r))
}

func TestGetPage_NegativeValue(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=100&page=-10", nil)
	assert.Nil(t, getPage(r))

	r = httptest.NewRequest("GET", "/?limit=100&page=-12.5", nil)
	assert.Nil(t, getPage(r))
}

func TestGetPage_Zero(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=100&page=0", nil)
	assert.Nil(t, getPage(r))
}

func TestGetPage_MultiPage(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=100&page=3&page=4", nil)
	page := getPage(r)
	assert.NotNil(t, page)
	assert.Equal(t, 3, *page)
}

func TestGetPage_Success(t *testing.T) {
	r := httptest.NewRequest("GET", "/?limit=100&page=3", nil)
	page := getPage(r)
	assert.NotNil(t, page)
	assert.Equal(t, 3, *page)
}

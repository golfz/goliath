package qrystr

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

// **********************************************
// hasKey
// **********************************************
func TestHasKey_NoAnyKey(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	assert.False(t, hasKey(r, "some_key"))
}

func TestHasKey_NoExpectedKey(t *testing.T) {
	r := httptest.NewRequest("GET", "/?other_key=1234", nil)
	assert.False(t, hasKey(r, "expected_key"))
}

func TestHasKey_ExpectedKeyWithoutValueIsNoKey(t *testing.T) {
	r := httptest.NewRequest("GET", "/?other_key=1234&expected_key", nil)
	assert.False(t, hasKey(r, "expected_key"))

	r = httptest.NewRequest("GET", "/?other_key=1234&expected_key=", nil)
	assert.False(t, hasKey(r, "expected_key"))
}

func TestHasKey_ThereIsExpectedKey(t *testing.T) {
	r := httptest.NewRequest("GET", "/?other_key=1234&expected_key=foo", nil)
	assert.True(t, hasKey(r, "expected_key"))

	r = httptest.NewRequest("GET", "/?other_key=1234&expected_key=foo&expected_key=bar", nil)
	assert.True(t, hasKey(r, "expected_key"))
}

// **********************************************
// keyToIntPtr
// **********************************************
func TestKeyToIntPtr_NoKey(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	assert.Nil(t, keyToIntPtr(r, "key1"))

	r = httptest.NewRequest("GET", "/?", nil)
	assert.Nil(t, keyToIntPtr(r, "key1"))
}

func TestKeyToIntPtr_KeyWithoutValue(t *testing.T) {
	r := httptest.NewRequest("GET", "/?key1", nil)
	assert.Nil(t, keyToIntPtr(r, "key1"))

	r = httptest.NewRequest("GET", "/?key1=", nil)
	assert.Nil(t, keyToIntPtr(r, "key1"))
}

func TestKeyToIntPtr_NoExpectedKey(t *testing.T) {
	r := httptest.NewRequest("GET", "/?key1=10", nil)
	assert.Nil(t, keyToIntPtr(r, "key2"))

	r = httptest.NewRequest("GET", "/?key1=10&key3=20", nil)
	assert.Nil(t, keyToIntPtr(r, "key2"))
}

func TestKeyToIntPtr_Success(t *testing.T) {
	r := httptest.NewRequest("GET", "/?key1=10", nil)
	key := keyToIntPtr(r, "key1")
	assert.NotNil(t, key)
	assert.Equal(t, 10, *key)
}

// **********************************************
// hasStringInSlice
// **********************************************
func TestHasStringInSlice_NotFound(t *testing.T) {
	arr := []string{"abc", "def", "ghi", "jkl"}
	i, has := hasStringInSlice(arr, "xyz")
	assert.False(t, has)
	assert.Equal(t, -1, i)
}

func TestHasStringInSlice_Found(t *testing.T) {
	arr := []string{"abc", "def", "ghi", "jkl"}
	i, has := hasStringInSlice(arr, "abc")
	assert.True(t, has)
	assert.Equal(t, 0, i)

	i, has = hasStringInSlice(arr, "jkl")
	assert.True(t, has)
	assert.Equal(t, 3, i)
}

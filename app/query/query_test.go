package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testSlice = []string{"hello", "world", "foobar"}
var predicateFunc = func(e string) bool {
	if e == "foobar" {
		return false
	} else {
		return true
	}
}

var predicateIndexFunc = func(e string) bool {
	if e == "foobar" {
		return true
	} else {
		return false
	}
}

func TestFilterSlice(t *testing.T) {

	resultSlice := Filter(testSlice, predicateFunc)
	assert.Equal(t, 2, len(resultSlice))

}

func TestFilterSlice_when_nil_is_passed(t *testing.T) {

	testSlice := Filter(nil, predicateFunc)
	assert.Equal(t, 0, len(testSlice))

}

func TestRemoveElementFromSlice(t *testing.T) {
	resultSlice := RemoveElement(testSlice, predicateIndexFunc)
	assert.Equal(t, 2, len(resultSlice))
}

func TestRemoveElementFromSliceElementNotFound(t *testing.T) {
	resultSlice := RemoveElement([]string{"hello", "world"}, predicateIndexFunc)
	assert.Equal(t, 2, len(resultSlice))
}

func TestRemoveElementFromSliceISNil(t *testing.T) {
	resultSlice := RemoveElement(nil, predicateIndexFunc)
	assert.Equal(t, 0, len(resultSlice))
}

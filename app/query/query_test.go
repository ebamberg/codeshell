package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterSlice(t *testing.T) {

	testSlice := []string{"hello", "world", "foobar"}
	testSlice = Filter(testSlice, func(e string) bool {
		if e == "foobar" {
			return false
		} else {
			return true
		}
	})
	assert.Equal(t, 2, len(testSlice))

}

func TestFilterSlice_when_nil_is_passed(t *testing.T) {

	testSlice := Filter(nil, func(e string) bool {
		if e == "foobar" {
			return false
		} else {
			return true
		}
	})
	assert.Equal(t, 0, len(testSlice))

}

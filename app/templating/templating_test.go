package templating

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testcontext struct {
	Test string
}

func TestProcessPlaceHolders(t *testing.T) {
	context := testcontext{"world"}
	assert.Equal(t, "hello world!", ProcessPlaceholders("hello {{.Test}}!", context))
}

package utils

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestFilterKeysByValue(t *testing.T) {
	in := map[string]bool{
		"foo": true,
		"bar": false,
		"baz": true,
	}

	filtered := FilterKeysByValue(in, true)

	assert.Equal(t, 2, len(filtered))
	assert.Equal(t, true, filtered["foo"])
	assert.Equal(t, true, filtered["baz"])
}

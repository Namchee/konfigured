package utils

import (
	"os"
	"testing"

	"gotest.tools/v3/assert"
)

func TestReadEnvString(t *testing.T) {
	key := "TOKEN"
	value := "access-token"

	os.Setenv(key, value)
	defer os.Clearenv()

	res := ReadEnvString(key)
	assert.Equal(t, value, res)
}

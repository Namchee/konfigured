package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadEnvString(t *testing.T) {
	key := "TOKEN"
	value := "access-token"

	os.Setenv(key, value)
	defer os.Clearenv()

	res := ReadEnvString(key)
	assert.Equal(t, value, res)
}

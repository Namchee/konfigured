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

func TestReadEnvBool(t *testing.T) {
	tests := []struct {
		name      string
		mockValue string
		want      bool
	}{
		{
			name:      "should read true correctly",
			mockValue: "true",
			want:      true,
		},
		{
			name:      "should read false correctly",
			mockValue: "false",
			want:      false,
		},
		{
			name:      "should fallback to false",
			mockValue: "bar",
			want:      false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("TEST", tc.mockValue)
			defer os.Unsetenv("TEST")

			got := ReadEnvBool("TEST")

			assert.Equal(t, got, tc.want)
		})
	}
}

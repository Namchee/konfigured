package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExtension(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "should return an empty string",
			args: "foobar",
			want: "",
		},
		{
			name: "should return yaml",
			args: "config.yaml",
			want: "yaml",
		},
		{
			name: "should handle dot-separated filename",
			args: "config.development.yaml",
			want: "yaml",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetExtension(tc.args)

			assert.Equal(t, tc.want, got)
		})
	}
}

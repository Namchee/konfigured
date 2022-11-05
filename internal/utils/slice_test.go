package utils

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		val  string
		want bool
	}{
		{
			name: "should return false",
			in:   []string{"a", "b", "c"},
			val:  "d",
			want: false,
		},
		{
			name: "should return true",
			in:   []string{"a", "b", "c"},
			val:  "c",
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Contains(tc.in, tc.val)

			assert.Equal(t, tc.want, got)
		})
	}
}

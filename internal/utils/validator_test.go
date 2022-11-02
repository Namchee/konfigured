package utils

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestIsValid(t *testing.T) {
	type args struct {
		ext     string
		content string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid JSON file",
			args: args{
				ext:     "json",
				content: "{",
			},
			want: false,
		},
		{
			name: "valid JSON file",
			args: args{
				ext:     "json",
				content: `{"foo": "bar"}`,
			},
			want: true,
		},
		{
			name: "invalid YAML file",
			args: args{
				ext: "yaml",
				content: `
foo:
  - bar
  baz`,
			},
			want: false,
		},
		{
			name: "valid YAML file",
			args: args{
				ext: "yaml",
				content: `
foo:
  - bar`,
			},
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValid(tc.args.ext, tc.args.content)

			assert.Equal(t, tc.want, got)
		})
	}
}

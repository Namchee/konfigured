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
			name: "invalid INI file",
			args: args{
				ext:     "ini",
				content: `key=:=''value"`,
			},
			want: false,
		},
		{
			name: "valid INI file",
			args: args{
				ext: "ini",
				content: `[ample]
key="value"
num=123`,
			},
			want: true,
		},
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
		{
			name: "invalid TOML file",
			args: args{
				ext: "toml",
				content: `key = "value"
[]`,
			},
			want: false,
		},
		{
			name: "valid TOML file",
			args: args{
				ext: "toml",
				content: `foo = "bar"

[map]
key = "value"
bilangan = 123

[a.b]
c = [1, 2, 3]

[[items]]

[[items]]
name = "eggs"`,
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

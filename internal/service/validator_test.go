package internal

import (
	"testing"

	"github.com/Namchee/konfigured/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/jarcoal/httpmock"
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
			name: "YML alias",
			args: args{
				ext: "yml",
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
			got := isValid(tc.args.ext, tc.args.content)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestValidateConfigurationFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mocks.NewMockGithub(ctrl)

	args := []*github.CommitFile{
		{
			Filename: github.String("foobar.json"),
			RawURL:   github.String("https://www.google.com"),
		},
		{
			Filename: github.String("sample.toml"),
			RawURL:   github.String("https://www.yahoo.com"),
		},
		{
			Filename: github.String("config.yaml"),
			RawURL:   github.String("https://www.facebook.com"),
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		"https://www.google.com",
		httpmock.NewBytesResponder(200, []byte("{")),
	)

	httpmock.RegisterResponder(
		"GET",
		"https://www.yahoo.com",
		httpmock.NewBytesResponder(200, []byte(`key = "value"`)),
	)

	httpmock.RegisterResponder(
		"GET",
		"https://www.facebook.com",
		httpmock.NewBytesResponder(200, []byte("foo: bar")),
	)

	result := ValidateConfigurationFiles(args)

	assert.Equal(t, 3, len(result))
	assert.Equal(t, 3, httpmock.GetTotalCallCount())
	assert.Equal(t, result["foobar.json"], false)
	assert.Equal(t, result["sample.toml"], true)
	assert.Equal(t, result["config.yaml"], true)
}

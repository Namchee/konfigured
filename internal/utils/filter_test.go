package utils

import (
	"testing"

	"github.com/google/go-github/v48/github"
	"gotest.tools/v3/assert"
)

func TestGetSupportedFiles(t *testing.T) {
	tests := []struct {
		name string
		args []*github.CommitFile
		want []*github.CommitFile
	}{
		{
			name: "none are supported",
			args: []*github.CommitFile{
				{
					Filename: github.String("foo.txt"),
				},
				{
					Filename: github.String("bar.jpg"),
				},
			},
			want: []*github.CommitFile{},
		},
		{
			name: "found supported files",
			args: []*github.CommitFile{
				{
					Filename: github.String("config.json"),
				},
				{
					Filename: github.String("bar.jpg"),
				},
				{
					Filename: github.String("configuration.yaml"),
				},
			},
			want: []*github.CommitFile{
				{
					Filename: github.String("config.json"),
				},
				{
					Filename: github.String("configuration.yaml"),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := GetSupportedFiles(tc.args)

			assert.DeepEqual(t, tc.want, got)
		})
	}
}

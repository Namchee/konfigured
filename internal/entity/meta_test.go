package entity

import (
	"testing"

	"github.com/Namchee/setel/internal/constant"
	"gotest.tools/v3/assert"
)

func TestCreateMeta(t *testing.T) {
	tests := []struct {
		name    string
		args    string
		want    *Meta
		wantErr error
	}{
		{
			name: "should be able to extract metadata",
			args: "foo/bar",
			want: &Meta{
				Name:  "bar",
				Owner: "foo",
			},
			wantErr: nil,
		},
		{
			name:    "should throw an error",
			args:    "fake_github_repository",
			want:    nil,
			wantErr: constant.ErrMalformedMetadata,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CreateMeta(tc.args)

			assert.DeepEqual(t, tc.want, got)
			assert.Equal(t, tc.wantErr, err)
		})
	}
}

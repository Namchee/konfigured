package internal

import (
	"context"
	"errors"
	"testing"

	"github.com/Namchee/konfigured/mocks/mock_client"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigurationValidator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := mock_client.NewMockGithubClient(ctrl)

	assert.NotPanics(t, func() {
		NewConfigurationValidator(client)
	})
}

func TestValidateConfigurationFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type response struct {
		content *github.RepositoryContent
		err     error
	}

	files := map[string]response{
		"foobar.json": {
			content: &github.RepositoryContent{
				Content: github.String("{"),
			},
			err: nil,
		},
		"sample.toml": {
			content: &github.RepositoryContent{
				Content: github.String(`key = "value"`),
			},
			err: nil,
		},
		"config.yaml": {
			content: &github.RepositoryContent{
				Content: github.String("key: value"),
			},
			err: nil,
		},
		"nested/config.yaml": {
			content: &github.RepositoryContent{
				Content: github.String(""),
			},
			err: errors.New("fail"),
		},
		"encoding.ini": {
			content: &github.RepositoryContent{
				Content:  github.String(""),
				Encoding: github.String("magic"),
			},
			err: nil,
		},
	}

	args := []*github.CommitFile{
		{
			Filename: github.String("foobar.json"),
		},
		{
			Filename: github.String("sample.toml"),
		},
		{
			Filename: github.String("config.yaml"),
		},
		{
			Filename: github.String("nested/config.yaml"),
		},
		{
			Filename: github.String("encoding.ini"),
		},
	}

	client := mock_client.NewMockGithubClient(ctrl)

	for filename, resp := range files {
		client.EXPECT().GetFileContent(gomock.Any(), filename).
			Return(resp.content, resp.err)
	}

	validator := &ConfigurationValidator{
		client: client,
	}

	got := validator.ValidateFiles(context.TODO(), args)

	assert.Equal(t, 5, len(got))
}

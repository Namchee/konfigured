package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Namchee/konfigured/internal/entity"
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
		NewConfigurationValidator(&entity.Configuration{}, client)
	})
}

func TestConfigurationValidator_GetSupportedFiles(t *testing.T) {
	files := []*github.CommitFile{
		{
			Filename: github.String("foobar.jpg"),
		},
		{
			Filename: github.String("test.toml"),
		},
		{
			Filename: github.String("baz.json"),
		},
		{
			Filename: github.String("README"),
		},
		{
			Filename: github.String("bb.ini"),
		},
	}

	validator := &ConfigurationValidator{}

	got := validator.GetSupportedFiles(files)

	assert.Equal(t, []*github.CommitFile{
		{
			Filename: github.String("test.toml"),
		},
		{
			Filename: github.String("baz.json"),
		},
		{
			Filename: github.String("bb.ini"),
		},
	}, got)
}

func TestConfigurationValidator_ValidateConfigurationFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type response struct {
		content *github.RepositoryContent
		err     error
	}

	files := map[string]response{
		"foobar.json": {
			content: &github.RepositoryContent{
				Content: github.String("{\n"),
			},
			err: nil,
		},
		"sample.toml": {
			content: &github.RepositoryContent{
				Content: github.String(`key = "value"
`),
			},
			err: nil,
		},
		"config.yaml": {
			content: &github.RepositoryContent{
				Content: github.String("key: value\n"),
			},
			err: nil,
		},
		"no-newline.yaml": {
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
			Filename: github.String("no-newline.yaml"),
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
		cfg: &entity.Configuration{
			Newline: true,
		},
		client: client,
	}

	got := validator.ValidateFiles(context.TODO(), args)

	assert.Equal(t, 6, len(got))

	invalids := entity.GetInvalidValidations(got)

	assert.Equal(t, 4, len(invalids))
}

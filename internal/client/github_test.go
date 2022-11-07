package client

import (
	"context"
	"errors"
	"testing"

	"github.com/Namchee/konfigured/internal/entity"
	"github.com/Namchee/konfigured/mocks/mock_client"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/v48/github"
	"gotest.tools/v3/assert"
)

func TestNewGithubClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	pullRequestService := mock_client.NewMockPullRequestService(ctrl)
	repositoryService := mock_client.NewMockRepositoryService(ctrl)

	_ = NewGithubClient("", &entity.Meta{}, pullRequestService, repositoryService)
}

func TestGithubClient_GetChangedFiles(t *testing.T) {
	type mockService struct {
		files []*github.CommitFile
		res   *github.Response
		err   error
	}

	tests := []struct {
		name        string
		mockService mockService
		want        []*github.CommitFile
		wantErr     bool
	}{
		{
			name: "error from service",
			mockService: mockService{
				files: []*github.CommitFile{},
				err:   errors.New("error"),
			},
			want:    []*github.CommitFile{},
			wantErr: true,
		},
		{
			name: "success",
			mockService: mockService{
				files: []*github.CommitFile{
					{
						Filename: github.String("commited"),
					},
				},
			},
			want: []*github.CommitFile{
				{
					Filename: github.String("commited"),
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			pullRequestService := mock_client.NewMockPullRequestService(ctrl)
			pullRequestService.EXPECT().ListFiles(
				gomock.Any(),
				gomock.Any(),
				gomock.Any(),
				gomock.Any(),
				gomock.Any(),
			).Return(
				tc.mockService.files,
				tc.mockService.res,
				tc.mockService.err,
			)

			svc := &githubClient{
				meta:               &entity.Meta{},
				pullRequestService: pullRequestService,
			}

			got, err := svc.GetChangedFiles(context.TODO(), 123)

			assert.DeepEqual(t, tc.want, got)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}

func TestGithubClient_GetFileContent(t *testing.T) {
	type mockService struct {
		file *github.RepositoryContent
		dir  []*github.RepositoryContent
		res  *github.Response
		err  error
	}

	tests := []struct {
		name        string
		mockService mockService
		want        *github.RepositoryContent
		wantErr     bool
	}{
		{
			name: "error from service",
			mockService: mockService{
				file: &github.RepositoryContent{},
				err:  errors.New("error"),
			},
			want:    &github.RepositoryContent{},
			wantErr: true,
		},
		{
			name: "success",
			mockService: mockService{
				file: &github.RepositoryContent{
					Name:    github.String("want"),
					Content: github.String("content"),
				},
				dir: []*github.RepositoryContent{
					{
						Name: github.String("content"),
					},
				},
			},
			want: &github.RepositoryContent{
				Name:    github.String("want"),
				Content: github.String("content"),
			},
			wantErr: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repositoryService := mock_client.NewMockRepositoryService(ctrl)

			repositoryService.EXPECT().GetContents(
				gomock.Any(),
				gomock.Any(),
				gomock.Any(),
				gomock.Any(),
				gomock.Any(),
			).Return(
				tc.mockService.file,
				tc.mockService.dir,
				tc.mockService.res,
				tc.mockService.err,
			)

			svc := &githubClient{
				meta:              &entity.Meta{},
				repositoryService: repositoryService,
			}

			got, err := svc.GetFileContent(context.TODO(), "")

			assert.DeepEqual(t, tc.want, got)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}

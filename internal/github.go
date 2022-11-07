package internal

import (
	"context"

	"github.com/google/go-github/v48/github"
)

type PullRequestService interface {
	ListFiles(
		ctx context.Context,
		owner string,
		name string,
		event int,
		opts *github.ListOptions,
	) ([]*github.CommitFile, *github.Response, error)
}

type RepositoryService interface {
	GetContents(
		ctx context.Context,
		owner string,
		name string,
		path string,
		opts *github.RepositoryContentGetOptions,
	) (*github.RepositoryContent, []*github.RepositoryContent, *github.Response, error)
}

type GithubClient interface {
	GetChangedFiles(
		ctx context.Context,
		event int,
	) ([]*github.CommitFile, error)
	GetFileContent(
		ctx context.Context,
		path string,
	) (*github.RepositoryContent, error)
}

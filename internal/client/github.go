package client

import (
	"context"

	"github.com/Namchee/konfigured/internal"
	"github.com/Namchee/konfigured/internal/entity"
	"github.com/google/go-github/v48/github"
)

type githubClient struct {
	ref  string
	meta *entity.Meta

	pullRequestService internal.PullRequestService
	repositoryService  internal.RepositoryService
}

func NewGithubClient(
	ref string,
	meta *entity.Meta,
	pullRequestService internal.PullRequestService,
	repositoryService internal.RepositoryService,
) internal.GithubClient {
	return &githubClient{
		ref:                ref,
		meta:               meta,
		pullRequestService: pullRequestService,
		repositoryService:  repositoryService,
	}
}

func (c *githubClient) GetChangedFiles(
	ctx context.Context,
	event int,
) ([]*github.CommitFile, error) {
	files, _, err := c.pullRequestService.ListFiles(
		ctx,
		c.meta.Owner,
		c.meta.Name,
		event,
		&github.ListOptions{},
	)

	return files, err
}

func (c *githubClient) GetFileContent(
	ctx context.Context,
	path string,
) (*github.RepositoryContent, error) {
	file, _, _, err := c.repositoryService.GetContents(
		ctx,
		c.meta.Owner,
		c.meta.Name,
		path,
		&github.RepositoryContentGetOptions{
			Ref: c.ref,
		},
	)

	return file, err
}

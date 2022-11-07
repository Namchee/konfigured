package mocks

import _ "github.com/golang/mock/mockgen/model"

//go:generate mockgen -destination=mock_client/mock_github.go -package=mock_client github.com/Namchee/konfigured/internal GithubClient,PullRequestService,RepositoryService

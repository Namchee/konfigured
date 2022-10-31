package main

import (
	"context"
	"log"
	"os"

	"github.com/Namchee/setel/internal/entity"
	"github.com/Namchee/setel/internal/utils"
	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

// Logger
var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lmsgprefix)
	errorLogger = log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lmsgprefix)
}

func main() {
	ctx := context.Background()

	config, err := entity.CreateConfiguration()
	if err != nil {
		errorLogger.Fatalln(err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Token},
	)

	oauth := oauth2.NewClient(ctx, ts)
	client := github.NewClient(oauth)

	event, err := entity.ReadEvent(os.DirFS("/"))

	if err != nil {
		errorLogger.Fatalf("Failed to read repository event: %s", err.Error())
	}

	meta, err := entity.CreateMeta(
		utils.ReadEnvString("GITHUB_REPOSITORY"),
	)

	if err != nil {
		errorLogger.Fatalf("Failed to read repository metadata: %s", err.Error())
	}

	pullRequest, _, err := client.PullRequests.Get(ctx, meta.Owner, meta.Name, event.Number)

	if err != nil {
		errorLogger.Fatalf("Failed to fetch pull request data: %s", err.Error())
	}
}

package main

import (
	"context"
	"log"
	"os"

	"github.com/Namchee/konfigured/internal"
	"github.com/Namchee/konfigured/internal/entity"
	"github.com/Namchee/konfigured/internal/utils"
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

	client.Repositories.GetContents()

	files, _, err := client.PullRequests.ListFiles(
		ctx,
		meta.Owner,
		meta.Name,
		event.Number,
		&github.ListOptions{},
	)

	client.Repositories.GetContents()

	if err != nil {
		errorLogger.Fatalf("Failed to fetch list of file changes: %s", err.Error())
	}

	supportedFiles := utils.GetSupportedFiles(files)

	infoLogger.Printf("Found %d supported configuration files\n", len(supportedFiles))

	if len(supportedFiles) == 0 {
		return
	}

	result := internal.ValidateConfigurationFiles(supportedFiles)
	invalids := utils.FilterKeysByValue(result, false)

	if len(invalids) == 0 {
		infoLogger.Println("All configuration files are valid!")
		os.Exit(0)
	}

	infoLogger.Printf("Found %d malformed configuration files\n", len(invalids))
	infoLogger.Println("Below are the list of malformed configuration files:")

	for name := range invalids {
		infoLogger.Println(name)
	}

	os.Exit(1)
}

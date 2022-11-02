package utils

import (
	"regexp"

	"github.com/google/go-github/v48/github"
)

var (
	extensions = regexp.MustCompile(`\.(ini|json|yaml|toml|hcl)$`)
)

// GetSuupportedFiles returns list of of supported configuration files
func GetSupportedFiles(files []*github.CommitFile) []*github.CommitFile {
	supportedFiles := []*github.CommitFile{}

	for _, file := range files {
		if extensions.Match([]byte(file.GetFilename())) {
			supportedFiles = append(supportedFiles, file)
		}
	}

	return supportedFiles
}

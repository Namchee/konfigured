package utils

import (
	"regexp"
	"strings"

	"github.com/google/go-github/v48/github"
)

var (
	extensions = regexp.MustCompile(`\.(ini|json|yaml|toml)$`)
)

// GetSupportedFiles returns list of of supported configuration files
func GetSupportedFiles(files []*github.CommitFile) []*github.CommitFile {
	supportedFiles := []*github.CommitFile{}

	for _, file := range files {
		if extensions.Match([]byte(file.GetFilename())) {
			supportedFiles = append(supportedFiles, file)
		}
	}

	return supportedFiles
}

// GetExtension returns file extension
func GetExtension(name string) string {
	token := strings.Split(name, ".")

	return token[len(token)-1]
}

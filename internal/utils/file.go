package utils

import (
	"path/filepath"
	"strings"

	"github.com/google/go-github/v48/github"
)

var (
	supportedExtensions = []string{
		"ini",
		"json",
		"yaml",
		"yml",
		"toml",
	}
)

// GetSupportedFiles returns list of of supported configuration files
func GetSupportedFiles(files []*github.CommitFile) []*github.CommitFile {
	supportedFiles := []*github.CommitFile{}

	for _, file := range files {
		ext := GetExtension(file.GetFilename())

		if Contains(supportedExtensions, ext) {
			supportedFiles = append(supportedFiles, file)
		}
	}

	return supportedFiles
}

// GetExtension returns file extension from its name
func GetExtension(name string) string {
	return strings.ReplaceAll(filepath.Ext(name), ".", "")
}

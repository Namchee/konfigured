package service

import (
	"bytes"
	"context"
	"sync"

	"github.com/Namchee/konfigured/internal"
	"github.com/Namchee/konfigured/internal/entity"
	"github.com/Namchee/konfigured/internal/utils"
	"github.com/google/go-github/v48/github"
	"github.com/spf13/viper"
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

// ConfigurationValidator is a service that validates configuration files
type ConfigurationValidator struct {
	client internal.GithubClient
}

func NewConfigurationValidator(
	client internal.GithubClient,
) *ConfigurationValidator {
	return &ConfigurationValidator{
		client: client,
	}
}

// ValidateConfigurationFiles returns a mapping of configuration file validity
func (v *ConfigurationValidator) ValidateFiles(
	ctx context.Context,
	files []*github.CommitFile,
) []entity.Validation {
	result := []entity.Validation{}

	pool := make(chan entity.Validation, len(files))

	wg := &sync.WaitGroup{}
	wg.Add(len(files))

	for _, file := range files {
		go func(ctx context.Context, f *github.CommitFile) {
			defer wg.Done()

			valid := v.isValid(ctx, f)

			pool <- entity.Validation{
				Filename: f.GetFilename(),
				Valid:    valid,
			}
		}(ctx, file)
	}

	go func() {
		wg.Wait()
		close(pool)
	}()

	for res := range pool {
		result = append(result, res)
	}

	return result
}

// GetSupportedFiles returns list of of supported configuration files
func (v *ConfigurationValidator) GetSupportedFiles(
	files []*github.CommitFile,
) []*github.CommitFile {
	supportedFiles := []*github.CommitFile{}

	for _, file := range files {
		ext := utils.GetExtension(file.GetFilename())

		if utils.Contains(supportedExtensions, ext) {
			supportedFiles = append(supportedFiles, file)
		}
	}

	return supportedFiles
}

// isValid verify the structure of the config file
func (v *ConfigurationValidator) isValid(
	ctx context.Context,
	file *github.CommitFile,
) bool {
	fileContent, err := v.client.GetFileContent(ctx, file.GetFilename())
	content, err := fileContent.GetContent()
	// always false if we are not able to test it
	if err != nil {
		return false
	}

	ext := utils.GetExtension(file.GetFilename())

	validator := viper.New()
	validator.SetConfigType(ext)

	err = validator.ReadConfig(bytes.NewBufferString(content))

	return err == nil
}

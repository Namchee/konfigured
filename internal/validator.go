package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/Namchee/setel/internal/utils"
	"github.com/google/go-github/v48/github"
	"github.com/spf13/viper"
)

type ValidationResult struct {
	name    string
	isValid bool
}

// ValidateConfigurationFiles returns a mapping of configuration file validity
func ValidateConfigurationFiles(files []*github.CommitFile) map[string]bool {
	filemap := map[string]bool{}

	pool := make(chan ValidationResult, len(files))

	wg := &sync.WaitGroup{}
	wg.Add(len(files))

	for _, file := range files {
		go func(f *github.CommitFile) {
			defer wg.Done()

			content := fetchFileContent(f.GetRawURL())
			extension := utils.GetExtension(f.GetFilename())

			valid := isValid(extension, content)

			pool <- ValidationResult{
				name:    f.GetFilename(),
				isValid: valid,
			}
		}(file)
	}

	go func() {
		wg.Wait()
		close(pool)
	}()

	for result := range pool {
		filemap[result.name] = result.isValid
	}

	return filemap
}

func fetchFileContent(url string) string {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := &http.Client{}

	response, _ := client.Do(req)

	body, _ := ioutil.ReadAll(response.Body)

	return string(body)
}

// isValid verify the structure of the config file
func isValid(ext string, content string) bool {
	viper.SetConfigType(ext)

	err := viper.ReadConfig(bytes.NewBuffer([]byte(content)))

	return err == nil
}

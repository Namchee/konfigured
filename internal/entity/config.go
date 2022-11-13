package entity

import (
	"github.com/Namchee/konfigured/internal/constant"
	"github.com/Namchee/konfigured/internal/utils"

	"github.com/bmatcuk/doublestar/v4"
)

const (
	defaultPattern = "**/*.{json,ini,yaml,yml,toml}"
)

type Configuration struct {
	Token   string
	Newline bool
	Include string
}

func CreateConfiguration() (*Configuration, error) {
	token := utils.ReadEnvString("INPUT_TOKEN")

	if token == "" {
		return nil, constant.ErrMissingToken
	}

	newline := utils.ReadEnvBool("INPUT_NEWLINE")
	include := utils.ReadEnvString("INPUT_INCLUDE")

	if !doublestar.ValidatePattern(include) {
		return nil, constant.ErrInvalidGlob
	}

	if include == "" {
		include = defaultPattern
	}

	return &Configuration{
		Token:   token,
		Newline: newline,
		Include: include,
	}, nil
}

package entity

import (
	"github.com/Namchee/konfigured/internal/constant"
	"github.com/Namchee/konfigured/internal/utils"
)

type Configuration struct {
	Token   string
	Newline bool
}

func CreateConfiguration() (*Configuration, error) {
	token := utils.ReadEnvString("INPUT_TOKEN")

	if token == "" {
		return nil, constant.ErrMissingToken
	}

	newline := utils.ReadEnvBool("INPUT_NEWLINE")

	return &Configuration{
		Token:   token,
		Newline: newline,
	}, nil
}

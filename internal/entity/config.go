package entity

import (
	"github.com/Namchee/atur/internal/constant"
	"github.com/Namchee/atur/internal/utils"
)

type Configuration struct {
	Token string
}

func CreateConfiguration() (*Configuration, error) {
	token := utils.ReadEnvString("TOKEN")

	if token == "" {
		return nil, constant.ErrMissingToken
	}

	return &Configuration{
		Token: token,
	}, nil
}

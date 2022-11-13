package entity

import (
	"fmt"
	"os"

	"github.com/Namchee/konfigured/internal/constant"
	"github.com/Namchee/konfigured/internal/utils"
)

type Configuration struct {
	Token   string
	Newline bool
	Include []string
}

func CreateConfiguration() (*Configuration, error) {
	token := utils.ReadEnvString("INPUT_TOKEN")

	if token == "" {
		return nil, constant.ErrMissingToken
	}

	newline := utils.ReadEnvBool("INPUT_NEWLINE")
	include := os.Getenv("INPUT_INCLUDE")

	fmt.Println(include)

	return &Configuration{
		Token:   token,
		Newline: newline,
	}, nil
}

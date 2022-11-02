package utils

import (
	"bytes"

	"github.com/spf13/viper"
)

// IsValid verify the structure of the config file
func IsValid(ext string, content string) bool {
	viper.SetConfigType(ext)

	err := viper.ReadConfig(bytes.NewBuffer([]byte(content)))

	return err == nil
}

package utils

import (
	"os"
	"strconv"
)

// ReadEnvString returns string value from environment variables
func ReadEnvString(key string) string {
	return os.Getenv(key)
}

// ReadEnvBool read and parse boolean environment variables.
// Will return `false` if the variable is not a `bool`
func ReadEnvBool(key string) bool {
	value := os.Getenv(key)
	parsed, err := strconv.ParseBool(value)

	if err != nil {
		return false
	}

	return parsed
}

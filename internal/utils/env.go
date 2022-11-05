package utils

import "os"

// ReadEnvString returns string value from environment variables
func ReadEnvString(key string) string {
	return os.Getenv(key)
}

package utils

import "os"

func ReadEnvString(key string) string {
	return os.Getenv(key)
}

package utils

import (
	"path/filepath"
	"strings"
)

// GetExtension returns file extension from its name
func GetExtension(name string) string {
	return strings.ReplaceAll(filepath.Ext(name), ".", "")
}

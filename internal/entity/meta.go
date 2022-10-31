package entity

import (
	"strings"

	"github.com/Namchee/setel/internal/constant"
)

// Meta is a struct that represents repository's metadata
type Meta struct {
	Name  string
	Owner string
}

// CreateMeta from a GitHub's repository string
func CreateMeta(name string) (*Meta, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) != 2 {
		return nil, constant.ErrMalformedMetadata
	}

	return &Meta{
		Name:  tokens[1],
		Owner: tokens[0],
	}, nil
}

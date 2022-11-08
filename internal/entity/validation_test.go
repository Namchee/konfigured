package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInvalidValidations(t *testing.T) {
	in := []Validation{
		{
			Filename: "foo.ini",
			Valid:    true,
		},
		{
			Filename: "bar.yaml",
			Valid:    false,
		},
		{
			Filename: "baz.toml",
			Valid:    false,
		},
		{
			Filename: "nested/one.json",
			Valid:    true,
		},
	}

	got := GetInvalidValidations(in)

	assert.Equal(t, []Validation{
		{
			Filename: "bar.yaml",
			Valid:    false,
		},
		{
			Filename: "baz.toml",
			Valid:    false,
		},
	}, got)
}

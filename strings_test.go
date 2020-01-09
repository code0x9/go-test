package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringsSplit(t *testing.T) {
	stringTests := []struct {
		src string
		res string
	}{
		{"", ""},
		{"Bearer ", ""},
		{"Bearer ABC", "ABC"},
	}

	for _, tt := range stringTests {
		t.Run(tt.src, func(t *testing.T) {
			res := strings.TrimPrefix(tt.src, "Bearer ")
			assert.Equal(t, tt.res, res)
		})
	}
}

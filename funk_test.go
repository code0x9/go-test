package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thoas/go-funk"
)

func TestFunkContains(t *testing.T) {
	assert.True(t, funk.Contains(map[int]string{1: "Florent"}, 1))
}

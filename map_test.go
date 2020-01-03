package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	m := map[interface{}]interface{}{
		"k1": 1,
		"k2": "aa",
	}

	v1, v1ok := m["k1"].(int)
	v2, v2ok := m["k2"].(int)
	v3, v3ok := m["k99"].(int)
	assert.Equal(t, 1, v1)
	assert.Equal(t, true, v1ok)
	assert.Equal(t, 0, v2)
	assert.Equal(t, false, v2ok)
	assert.Equal(t, 0, v3)
	assert.Equal(t, false, v3ok)
}

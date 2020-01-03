package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func loop() (results []int) {
	a := []int{1, 2, 3}
	b := []int{10, 20, 30}

Done:
	for _, x := range a {
		results = append(results, x)
		for _, y := range b {
			results = append(results, y)
			if x == 2 && y == 20 {
				break Done
			}
		}
	}
	return
}

func TestLabelBreak(t *testing.T) {
	results := loop()
	assert.Equal(t, []int{1, 10, 20, 30, 2, 10, 20}, results)
}

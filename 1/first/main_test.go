package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountIncreases(t *testing.T) {
	input := []int {
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	actual := countIncreases(input)
	expected := 7
	assert.Equal(t, expected, actual)
}
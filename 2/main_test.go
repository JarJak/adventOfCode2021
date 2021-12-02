package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountPosition(t *testing.T) {
	input := []string {
		`forward 5`,
		`down 5`,
		`forward 8`,
		`up 3`,
		`down 8`,
		`forward 2`,
	}
	horizontal, depth := countPosition(input)
	expectedHorizontal, expectedDepth := 15, 10
	assert.Equal(t, horizontal, expectedHorizontal)
	assert.Equal(t, depth, expectedDepth)
}

func TestCountPosition2(t *testing.T) {
	input := []string {
		`forward 5`,
		`down 5`,
		`forward 8`,
		`up 3`,
		`down 8`,
		`forward 2`,
	}
	horizontal, depth := countPosition2(input)
	expectedHorizontal, expectedDepth := 15, 60
	assert.Equal(t, horizontal, expectedHorizontal)
	assert.Equal(t, depth, expectedDepth)
}
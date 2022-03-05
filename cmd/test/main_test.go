package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare(t *testing.T) {
	ass := assert.New(t)
	expected := 81
	ass.Equal(expected, square(9), "square(9) should be 81")
}

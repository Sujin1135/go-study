package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare2(t *testing.T) {
	ass := assert.New(t)
	ass.Equal(9, square(3), "square(3) should be 9")
}

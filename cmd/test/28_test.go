package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare2(t *testing.T) {
	ass := assert.New(t)
	ass.Equal(9, square(3), "square(3) should be 9")
}

func TestAtoI(t *testing.T) {
	ass := assert.New(t)
	rst, err := AtoI("13")

	ass.Nilf(err, "It was not occurred an error")
	ass.Equal(13, rst, "string(13) should be int(13)")
}

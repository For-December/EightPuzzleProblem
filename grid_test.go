package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMove(t *testing.T) {
	g := new(Grid)
	g.CurState = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}

	assert.Equal(t, g.up().CurState[1][2], 0, "err")
	assert.Nil(t, g.down(), "err")
	assert.Equal(t, g.left().CurState[2][1], 0, "err")
	assert.Nil(t, g.right(), "err")

}

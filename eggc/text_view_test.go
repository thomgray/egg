package eggc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitInts(t *testing.T) {
	ints := []int{1, 2, 3, 10, 100}

	before, after := splitIntSliceBetween(ints, 7)
	assert.Equal(t, ints[:3], before)
	assert.Equal(t, ints[3:], after)

	before1, after1 := splitIntSliceBetween(ints, 1000)
	assert.Equal(t, ints, before1)
	assert.Equal(t, ints[5:], after1)

	before2, after2 := splitIntSliceBetween(ints, 0)
	assert.Equal(t, ints[:0], before2)
	assert.Equal(t, ints[0:], after2)
}

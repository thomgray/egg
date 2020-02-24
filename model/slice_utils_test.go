package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuneInsert(t *testing.T) {
	arr := []rune{'h', 'e', 'l', 'l', 'o'}

	assert.Equal(t, []rune("shello"), runeSliceInsert(arr, 0, 's'))
	assert.Equal(t, []rune("hellos"), runeSliceInsert(arr, 5, 's'))
	assert.Equal(t, []rune("helslo"), runeSliceInsert(arr, 3, 's'))
}

func TestRuneGrowDoublingStrategy(t *testing.T) {
	// doubles if needs to grow
	arr := make([]rune, 10, 10)
	doubled := runeSliceGrowDoublingStrategy(arr)
	assert.Equal(t, 20, cap(doubled))
	assert.Equal(t, 10, len(doubled))

	// doesn't change if it has space capacity
	arr2 := make([]rune, 1, 10)
	notDoubled := runeSliceGrowDoublingStrategy(arr2)
	assert.Equal(t, 10, cap(notDoubled))
	assert.Equal(t, 1, len(notDoubled))
}

func TestRuneShrinkHalvingStrategy(t *testing.T) {
	arr := make([]rune, 3, 12)
	halved := runeSliceShrinkHalvingStrategy(arr)
	assert.Equal(t, 6, cap(halved))
	assert.Equal(t, 3, len(halved))

	arr2 := make([]rune, 4, 12)
	notHalved := runeSliceShrinkHalvingStrategy(arr2)
	assert.Equal(t, 12, cap(notHalved))
	assert.Equal(t, 4, len(notHalved))
}

func TestRuneSliceRemove(t *testing.T) {
	assert.Equal(t, "helo there", string(runeSliceRemove([]rune("hello there"), 2)))
	assert.Equal(t, "hellothere", string(runeSliceRemove([]rune("hello there"), 5)))
	assert.Equal(t, "hello ther", string(runeSliceRemove([]rune("hello there"), 10)))
}

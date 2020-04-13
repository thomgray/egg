package eggc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesSliceGrow(t *testing.T) {
	arr := []byte("string")

	assert.Equal(t, 6, len(arr))
	assert.Equal(t, 6, cap(arr))
	grown := byteSliceGrow(arr, 100)
	assert.Equal(t, 6, len(grown))
	assert.Equal(t, 100, cap(grown))
}

func TestByeSliceInsert(t *testing.T) {
	arr := []byte("hello carl")
	toAdd := []byte(" there")
	arr2 := byteSliceInsert(arr, 5, toAdd)

	assert.Equal(t, 16, len(arr2))
	assert.Equal(t, []byte("hello there carl"), arr2)
}

func TestByteSliceSplit(t *testing.T) {
	arr := []byte("a dream is a wish")
	split := byteSliceSplit(arr, ' ')

	assert.Equal(t, 5, len(split))
	assert.Equal(t, "a", string(split[0]))
	assert.Equal(t, "dream", string(split[1]))
	assert.Equal(t, "is", string(split[2]))
	assert.Equal(t, "a", string(split[3]))
	assert.Equal(t, "wish", string(split[4]))
}

func TestByteSliceSplitWithDelimitersAtEnds(t *testing.T) {
	arr := []byte(" a dream is a wish ")
	split := byteSliceSplit(arr, ' ')

	assert.Equal(t, 7, len(split))
	assert.Equal(t, "", string(split[0]))
	assert.Equal(t, "a", string(split[1]))
	assert.Equal(t, "dream", string(split[2]))
	assert.Equal(t, "is", string(split[3]))
	assert.Equal(t, "a", string(split[4]))
	assert.Equal(t, "wish", string(split[5]))
	assert.Equal(t, "", string(split[6]))
}

func TestByteSliceIndexes(t *testing.T) {
	arr := []byte("a dream is a wish")
	split := byteSliceFindIndexes(arr, ' ')

	assert.Equal(t, 4, len(split))
	assert.Equal(t, 1, split[0])
	assert.Equal(t, 7, split[1])
	assert.Equal(t, 10, split[2])
	assert.Equal(t, 12, split[3])
}

func TestByteSliceIndexesWithDelimiterOnEdges(t *testing.T) {
	arr := []byte(" a dream is a wish ")
	split := byteSliceFindIndexes(arr, ' ')

	assert.Equal(t, 6, len(split))

	assert.Equal(t, 0, split[0])
	assert.Equal(t, 2, split[1])
	assert.Equal(t, 8, split[2])
	assert.Equal(t, 11, split[3])
	assert.Equal(t, 13, split[4])
	assert.Equal(t, 18, split[5])
}

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

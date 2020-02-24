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

package eggc

import "unicode/utf8"

// grow the slice to desired capacity if not already at capacity
func byteSliceGrow(bytes []byte, desiredCapacity int) []byte {
	if cap(bytes) < desiredCapacity {
		newSlice := make([]byte, len(bytes), desiredCapacity)
		copy(newSlice, bytes)
		return newSlice
	}
	return bytes
}

func byteSliceInsert(bytes []byte, bytesPos int, bytesToInsert []byte) []byte {
	desireCapacity := len(bytes) + len(bytesToInsert)
	lengthOfInsertion := len(bytesToInsert)
	bytesAtCapacity := byteSliceGrow(bytes, desireCapacity)
	bytesAtCapacity = bytesAtCapacity[:desireCapacity]
	copy(bytesAtCapacity[bytesPos+lengthOfInsertion:], bytesAtCapacity[bytesPos:])
	copy(bytesAtCapacity[bytesPos:], bytesToInsert)
	return bytesAtCapacity
}

func byteSliceSplit(bytes []byte, delimiter rune) [][]byte {
	remaining := bytes
	result := make([][]byte, 0)
	buff := make([]byte, 0, len(remaining))
	for len(remaining) > 0 {
		r, l := utf8.DecodeRune(remaining)
		runeBytes := remaining[:l]
		if r == delimiter {
			result = append(result, buff[:len(buff)])
			buff = make([]byte, 0, len(remaining))
		} else {
			bl := len(buff)
			buff = buff[:len(buff)+len(runeBytes)]
			copy(buff[bl:], runeBytes)
		}
		remaining = remaining[l:]
	}
	result = append(result, buff[:len(buff)])
	return result
}

func byteSliceFindIndexes(bytes []byte, delimiter rune) []int {
	result := make([]int, 0)
	remaining := bytes
	i := 0
	for len(remaining) > 0 {
		r, l := utf8.DecodeRune(remaining)
		runeBytes := remaining[:l]
		if r == delimiter {
			result = append(result, i)
		}
		i += len(runeBytes)
		remaining = remaining[l:]
	}

	return result
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func runeSliceGrowDoublingStrategy(runes []rune) []rune {
	capacity := cap(runes)
	length := len(runes)
	if capacity <= length {
		newLen := max((capacity * 2), 5)
		newSlice := make([]rune, len(runes), newLen)
		copy(newSlice, runes)
		return newSlice
	}
	return runes
}

func runeSliceShrinkHalvingStrategy(runes []rune) []rune {
	capacity := cap(runes)
	length := len(runes)
	if length <= capacity/4 {
		newSlice := make([]rune, len(runes), capacity/2)
		copy(newSlice, runes)
		return newSlice
	}
	return runes
}

func runeSliceGrow(runes []rune, desiredCapacity int) []rune {
	if cap(runes) < desiredCapacity {
		newSlice := make([]rune, len(runes), desiredCapacity)
		copy(newSlice, runes)
		return newSlice
	}
	return runes
}

func runeSliceInsert(runes []rune, pos int, runeToInsert rune) []rune {
	desiredCapacity := len(runes) + 1
	newRunes := runeSliceGrowDoublingStrategy(runes)
	newRunes = newRunes[:desiredCapacity]

	copy(newRunes[pos+1:], newRunes[pos:])
	newRunes[pos] = runeToInsert
	return newRunes
}

func runeSliceRemove(runes []rune, pos int) []rune {
	newRunes := runeSliceShrinkHalvingStrategy(runes)
	newLength := len(newRunes) - 1
	copy(newRunes[pos:], newRunes[pos+1:])
	newRunes = newRunes[:newLength]
	return newRunes
}

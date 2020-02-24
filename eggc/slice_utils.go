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

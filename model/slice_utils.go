package model

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

package eggc

type (
	// AlignmentHorizontal - horizontal alignment
	AlignmentHorizontal uint8
	// AlignmentVertical - vertical alignment
	AlignmentVertical uint8
)

// Alignment ..
const (
	AlignedLeft             AlignmentHorizontal = 0
	AlignedRight                                = 1
	AlignedCenterHorizontal                     = 2
	AlignedTop              AlignmentVertical   = 0
	AlignedBottom                               = 1
	AlignedCenterVertical                       = 2
)

type aligned struct {
	h AlignmentHorizontal
	v AlignmentVertical
}

package egg

// Origin - represents an x-y coordinate ancohored at the top-left
type Origin struct {
	X int
	Y int
}

// Size - represents a width/height
type Size struct {
	Width  int
	Height int
}

// Bounds - composit of Origin and Size, representing a position and extension
type Bounds struct {
	Origin
	Size
}

// Contains - returns if this bounds contains the parameter x/y coordinate
func (b Bounds) Contains(x, y int) bool {
	return b.X <= x && b.Y <= y && b.X+b.Width > x && b.Y+b.Height > y
}

// MakeBounds - construct a bounds with the given x, y coordinates and width/height dimentions
func MakeBounds(x, y, width, height int) Bounds {
	return Bounds{
		Origin{
			X: x,
			Y: y,
		},
		Size{
			Width:  width,
			Height: height,
		},
	}
}

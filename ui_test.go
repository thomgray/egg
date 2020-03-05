package egg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbsoluteBounds(t *testing.T) {
	v := MakeView()
	v.SetBounds(MakeBounds(1, 1, 10, 10))

	assert.Equal(t, MakeBounds(1, 1, 10, 10), absoluteBounds(v))
	assert.Equal(t, MakeBounds(1, 1, 10, 10), absoluteBounds2(v.bounds, v))
}

func TestAbsoluteBoundsInView(t *testing.T) {
	v := MakeView()
	v.SetBounds(MakeBounds(1, 1, 10, 10))

	v2 := MakeView()
	v2.SetBounds(MakeBounds(2, 2, 18, 18))

	v2.AddSubView(v)

	assert.Equal(t, MakeBounds(3, 3, 10, 10), absoluteBounds(v))
	assert.Equal(t, MakeBounds(3, 3, 10, 10), absoluteBounds2(v.bounds, v))
}

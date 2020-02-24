package egg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeView(t *testing.T) {
	v1 := MakeView()

	assert.True(t, v1.focusable)
	assert.Equal(t, []*View{}, v1.subViews)
}

func TestAddSubview(t *testing.T) {
	v1 := MakeView()
	v2 := View{}

	v1.AddSubView(&v2)

	assert.Equal(t, &v2, v1.subViews[0])
	assert.Equal(t, v1, v2.superView)
}

func TestUnmount(t *testing.T) {
	v1 := MakeView()
	v2 := MakeView()

	v1.AddSubView(v2)
	v2.Unmount()

	assert.Equal(t, 0, len(v1.subViews))
	assert.Nil(t, v2.superView)
}

package eggc

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBatcher(t *testing.T) {
	called := false
	b := MakeBatcher()

	b.Receive(func(i []interface{}) {
		assert.Equal(t, []interface{}{1, 2, 3}, i)
		called = true
	})

	for i := 1; i <= 3; i++ {
		b.Send(i)
	}

	time.Sleep(150 * time.Millisecond)

	assert.True(t, called)
}

func TestBatcherAfterInterval(t *testing.T) {
	called := false
	b := MakeBatcher()
	batches := make([]interface{}, 0, 2)

	b.Receive(func(i []interface{}) {
		called = true
		batches = append(batches, i)
	})

	for i := 1; i <= 4; i++ {
		if i == 4 {
			time.Sleep(120 * time.Millisecond)
		}
		b.Send(i)
	}

	time.Sleep(200 * time.Millisecond)

	assert.True(t, called)
	assert.Equal(t, 2, len(batches))
	assert.Equal(t, []interface{}{1, 2, 3}, batches[0])
	assert.Equal(t, []interface{}{4}, batches[1])
}

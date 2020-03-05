package egg

import (
	"sync"
	"time"
)

// Debouncer - batches messages based on an optional maximum batch threshold and wait time.
// this can be useful for implementing event handlers that would rather batch multiple simulataneous events.
// Use MakeDebouncer to construct an instance
type Debouncer struct {
	items   []interface{}
	batchF  func([]interface{})
	channel chan interface{}
	stop    chan bool
	waiting bool
	mux     sync.Mutex
	timer   *batchTimer
}

type batchTimer struct {
	channel chan bool
	delay   int
}

// MakeDebouncer ...
func MakeDebouncer() *Debouncer {
	return &Debouncer{
		items:   make([]interface{}, 0),
		batchF:  func([]interface{}) {},
		channel: make(chan interface{}, 1),
		stop:    make(chan bool),
		timer: &batchTimer{
			delay:   100,
			channel: make(chan bool),
		},
	}
}

func (d *Debouncer) Receive(batchF func([]interface{})) {
	d.batchF = batchF
}

func (d *Debouncer) Send(e interface{}) {
	if d.waiting {
		d.channel <- 0
	} else {
		d.mux.Lock()
		d.waiting = true
		d.mux.Unlock()
	}
	d.items = append(d.items, e)
	// d.items = append(d.items, e)
	// d.channel <- e
	go d.waitReceive()
}

func (d *Debouncer) waitReceive() {
	select {
	case <-d.channel:
		// do nothing
		// fmt.Println("Received an item")
		// d.items = append(d.items, e)
	case <-time.After(1 * time.Millisecond):
		d.mux.Lock()
		d.waiting = false
		d.mux.Unlock()

		d.batchF(d.items)
		d.items = make([]interface{}, 0)
	}
}

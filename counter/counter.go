package counter

import (
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

var globalCounter = &Counter{}

func Increment() int {
	globalCounter.mu.Lock()
	defer globalCounter.mu.Unlock()
	globalCounter.value++
	return globalCounter.value
}

func Get() int {
	globalCounter.mu.Lock()
	defer globalCounter.mu.Unlock()
	return globalCounter.value
}

func Reset() {
	globalCounter.mu.Lock()
	defer globalCounter.mu.Unlock()
	globalCounter.value = 0
}

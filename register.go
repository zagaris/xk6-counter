package xk6_entry

import (
	"github.com/zagaris/xk6-counter/counter"
	"go.k6.io/k6/js/modules"
)

type CounterModule struct{}

func (m *CounterModule) Increment() int {
	return counter.Increment()
}

func (m *CounterModule) Get() int {
	return counter.Get()
}

func (m *CounterModule) Reset() {
	counter.Reset()
}

func init() {
	modules.Register("k6/x/counter", &CounterModule{})
}

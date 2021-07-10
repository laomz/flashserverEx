package common

import (
	"sync"
)

type SafeBool struct {
	b    bool
	lock sync.RWMutex
}

func (sb *SafeBool) Set(v bool) {
	sb.lock.Lock()
	sb.b = v
	sb.lock.Unlock()
}

func (sb *SafeBool) Get() bool {
	sb.lock.RLock()
	b := sb.b
	sb.lock.RUnlock()
	return b
}

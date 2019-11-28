package mutex

import "sync"

// Mutex hides complexity of the sync.RWMutex
type Mutex struct {
	sync.RWMutex
}

// SoftRead acquires lock by calling RLock()
func (m *Mutex) SoftRead(f func()) {
	m.RLock()
	defer m.RUnlock()

	f()
}

// HardRead acquires lock by calling Lock()
func (m *Mutex) HardRead(f func()) {
	m.Lock()
	defer m.Unlock()

	f()
}

// Write acquires lock by calling Lock()
func (m *Mutex) Write(f func()) {
	m.Lock()
	defer m.Unlock()

	f()
}

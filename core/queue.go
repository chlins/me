// Package core provides some core libs.
package core

import "sync"

// BaseQueue is a simple queue interface
type BaseQueue interface {
	Push(interface{}) error
	Pop() interface{}
	Empty() bool
}

// MeQueue is my queue struct
type MeQueue struct {
	data []interface{}
	mu   sync.Mutex
}

// NewMeQueue makes a new meQueue
func NewMeQueue() *MeQueue {
	return &MeQueue{
		data: make([]interface{}, 0),
	}
}

// Push a element to end
func (m *MeQueue) Push(v interface{}) error {
	m.mu.Lock()
	m.data = append(m.data, v)
	m.mu.Unlock()
	return nil
}

// Pop a front element
func (m *MeQueue) Pop() interface{} {
	m.mu.Lock()
	if m.Empty() {
		return nil
	}
	defer func() {
		m.data = m.data[1:]
		m.mu.Unlock()
	}()
	return m.data[0]
}

// Empty checks data length
func (m *MeQueue) Empty() bool {
	return len(m.data) == 0
}

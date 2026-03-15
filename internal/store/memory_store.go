package store

import (
	"sync"
	"arna/internal/models"
)

type TaskStore interface {
	Create(task models.Task)
	List() []models.Task
}

type MemoryStore struct {
	mu    sync.RWMutex
	tasks []models.Task
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tasks: []models.Task{},
	}
}

func (m *MemoryStore) Create(task models.Task) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.tasks = append(m.tasks, task)
}

func (m *MemoryStore) List() []models.Task {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.tasks
}
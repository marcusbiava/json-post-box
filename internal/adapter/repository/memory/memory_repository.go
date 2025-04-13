package memory

import (
	"fmt"
	"sync"

	"github.com/marcusbiava/json-post-box/internal/domain"
)

// MemoryJSONRepository implements domain.JSONRepository interface with in-memory storage
type MemoryJSONRepository struct {
	data    map[string]*domain.JSON
	mu      sync.RWMutex
	counter int
}

// NewJSONRepository creates a new MemoryJSONRepository instance
func NewJSONRepository() domain.JSONRepository {
	return &MemoryJSONRepository{
		data:    make(map[string]*domain.JSON),
		counter: 0,
	}
}

// Store saves a JSON document in memory
func (r *MemoryJSONRepository) Store(data *domain.JSON) error {
	if data == nil {
		return domain.ErrInvalidJSON
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.counter++
	data.ID = fmt.Sprintf("%d", r.counter)
	r.data[data.ID] = data
	return nil
}

// FindByID retrieves a JSON document from memory by its ID
func (r *MemoryJSONRepository) FindByID(id string) (*domain.JSON, error) {
	if id == "" {
		return nil, domain.ErrInvalidJSON
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	data, exists := r.data[id]
	if !exists {
		return nil, domain.ErrJSONNotFound
	}

	return data, nil
}

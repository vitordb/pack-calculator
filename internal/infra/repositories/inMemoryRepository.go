package repositories

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"pack-calculator/internal/domain/ports"
)

// InMemoryResultRepository implements the DBInterface defined in the domain.
type InMemoryResultRepository struct {
	mu   sync.RWMutex
	data map[string]*ports.Result
}

// NewInMemoryResultRepository creates a new in-memory repository.
func NewInMemoryResultRepository() *InMemoryResultRepository {
	return &InMemoryResultRepository{
		data: make(map[string]*ports.Result),
	}
}

// SaveResult saves the calculation result. If no ID is provided, a new one is generated.
func (repo *InMemoryResultRepository) SaveResult(result *ports.Result) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	if result.ID == "" {
		result.ID = uuid.New().String()
	}
	repo.data[result.ID] = result
	return nil
}

// GetResultByID retrieves a saved result by its ID.
func (repo *InMemoryResultRepository) GetResultByID(id string) (*ports.Result, error) {
	repo.mu.RLock()
	defer repo.mu.RUnlock()

	if res, ok := repo.data[id]; ok {
		return res, nil
	}
	return nil, errors.New("result not found")
}

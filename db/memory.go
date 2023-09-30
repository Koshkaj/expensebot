package db

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/koshkaj/expensebot/config"
	"github.com/koshkaj/expensebot/types"
)

type MemoryDatabase struct {
	*config.MemoryConfig
	mu   sync.RWMutex
	data map[string]*types.Document
}

func NewMemoryDatabase() Databaser {
	return &MemoryDatabase{
		data: make(map[string]*types.Document),
	}
}

func (m *MemoryDatabase) Get(id uuid.UUID) (types.Document, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	doc, ok := m.data[id.String()]
	if !ok {
		return types.Document{}, fmt.Errorf("document with ID %s not found", id)
	}

	return *doc, nil
}

func (m *MemoryDatabase) Create(doc *types.Document) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.data[doc.Id]; ok {
		return fmt.Errorf("document with ID %s already exists", doc.Id)
	}

	m.data[doc.Id] = doc
	return nil
}

package db

import (
	"github.com/nguyenbatam/op-contract-test/backend/models"
	"sync"
)

type MemoryDB struct {
	docHash map[uint]*models.Document
	mu      sync.RWMutex
}

func (s *MemoryDB) SaveDocument(info *models.Document) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.docHash[info.Id] = info
}
func (s *MemoryDB) GetDocument(id uint) *models.Document {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.docHash[id]
}

func (s *MemoryDB) FindAllDocumentWaiting() []*models.Document {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var listWaiting []*models.Document
	for _, info := range s.docHash {
		if !info.ProofStatus {
			listWaiting = append(listWaiting, info)
		}
	}
	return listWaiting
}

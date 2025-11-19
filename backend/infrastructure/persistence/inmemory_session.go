package persistence

import (
	"errors"
	"grease_trace/domain/model"
	"grease_trace/domain/repository"
	"sync"
)

// InMemorySessionRepository は repository.SessionRepository の実装
type InMemorySessionRepository struct {
	sessions map[string]string // map[SessionID]PeerID
	mu       sync.RWMutex
}

func NewInMemorySessionRepository() repository.SessionRepository {
	return &InMemorySessionRepository{
		sessions: make(map[string]string),
	}
}

func (r *InMemorySessionRepository) Save(session *model.Session) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.sessions[session.ID] = session.SupporterPeerID
	return nil
}

func (r *InMemorySessionRepository) FindByID(id string) (*model.Session, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	peerID, ok := r.sessions[id]
	if !ok {
		return nil, errors.New("not found")
	}

	return &model.Session{
		ID:              id,
		SupporterPeerID: peerID,
	}, nil
}

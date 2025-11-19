package usecase

import (
	"errors"
	"grease_trace/domain/model"
	"grease_trace/domain/repository"

	"github.com/google/uuid"
)

// SessionInteractor はユースケースの実装
type SessionInteractor struct {
	Repo repository.SessionRepository
}

func NewSessionInteractor(repo repository.SessionRepository) SessionUsecase {
	return &SessionInteractor{
		Repo: repo,
	}
}

// CreateSession: IDを生成して保存する
func (i *SessionInteractor) CreateSession(peerID string) (string, error) {
	// UUIDの生成
	sessionID := uuid.New().String()

	session := &model.Session{
		ID:              sessionID,
		SupporterPeerID: peerID,
	}

	// リポジトリを使って保存
	if err := i.Repo.Save(session); err != nil {
		return "", err
	}

	return sessionID, nil
}

// ValidateSession: IDからPeerIDを探す
func (i *SessionInteractor) ValidateSession(sessionID string) (string, error) {
	session, err := i.Repo.FindByID(sessionID)
	if err != nil {
		return "", err
	}
	if session == nil {
		return "", errors.New("session not found")
	}

	return session.SupporterPeerID, nil
}

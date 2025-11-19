package usecase

type SessionUsecase interface {
	CreateSession(peerID string) (string, error)
	ValidateSession(sessionID string) (string, error)
}

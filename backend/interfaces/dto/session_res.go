package dto

type CreateSessionResponse struct {
	SessionID string `json:"sessionId"`
}

type ValidateSessionResponse struct {
	SupporterPeerID string `json:"supporterPeerId"`
}

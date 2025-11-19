package controller

import (
	"grease_trace/application/usecase"
	"grease_trace/interfaces/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionController struct {
	Usecase usecase.SessionUsecase
}

func NewSessionController(u usecase.SessionUsecase) *SessionController {
	return &SessionController{Usecase: u}
}

// Create ハンドラ
func (ctrl *SessionController) Create(c *gin.Context) {
	var req dto.CreateSessionRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.PeerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "peerId is required"})
		return
	}

	// ユースケースを実行
	sessionID, err := ctrl.Usecase.CreateSession(req.PeerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := dto.CreateSessionResponse{SessionID: sessionID}
	c.JSON(http.StatusOK, res)
}

// Validate ハンドラ
func (ctrl *SessionController) Validate(c *gin.Context) {
	sessionID := c.Param("sessionId")

	// ユースケースを実行
	peerID, err := ctrl.Usecase.ValidateSession(sessionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found or invalid"})
		return
	}

	res := dto.ValidateSessionResponse{SupporterPeerID: peerID}
	c.JSON(http.StatusOK, res)
}

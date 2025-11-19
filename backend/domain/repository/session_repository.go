package repository

import "grease_trace/domain/model"

// SessionRepository はデータの保存・取得のルール（インターフェース）
type SessionRepository interface {
	Save(session *model.Session) error
	FindByID(id string) (*model.Session, error)
}

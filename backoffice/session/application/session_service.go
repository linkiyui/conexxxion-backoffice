package application

import (
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/session/domain"
)

type ISessionService interface {
	Create(userID, deviceInfo, ip string) (*domain.Session, error)
	DeleteUserSessions(userID string) error
	GetByID(sessID string) (*domain.Session, error)
	Save(session domain.Session) error
}

type SessionService struct {
	sessionRepo domain.ISessionRepository
}

func NewSessionService(repo domain.ISessionRepository) *SessionService {
	return &SessionService{
		sessionRepo: repo,
	}
}

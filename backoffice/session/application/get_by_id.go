package application

import (
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/session/domain"
)

func (ss *SessionService) GetByID(sessID string) (*domain.Session, error) {
	return ss.sessionRepo.GetByID(sessID)
}

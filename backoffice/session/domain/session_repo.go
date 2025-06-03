package domain

import (
	uow "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/uow"
)

type ISessionRepository interface {
	GetByID(string) (*Session, error)
	GetByUserID(string) (*Session, error)
	Create(Session) error
	Delete(string) error
	DeleteUserSessions(string) error

	WithUnitOfWork(uow.UnitOfWork) ISessionRepository
}

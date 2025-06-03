package application

import "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/session/domain"

func (ss *SessionService) Save(session domain.Session) error {
	return ss.sessionRepo.Create(session)
}

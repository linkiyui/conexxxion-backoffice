package di_container

import (
	session_app "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/session/application"
	session_infra "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/session/infra"
	uow "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/uow"
)

func SessionService() *session_app.SessionService {
	repo := session_infra.NewSessionPostgreRepository(PostgreRepository())
	return session_app.NewSessionService(repo)
}

func SessionServiceUOW(uow *uow.PostgresUnitOfWork) *session_app.SessionService {
	repo := session_infra.NewSessionPostgreRepository(PostgreRepositoryWithUOW(uow))
	return session_app.NewSessionService(repo)
}

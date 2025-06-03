package di_container

import (
	auth_service "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/auth/application"
	uow "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/uow"
)

func AuthService() *auth_service.AuthService {
	sessSrv := SessionService()
	userSrv := UserService()

	return auth_service.NewAuthService(userSrv, sessSrv)

}

func AuthServiceUOW(uow *uow.PostgresUnitOfWork) *auth_service.AuthService {
	sessSrv := SessionServiceUOW(uow)
	userSrv := UserServiceUOW(uow)

	return auth_service.NewAuthService(userSrv, sessSrv)
}

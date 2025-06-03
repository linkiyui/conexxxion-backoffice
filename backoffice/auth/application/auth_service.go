package application

import (
	session_app "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/session/application"
	user_service "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/application"
)

type AuthService struct {
	userSrv    user_service.IUserService
	sessionSrv *session_app.SessionService
}

func NewAuthService(

	userSrv user_service.IUserService,
	sessionSrv *session_app.SessionService) *AuthService {

	return &AuthService{

		userSrv:    userSrv,
		sessionSrv: sessionSrv,
	}
}

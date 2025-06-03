package application

import (
	domain_errors "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/errors"
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/domain"
	"gitlab.com/conexxxion/conexxxion-backoffice/utils"
)

func (as *AuthService) Login(email, password string) (*domain.User, error) {
	user, err := as.userSrv.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPasswordHashFunc(password, user.Password) {
		return nil, domain_errors.ErrPasswordNotmatch
	}
	return user, nil
}

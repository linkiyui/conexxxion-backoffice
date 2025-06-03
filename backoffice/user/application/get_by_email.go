package application

import (
	domain_errors "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/errors"
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/domain"
)

func (us *UserService) GetByEmail(email string) (*domain.User, error) {
	user, err := us.userRepo.GetByEmail(email)

	if user != nil {
		return user, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, domain_errors.ErrUserNotFound

}

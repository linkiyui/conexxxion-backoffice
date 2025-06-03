package application

import (
	domain_errors "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/errors"
)

func (us *UserService) CheckByUsername(username string) error {
	user, err := us.userRepo.GetByUsername(username)

	if user != nil {
		return domain_errors.ErrUsernameExists
	}

	if err != nil {
		return err
	}

	return nil

}

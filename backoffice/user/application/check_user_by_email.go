package application

import (
	domain_errors "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/errors"
)

func (us *UserService) CheckByEmail(email string) error {
	user, err := us.userRepo.GetByEmail(email)

	if user != nil {
		return domain_errors.ErrEmailExists
	}

	if err != nil {
		return err
	}
	return nil

}

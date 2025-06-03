package application

import (
	domain_errors "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/errors"
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/domain"
	"gitlab.com/conexxxion/conexxxion-backoffice/utils"
)

func (us *UserService) UpdateUser(user *domain.User) error {
	// TODO: check if username is valid
	is_valid_username := utils.IsValidUsername(user.Username)
	if !is_valid_username {
		return domain_errors.ErrUsernameNotValid
	}

	// TODO: check if username is already taken
	err := us.CheckByUsername(user.Username)
	if err != nil {
		return err
	}

	// TODO: check if email is valid
	is_valid_email := utils.IsValidEmail(user.Email)
	if !is_valid_email {
		return domain_errors.ErrEmailNotValid
	}

	// TODO: check if email is already taken
	err = us.CheckByEmail(user.Email)
	if err != nil {
		return err
	}

	// TODO: update user
	err = us.userRepo.Save(user)
	if err != nil {
		return err
	}

	return nil
}

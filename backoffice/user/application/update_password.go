package application

import (
	domain_errors "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/errors"
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/domain"
	"gitlab.com/conexxxion/conexxxion-backoffice/utils"
)

func (us *UserService) UpdatePassword(user *domain.User, password string) error {
	// TODO: check if password is valid
	is_valid_password := utils.IsValidPassword(password)
	if !is_valid_password {
		return domain_errors.ErrPasswordNotValid
	}
	// encrypt password
	enc_pass := utils.HashingPasswordFunc(password)

	// TODO: update password
	err := us.userRepo.UpdatePassword(user.ID, enc_pass)
	if err != nil {
		return err
	}

	return nil
}

package application

import (
	"time"

	domain_errors "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/errors"
	user_domain "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/domain"
	"gitlab.com/conexxxion/conexxxion-backoffice/utils"
)

func (as *AuthService) CreateUser(user user_domain.User) error {

	// TODO: check if username is already taken
	err := as.userSrv.CheckByUsername(user.Username)
	if err != nil {
		return err
	}

	// TODO: check if email is valid
	is_valid_email := utils.IsValidEmail(user.Email)
	if !is_valid_email {
		return domain_errors.ErrEmailNotValid
	}

	// TODO: check if email is already taken
	err = as.userSrv.CheckByEmail(user.Email)
	if err != nil {
		return err
	}
	// TODO: check if password is valid
	is_valid_password := utils.IsValidPassword(user.Password)
	if !is_valid_password {
		return domain_errors.ErrPasswordNotValid
	}
	// encrypt password
	enc_pass := utils.HashingPasswordFunc(user.Password)

	// TODO: create user

	id, err := utils.GenerateUUIDv7()
	if err != nil {
		return err
	}

	// TODO: check if username is valid
	is_valid_username := utils.IsValidUsername(user.Username)
	if !is_valid_username {
		return domain_errors.ErrUsernameNotValid
	}

	user_to_repo := user_domain.User{
		ID:       id,
		Email:    user.Email,
		Password: enc_pass,
		Username: user.Username,
		Name:     user.Name,
		LastName: user.LastName,
		Role:     user.Role,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	err = as.userSrv.CreateUser(&user_to_repo)
	if err != nil {
		return err
	}

	return nil
}

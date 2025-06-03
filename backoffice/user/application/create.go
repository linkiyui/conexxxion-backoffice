package application

import "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/domain"

func (us *UserService) CreateUser(user *domain.User) error {
	return us.userRepo.Create(user)
}

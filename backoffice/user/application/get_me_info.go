package application

import "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/domain"

func (us *UserService) GetMeInfo(user_id string) (*domain.User, error) {
	return us.userRepo.GetByID(user_id)
}

package application

import (
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/domain"
)

type IUserService interface {
	CheckByEmail(email string) error
	CheckByUsername(username string) error
	CreateUser(user *domain.User) error
	GetByEmail(email string) (*domain.User, error)
	GetMeInfo(user_id string) (*domain.User, error)
	UpdatePassword(user *domain.User, password string) error
	// CreateUser(user *domain.User) error
	UpdateUser(user *domain.User) error
	DeleteUser(id string) error
}

type UserService struct {
	userRepo domain.IUserRepository
}

func NewUserService(repo domain.IUserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

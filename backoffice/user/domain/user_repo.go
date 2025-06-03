package domain

type IUserRepository interface {
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	Create(user *User) error
	UpdatePassword(id string, password string) error
	Save(user *User) error
	// Update(user *User) error
	Delete(id string) error
}

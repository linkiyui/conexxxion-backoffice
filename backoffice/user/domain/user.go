package domain

import "time"

type User struct {
	ID       string    `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Password string    `json:"password"`
	Role     Role      `json:"role"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type Role string

const (
	SupportRole    Role = "support"
	AdminRole      Role = "admin"
	SuperAdminRole Role = "super_admin"
)

package domain

import "time"

type Session struct {
	ID         string
	UserID     string
	DeviceInfo string
	IP         string
	LastAccess time.Time
	CreatedAt  time.Time
}

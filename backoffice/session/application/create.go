package application

import (
	"time"

	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/session/domain"
	"gitlab.com/conexxxion/conexxxion-backoffice/utils"
)

func (ss *SessionService) Create(userID string, deviceInfo string, ip string) (*domain.Session, error) {
	sessionID, err := utils.GenerateUUIDv7()
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()

	sess := domain.Session{
		ID:         sessionID,
		UserID:     userID,
		DeviceInfo: deviceInfo,
		IP:         ip,
		LastAccess: now,
		CreatedAt:  now,
	}

	if err = ss.sessionRepo.Create(sess); err != nil {
		return nil, err
	}

	return &sess, nil
}

package application

func (ss *SessionService) DeleteUserSessions(userID string) error {
	return ss.sessionRepo.DeleteUserSessions(userID)
}

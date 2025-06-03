package application

func (us *UserService) DeleteUser(id string) error {
	// TODO: check if user is blocked
	err := us.userRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

package di_container

import (
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/uow"
	user_service "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/application"
	user_infra "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/infra"
)

func UserService() user_service.IUserService {
	repo := user_infra.NewUserPostgreRepository(PostgreRepository())
	return user_service.NewUserService(repo)
}

func UserServiceUOW(uow *uow.PostgresUnitOfWork) user_service.IUserService {
	repo := user_infra.NewUserPostgreRepository(PostgreRepositoryWithUOW(uow))
	return user_service.NewUserService(repo)
}

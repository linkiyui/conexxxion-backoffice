package di_container

import (
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/shared/infra"
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/uow"
	"gitlab.com/conexxxion/conexxxion-backoffice/database"
)

func PostgreRepository() *infra.PostgreRepository {
	return infra.NewPostgreRepository(database.Get())
}

func PostgreRepositoryWithUOW(uow uow.UnitOfWork) *infra.PostgreRepository {
	return infra.NewPostgreRepository(database.Get()).WithUnitOfWork(uow)
}

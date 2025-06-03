package di_container

import (
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/uow"
	"gitlab.com/conexxxion/conexxxion-backoffice/database"
)

func PostgreUOW() *uow.PostgresUnitOfWork {
	return uow.NewPostgresUOW(database.Get())
}

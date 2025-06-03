package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/conexxxion/conexxxion-backoffice/di_container"
)

func DeleteUser(ctx *gin.Context) {
	var (
		err     error
		user_id string
	)

	user_id = ctx.Param("id")

	uow := di_container.PostgreUOW()
	if err := uow.BeginTransaction(ctx); err != nil {
		ctx.AbortWithStatus(500)
		return
	}
	defer uow.Rollback(ctx)

	user_service := di_container.UserServiceUOW(uow)

	err = user_service.DeleteUser(user_id)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	if err := uow.Commit(ctx); err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	ctx.AbortWithStatus(200)
}

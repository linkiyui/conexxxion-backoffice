package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/conexxxion/conexxxion-backoffice/di_container"
)

func DeleteMeUser(ctx *gin.Context) {
	var (
		uid     any
		exists  bool
		err     error
		user_id string
	)

	uid, exists = ctx.Get("user_id")
	if !exists {
		ctx.AbortWithStatus(400)
		return
	}

	if uid != nil {
		user_id = uid.(string)
	}

	uow := di_container.PostgreUOW()
	if err := uow.BeginTransaction(ctx); err != nil {
		ctx.AbortWithStatus(500)
		return
	}
	defer uow.Rollback(ctx)

	user_service := di_container.UserServiceUOW(uow)

	u, err := user_service.GetMeInfo(user_id)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	err = user_service.DeleteUser(u.ID)
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

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gitlab.com/conexxxion/conexxxion-backoffice/di_container"
)

type UpdatePasswordRequest struct {
	Password string `json:"password"`
}

func UpdatePassword(ctx *gin.Context) {
	var req UpdatePasswordRequest

	var (
		uid     any
		err     error
		user_id string
		exists  bool
	)

	uid, exists = ctx.Get("user_id")
	if !exists {
		ctx.AbortWithStatus(400)
		return
	}

	if uid != nil {
		user_id = uid.(string)
	}

	user, err := di_container.UserService().GetMeInfo(user_id)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	if err := ctx.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		ctx.AbortWithStatus(400)
		return
	}

	uow := di_container.PostgreUOW()
	if err := uow.BeginTransaction(ctx); err != nil {
		ctx.AbortWithStatus(500)
		return
	}
	defer uow.Rollback(ctx)

	user_service := di_container.UserServiceUOW(uow)

	err = user_service.UpdatePassword(user, req.Password)
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

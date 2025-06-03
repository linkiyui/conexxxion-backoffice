package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gitlab.com/conexxxion/conexxxion-backoffice/di_container"
)

type UpdateUserRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

func UpdateUser(ctx *gin.Context) {
	var req UpdateUserRequest

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

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}

	err = user_service.UpdateUser(user)
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

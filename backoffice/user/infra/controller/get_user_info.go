package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/conexxxion/conexxxion-backoffice/di_container"
)

func GetUserInfo(ctx *gin.Context) {
	var (
		err     error
		user_id string
	)

	user_id = ctx.Param("id")
	user_service := di_container.UserService()

	u, err := user_service.GetMeInfo(user_id)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	user_info := user_info{
		ID:       u.ID,
		Email:    u.Email,
		Username: u.Username,
		Name:     u.Name,
		LastName: u.LastName,
		Role:     string(u.Role),
		CreateAt: u.CreateAt.Format("2006-01-02 15:04:05"),
		UpdateAt: u.UpdateAt.Format("2006-01-02 15:04:05"),
	}

	ctx.JSON(200, gin.H{
		"user": user_info,
	})

}

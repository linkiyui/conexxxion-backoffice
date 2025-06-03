package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/conexxxion/conexxxion-backoffice/di_container"
)

func GetMeInfo(ctx *gin.Context) {
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

type user_info struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Role     string `json:"role"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

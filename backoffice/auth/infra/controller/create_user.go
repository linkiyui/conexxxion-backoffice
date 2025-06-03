package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	domain_errors "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/errors"
	user_domain "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/user/domain"
	"gitlab.com/conexxxion/conexxxion-backoffice/di_container"
	clog "gitlab.com/conexxxion/conexxxion-backoffice/logger"
)

type SignupRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func CreateUser(ctx *gin.Context) {

	var req SignupRequest

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		clog.ErrorCtx(ctx, "binding request body | error: "+err.Error(), nil)
		de := domain_errors.IsDomainError(err)
		if de != nil {
			ctx.JSON(de.Code, gin.H{"error": de.Message})
			return
		}
		ctx.AbortWithStatus(500)
		return
	}

	uow := di_container.PostgreUOW()
	if err := uow.BeginTransaction(ctx); err != nil {
		clog.ErrorCtx(ctx, "starting transaction | error: "+err.Error(), nil)
		de := domain_errors.IsDomainError(err)
		if de != nil {
			ctx.JSON(de.Code, gin.H{"error": de.Message})
			return
		}
		ctx.AbortWithStatus(500)
		return

	}
	defer uow.Rollback(ctx)

	auth_service := di_container.AuthServiceUOW(uow)

	user := user_domain.User{
		Email:    req.Email,
		Username: req.Username,
		Name:     req.Name,
		LastName: req.LastName,
		Password: req.Password,
		Role:     user_domain.Role(req.Role),
	}

	err := auth_service.CreateUser(user)
	if err != nil {
		clog.ErrorCtx(ctx, "signup | error: "+err.Error(), nil)
		de := domain_errors.IsDomainError(err)
		if de != nil {
			ctx.JSON(de.Code, gin.H{"error": de.Message})
			return
		}
		fmt.Println(err)
		ctx.AbortWithStatus(500)
		return
	}

	if err := uow.Commit(ctx); err != nil {
		clog.ErrorCtx(ctx, "committing transaction | error: "+err.Error(), nil)
		de := domain_errors.IsDomainError(err)
		if de != nil {
			ctx.JSON(de.Code, gin.H{"error": de.Message})
			return
		}
		fmt.Println(err)
		ctx.AbortWithStatus(500)
		return
	}

	ctx.AbortWithStatus(200)

}

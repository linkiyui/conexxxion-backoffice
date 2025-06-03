package controller

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/conexxxion/conexxxion-backoffice/backoffice/auth/infra/auth_token"
	domain_errors "gitlab.com/conexxxion/conexxxion-backoffice/backoffice/errors"
	"gitlab.com/conexxxion/conexxxion-backoffice/di_container"
	clog "gitlab.com/conexxxion/conexxxion-backoffice/logger"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {

	var req LoginRequest

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		clog.ErrorCtx(ctx, "binding request body | error: "+err.Error(), nil)
		de := domain_errors.IsDomainError(err)
		if de != nil {
			ctx.JSON(de.Code, gin.H{"error": de.Message})
			return
		}
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
		return

	}
	defer func() {
		uow.Rollback(ctx)
	}()

	auth_service := di_container.AuthServiceUOW(uow)

	user, err := auth_service.Login(req.Email, req.Password)
	if err != nil {
		clog.ErrorCtx(ctx, "signup | error: "+err.Error(), nil)
		de := domain_errors.IsDomainError(err)
		if de != nil {
			ctx.JSON(de.Code, gin.H{"error": de.Message})
			return
		}
		return
	}

	role := string(user.Role)

	toekn, err := auth_token.GenerateLoginToken(user.ID, role)
	if err != nil {
		clog.ErrorCtx(ctx, "signup | error: "+err.Error(), nil)
		de := domain_errors.IsDomainError(err)
		if de != nil {
			ctx.JSON(de.Code, gin.H{"error": de.Message})
			return
		}
		return
	}

	refresh_token, err := auth_token.GenerateRefreshToken(user.ID, role)
	if err != nil {
		clog.ErrorCtx(ctx, "signup | error: "+err.Error(), nil)
		de := domain_errors.IsDomainError(err)
		if de != nil {
			ctx.JSON(de.Code, gin.H{"error": de.Message})
			return
		}
		return
	}

	if err := uow.Commit(ctx); err != nil {
		clog.ErrorCtx(ctx, "committing transaction | error: "+err.Error(), nil)
		de := domain_errors.IsDomainError(err)
		if de != nil {
			ctx.JSON(de.Code, gin.H{"error": de.Message})
			return
		}
		return
	}

	ctx.JSON(200, gin.H{
		"token":         toekn,
		"refresh_token": refresh_token,
	})

}

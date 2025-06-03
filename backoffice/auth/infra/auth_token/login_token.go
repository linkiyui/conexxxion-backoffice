package auth_token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/conexxxion/conexxxion-backoffice/config"
)

func GenerateLoginToken(user_id string, Role string) (string, error) {
	exp := config.GetConfig().JWTOptions.LoginExpirationTime
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"role":    Role,
		"exp":     time.Now().Add(time.Hour * time.Duration(exp)).Unix(),
		"iat":     time.Now().Unix(),
		"sub":     "login",
	})
	secret := config.GetConfig().JWTOptions.Secret
	return tkn.SignedString([]byte(secret))
}

func ValidateLoginToken(token string) (*LoginClaims, error) {

	secret := config.GetConfig().JWTOptions.Secret

	jwtKeyFunc := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	}

	tkn, err := jwt.ParseWithClaims(token, &LoginClaims{}, jwtKeyFunc, jwt.WithExpirationRequired(), jwt.WithSubject("login"))

	if err != nil {
		return nil, err
	}

	claims, ok := tkn.Claims.(*LoginClaims)
	if !ok {
		return nil, fmt.Errorf("can't parse claims: %v", tkn.Claims)
	}

	return claims, nil
}

type LoginClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`

	jwt.RegisteredClaims
}

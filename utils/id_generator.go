package utils

import (
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	clog "gitlab.com/conexxxion/conexxxion-backoffice/logger"
)

func GenerateULID() string {
	return ulid.Make().String()
}

func GenerateUUIDv4() (string, error) {
	res, err := uuid.NewRandom()
	if err != nil {
		clog.Error(err.Error(), nil)
		return "", err
	}
	return res.String(), nil
}

func GenerateUUIDv7() (string, error) {
	res, err := uuid.NewV7()
	if err != nil {
		clog.Error(err.Error(), nil)
		return "", err
	}
	return res.String(), nil
}

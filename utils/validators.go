package utils

import (
	"regexp"

	log "gitlab.com/conexxxion/conexxxion-backoffice/logger"
)

var usernameRegex *regexp.Regexp

func init() {
	var err error
	usernameRegex, err = regexp.Compile("^[A-Za-z]([A-Za-z0-9_]){2,15}[a-zA-Z0-9]$")
	if err != nil {
		log.Fatal("compiling username regexp: "+err.Error(), nil)
	}

}

func IsValidPhoneNumber(phone string) bool {
	r, _ := regexp.Compile("^[0-9]{3,15}$")
	return r.Match([]byte(phone))
}

func IsValidEmail(email string) bool {
	r, _ := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return r.Match([]byte(email))
}

func IsValidUsername(username string) bool {
	// Empieza por letra, puede contener letras, numeros y underscore
	// no puede terminar en underscore, maximo 15 caracteres
	return usernameRegex.Match([]byte(username))
}

func IsValidPassword(password string) bool {

	if len(password) < 8 {
		return false
	}
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*]`).MatchString(password)
	return hasLower && hasUpper && hasNumber && hasSpecial
}

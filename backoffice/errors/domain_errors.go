package errors

var errors = map[string]*DomainError{}

type DomainError struct {
	Code        int
	Message     string
	Description string
}

func (de DomainError) Error() string {
	return de.Message
}

func IsDomainError(e any) *DomainError {
	if de, ok := e.(DomainError); ok {
		return &de
	}
	if de, ok := e.(*DomainError); ok {
		return de
	}
	return nil
}

func (de DomainError) Equal(e *DomainError) bool {
	return de.Message == e.Message
}

func newDE(code int, message, description string) DomainError {
	de := &DomainError{
		Code:        code,
		Message:     message,
		Description: description,
	}
	errors[message] = de
	return *de
}

var (
	ErrBadRequest            = newDE(400, "ERR_BAD_REQUEST", "Bad request, check your request")
	ErrForbidden             = newDE(403, "ERR_FORBIDDEN", "Forbidden user")
	ErrPhoneNumberNotValid   = newDE(400, "ERR_PHONE_NUMBER_NOT_VALID", "phone number not valid , check your request")
	ErrCodeAlreadyRequested  = newDE(400, "ERR_CODE_ALREADY_REQUESTED", "code alredy requested, wait a moment")
	ErrPhoneNumberExists     = newDE(400, "ERR_PHONE_NUMBER_EXISTS", "Phone number exists in database, you can't use it")
	ErrUsernameExists        = newDE(400, "ERR_USERNAME_EXISTS", "username exists in database, you can't use it")
	ErrUsernameCantChangeYet = newDE(400, "ERR_USERNAME_CANT_CHANGE_YET", "Username alredy changed, wait a moment")
	ErrCodeExpired           = newDE(400, "ERR_CODE_EXPIRED_OR_NOT_VALID", "your code expired or is not valid")
	ErrEmailExists           = newDE(400, "ERR_EMAIL_EXISTS", "Email exists sin database, you can't use it")
	ErrEmailNotFound         = newDE(404, "ERR_EMAIL_NOT_FOUND", "Email not found ")
	ErrEmailNotValid         = newDE(400, "ERR_EMAIL_NOT_VALID", "Email is not valid")
	ErrPasswordNotValid      = newDE(400, "ERR_PASSWORD_NOT_VALID", "Password is not valid")
	ErrPasswordNotmatch      = newDE(400, "ERR_PASSWORD_NOT_MATCH", "Password not match")

	ErrUserAlredyInTheGroupWeCameFor = newDE(400, "ERR_USER_ALREDY_IN_THE_GROUP_WHAT_WE_CAME_FOR", "User alredy in the group whatwe came for")
	ErrUserIsNotInTheGroupWeCameFor  = newDE(400, "ERR_USER_IS_NOT_IN_THE_GROUP_WHAT_WE_CAME_FOR", "User is not in the group whatwe came for")

	ErrProfileTypeNotValidToBeUnicorn = newDE(400, "ERR_PROFILE_TYPE_NOT_VALID_TO_BE_UNICORN", "Your profile type is not valid to be an unicorn")
	ErrUserAlredyUnicorn              = newDE(400, "ERR_USER_ALREDY_UNICORN", "User alredy is an unicorn")
	ErrUserNotUnicorn                 = newDE(400, "ERR_USER_NOT_UNICORN", "User is not aan unicorn")
	ErrGenderNotValidToBeUnicorn      = newDE(400, "ERR_GENDER_NOT_VALID_TO_BE_UNICORN", "Your gender is not valid to be an unicorn")

	ErrProfileTypeNotValidToBeBull = newDE(400, "ERR_PROFILE_TYPE_NOT_VALID_TO_BE_BULL", "Your profile type is not valid to be a bull")
	ErrUserAlredyBull              = newDE(400, "ERR_USER_ALREDY_BULL", "User alredy is a bull")
	ErrUserNotBull                 = newDE(400, "ERR_USER_NOT_BULL", "User is not a bull")
	ErrGenderNotValidToBeBull      = newDE(400, "ERR_GENDER_NOT_VALID_TO_BE_BULL", "Your gender is not valid to be aa bull")

	ErrFreeMembership     = newDE(400, "ERR_FREE_MEMBERSHIP", "You are using a free membership")
	ErrNotMembershipGold  = newDE(400, "ERR_NOT_MEMBERSHIP_GOLD", "Is not a gold membership")
	ErrInvalidMembership  = newDE(400, "ERR_INVALID_MEMBERSHIP", "invalid membership")
	ErrMembershipNotFound = newDE(404, "ERR_MEMBERSHIP_NOT_FOUND", "membership not found")

	ErrCantMakeConection   = newDE(400, "ERR_CANT_MAKE_CONECTION", "you can't make the connection")
	ErrCantcancelConection = newDE(400, "ERR_USER_CANT_CANCEL_CONECTION", "you can't cancel the connection ")
	ErrCantAcceptConection = newDE(400, "ERR_USER_CANT_ACCEPT_CONECTION", "you can't accept the connection ")
	ErrConectionNotExist   = newDE(404, "ERR_CONECTION_NOT_EXIST", "connection don't exists")

	ErrGenderNotValid             = newDE(400, "ERR_GENDER_NOT_VALID", "gender not valid")
	ErrProfileTypeNotValid        = newDE(400, "ERR_PROFILE_TYPE_NOT_VALID", "Profile type not valid")
	ErrBodyTypeNotValid           = newDE(400, "ERR_BODY_TYPE_NOT_VALID", "Body type not valid")
	ErrAgeNotValid                = newDE(400, "ERR_AGE_NOT_VALID", "Age not valid")
	ErrUsernameNotValid           = newDE(400, "ERR_USERNAME_NOT_VALID", "Username is not valid")
	ErrUsernameTooLongDescription = newDE(400, "ERR_USERNAME_TOO_LONG_DESCRIPTION", "Username too long Description")
	ErrNameNotValid               = newDE(400, "ERR_NAME_NOT_VALID", "name not valid")
	ErrHeightNotValid             = newDE(400, "ERR_HEIGHT_NOT_VALID", "Height not valid")
	ErrNotEnoughPoints            = newDE(400, "ERR_NOT_ENOUGH_SUPER_HOT_POINTS", "You have not enough superhot points to gift")
	ErrCantPoints                 = newDE(400, "ERR_CANT_POINTS", "Wrong cant of points")

	ErrIsYourSelf              = newDE(400, "ERR_IS_YOUR_OWN_USER", "you can't give you superhot points or make a connection with your own profile")
	ErrUserAlredyBlockedForYou = newDE(400, "ERR_USER_ALREDY_BLOCKED_FOR_YOU", "this user is blocked for you, yu can't block him again")
	ErrUserIsNotBlockedForYou  = newDE(400, "ERR_USER_IS_NOT_BLOCKED_FOR_YOU", "this user is not blocked for you, yu can't unblock him ")
	ErrBlockdUser              = newDE(403, "ERR_BLOCKED_USER", "You blocked this user")
	ErrBlockedByUser           = newDE(403, "ERR_BLOCKED_BY_USER", "You are blocked by this user")

	ErrUserNotFound               = newDE(404, "ERR_USER_NOT_FOUND", "user not found")
	ErrUserBlocked                = newDE(403, "ERR_USER_BLOCKED", "user are blocked")
	ErrUserProfileImageWrongOrder = newDE(400, "ERR_USER_PROFILE_IMAGE_WRONG_ORDER", "Wrong order of profile image")
	ErrUserProfileImageTooBig     = newDE(400, "ERR_USER_PROFILE_IMAGE_TOO_BIG", "Profile image too big")
	ErrSendSMSCode                = newDE(403, "ERR_SEND_SMS_CODE", "Error to send sms code")
	ErrTooManyAttempts            = newDE(429, "ERR_TOO_MANY_ATTEMPTS", "Too many attempts. Try later")
)

func GetAllDomainErrors() []DomainError {
	e := make([]DomainError, 0, len(errors))
	for _, v := range errors {
		e = append(e, *v)
	}
	return e
}

package helper

import "errors"

func HttpStatusErr(err error) int {
	switch {
	case errors.Is(err, ErrRegistration):
		return 400
	case errors.Is(err, ErrEmailOrPasswordWrong):
		return 401
	default:
		return 500
	}
}

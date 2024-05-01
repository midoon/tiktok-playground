package helper

import "errors"

var ErrRegistration = errors.New("registration error, please check your username, email, or password form")
var ErrEmailOrPasswordWrong = errors.New("login  error, email or password wrong")

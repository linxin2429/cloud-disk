package utils

import (
	"errors"
	"fmt"
)

var (
	ErrUsernameOrPasswd = errors.New("username or passwd error")
	ErrUserNotFound = errors.New("user not found")
	ErrEmailRegistered = errors.New("the email already registered")
	ErrCaptchaRepeated = errors.New("Send the captcha code repeatedly")
	ErrCaptchaNotMatch = errors.New("the captcha not match")
	ErrUserExist = errors.New("username already exist")
)

type ErrWrapper struct {
	Err error
	Opt string
}

func NewErrWrapper(err error, opt string) *ErrWrapper {
	return &ErrWrapper{
		Err: err,
		Opt: opt,
	}
}

func (e *ErrWrapper) Error() string {
	return fmt.Sprintf("%s: Error: %s", e.Opt, e.Err.Error())
}

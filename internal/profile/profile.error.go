package profile

import (
	"errors"
	"profile-template/internal/app/err"
)

var (
	ErrNotFound           = errors.New("user not found")
	ErrDuplicateTag       = errors.New("tag already exists")
	ErrInvalidUsername    = errors.New("invalid username")
	ErrTagAlreadyAssigned = errors.New("tag already assigned")
	ErrInvalidInput       = errors.New("invalid input profile")
)

const (
	CodeUserNotFound       err.ErrCode = "USER_NOT_FOUND"
	CodeTagAlreadyAssigned err.ErrCode = "TAG_ALREADY_ASSIGNED"
	CodeInvalidUsername    err.ErrCode = "INVALID_USERNAME"
	CodeInvalidInput       err.ErrCode = "INVALID_INPUT"
)

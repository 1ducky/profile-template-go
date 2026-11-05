package article

import (
	"errors"
)

var (
	ErrNotFound        = errors.New("article: not found")
	ErrInvalidID       = errors.New("article: id is required")
	ErrDuplicateSlug   = errors.New("article: slug already exists")
	ErrNothingToUpdate = errors.New("article: nothing to update")
	ErrUnauthenticated = errors.New("article: author id is required")
	ErrForbidden       = errors.New("article: not the owner of this article")
	ErrInvalidInput    = errors.New("article: invalid input")
	ErrBadRequest      = errors.New("article: Bad Request")
)

// const (
// 	CodeNotFound err.ErrCode = "ARTICLE_NOT_FOUND"
// 	CodeInvalidID err.ErrCode = "AUTHOR_IDRE"
// 	CodeNotFound err.ErrCode = "ARTICLE_NOT_FOUND"
// 	CodeNotFound err.ErrCode = "ARTICLE_NOT_FOUND"
// 	CodeNotFound err.ErrCode = "ARTICLE_NOT_FOUND"
// )

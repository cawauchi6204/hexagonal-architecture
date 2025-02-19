package base

import "errors"

var (
	// User関連のエラー
	ErrEmptyUsername = errors.New("username cannot be empty")
	ErrEmptyEmail    = errors.New("email cannot be empty")
	ErrEmptyPassword = errors.New("password cannot be empty")
	ErrInvalidEmail  = errors.New("invalid email format")

	// Thread関連のエラー
	ErrEmptyTitle    = errors.New("thread title cannot be empty")
	ErrInvalidUserID = errors.New("invalid user id")

	// Post関連のエラー
	ErrEmptyContent    = errors.New("post content cannot be empty")
	ErrInvalidThreadID = errors.New("invalid thread id")

	// Comment関連のエラー
	ErrEmptyCommentContent = errors.New("comment content cannot be empty")
	ErrInvalidPostID       = errors.New("invalid post id")

	// Tag関連のエラー
	ErrEmptyTagName = errors.New("tag name cannot be empty")
)

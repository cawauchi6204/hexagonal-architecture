package base

import "time"

type User struct {
	ID           string
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

func NewUser(username, email, passwordHash string) *User {
	return &User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}
}

func (u *User) Validate() error {
	if u.Username == "" {
		return ErrEmptyUsername
	}
	if u.Email == "" {
		return ErrEmptyEmail
	}
	if u.PasswordHash == "" {
		return ErrEmptyPassword
	}
	return nil
}

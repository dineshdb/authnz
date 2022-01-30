package user

import (
	"errors"
)

type User struct {
	ID           int64  `json:"id,omitempty"`
	Email        string `json:"email,omitempty"`
	Name         string `json:"name,omitempty"`
	Location     string `json:"location,omitempty"`
	About        string `json:"about,omitempty"`
	PasswordHash string `json:"-"`
}

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExists    = errors.New("row not exists")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

package users

import (
	authModel "tinderleaked/models/auth"
	usersModel "tinderleaked/models/users"
)

type UsecaseUsers interface {
	ListAnotherUsers(userId int) (*[]authModel.ResponseUser, error)
	SubmitActionUser(userId int, req usersModel.SubmitActionUser) error
}

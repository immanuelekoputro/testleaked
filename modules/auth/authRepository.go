package auth

import authModel "tinderleaked/models/auth"

type RepositoryAuth interface {
	SaveRegisterUsers(request *authModel.Users) error
	FetchUserByAlias(email string) (*authModel.Users, error)
	FetchUserByID(userId int) (*authModel.ResponseUser, error)
}

package auth

import authModel "tinderleaked/models/auth"

type UsecaseAuth interface {
	SubmitRegisterAccount(req *authModel.RegisterRequest) error
	Login(req *authModel.LoginRequest) (*authModel.ResponseLogin, error)
}

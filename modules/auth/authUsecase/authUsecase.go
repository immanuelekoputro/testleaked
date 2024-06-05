package authUsecase

import (
	"errors"
	"time"
	authModel "tinderleaked/models/auth"
	"tinderleaked/modules/auth"
	"tinderleaked/tools/generator/pwd"
	"tinderleaked/tools/jwt"
)

type authUsecase struct {
	authRepository auth.RepositoryAuth
}

func NewAuthUsecase(authRepository auth.RepositoryAuth) auth.UsecaseAuth {
	return &authUsecase{
		authRepository: authRepository,
	}
}

func (usecase *authUsecase) SubmitRegisterAccount(req *authModel.RegisterRequest) error {

	//	Check Date Format
	_, err := time.Parse("2006-01-02 15:04:05", req.DateOfBirthday)
	if err != nil {
		return errors.New("failed, DOB format not as expected. required format YY-MM-DD HH:MM:SS")
	}

	// Generate HASH
	hashedPwd, errHashedPwd := pwd.GeneratePasswordHash([]byte(req.Password))
	if errHashedPwd != nil {
		return errHashedPwd
	}

	users := authModel.Users{
		Name:           req.Name,
		Email:          req.Email,
		Password:       hashedPwd,
		Gender:         req.Gender,
		DateOfBirthday: req.DateOfBirthday,
	}

	errSubmit := usecase.authRepository.SaveRegisterUsers(&users)
	if errSubmit != nil {
		return errSubmit
	}

	return nil
}

func (usecase *authUsecase) Login(req *authModel.LoginRequest) (*authModel.ResponseLogin, error) {
	getUserDetail, errGetUserDetail := usecase.authRepository.FetchUserByAlias(req.Email)
	if errGetUserDetail != nil {
		return nil, errGetUserDetail
	}

	// test password
	checkPassword := pwd.CheckPasswordHash(req.Password, getUserDetail.Password)
	if !checkPassword {
		return nil, errors.New("wrong authentication")
	}

	token, errToken := jwt.CreateTokenJWT(getUserDetail.Email, int(getUserDetail.ID))
	if errToken != nil {
		return nil, errToken
	}

	res := authModel.ResponseLogin{
		Name:           getUserDetail.Name,
		Email:          getUserDetail.Email,
		Gender:         getUserDetail.Gender,
		DateOfBirthday: getUserDetail.DateOfBirthday,
		Token:          token,
	}

	return &res, nil
}

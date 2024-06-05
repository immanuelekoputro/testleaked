package authRepository

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	authModel "tinderleaked/models/auth"
	"tinderleaked/modules/auth"
)

type sqlRepository struct {
	Conn *gorm.DB
}

func NewAuthRepository(Conn *gorm.DB) auth.RepositoryAuth {
	return &sqlRepository{Conn: Conn}
}

func (repo *sqlRepository) SaveRegisterUsers(request *authModel.Users) error {
	err := repo.Conn.Create(&request).Error
	if err != nil {
		log.Error().Msg("[Error] : " + err.Error())
		return nil
	}

	return nil
}

func (repo *sqlRepository) FetchUserByAlias(email string) (*authModel.Users, error) {
	var user authModel.Users
	err := repo.Conn.First(&user, "email = ?", email).Error
	if err != nil {
		log.Error().Msg("[Error] : " + err.Error())
		return nil, err
	}

	return &user, nil
}

func (repo *sqlRepository) FetchUserByID(userId int) (*authModel.ResponseUser, error) {
	var user authModel.Users
	err := repo.Conn.First(&user, "id = ?", userId).Error
	if err != nil {
		log.Error().Msg("[Error] : " + err.Error())
		return nil, err
	}

	res := &authModel.ResponseUser{
		Name:           user.Name,
		Email:          user.Email,
		Gender:         user.Gender,
		DateOfBirthday: user.DateOfBirthday,
	}

	return res, nil
}
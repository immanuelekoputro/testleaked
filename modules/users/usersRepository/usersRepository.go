package usersRepository

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	authModel "tinderleaked/models/auth"
	usersModel "tinderleaked/models/users"
	"tinderleaked/modules/users"
)

type sqlRepository struct {
	Conn *gorm.DB
}

func NewUsersRepository(Conn *gorm.DB) users.RepositoryUsers {
	return &sqlRepository{Conn: Conn}
}

func (repo *sqlRepository) UserPackage(userId int) error {
	return nil
}

func (repo *sqlRepository) CountViewed(userId int) (int64, error) {
	var count int64
	err := repo.Conn.Model(&usersModel.UserViewHistories{}).Where("visitor_user_id = ? and DATE(created_at) = DATE(NOW())", userId).Count(&count).Error
	if err != nil {
		log.Error().Msg("[Error] : " + err.Error())
		return 0, err
	}

	return count, nil
}

func (repo *sqlRepository) FetchUserHasNotViewed(userId, limit int) (*[]authModel.ResponseUser, error) {
	var user []authModel.ResponseUser

	//query := fmt.Sprintf("select users.id as id, users.name as name, users.email as email, users.gender as gender, DATE_FORMAT(users.date_of_birthday,%s) as date_of_birthday from users left join user_view_histories uvh on users.id = uvh.visitor_user_id where users.id not in (select host_user_id as id from user_view_histories where user_view_histories.visitor_user_id = %d and DATE(created_at) = DATE(NOW())) and users.id != %d limit %d", "'%Y-%m-%d %H:%i:%s'", userId, userId, limit)
	query := fmt.Sprintf("select users.id as id, users.name as name, users.email as email, users.gender as gender, DATE_FORMAT(users.date_of_birthday,%s) as date_of_birthday, case when (select vus.package_id from vw_user_subscribes vus where vus.subscribe_status = 1 and vus.user_id = users.id and vus.package_id = 2) then true else false end as is_verified_user from users left join user_view_histories uvh on users.id = uvh.visitor_user_id where users.id not in (select host_user_id as id from user_view_histories where user_view_histories.visitor_user_id = %d and DATE(created_at) = DATE(NOW())) and users.id != %d limit %d", "'%Y-%m-%d %H:%i:%s'", userId, userId, limit)
	err := repo.Conn.Raw(query).Scan(&user).Error
	if err != nil {
		log.Error().Msg("[Error] : " + err.Error())
		return nil, err
	}

	return &user, nil
}

func (repo *sqlRepository) UserActionToView(viewerId, hostId int, visitorAction string, superlike bool) error {
	req := usersModel.UserViewHistories{
		VisitorUserId: uint(viewerId),
		HostUserId:    uint(hostId),
		VisitorAction: visitorAction,
		IsSuperlike:   false,
	}

	err := repo.Conn.Create(&req).Error
	if err != nil {
		log.Error().Msg("[Error] : " + err.Error())
		return nil
	}

	return nil
}

package packagesRepository

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"time"
	packagesModel "tinderleaked/models/packages"
	usersModel "tinderleaked/models/users"
	"tinderleaked/modules/packages"
)

type sqlRepository struct {
	Conn *gorm.DB
}

func NewPackagesRepository(Conn *gorm.DB) packages.RepositoryPackages {
	return &sqlRepository{Conn: Conn}
}

func (repo *sqlRepository) FetchPackages() (*[]packagesModel.Packages, error) {
	var res []packagesModel.Packages

	// Get all matched records
	queries := repo.Conn.Where("status != ?", "false").Find(&res)

	if queries.Error != nil {
		return nil, queries.Error
	}

	return &res, nil
}

func (repo *sqlRepository) FetchActivePackageByUserID(userId int) (*usersModel.VwUserSubscribes, error) {
	var user usersModel.VwUserSubscribes
	err := repo.Conn.First(&user, "user_id = ? and subscribe_status = 1", userId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		log.Error().Msg("[Error] : " + err.Error())
		return nil, err
	}

	return &user, nil
}

func (repo *sqlRepository) SaveUserPackages(userId, packageId int) error {
	currentTime := time.Now()
	startTime := currentTime.Format("2006-01-02 15:04:05")

	futureTime := currentTime.AddDate(0, 0, 30)
	endTime := futureTime.Format("2006-01-02 15:04:05")

	req := usersModel.UserSubscribes{
		UserId:    uint(userId),
		PackageId: uint(packageId),
		StartDate: startTime,
		EndDate:   endTime,
	}

	err := repo.Conn.Create(&req).Error
	if err != nil {
		log.Error().Msg("[Error] : " + err.Error())
		return nil
	}

	return nil
}

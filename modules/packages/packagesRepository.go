package packages

import (
	packagesModel "tinderleaked/models/packages"
	usersModel "tinderleaked/models/users"
)

type RepositoryPackages interface {
	FetchPackages() (*[]packagesModel.Packages, error)
	FetchActivePackageByUserID(userId int) (*usersModel.VwUserSubscribes, error)
	SaveUserPackages(userId, packageId int) error
}

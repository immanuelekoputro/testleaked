package packages

import (
	packagesModel "tinderleaked/models/packages"
	usersModel "tinderleaked/models/users"
)

type UsecasePackages interface {
	GetPackages() (*[]packagesModel.Packages, error)
	AssignPackageByLoggedUser(userID, packageID int) error
	GetLoggedUserPackage(userID int) (*usersModel.VwUserSubscribes, error)
}

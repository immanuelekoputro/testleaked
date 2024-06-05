package packagesUsecase

import (
	"errors"
	packagesModel "tinderleaked/models/packages"
	usersModel "tinderleaked/models/users"
	"tinderleaked/modules/packages"
)

type packagesUsecase struct {
	packagesRepository packages.RepositoryPackages
}

func NewPackagesUsecase(packagesRepository packages.RepositoryPackages) packages.UsecasePackages {
	return &packagesUsecase{
		packagesRepository: packagesRepository,
	}
}

func (usecase *packagesUsecase) GetPackages() (*[]packagesModel.Packages, error) {
	allPackage, errAllPackage := usecase.packagesRepository.FetchPackages()
	if errAllPackage != nil {
		return nil, errAllPackage
	}

	return allPackage, nil
}

func (usecase *packagesUsecase) GetLoggedUserPackage(userID int) (*usersModel.VwUserSubscribes, error) {
	hasActivePackage, errHasActivePackage := usecase.packagesRepository.FetchActivePackageByUserID(userID)
	if errHasActivePackage != nil {
		return nil, errHasActivePackage
	}

	return hasActivePackage, nil

}

func (usecase *packagesUsecase) AssignPackageByLoggedUser(userID, packageID int) error {
	hasActivePackage, errHasActivePackage := usecase.packagesRepository.FetchActivePackageByUserID(userID)
	if errHasActivePackage != nil {
		return errHasActivePackage
	}

	if hasActivePackage != nil {
		return errors.New("user have active package")
	}

	errSetPackage := usecase.packagesRepository.SaveUserPackages(userID, packageID)
	if errSetPackage != nil {
		return errHasActivePackage
	}

	return nil
}

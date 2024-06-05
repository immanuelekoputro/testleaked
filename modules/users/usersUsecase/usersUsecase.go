package usersUsecase

import (
	authModel "tinderleaked/models/auth"
	usersModel "tinderleaked/models/users"
	"tinderleaked/modules/packages"
	"tinderleaked/modules/users"
)

type usersUsecase struct {
	usersRepository    users.RepositoryUsers
	packagesRepository packages.RepositoryPackages
}

func NewUserUsecase(usersRepository users.RepositoryUsers, packagesRepository packages.RepositoryPackages) users.UsecaseUsers {
	return &usersUsecase{
		usersRepository:    usersRepository,
		packagesRepository: packagesRepository,
	}
}

var limit int

func (usecase *usersUsecase) ListAnotherUsers(userId int) (*[]authModel.ResponseUser, error) {
	// get user package
	userPackage, errUserPackage := usecase.packagesRepository.FetchActivePackageByUserID(userId)
	if errUserPackage != nil {
		return nil, errUserPackage
	}

	isUnlimitedSwipe := false
	if userPackage != nil {
		if userPackage.PackageID == 1 {
			isUnlimitedSwipe = true
		}
	}

	limit = 10

	if !isUnlimitedSwipe {
		getRestLimit, errGetRestLimit := usecase.usersRepository.CountViewed(userId)
		if errGetRestLimit != nil {
			return nil, errGetRestLimit
		}

		limit = limit - int(getRestLimit)
	}

	getUserAvailableForView, errGetUserAvailableForView := usecase.usersRepository.FetchUserHasNotViewed(userId, limit)
	if errGetUserAvailableForView != nil {
		return nil, errGetUserAvailableForView
	}

	return getUserAvailableForView, nil
}

func (usecase *usersUsecase) SubmitActionUser(userId int, req usersModel.SubmitActionUser) error {
	errSubmitUserAction := usecase.usersRepository.UserActionToView(userId, int(req.HostID), req.VisitorAction, false)
	if errSubmitUserAction != nil {
		return errSubmitUserAction
	}

	return nil
}

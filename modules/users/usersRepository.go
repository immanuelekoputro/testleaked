package users

import authModel "tinderleaked/models/auth"

type RepositoryUsers interface {
	CountViewed(userId int) (int64, error)
	FetchUserHasNotViewed(userId, limit int) (*[]authModel.ResponseUser, error)
	UserActionToView(viewerId, hostId int, visitorAction string, superlike bool) error
}

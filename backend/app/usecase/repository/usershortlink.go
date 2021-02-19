package repository

import "github.com/cross-team/clublink/backend/app/entity"

// UserShortLink accesses User-ShortLink relationship from storage, such as database.
type UserShortLink interface {
	CreateRelation(user entity.User, shortLinkInput entity.ShortLinkInput) error
	GetUserByShortLink(shortLinkID string) (entity.User, error)
	FindAliasesByUser(user entity.User) ([]string, error)
	HasMapping(user entity.User, alias string) (bool, error)
}

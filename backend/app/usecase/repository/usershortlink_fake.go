package repository

import (
	"errors"
	"fmt"

	"github.com/short-d/short/backend/app/entity"
)

var _ UserShortLink = (*UserShortLinkFake)(nil)

// UserShortLinkFake represents in memory implementation of User-ShortLink relationship accessor.
type UserShortLinkFake struct {
	users      []entity.User
	shortLinks []entity.ShortLink
}

// CreateRelation creates many to many relationship between User and ShortLink.
func (u *UserShortLinkFake) CreateRelation(user entity.User, shortLinkInput entity.ShortLinkInput) error {
	if shortLinkInput.CustomAlias == nil {
		return errors.New("empty alias")
	}
	customAlias := shortLinkInput.GetCustomAlias("")
	isExist, err := u.HasMapping(user, customAlias)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("relationship exists")
	}
	u.users = append(u.users, user)
	u.shortLinks = append(u.shortLinks, entity.ShortLink{
		Alias:     customAlias,
		LongLink:  shortLinkInput.GetLongLink(""),
		ExpireAt:  shortLinkInput.ExpireAt,
		CreatedAt: shortLinkInput.CreatedAt,
	})
	return nil
}

// FindAliasesByUser fetches the aliases of all the ShortLinks created by the given user.
func (u UserShortLinkFake) FindAliasesByUser(user entity.User) ([]string, error) {
	var aliases []string
	for idx, currUser := range u.users {
		if currUser.ID != user.ID {
			continue
		}
		aliases = append(aliases, u.shortLinks[idx].Alias)
	}
	return aliases, nil
}

// HasMapping checks whether a given short link belongs to a user.
func (u UserShortLinkFake) HasMapping(user entity.User, alias string) (bool, error) {
	for idx, currUser := range u.users {
		if currUser.ID == user.ID && u.shortLinks[idx].Alias == alias {
			return true, nil
		}
	}
	return false, nil
}

// UpdateAliasCascade updates user-shortlink relationships to reflect changes to alias.
// TODO(issue#958) use eventbus for propagating short link change to all related repos
func (u *UserShortLinkFake) UpdateAliasCascade(oldAlias string, shortLinkInput entity.ShortLinkInput) error {
	for idx := range u.users {
		if u.shortLinks[idx].Alias == oldAlias {
			u.shortLinks[idx] = entity.ShortLink{
				Alias:     shortLinkInput.GetCustomAlias(""),
				LongLink:  shortLinkInput.GetLongLink(""),
				ExpireAt:  shortLinkInput.ExpireAt,
				CreatedAt: shortLinkInput.CreatedAt,
			}
			return nil
		}
	}
	return fmt.Errorf("no relationships with alias '%s' exist", oldAlias)
}

// NewUserShortLinkRepoFake creates UserShortLinkFake
func NewUserShortLinkRepoFake(users []entity.User, shortLinks []entity.ShortLink) UserShortLinkFake {
	return UserShortLinkFake{
		users:      users,
		shortLinks: shortLinks,
	}
}

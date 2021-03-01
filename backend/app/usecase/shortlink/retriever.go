package shortlink

import (
	"time"

	"github.com/cross-team/clublink/backend/app/entity"
	"github.com/cross-team/clublink/backend/app/usecase/repository"
)

var _ Retriever = (*RetrieverPersist)(nil)

// Retriever represents ShortLink retriever
type Retriever interface {
	GetShortLink(ID string) (entity.ShortLink, error)
	GetActiveShortLink(alias string, expiringAt *time.Time) (entity.ShortLink, error)
	GetShortLinksByUser(user entity.User) ([]entity.ShortLink, error)
}

// RetrieverPersist represents ShortLink retriever that fetches ShortLink from persistent
// storage, such as database
type RetrieverPersist struct {
	shortLinkRepo     repository.ShortLink
	userShortLinkRepo repository.UserShortLink
}

// GetShortLink retrieves ShortLink from persistent storage given alias
func (r RetrieverPersist) GetShortLink(ID string) (entity.ShortLink, error) {
	shortLink, err := r.shortLinkRepo.GetShortLinkByID(ID)
	if err != nil {
		return entity.ShortLink{}, err
	}

	return shortLink, nil
}

// GetShortLink retrieves ShortLink from persistent storage given alias
func (r RetrieverPersist) GetActiveShortLink(alias string, expiringAt *time.Time) (entity.ShortLink, error) {
	if expiringAt == nil {
		return r.getActiveShortLink(alias, time.Now())
	}
	return r.getActiveShortLink(alias, *expiringAt)
}

func (r RetrieverPersist) getActiveShortLink(alias string, expiringAt time.Time) (entity.ShortLink, error) {
	shortLink, err := r.shortLinkRepo.GetShortLinkByAlias(alias, expiringAt)
	if err != nil {
		return entity.ShortLink{}, err
	}

	return shortLink, nil
}

// GetShortLinksByUser retrieves ShortLinks created by given user from persistent storage
func (r RetrieverPersist) GetShortLinksByUser(user entity.User) ([]entity.ShortLink, error) {
	aliases, err := r.userShortLinkRepo.FindAliasesByUser(user)
	if err != nil {
		return []entity.ShortLink{}, err
	}

	return r.shortLinkRepo.GetShortLinksByAliases(aliases)
}

// NewRetrieverPersist creates persistent ShortLink retriever
func NewRetrieverPersist(shortLinkRepo repository.ShortLink, userShortLinkRepo repository.UserShortLink) RetrieverPersist {
	return RetrieverPersist{
		shortLinkRepo:     shortLinkRepo,
		userShortLinkRepo: userShortLinkRepo,
	}
}

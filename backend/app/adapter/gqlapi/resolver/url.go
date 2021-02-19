package resolver

import (
	"github.com/cross-team/clublink/backend/app/adapter/gqlapi/scalar"
	"github.com/cross-team/clublink/backend/app/entity"
)

// ShortLink retrieves requested fields of ShortLink entity.
type ShortLink struct {
	shortLink entity.ShortLink
}

// Alias retrieves the alias of ShortLink entity.
func (s ShortLink) Alias() *string {
	return &s.shortLink.Alias
}

// LongLink retrieves the long link of ShortLink entity.
func (s ShortLink) LongLink() *string {
	return &s.shortLink.LongLink
}

// Room retrieves the room description of ShortLink entity.
func (s ShortLink) Room() *string {
	return &s.shortLink.Room
}

// ID retrieves the ID of ShortLink entity.
func (s ShortLink) ID() *string {
	return &s.shortLink.ID
}

// ExpireAt retrieves the expiration time of ShortLink entity.
func (s ShortLink) ExpireAt() *scalar.Time {
	if s.shortLink.ExpireAt == nil {
		return nil
	}

	return &scalar.Time{Time: *s.shortLink.ExpireAt}
}

func newShortLink(shortLink entity.ShortLink) ShortLink {
	return ShortLink{shortLink: shortLink}
}

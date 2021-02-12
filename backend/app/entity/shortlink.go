package entity

import (
	"time"

	"github.com/cross-team/clublink/backend/app/entity/metatag"
)

// ShortLink represents a short link.
type ShortLink struct {
	Alias         string
	LongLink      string
	Room      		string
	ID						string
	ExpireAt      *time.Time
	CreatedBy     *User
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	OpenGraphTags metatag.OpenGraph
	TwitterTags   metatag.Twitter
}

// ShortLinkInput represents possible ShortLink attributes for a short link.
type ShortLinkInput struct {
	LongLink    *string
	CustomAlias *string
	Username    *string
	Room 				*string
	ID					*string
	ExpireAt    *time.Time
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

// GetLongLink fetches LongLink for ShortLinkInput with default value.
func (s *ShortLinkInput) GetLongLink(defaultVal string) string {
	if s.LongLink == nil {
		return defaultVal
	}
	return *s.LongLink
}

// GetCustomAlias fetches CustomAlias for ShortLinkInput with default value.
func (s *ShortLinkInput) GetCustomAlias(defaultVal string) string {
	if s.CustomAlias == nil {
		return defaultVal
	}
	return *s.CustomAlias
}

// GetUsername fetches Username for ShortLinkInput with default value.
func (s *ShortLinkInput) GetUsername(defaultVal string) string {
	if s.Username == nil {
		return defaultVal
	}
	return *s.Username
}

// GetUsername fetches ID for ShortLinkInput with default value.
func (s *ShortLinkInput) GetID(defaultVal string) string {
	if s.ID == nil {
		return defaultVal
	}
	return *s.ID
}

// GetUsername fetches ID for ShortLinkInput with default value.
func (s *ShortLinkInput) GetRoom(defaultVal string) string {
	if s.Room == nil {
		return defaultVal
	}
	return *s.Room
}
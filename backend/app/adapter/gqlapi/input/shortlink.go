package input

import (
	"time"

	"github.com/cross-team/clublink/backend/app/entity"
)

// ShortLinkInput represents possible ShortLink attributes
type ShortLinkInput struct {
	LongLink    *string
	CustomAlias *string
	Username    *string
	Room        *string
	ExpireAt    *time.Time
}

// CreateShortLinkInput converts GraphQL ShortLinkInput into consumable entity for use cases.
func (s ShortLinkInput) CreateShortLinkInput() entity.ShortLinkInput {
	return entity.ShortLinkInput{
		LongLink:    s.LongLink,
		CustomAlias: s.CustomAlias,
		Username:    s.Username,
		Room: 			 s.Room,
		ExpireAt:    s.ExpireAt,
	}
}

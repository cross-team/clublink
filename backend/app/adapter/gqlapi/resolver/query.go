package resolver

import (
	"github.com/short-d/app/fw/logger"
	"github.com/cross-team/clublink/backend/app/usecase/authenticator"
	"github.com/cross-team/clublink/backend/app/usecase/changelog"
	"github.com/cross-team/clublink/backend/app/usecase/repository"
	"github.com/cross-team/clublink/backend/app/usecase/shortlink"
)

// Query represents GraphQL query resolver
type Query struct {
	logger             logger.Logger
	authenticator      authenticator.Authenticator
	changeLog          changelog.ChangeLog
	shortLinkRetriever shortlink.Retriever
	userShortLinkRepo  repository.UserShortLink
}

// AuthQueryArgs represents possible parameters for AuthQuery endpoint
type AuthQueryArgs struct {
	AuthToken *string
}

// AuthQuery extracts user information from authentication token
func (q Query) AuthQuery(args *AuthQueryArgs) (*AuthQuery, error) {
	authQuery := newAuthQuery(args.AuthToken, q.authenticator, q.changeLog, q.shortLinkRetriever, q.userShortLinkRepo)
	return &authQuery, nil
}

func newQuery(
	logger logger.Logger,
	authenticator authenticator.Authenticator,
	changeLog changelog.ChangeLog,
	shortLinkRetriever shortlink.Retriever,
	userShortLinkRepo repository.UserShortLink,
) Query {
	return Query{
		logger:             logger,
		authenticator:      authenticator,
		changeLog:          changeLog,
		shortLinkRetriever: shortLinkRetriever,
		userShortLinkRepo:  userShortLinkRepo,
	}
}

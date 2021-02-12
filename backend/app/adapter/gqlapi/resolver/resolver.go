package resolver

import (
	"github.com/short-d/app/fw/logger"
	"github.com/cross-team/clublink/backend/app/usecase/authenticator"
	"github.com/cross-team/clublink/backend/app/usecase/changelog"
	"github.com/cross-team/clublink/backend/app/usecase/requester"
	"github.com/cross-team/clublink/backend/app/usecase/shortlink"
	"github.com/cross-team/clublink/backend/app/usecase/repository"
	"github.com/cross-team/clublink/backend/app/usecase/keygen"
)

// Resolver contains GraphQL request handlers.
type Resolver struct {
	Query
	Mutation
}

// NewResolver creates a new GraphQL resolver.
func NewResolver(
	logger logger.Logger,
	shortLinkRetriever shortlink.Retriever,
	shortLinkCreator shortlink.Creator,
	shortLinkUpdater shortlink.Updater,
	changeLog changelog.ChangeLog,
	requesterVerifier requester.Verifier,
	authenticator authenticator.Authenticator,
	userRepo repository.User,
	keyGen keygen.KeyGenerator,
) Resolver {
	return Resolver{
		Query: newQuery(logger, authenticator, changeLog, shortLinkRetriever),
		Mutation: newMutation(
			logger,
			changeLog,
			shortLinkCreator,
			shortLinkUpdater,
			requesterVerifier,
			authenticator,
			userRepo,
			keyGen,
		),
	}
}

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

// Mutation represents GraphQL mutation resolver
type Mutation struct {
	logger            logger.Logger
	shortLinkCreator  shortlink.Creator
	shortLinkUpdater  shortlink.Updater
	requesterVerifier requester.Verifier
	authenticator     authenticator.Authenticator
	changeLog         changelog.ChangeLog
	userRepo 				  repository.User
	keyGen   				  keygen.KeyGenerator
}

// AuthMutationArgs represents possible parameters for AuthMutation endpoint
type AuthMutationArgs struct {
	AuthToken       *string
	CaptchaResponse string
}

type NoAuthMutationArgs struct {
	CaptchaResponse string
}

func (m Mutation) NoAuthMutation(args *NoAuthMutationArgs) (*NoAuthMutation, error) {
	isHuman, err := m.requesterVerifier.IsHuman(args.CaptchaResponse)

	if err != nil {
		return nil, ErrUnknown{}
	}

	if !isHuman {
		return nil, ErrNotHuman{}
	}

	noAuthMutation := newNoAuthMutation(
		m.changeLog,
		m.shortLinkCreator,
		m.shortLinkUpdater,
		m.userRepo,
		m.keyGen,
	)
	return &noAuthMutation, nil
}

// AuthMutation extracts user information from authentication token
func (m Mutation) AuthMutation(args *AuthMutationArgs) (*AuthMutation, error) {
	isHuman, err := m.requesterVerifier.IsHuman(args.CaptchaResponse)

	if err != nil {
		return nil, ErrUnknown{}
	}

	if !isHuman {
		return nil, ErrNotHuman{}
	}

	authMutation := newAuthMutation(
		args.AuthToken,
		m.authenticator,
		m.changeLog,
		m.shortLinkCreator,
		m.shortLinkUpdater,
	)
	return &authMutation, nil
}

func newMutation(
	logger logger.Logger,
	changeLog changelog.ChangeLog,
	shortLinkCreator shortlink.Creator,
	shortLinkUpdater shortlink.Updater,
	requesterVerifier requester.Verifier,
	authenticator authenticator.Authenticator,
	userRepo repository.User,
	keyGen keygen.KeyGenerator,
) Mutation {
	return Mutation{
		logger:            logger,
		changeLog:         changeLog,
		shortLinkCreator:  shortLinkCreator,
		shortLinkUpdater:  shortLinkUpdater,
		requesterVerifier: requesterVerifier,
		authenticator:     authenticator,
		userRepo:					 userRepo,
		keyGen:				 		 keyGen,
	}
}

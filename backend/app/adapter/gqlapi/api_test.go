// +build !integration all

package gqlapi

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/graphql"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/adapter/gqlapi/resolver"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/fw/filesystem"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/authorizer"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/role"
	"github.com/short-d/short/backend/app/usecase/changelog"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/requester"
	"github.com/short-d/short/backend/app/usecase/risk"
	"github.com/short-d/short/backend/app/usecase/shortlink"
	"github.com/short-d/short/backend/app/usecase/validator"
)

func TestGraphQlAPI(t *testing.T) {
	t.Parallel()
	now := time.Now()
	blockedURLs := map[string]bool{}
	blacklist := risk.NewBlackListFake(blockedURLs)

	shortLinkRepo := repository.NewShortLinkFake(nil, map[string]entity.ShortLink{})
	userShortLinkRepo := repository.NewUserShortLinkRepoFake([]entity.User{}, []entity.ShortLink{})
	retriever := shortlink.NewRetrieverPersist(&shortLinkRepo, &userShortLinkRepo)
	keyFetcher := keygen.NewKeyFetcherFake([]keygen.Key{})
	keyGen, err := keygen.NewKeyGenerator(2, &keyFetcher)
	assert.Equal(t, nil, err)

	longLinkValidator := validator.NewLongLink()
	customAliasValidator := validator.NewCustomAlias()
	tm := timer.NewStub(now)
	riskDetector := risk.NewDetector(blacklist)

	creator := shortlink.NewCreatorPersist(
		&shortLinkRepo,
		&userShortLinkRepo,
		keyGen,
		longLinkValidator,
		customAliasValidator,
		tm,
		riskDetector,
	)

	updater := shortlink.NewUpdaterPersist(
		&shortLinkRepo,
		&userShortLinkRepo,
		longLinkValidator,
		customAliasValidator,
		tm,
		riskDetector,
	)

	s := requester.NewReCaptchaFake(requester.VerifyResponse{})
	verifier := requester.NewReCaptchaVerifier(s)
	auth := authenticator.NewAuthenticatorFake(time.Now(), time.Hour)

	entryRepo := logger.NewEntryRepoFake()
	lg, err := logger.NewFake(logger.LogOff, &entryRepo)
	assert.Equal(t, nil, err)

	changeLogRepo := repository.NewChangeLogFake([]entity.Change{})
	userChangeLogRepo := repository.NewUserChangeLogFake(map[string]time.Time{})
	fakeRolesRepo := repository.NewUserRoleFake(map[string][]role.Role{})
	rb := rbac.NewRBAC(fakeRolesRepo)
	au := authorizer.NewAuthorizer(rb)
	changeLog := changelog.NewPersist(keyGen, tm, &changeLogRepo, &userChangeLogRepo, au)
	r := resolver.NewResolver(lg, retriever, creator, updater, changeLog, verifier, auth)

	schema := "schema.graphql"
	fileSystem := filesystem.NewLocal()
	graphqlAPI, err := NewShort(schema, fileSystem, r)
	assert.Equal(t, nil, err)
	assert.Equal(t, true, graphql.IsGraphQlAPIValid(graphqlAPI))
}

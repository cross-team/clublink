// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package dep

import (
	"database/sql"
	"github.com/cross-team/clublink/backend/app/adapter/facebook"
	"github.com/cross-team/clublink/backend/app/adapter/github"
	"github.com/cross-team/clublink/backend/app/adapter/google"
	"github.com/cross-team/clublink/backend/app/adapter/gqlapi/resolver"
	"github.com/cross-team/clublink/backend/app/adapter/grpcapi"
	"github.com/cross-team/clublink/backend/app/adapter/kgs"
	"github.com/cross-team/clublink/backend/app/adapter/request"
	"github.com/cross-team/clublink/backend/app/adapter/sqldb"
	"github.com/cross-team/clublink/backend/app/fw/filesystem"
	"github.com/cross-team/clublink/backend/app/usecase/authorizer"
	"github.com/cross-team/clublink/backend/app/usecase/authorizer/rbac"
	"github.com/cross-team/clublink/backend/app/usecase/changelog"
	"github.com/cross-team/clublink/backend/app/usecase/keygen"
	"github.com/cross-team/clublink/backend/app/usecase/repository"
	"github.com/cross-team/clublink/backend/app/usecase/risk"
	"github.com/cross-team/clublink/backend/app/usecase/shortlink"
	"github.com/cross-team/clublink/backend/app/usecase/sso"
	"github.com/cross-team/clublink/backend/app/usecase/validator"
	"github.com/cross-team/clublink/backend/dep/provider"
	"github.com/cross-team/clublink/backend/tool"
	"github.com/google/wire"
	"github.com/short-d/app/fw/analytics"
	"github.com/short-d/app/fw/cli"
	"github.com/short-d/app/fw/db"
	"github.com/short-d/app/fw/env"
	"github.com/short-d/app/fw/graphql"
	"github.com/short-d/app/fw/io"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/metrics"
	"github.com/short-d/app/fw/network"
	"github.com/short-d/app/fw/runtime"
	"github.com/short-d/app/fw/security"
	"github.com/short-d/app/fw/service"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/app/fw/webreq"
)

// Injectors from wire.go:

func InjectCommandFactory() cli.CommandFactory {
	cobraFactory := cli.NewCobraFactory()
	return cobraFactory
}

func InjectDBConnector() db.Connector {
	postgresConnector := db.NewPostgresConnector()
	return postgresConnector
}

func InjectDBMigrationTool() db.MigrationTool {
	postgresMigrationTool := db.NewPostgresMigrationTool()
	return postgresMigrationTool
}

func InjectEnv() env.Env {
	goDotEnv := env.NewGoDotEnv()
	return goDotEnv
}

func InjectGRPCService(runtime2 env.Runtime, prefix provider.LogPrefix, logLevel logger.LogLevel, sqlDB *sql.DB, securityPolicy security.Policy, dataDogAPIKey provider.DataDogAPIKey) (service.GRPC, error) {
	system := timer.NewSystem()
	program := runtime.NewProgram()
	deployment := env.NewDeployment(runtime2)
	stdOut := io.NewStdOut()
	client := webreq.NewHTTPClient()
	http := webreq.NewHTTP(client)
	entryRepository := provider.NewEntryRepositorySwitch(runtime2, deployment, stdOut, dataDogAPIKey, http)
	loggerLogger := provider.NewLogger(prefix, logLevel, system, program, entryRepository)
	shortLinkSQL := sqldb.NewShortLinkSQL(sqlDB)
	metaTagPersist := shortlink.NewMetaTagPersist(shortLinkSQL)
	metaTagServiceServer := grpcapi.NewMetaTagServer(metaTagPersist)
	short := grpcapi.NewShort(metaTagServiceServer)
	grpc, err := service.NewGRPC(loggerLogger, short, securityPolicy)
	if err != nil {
		return service.GRPC{}, err
	}
	return grpc, nil
}

func InjectGraphQLService(runtime2 env.Runtime, prefix provider.LogPrefix, logLevel logger.LogLevel, sqlDB *sql.DB, graphqlSchemaPath provider.GraphQLSchemaPath, graphqlPath provider.GraphQLPath, graphiQLDefaultQuery provider.GraphiQLDefaultQuery, secret provider.ReCaptchaSecret, jwtSecret provider.JwtSecret, bufferSize provider.KeyGenBufferSize, kgsRPCConfig provider.KgsRPCConfig, tokenValidDuration provider.TokenValidDuration, dataDogAPIKey provider.DataDogAPIKey, segmentAPIKey provider.SegmentAPIKey, ipStackAPIKey provider.IPStackAPIKey, googleAPIKey provider.GoogleAPIKey) (service.GraphQL, error) {
	local := filesystem.NewLocal()
	system := timer.NewSystem()
	program := runtime.NewProgram()
	deployment := env.NewDeployment(runtime2)
	stdOut := io.NewStdOut()
	client := webreq.NewHTTPClient()
	http := webreq.NewHTTP(client)
	entryRepository := provider.NewEntryRepositorySwitch(runtime2, deployment, stdOut, dataDogAPIKey, http)
	loggerLogger := provider.NewLogger(prefix, logLevel, system, program, entryRepository)
	shortLinkSQL := sqldb.NewShortLinkSQL(sqlDB)
	userShortLinkSQL := sqldb.NewUserShortLinkSQL(sqlDB)
	retrieverPersist := shortlink.NewRetrieverPersist(shortLinkSQL, userShortLinkSQL)
	rpc, err := provider.NewKgsRPC(kgsRPCConfig)
	if err != nil {
		return service.GraphQL{}, err
	}
	keyGenerator, err := provider.NewKeyGenerator(bufferSize, rpc)
	if err != nil {
		return service.GraphQL{}, err
	}
	longLink := validator.NewLongLink()
	customAlias := validator.NewCustomAlias()
	safeBrowsing := provider.NewSafeBrowsing(googleAPIKey, http)
	detector := risk.NewDetector(safeBrowsing)
	creatorPersist := shortlink.NewCreatorPersist(shortLinkSQL, userShortLinkSQL, keyGenerator, longLink, customAlias, system, detector)
	updaterPersist := shortlink.NewUpdaterPersist(shortLinkSQL, userShortLinkSQL, longLink, customAlias, system, detector)
	changeLogSQL := sqldb.NewChangeLogSQL(sqlDB)
	userChangeLogSQL := sqldb.NewUserChangeLogSQL(sqlDB)
	userRoleSQL := sqldb.NewUserRoleSQL(sqlDB)
	rbacRBAC := rbac.NewRBAC(userRoleSQL)
	authorizerAuthorizer := authorizer.NewAuthorizer(rbacRBAC)
	persist := changelog.NewPersist(keyGenerator, system, changeLogSQL, userChangeLogSQL, authorizerAuthorizer)
	reCaptcha := provider.NewReCaptchaService(http, secret)
	verifier := provider.NewVerifier(deployment, reCaptcha)
	tokenizer := provider.NewJwtGo(jwtSecret)
	authenticator := provider.NewAuthenticator(tokenizer, system, tokenValidDuration)
	userSQL := sqldb.NewUserSQL(sqlDB)
	resolverResolver := resolver.NewResolver(loggerLogger, retrieverPersist, creatorPersist, updaterPersist, persist, verifier, authenticator, userSQL, userShortLinkSQL, keyGenerator)
	api, err := provider.NewShortGraphQLAPI(graphqlSchemaPath, local, resolverResolver)
	if err != nil {
		return service.GraphQL{}, err
	}
	graphGopherHandler := graphql.NewGraphGopherHandler(api)
	graphiQL := provider.NewGraphiQL(graphqlPath, graphiQLDefaultQuery)
	graphQL := provider.NewGraphQLService(graphqlPath, graphGopherHandler, graphiQL, loggerLogger)
	return graphQL, nil
}

func InjectRoutingService(runtime2 env.Runtime, prefix provider.LogPrefix, logLevel logger.LogLevel, sqlDB *sql.DB, githubClientID provider.GithubClientID, githubClientSecret provider.GithubClientSecret, facebookClientID provider.FacebookClientID, facebookClientSecret provider.FacebookClientSecret, facebookRedirectURI provider.FacebookRedirectURI, googleClientID provider.GoogleClientID, googleClientSecret provider.GoogleClientSecret, googleRedirectURI provider.GoogleRedirectURI, jwtSecret provider.JwtSecret, bufferSize provider.KeyGenBufferSize, kgsRPCConfig provider.KgsRPCConfig, webFrontendURL provider.WebFrontendURL, tokenValidDuration provider.TokenValidDuration, searchTimeout provider.SearchTimeout, swaggerUIDir provider.SwaggerUIDir, openAPISpecPath provider.OpenAPISpecPath, dataDogAPIKey provider.DataDogAPIKey, segmentAPIKey provider.SegmentAPIKey, ipStackAPIKey provider.IPStackAPIKey) (service.Routing, error) {
	system := timer.NewSystem()
	program := runtime.NewProgram()
	deployment := env.NewDeployment(runtime2)
	stdOut := io.NewStdOut()
	client := webreq.NewHTTPClient()
	http := webreq.NewHTTP(client)
	entryRepository := provider.NewEntryRepositorySwitch(runtime2, deployment, stdOut, dataDogAPIKey, http)
	loggerLogger := provider.NewLogger(prefix, logLevel, system, program, entryRepository)
	dataDog := provider.NewDataDogMetrics(dataDogAPIKey, http, system, runtime2)
	segment := provider.NewSegment(segmentAPIKey, system, loggerLogger)
	rpc, err := provider.NewKgsRPC(kgsRPCConfig)
	if err != nil {
		return service.Routing{}, err
	}
	keyGenerator, err := provider.NewKeyGenerator(bufferSize, rpc)
	if err != nil {
		return service.Routing{}, err
	}
	proxy := network.NewProxy()
	ipStack := provider.NewIPStack(ipStackAPIKey, http, loggerLogger)
	requestClient := request.NewClient(proxy, ipStack)
	instrumentationFactory := request.NewInstrumentationFactory(loggerLogger, system, dataDog, segment, keyGenerator, requestClient)
	shortLinkSQL := sqldb.NewShortLinkSQL(sqlDB)
	userShortLinkSQL := sqldb.NewUserShortLinkSQL(sqlDB)
	retrieverPersist := shortlink.NewRetrieverPersist(shortLinkSQL, userShortLinkSQL)
	featureToggleSQL := sqldb.NewFeatureToggleSQL(sqlDB)
	userRoleSQL := sqldb.NewUserRoleSQL(sqlDB)
	rbacRBAC := rbac.NewRBAC(userRoleSQL)
	authorizerAuthorizer := authorizer.NewAuthorizer(rbacRBAC)
	decisionMakerFactory := provider.NewFeatureDecisionMakerFactorySwitch(deployment, featureToggleSQL, authorizerAuthorizer)
	tokenizer := provider.NewJwtGo(jwtSecret)
	authenticator := provider.NewAuthenticator(tokenizer, system, tokenValidDuration)
	factory := sso.NewFactory(authenticator)
	userSQL := sqldb.NewUserSQL(sqlDB)
	accountLinkerFactory := sso.NewAccountLinkerFactory(keyGenerator, userSQL)
	githubSSOSql := sqldb.NewGithubSSOSql(sqlDB, loggerLogger)
	accountLinker := provider.NewGithubAccountLinker(accountLinkerFactory, githubSSOSql)
	identityProvider := provider.NewGithubIdentityProvider(http, githubClientID, githubClientSecret)
	clientFactory := graphql.NewClientFactory(http)
	account := github.NewAccount(clientFactory)
	singleSignOn := provider.NewGithubSSO(factory, accountLinker, identityProvider, account)
	facebookIdentityProvider := provider.NewFacebookIdentityProvider(http, facebookClientID, facebookClientSecret, facebookRedirectURI)
	facebookAccount := facebook.NewAccount(http)
	facebookSSOSql := sqldb.NewFacebookSSOSql(sqlDB, loggerLogger)
	facebookAccountLinker := provider.NewFacebookAccountLinker(accountLinkerFactory, facebookSSOSql)
	facebookSingleSignOn := provider.NewFacebookSSO(factory, facebookIdentityProvider, facebookAccount, facebookAccountLinker)
	googleIdentityProvider := provider.NewGoogleIdentityProvider(http, googleClientID, googleClientSecret, googleRedirectURI)
	googleAccount := google.NewAccount(http)
	googleSSOSql := sqldb.NewGoogleSSOSql(sqlDB, loggerLogger)
	googleAccountLinker := provider.NewGoogleAccountLinker(accountLinkerFactory, googleSSOSql)
	googleSingleSignOn := provider.NewGoogleSSO(factory, googleIdentityProvider, googleAccount, googleAccountLinker)
	search := provider.NewSearch(loggerLogger, shortLinkSQL, userShortLinkSQL, searchTimeout)
	v := provider.NewShortRoutes(instrumentationFactory, webFrontendURL, system, retrieverPersist, decisionMakerFactory, singleSignOn, facebookSingleSignOn, googleSingleSignOn, authenticator, search, swaggerUIDir, openAPISpecPath)
	routing := service.NewRouting(loggerLogger, v)
	return routing, nil
}

func InjectDataTool(prefix provider.LogPrefix, logLevel logger.LogLevel, dbConfig db.Config, dbConnector db.Connector, bufferSize provider.KeyGenBufferSize, kgsRPCConfig provider.KgsRPCConfig) (tool.Data, error) {
	rpc, err := provider.NewKgsRPC(kgsRPCConfig)
	if err != nil {
		return tool.Data{}, err
	}
	keyGenerator, err := provider.NewKeyGenerator(bufferSize, rpc)
	if err != nil {
		return tool.Data{}, err
	}
	system := timer.NewSystem()
	program := runtime.NewProgram()
	stdOut := io.NewStdOut()
	local := provider.NewLocalEntryRepo(stdOut)
	loggerLogger := provider.NewLogger(prefix, logLevel, system, program, local)
	data, err := tool.NewData(dbConfig, dbConnector, keyGenerator, loggerLogger)
	if err != nil {
		return tool.Data{}, err
	}
	return data, nil
}

// wire.go:

var authenticatorSet = wire.NewSet(provider.NewJwtGo, provider.NewAuthenticator)

var authorizerSet = wire.NewSet(wire.Bind(new(repository.UserRole), new(sqldb.UserRoleSQL)), sqldb.NewUserRoleSQL, rbac.NewRBAC, authorizer.NewAuthorizer)

var observabilitySet = wire.NewSet(wire.Bind(new(io.Output), new(io.StdOut)), wire.Bind(new(runtime.Runtime), new(runtime.Program)), wire.Bind(new(metrics.Metrics), new(metrics.DataDog)), wire.Bind(new(analytics.Analytics), new(analytics.Segment)), wire.Bind(new(network.Network), new(network.Proxy)), io.NewStdOut, provider.NewEntryRepositorySwitch, provider.NewLogger, runtime.NewProgram, provider.NewDataDogMetrics, provider.NewSegment, network.NewProxy, request.NewClient, request.NewInstrumentationFactory)

var githubAPISet = wire.NewSet(provider.NewGithubIdentityProvider, github.NewAccount, github.NewAPI)

var facebookAPISet = wire.NewSet(provider.NewFacebookIdentityProvider, facebook.NewAccount, facebook.NewAPI)

var googleAPISet = wire.NewSet(provider.NewGoogleIdentityProvider, google.NewAccount, google.NewAPI)

var keyGenSet = wire.NewSet(wire.Bind(new(keygen.KeyFetcher), new(kgs.RPC)), provider.NewKgsRPC, provider.NewKeyGenerator)

var featureDecisionSet = wire.NewSet(wire.Bind(new(repository.FeatureToggle), new(sqldb.FeatureToggleSQL)), sqldb.NewFeatureToggleSQL, provider.NewFeatureDecisionMakerFactorySwitch)

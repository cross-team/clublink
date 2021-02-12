package provider

import (
	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/timer"
	"github.com/cross-team/clublink/backend/app/adapter/facebook"
	"github.com/cross-team/clublink/backend/app/adapter/github"
	"github.com/cross-team/clublink/backend/app/adapter/google"
	"github.com/cross-team/clublink/backend/app/adapter/request"
	"github.com/cross-team/clublink/backend/app/adapter/routing"
	"github.com/cross-team/clublink/backend/app/usecase/authenticator"
	"github.com/cross-team/clublink/backend/app/usecase/feature"
	"github.com/cross-team/clublink/backend/app/usecase/search"
	"github.com/cross-team/clublink/backend/app/usecase/shortlink"
)

// WebFrontendURL represents the URL of the web frontend
type WebFrontendURL string

// SwaggerUIDir represents the root directory of Swagger UI static assets.
type SwaggerUIDir string

// OpenAPISpecPath represents the location of OpenAPI specification.
type OpenAPISpecPath string

// NewShortRoutes creates HTTP routes for Short API with WwwRoot to uniquely identify WwwRoot during dependency injection.
func NewShortRoutes(
	instrumentationFactory request.InstrumentationFactory,
	webFrontendURL WebFrontendURL,
	timer timer.Timer,
	shortLinkRetriever shortlink.Retriever,
	featureDecisionMakerFactory feature.DecisionMakerFactory,
	githubSSO github.SingleSignOn,
	facebookSSO facebook.SingleSignOn,
	googleSSO google.SingleSignOn,
	authenticator authenticator.Authenticator,
	search search.Search,
	swaggerUIDir SwaggerUIDir,
	openAPISpecPath OpenAPISpecPath,
) []router.Route {
	return routing.NewShort(
		instrumentationFactory,
		string(webFrontendURL),
		timer,
		shortLinkRetriever,
		featureDecisionMakerFactory,
		githubSSO,
		facebookSSO,
		googleSSO,
		authenticator,
		search,
		string(swaggerUIDir),
		string(openAPISpecPath),
	)
}

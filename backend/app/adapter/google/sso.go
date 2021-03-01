package google

import "github.com/cross-team/clublink/backend/app/usecase/sso"

// AccountLinker links user's Google account with Short account.
type AccountLinker sso.AccountLinker

// SingleSignOn enables users to sign in through their Google account.
type SingleSignOn sso.SingleSignOn

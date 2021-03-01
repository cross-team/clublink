package github

import "github.com/cross-team/clublink/backend/app/usecase/sso"

// AccountLinker links user's Github account with Short account.
type AccountLinker sso.AccountLinker

// SingleSignOn enables users to sign in through their Github account.
type SingleSignOn sso.SingleSignOn

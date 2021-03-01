package facebook

import "github.com/cross-team/clublink/backend/app/usecase/sso"

// AccountLinker links user's Facebook account with Short account.
type AccountLinker sso.AccountLinker

// SingleSignOn enables users to sign in through their Facebook account.
type SingleSignOn sso.SingleSignOn

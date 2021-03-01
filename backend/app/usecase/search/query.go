package search

import "github.com/cross-team/clublink/backend/app/entity"

// Query represents a user query.
type Query struct {
	Query string
	User  *entity.User
}

package resolver

import (
	"github.com/cross-team/clublink/backend/app/entity"
)

// User retrieves requested fields of User entity.
type User struct {
	user entity.User
}

// ID retrieves the ID of User entity.
func (u User) ID() *string {
	return &u.user.ID
}

// ID retrieves the name of User entity.
func (u User) Name() *string {
	return &u.user.Name
}

// ID retrieves the email of User entity.
func (u User) Email() *string {
	return &u.user.Email
}

func newUser(user entity.User) User {
	return User{user: user}
}

package entity

import "time"

// User contains basic user information such as, user ID, name, and email.
type User struct {
	ID             string
	Name           string
	Email          string
	LastSignedInAt *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

// UserInput represents possible User attributes for a user.
type UserInput struct {
	ID             *string
	Name           *string
	Email          *string
	LastSignedInAt *time.Time
	CreatedAt      *time.Time
	UpdatedAt      *time.Time
}

// GetID fetches ID for UserInput with default value.
func (u *UserInput) GetID(defaultVal string) string {
	if u.ID == nil {
		return defaultVal
	}
	return *u.ID
}

// GetName fetches name for UserInput with default value.
func (u *UserInput) GetName(defaultVal string) string {
	if u.Name == nil {
		return defaultVal
	}
	return *u.Name
}

// GetEmail fetches email for UserInput with default value.
func (u *UserInput) GetEmail(defaultVal string) string {
	if u.Email == nil {
		return defaultVal
	}
	return *u.Email
}
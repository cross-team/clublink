package repository

import (
	"github.com/cross-team/clublink/backend/app/entity"
)

// APIKey accesses API keys for third party apps from persistent storage, such as database.
type APIKey interface {
	GetAPIKey(appID string, key string) (entity.APIKey, error)
	CreateAPIKey(input entity.APIKeyInput) (entity.APIKey, error)
}

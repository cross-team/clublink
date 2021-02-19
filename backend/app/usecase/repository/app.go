package repository

import "github.com/cross-team/clublink/backend/app/entity"

// App accesses third party app info from persistent storage, such as database.
type App interface {
	GetAppByID(id string) (entity.App, error)
}

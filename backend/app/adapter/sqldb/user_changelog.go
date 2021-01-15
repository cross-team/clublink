package sqldb

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.UserChangeLog = (*UserChangeLogSQL)(nil)

// UserChangeLogSQL accesses UserChangeLog information in user_changelog table through SQL.
type UserChangeLogSQL struct {
	db *sql.DB
}

// GetLastViewedAt retrieves LastViewedAt for a given user.
func (u UserChangeLogSQL) GetLastViewedAt(user entity.User) (time.Time, error) {
	statement := fmt.Sprintf(`
SELECT "%s" 
FROM "%s"
WHERE "%s"=$1;`,
		table.UserChangeLog.ColumnLastViewedAt,
		table.UserChangeLog.TableName,
		table.UserChangeLog.ColumnUserID,
	)

	row := u.db.QueryRow(statement, user.ID)
	lastViewedAt := time.Time{}
	err := row.Scan(&lastViewedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return lastViewedAt, repository.ErrEntryNotFound("user not found")
	}

	return lastViewedAt.UTC(), err
}

// UpdateLastViewedAt updates LastViewedAt for the given user.
func (u UserChangeLogSQL) UpdateLastViewedAt(user entity.User, currentTime time.Time) (time.Time, error) {
	statement := fmt.Sprintf(`
UPDATE "%s"
SET %s=$1
WHERE %s=$2;
`,
		table.UserChangeLog.TableName,
		table.UserChangeLog.ColumnLastViewedAt,
		table.UserChangeLog.ColumnUserID,
	)

	result, err := u.db.Exec(
		statement,
		currentTime,
		user.ID,
	)
	if err != nil {
		return time.Time{}, err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return time.Time{}, err
	}

	if rowsUpdated == 0 {
		return time.Time{}, repository.ErrEntryNotFound("user not found")
	}

	return currentTime, nil
}

// CreateRelation inserts a new entry into user_changelog table.
func (u UserChangeLogSQL) CreateRelation(user entity.User, currentTime time.Time) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s","%s")
VALUES ($1, $2);
`,
		table.UserChangeLog.TableName,
		table.UserChangeLog.ColumnUserID,
		table.UserChangeLog.ColumnLastViewedAt,
	)

	_, err := u.db.Exec(
		statement,
		user.ID,
		currentTime,
	)

	return err
}

// NewUserChangeLogSQL creates UserChangeLogSQL
func NewUserChangeLogSQL(db *sql.DB) UserChangeLogSQL {
	return UserChangeLogSQL{
		db: db,
	}
}

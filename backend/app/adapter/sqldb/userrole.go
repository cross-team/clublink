package sqldb

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/role"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.UserRole = (*UserRoleSQL)(nil)

// UserRoleSQL accesses a user's role information in user_role table through SQL.
type UserRoleSQL struct {
	db *sql.DB
}

// GetRoles fetches all the roles for a given user
func (u UserRoleSQL) GetRoles(user entity.User) ([]role.Role, error) {
	statement := fmt.Sprintf(`
SELECT "%s" 
FROM "%s" 
WHERE "%s"=$1;
`,
		table.UserRole.ColumnRole,
		table.UserRole.TableName,
		table.UserRole.ColumnUserID,
	)

	roles := []role.Role{}
	rows, err := u.db.Query(statement, user.ID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrEntryNotFound("user not found")
	}
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var r role.Role

		if err = rows.Scan(&r); err != nil {
			return nil, err
		}

		roles = append(roles, r)
	}

	return roles, nil
}

// AddRole assigns the given role to the user
func (u UserRoleSQL) AddRole(user entity.User, r role.Role) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s", "%s")
VALUES ($1, $2);
`,
		table.UserRole.TableName,
		table.UserRole.ColumnUserID,
		table.UserRole.ColumnRole,
	)

	_, err := u.db.Exec(statement, user.ID, r)
	return err
}

// DeleteRole removes the given role from the user
func (u UserRoleSQL) DeleteRole(user entity.User, r role.Role) error {
	statement := fmt.Sprintf(`
DELETE FROM "%s"
WHERE "%s"=$1 AND "%s"=$2;
`,
		table.UserRole.TableName,
		table.UserRole.ColumnUserID,
		table.UserRole.ColumnRole,
	)

	_, err := u.db.Exec(statement, user.ID, r)
	return err
}

// NewUserRoleSQL creates UserRoleSQL
func NewUserRoleSQL(db *sql.DB) UserRoleSQL {
	return UserRoleSQL{db: db}
}

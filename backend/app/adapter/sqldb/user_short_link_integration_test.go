// +build integration all

package sqldb_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/db/dbtest"
	"github.com/short-d/short/backend/app/adapter/sqldb"
	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/fw/must"
	"github.com/short-d/short/backend/app/fw/ptr"
)

var insertUserShortLinkRowSQL = fmt.Sprintf(`
INSERT INTO %s (%s, %s)
VALUES ($1, $2)`,
	table.UserShortLink.TableName,
	table.UserShortLink.ColumnShortLinkAlias,
	table.UserShortLink.ColumnUserID,
)

type userShortLinkTableRow struct {
	alias  string
	userID string
}

func TestListShortLinkSql_FindAliasesByUser(t *testing.T) {
	testCases := []struct {
		name               string
		userTableRows      []userTableRow
		shortLinkTableRows []shortLinkTableRow
		relationTableRows  []userShortLinkTableRow
		user               entity.User
		hasErr             bool
		expectedAliases    []string
	}{
		{
			name:               "no alias found",
			userTableRows:      []userTableRow{},
			shortLinkTableRows: []shortLinkTableRow{},
			relationTableRows:  []userShortLinkTableRow{},
			user: entity.User{
				ID:             "test",
				Name:           "mockedUser",
				Email:          "test@example.com",
				LastSignedInAt: ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				CreatedAt:      ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				UpdatedAt:      ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
			},
			hasErr:          false,
			expectedAliases: nil,
		},
		{
			name: "aliases found",
			userTableRows: []userTableRow{
				{
					id:    "test",
					name:  "mockedUser",
					email: "test@example.com",
				},
			},
			shortLinkTableRows: []shortLinkTableRow{
				{alias: "abcd-123-xyz"},
			},
			relationTableRows: []userShortLinkTableRow{
				{
					alias:  "abcd-123-xyz",
					userID: "test",
				},
			},
			user: entity.User{
				ID:             "test",
				Name:           "mockedUser",
				Email:          "test@example.com",
				LastSignedInAt: ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				CreatedAt:      ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				UpdatedAt:      ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
			},
			hasErr: false,
			expectedAliases: []string{
				"abcd-123-xyz",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbtest.AccessTestDB(
				dbConnector,
				dbMigrationTool,
				dbMigrationRoot,
				dbConfig,
				func(sqlDB *sql.DB) {
					insertUserTableRows(t, sqlDB, testCase.userTableRows)
					insertShortLinkTableRows(t, sqlDB, testCase.shortLinkTableRows)
					insertUserShortLinkTableRows(t, sqlDB, testCase.relationTableRows)

					userShortLinkRepo := sqldb.NewUserShortLinkSQL(sqlDB)
					result, err := userShortLinkRepo.FindAliasesByUser(testCase.user)

					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}
					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedAliases, result)
				})
		})
	}
}

func TestListShortLinkSql_HasMapping(t *testing.T) {
	testCases := []struct {
		name               string
		userTableRows      []userTableRow
		shortLinkTableRows []shortLinkTableRow
		relationTableRows  []userShortLinkTableRow
		alias              string
		user               entity.User
		expectIsFound      bool
	}{
		{
			name: "alias does not exist",
			userTableRows: []userTableRow{
				{
					id:           "test",
					email:        "test@example.com",
					name:         "mockedUser",
					lastSignedIn: ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
					createdAt:    ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
					updatedAt:    ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				},
			},
			shortLinkTableRows: []shortLinkTableRow{},
			relationTableRows:  []userShortLinkTableRow{},
			alias:              "fizzbuzz",
			user: entity.User{
				ID:             "test",
				Name:           "mockedUser",
				Email:          "test@example.com",
				LastSignedInAt: ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				CreatedAt:      ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				UpdatedAt:      ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
			},
			expectIsFound: false,
		},
		{
			name: "alias does not belong to the user",
			userTableRows: []userTableRow{
				{
					id:           "test",
					email:        "test@example.com",
					name:         "mockedUser",
					lastSignedIn: ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
					createdAt:    ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
					updatedAt:    ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				},
				{
					id:           "test2",
					email:        "test2@example.com",
					name:         "mockedUser2",
					lastSignedIn: ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
					createdAt:    ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
					updatedAt:    ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				},
			},
			shortLinkTableRows: []shortLinkTableRow{
				{
					alias: "fizzbuzz",
				},
			},
			relationTableRows: []userShortLinkTableRow{
				{
					alias:  "fizzbuzz",
					userID: "test2",
				},
			},
			alias: "fizzbuzz",
			user: entity.User{
				ID:             "test",
				Name:           "mockedUser",
				Email:          "test@example.com",
				LastSignedInAt: ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				CreatedAt:      ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				UpdatedAt:      ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
			},
			expectIsFound: false,
		},
		{
			name: "alias belongs to the user",
			userTableRows: []userTableRow{
				{
					id:           "test",
					email:        "test@example.com",
					name:         "mockedUser",
					lastSignedIn: ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
					createdAt:    ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
					updatedAt:    ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				},
			},
			shortLinkTableRows: []shortLinkTableRow{
				{
					alias: "fizzbuzz",
				},
			},
			relationTableRows: []userShortLinkTableRow{
				{
					alias:  "fizzbuzz",
					userID: "test",
				},
			},
			alias: "fizzbuzz",
			user: entity.User{
				ID:             "test",
				Name:           "mockedUser",
				Email:          "test@example.com",
				LastSignedInAt: ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				CreatedAt:      ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				UpdatedAt:      ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
			},
			expectIsFound: true,
		},
		{
			name: "user does not exist",
			userTableRows: []userTableRow{
				{
					id:           "test",
					email:        "test@example.com",
					name:         "mockedUser",
					lastSignedIn: ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
					createdAt:    ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
					updatedAt:    ptr.Time(must.Time(t, "2019-05-01T08:02:16Z")),
				},
			},
			shortLinkTableRows: []shortLinkTableRow{
				{
					alias: "fizzbuzz",
				},
			},
			relationTableRows: []userShortLinkTableRow{
				{
					alias:  "fizzbuzz",
					userID: "test",
				},
			},
			alias:         "fizzbuzz",
			user:          entity.User{},
			expectIsFound: false,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbtest.AccessTestDB(
				dbConnector,
				dbMigrationTool,
				dbMigrationRoot,
				dbConfig,
				func(sqlDB *sql.DB) {
					insertUserTableRows(t, sqlDB, testCase.userTableRows)
					insertShortLinkTableRows(t, sqlDB, testCase.shortLinkTableRows)
					insertUserShortLinkTableRows(t, sqlDB, testCase.relationTableRows)

					userShortLinkRepo := sqldb.NewUserShortLinkSQL(sqlDB)
					result, err := userShortLinkRepo.HasMapping(testCase.user, testCase.alias)
					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectIsFound, result)
				})
		})
	}
}

func insertUserShortLinkTableRows(
	t *testing.T,
	sqlDB *sql.DB,
	tableRows []userShortLinkTableRow,
) {
	for _, tableRow := range tableRows {
		_, err := sqlDB.Exec(
			insertUserShortLinkRowSQL,
			tableRow.alias,
			tableRow.userID,
		)
		assert.Equal(t, nil, err)
	}
}

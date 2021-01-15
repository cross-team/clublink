package authorizer

import (
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac"
	"github.com/short-d/short/backend/app/usecase/authorizer/rbac/role"
	"github.com/short-d/short/backend/app/usecase/repository"
)

func TestAuthorizer_hasPermission(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		roles    map[string][]role.Role
		user     entity.User
		hasRight bool
	}{
		{
			name: "permission denied",
			roles: map[string][]role.Role{
				"id": {role.Basic},
			},
			user: entity.User{
				ID: "id",
			},
			hasRight: false,
		},
		{
			name: "permission granted",
			roles: map[string][]role.Role{
				"id": {role.Admin},
			},
			user: entity.User{
				ID: "id",
			},
			hasRight: true,
		},
		{
			name: "multiple roles grant the permission",
			roles: map[string][]role.Role{
				"id": {role.Basic, role.Admin},
			},
			user: entity.User{
				ID: "id",
			},
			hasRight: true,
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			fakeRolesRepo := repository.NewUserRoleFake(testCase.roles)
			ac := rbac.NewRBAC(fakeRolesRepo)
			authorizer := NewAuthorizer(ac)

			canChange, err := authorizer.CanCreateChange(testCase.user)
			assert.Equal(t, nil, err)

			assert.Equal(t, testCase.hasRight, canChange)
		})
	}
}

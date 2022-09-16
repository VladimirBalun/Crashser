package mongodb_repository

import (
	"errors"
	mock_repositories "server/internal/app/repositories/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetApplicationNames(t *testing.T) {
	many := []string{"app1", "app2", "app3"}
	one := []string{"app1"}
	empty := []string{}

	c := gomock.NewController(t)
	defer c.Finish()

	tests := []struct {
		name  string
		stubs func(store *mock_repositories.MockApplicationsRepository, slice []string)
		slice []string
		error error
	}{
		{
			name: "ok - many elements",
			stubs: func(r *mock_repositories.MockApplicationsRepository, slice []string) {
				r.EXPECT().GetApplicationNames().Return(slice, nil)
			},
			slice: many,
			error: nil,
		},
		{
			name: "ok - one app",
			stubs: func(r *mock_repositories.MockApplicationsRepository, slice []string) {
				r.EXPECT().GetApplicationNames().Return(slice, nil)
			},
			slice: one,
			error: nil,
		},
		{
			name: "ok - empty slice",
			stubs: func(r *mock_repositories.MockApplicationsRepository, slice []string) {
				r.EXPECT().GetApplicationNames().Return(slice, nil)
			},
			slice: empty,
			error: nil,
		},
		{
			name: "error",
			stubs: func(r *mock_repositories.MockApplicationsRepository, slice []string) {
				r.EXPECT().GetApplicationNames().Return(slice, errors.New("error getting application names"))
			},
			slice: empty,
			error: errors.New("error getting application names"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := mock_repositories.NewMockApplicationsRepository(c)
			test.stubs(r, test.slice)

			result, err := r.GetApplicationNames()
			require.Equal(t, test.error, err)
			require.Equal(t, len(test.slice), len(result))
		})
	}
}

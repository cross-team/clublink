package search

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/search/order"
)

type shortLinks = map[string]entity.ShortLink

func TestSearch(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name               string
		shortLinks         shortLinks
		Query              Query
		maxResults         int
		resources          []Resource
		orders             []order.By
		relationUsers      []entity.User
		relationShortLinks []entity.ShortLink
		expectedResult     Result
	}{
		{
			name: "search without user",
			shortLinks: shortLinks{
				"git-google": entity.ShortLink{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				"google": entity.ShortLink{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				"short": entity.ShortLink{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				"facebook": entity.ShortLink{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			Query: Query{
				Query: "http google",
			},
			maxResults: 2,
			resources:  []Resource{ShortLink},
			orders:     []order.By{order.ByCreatedTimeASC},
			relationUsers: []entity.User{
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "beta",
					Email: "beta@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			relationShortLinks: []entity.ShortLink{
				{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			expectedResult: Result{},
		},
		{
			name: "search without query",
			shortLinks: shortLinks{
				"git-google": entity.ShortLink{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				"google": entity.ShortLink{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				"short": entity.ShortLink{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				"facebook": entity.ShortLink{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			Query: Query{
				User: &entity.User{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			maxResults: 2,
			resources:  []Resource{ShortLink},
			orders:     []order.By{order.ByCreatedTimeASC},
			relationUsers: []entity.User{
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "beta",
					Email: "beta@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			relationShortLinks: []entity.ShortLink{
				{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			expectedResult: Result{
				ShortLinks: []entity.ShortLink{
					{
						Alias:    "facebook",
						LongLink: "https://facebook.com",
					},
					{
						Alias:    "short",
						LongLink: "https://short-d.com",
					},
				},
			},
		},
		{
			name: "valid search",
			shortLinks: shortLinks{
				"git-google": entity.ShortLink{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				"google": entity.ShortLink{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				"short": entity.ShortLink{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				"facebook": entity.ShortLink{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			Query: Query{
				Query: "http google",
				User: &entity.User{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			maxResults: 2,
			resources:  []Resource{ShortLink},
			orders:     []order.By{order.ByCreatedTimeASC},
			relationUsers: []entity.User{
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "beta",
					Email: "beta@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			relationShortLinks: []entity.ShortLink{
				{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			expectedResult: Result{
				ShortLinks: []entity.ShortLink{
					{
						Alias:    "google",
						LongLink: "https://google.com",
					},
					{
						Alias:    "git-google",
						LongLink: "http://github.com/google",
					},
				},
				Users: nil,
			},
		},
		{
			name: "query no match",
			shortLinks: shortLinks{
				"git-google": entity.ShortLink{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				"google": entity.ShortLink{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				"short": entity.ShortLink{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				"facebook": entity.ShortLink{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			Query: Query{
				Query: "non existent",
				User: &entity.User{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			maxResults: 2,
			resources:  []Resource{ShortLink},
			orders:     []order.By{order.ByCreatedTimeASC},
			relationUsers: []entity.User{
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "beta",
					Email: "beta@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			relationShortLinks: []entity.ShortLink{
				{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			expectedResult: Result{
				ShortLinks: nil,
				Users:      nil,
			},
		},
		{
			name: "less matches than maxResults",
			shortLinks: shortLinks{
				"git-google": entity.ShortLink{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				"google": entity.ShortLink{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				"short": entity.ShortLink{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				"facebook": entity.ShortLink{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			Query: Query{
				Query: "short",
				User: &entity.User{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			maxResults: 2,
			resources:  []Resource{ShortLink},
			orders:     []order.By{order.ByCreatedTimeASC},
			relationUsers: []entity.User{
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "beta",
					Email: "beta@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			relationShortLinks: []entity.ShortLink{
				{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			expectedResult: Result{
				ShortLinks: []entity.ShortLink{
					{
						Alias:    "short",
						LongLink: "https://short-d.com",
					},
				},
				Users: nil,
			},
		},
		{
			name: "search more than one resource",
			shortLinks: shortLinks{
				"git-google": entity.ShortLink{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				"google": entity.ShortLink{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				"short": entity.ShortLink{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				"facebook": entity.ShortLink{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			Query: Query{
				Query: "short",
				User: &entity.User{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			maxResults: 2,
			resources:  []Resource{ShortLink, User, Unknown},
			orders:     []order.By{order.ByCreatedTimeASC, order.ByUnsorted, order.ByCreatedTimeASC},
			relationUsers: []entity.User{
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "beta",
					Email: "beta@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
				{
					ID:    "alpha",
					Email: "alpha@example.com",
				},
			},
			relationShortLinks: []entity.ShortLink{
				{
					Alias:    "git-google",
					LongLink: "http://github.com/google",
				},
				{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				{
					Alias:    "google",
					LongLink: "https://google.com",
				},
				{
					Alias:    "short",
					LongLink: "https://short-d.com",
				},
				{
					Alias:    "facebook",
					LongLink: "https://facebook.com",
				},
			},
			expectedResult: Result{
				ShortLinks: []entity.ShortLink{
					{
						Alias:    "short",
						LongLink: "https://short-d.com",
					},
				},
				Users: nil,
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			userShortLinkRepo := repository.NewUserShortLinkRepoFake(testCase.relationUsers, testCase.relationShortLinks)
			shortLinkRepo := repository.NewShortLinkFake(nil, testCase.shortLinks)
			timeout := time.Second

			entryRepo := logger.NewEntryRepoFake()
			lg, err := logger.NewFake(logger.LogOff, &entryRepo)
			assert.Equal(t, nil, err)

			search := NewSearch(lg, &shortLinkRepo, &userShortLinkRepo, timeout)

			filter, err := NewFilter(testCase.maxResults, testCase.resources, testCase.orders)
			assert.Equal(t, nil, err)

			result, err := search.Search(testCase.Query, filter)

			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedResult, result)
		})
	}
}

package repokit_test

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/imantung/boilerplate-go-backend/pkg/repokit"
	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {
	testcases := []struct {
		testName      string
		sorts         repokit.Sorts
		builder       sq.SelectBuilder
		expectedQuery string
		expectedArgs  []interface{}
	}{
		{
			sorts:         repokit.Sorts{},
			builder:       sq.Select("col1", "col2", "col3").From("sometables"),
			expectedQuery: "SELECT col1, col2, col3 FROM sometables",
		},
		{
			sorts:         repokit.Sorts{"col1", "+col2", "-col3"},
			builder:       sq.Select("col1", "col2", "col3").From("sometables"),
			expectedQuery: "SELECT col1, col2, col3 FROM sometables ORDER BY col1 ASC, col2 ASC, col3 DESC",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			query, args, _ := tt.sorts.CompileSelect(tt.builder).ToSql()
			require.Equal(t, tt.expectedQuery, query)
			require.Equal(t, tt.expectedArgs, args)
		})
	}
}

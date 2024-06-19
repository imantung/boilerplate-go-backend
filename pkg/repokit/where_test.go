package repokit_test

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/imantung/boilerplate-go-backend/pkg/repokit"
	"github.com/stretchr/testify/require"
)

func TestWhere_CompileSelect(t *testing.T) {
	testcases := []struct {
		testName string
		repokit.Where
		base          sq.SelectBuilder
		expectedQuery string
		expectedArgs  []interface{}
	}{
		{
			base:          sq.Select("name", "version").From("some-table"),
			expectedQuery: "SELECT name, version FROM some-table",
		},
		{
			Where: repokit.Where{
				sq.Eq{"name": "dummy-name"},
				sq.GtOrEq{"version": 1},
			},
			base:          sq.Select("name", "version").From("some-table"),
			expectedQuery: "SELECT name, version FROM some-table WHERE name = ? AND version >= ?",
			expectedArgs:  []interface{}{"dummy-name", 1},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			query, args, err := tt.CompileSelect(tt.base).ToSql()
			require.NoError(t, err)
			require.Equal(t, tt.expectedQuery, query)
			require.Equal(t, tt.expectedArgs, args)
		})
	}
}

func TestWhere_CompileUpdate(t *testing.T) {
	testcases := []struct {
		testName string
		repokit.Where
		base          sq.UpdateBuilder
		expectedQuery string
		expectedArgs  []interface{}
	}{
		{
			Where: repokit.Where{
				sq.Eq{"name": "dummy-name"},
				sq.LtOrEq{"version": 2},
			},
			base:          sq.Update("some-table").Set("column", "column-value"),
			expectedQuery: "UPDATE some-table SET column = ? WHERE name = ? AND version <= ?",
			expectedArgs:  []interface{}{"column-value", "dummy-name", 2},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			query, args, _ := tt.CompileUpdate(tt.base).ToSql()
			require.Equal(t, tt.expectedQuery, query)
			require.Equal(t, tt.expectedArgs, args)
		})
	}

}

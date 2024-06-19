package repokit_test

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/imantung/boilerplate-go-backend/pkg/repokit"
	"github.com/stretchr/testify/require"
)

func TestSelectOption(t *testing.T) {
	expected := sq.Select("")
	selectOpt := repokit.NewSelectOption(func(sq.SelectBuilder) sq.SelectBuilder {
		return expected
	})
	require.Equal(t, expected, selectOpt.CompileSelect(sq.Select("")))
}

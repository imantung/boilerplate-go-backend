package repokit_test

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/imantung/boilerplate-go-backend/pkg/repokit"
	"github.com/stretchr/testify/require"
)

func TestUpdateOption(t *testing.T) {
	expected := sq.Update("")
	selectOpt := repokit.NewUpdateOption(func(sq.UpdateBuilder) sq.UpdateBuilder {
		return expected
	})

	require.Equal(t, expected, selectOpt.CompileUpdate(sq.Update("")))

}

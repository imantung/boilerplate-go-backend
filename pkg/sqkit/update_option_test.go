package sqkit_test

import (
	"testing"

	sq "github.com/Masterminds/squirrel"
	"github.com/imantung/boilerplate-go-backend/pkg/sqkit"
	"github.com/stretchr/testify/require"
)

func TestUpdateOption(t *testing.T) {
	expected := sq.Update("")
	selectOpt := sqkit.NewUpdateOption(func(sq.UpdateBuilder) sq.UpdateBuilder {
		return expected
	})

	require.Equal(t, expected, selectOpt.CompileUpdate(sq.Update("")))

}

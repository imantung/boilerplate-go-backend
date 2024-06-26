package repokit

import sq "github.com/Masterminds/squirrel"

type (
	// Where conditions as squirel query builder
	Where []interface{}
)

var _ SelectOption = (Where)(nil)
var _ UpdateOption = (Where)(nil)

// CompileSelect to compile select query for filtering
func (e Where) CompileSelect(base sq.SelectBuilder) sq.SelectBuilder {
	for _, cond := range e {
		base = base.Where(cond)
	}
	return base
}

// CompileUpdate to compile update query for filtering
func (e Where) CompileUpdate(base sq.UpdateBuilder) sq.UpdateBuilder {
	for _, cond := range e {
		base = base.Where(cond)
	}
	return base
}

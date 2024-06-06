package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertToFieldName(t *testing.T) {
	require.Equal(t, "ID", convertToFieldName("id"))
	require.Equal(t, "EmployeeID", convertToFieldName("employee_id"))
	require.Equal(t, "TableIdent", convertToFieldName("table_ident"))
	require.Equal(t, "SomeOtherField", convertToFieldName("some_other_field"))
}

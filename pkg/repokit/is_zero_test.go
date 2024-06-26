package repokit_test

import (
	"testing"

	"github.com/imantung/boilerplate-go-backend/pkg/repokit"
)

func TestIsZero(t *testing.T) {
	testcases := []struct {
		TestName string
		V        interface{}
		Expected bool
	}{
		{V: "", Expected: true},
		{V: "asdf", Expected: false},
		{V: 0, Expected: true},
		{V: 1234, Expected: false},
		{V: false, Expected: true},
		{V: nil, Expected: true},
		{V: TestIsZero, Expected: false},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			repokit.IsZero(tt.V)
		})
	}

}

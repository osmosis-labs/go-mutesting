package branch

import (
	"testing"

	"github.com/osmosis-labs/go-mutesting/test"
)

func TestMutatorIf(t *testing.T) {
	test.Mutator(
		t,
		MutatorIf,
		"../../testdata/branch/mutateif.go",
		2,
	)
}

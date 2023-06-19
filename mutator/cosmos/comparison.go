package cosmos

import (
	"go/ast"
	"go/types"

	"github.com/osmosis-labs/go-mutesting/mutator"
)

func init() {
	mutator.Register("cosmos/comparison", MutatorComparisonCosmos)
}

var comparisonMutations = map[string][]string{
	"GT":  {"LTE", "LT", "GTE"},
	"LT":  {"GTE", "GT", "LTE"},
	"GTE": {"LT", "GT", "LTE"},
	"LTE": {"GT", "GTE", "LT"},
}

// MutatorComparisonCosmos implements a mutator to change Cosmos SDK comparisons.
func MutatorComparisonCosmos(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.Ident)
	if !ok {
		return nil
	}

	// ensure node has a valid SDK comparison operator
	mutations, ok := comparisonMutations[n.Name]
	if !ok || len(mutations) == 0 {
		return nil
	}

	o := n.Name
	var mutationsResult []mutator.Mutation

	for _, mutation := range mutations {
		r := mutation

		mutationsResult = append(mutationsResult, mutator.Mutation{
			Change: func() {
				n.Name = r
			},
			Reset: func() {
				n.Name = o
			},
		})
	}

	return mutationsResult
}

package cosmos

import (
	"go/ast"
	"go/types"

	"github.com/osmosis-labs/go-mutesting/mutator"
)

func init() {
	mutator.Register("cosmos/cl", MutatorClCosmos)
}

var clMutations = map[string]string{
	"Liquidity0":       "Liquidity1",
	"Liquidity1":       "Liquidity0",
	"CalcAmount0Delta": "CalcAmount1Delta",
	"CalcAmount1Delta": "CalcAmount0Delta",
}

// MutatorClCosmos implements a mutator to change concentrated liquidity helpers.
func MutatorClCosmos(pkg *types.Package, info *types.Info, node ast.Node) []mutator.Mutation {
	n, ok := node.(*ast.Ident)
	if !ok {
		return nil
	}

	if _, ok := clMutations[n.Name]; !ok {
		return nil
	}

	o := n.Name
	r, ok := clMutations[n.Name]
	if !ok {
		return nil
	}

	return []mutator.Mutation{
		{
			Change: func() {
				n.Name = r
			},
			Reset: func() {
				n.Name = o
			},
		},
	}
}

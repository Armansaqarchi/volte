package circuits

import (
	"flag"
	"github.com/consensys/gnark/frontend"
)

var maxVoteValues = flag.Int("max_vote_values", 100, "Maximum possible values for vote.")

type VoteCircuit struct {
	Vote frontend.Variable
}

func (c *VoteCircuit) Define(api frontend.API) error {
	// To check 0 <= v < m equality, we are going to check (v)(v-2)(v-3)...(v-m+1) = 0 instead.

	equation := c.Vote
	for i := 1; i < *maxVoteValues; i++ {
		equation = api.Mul(equation, api.Sub(c.Vote, i))
	}

	api.AssertIsEqual(equation, 0)
	return nil
}

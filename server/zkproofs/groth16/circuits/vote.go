package circuits

import "github.com/consensys/gnark/frontend"
import "volte/backend/consts"

type VoteCircuit struct {
	Vote frontend.Variable
}

func (c *VoteCircuit) Define(api frontend.API) error {
	// To check 0 <= v < m equality, we are going to check (v)(v-2)(v-3)...(v-m+1) = 0 instead.

	equation := c.Vote
	for i := 1; i < consts.MAX_TOTAL_VOTES; i++ {
		equation = api.Mul(equation, api.Sub(c.Vote, i))
	}

	api.AssertIsEqual(equation, 0)
	return nil
}

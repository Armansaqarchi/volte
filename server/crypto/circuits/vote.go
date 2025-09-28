package circuits

import (
	"flag"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/poseidon2"
)

var maxVoteValues = flag.Int("max_vote_values", 100, "Maximum possible values for vote.")

type BallotCircuit struct {
	Vote          frontend.Variable
	EncryptedVote frontend.Variable `gnark:",public"`
	EncryptionKey frontend.Variable `gnark:",public"`
}

func (c *BallotCircuit) Define(api frontend.API) error {
	// To check 0 <= v < m equality, we are going to check (v)(v-2)(v-3)...(v-m+1) = 0 instead.

	equation := c.Vote
	for i := 1; i < *maxVoteValues; i++ {
		equation = api.Mul(equation, api.Sub(c.Vote, i))
	}

	// Creating new poseidon hashing circuit.
	h, err := poseidon2.NewMerkleDamgardHasher(api)
	if err != nil {
		return err
	}
	// Absorb inputs into the poseidon sponge state.
	h.Write(c.Vote, c.EncryptionKey)
	// Finalize the poseidon sponge state and return the output.
	// This returns the rate coefficients from the state s = [r1, r2, r3, .... rs, c1, c2, c3, .... cm].
	digest := h.Sum()
	// Make sure the encrypted key is correctly calculated.
	api.AssertIsEqual(digest, c.EncryptedVote)
	api.AssertIsEqual(equation, 0)
	return nil
}

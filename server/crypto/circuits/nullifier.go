package circuits

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/poseidon2"
)

type NullifierCircuit struct {
	SecretKey frontend.Variable
	EventID   frontend.Variable `gnark:",public"`
	Nullifier frontend.Variable `gnark:",public"`
}

func (c *NullifierCircuit) Define(api frontend.API) error {
	// Creating new poseidon hashing circuit.
	h, err := poseidon2.NewMerkleDamgardHasher(api)
	if err != nil {
		return err
	}
	// Absorb inputs into the poseidon sponge state.
	h.Write(c.SecretKey, c.EventID)
	// Finalize the poseidon sponge state and return the output.
	// This returns the rate coefficients from the state s = [r1, r2, r3, .... rs, c1, c2, c3, .... cm].
	digest := h.Sum()
	api.AssertIsEqual(digest, c.Nullifier)
	return nil
}

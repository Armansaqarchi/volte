package circuits

import (
	"fmt"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
	"log/slog"
)

type MerkleCircuit struct {
	// We keep merkleRoot public because it's supposed to be stored on chain.
	MerkleRoot frontend.Variable `gnark:",public"`
	LeafValue  frontend.Variable `gnark:",public"`
	// List of siblings alongside the leaf's path up to the root.
	MerklePath []frontend.Variable `gnark:",secret"`
	// List of bits indicating child position at each index (0 indicates left, 1 indicates right). Without this,
	// positions can be mistaken and result in wrong hashing.
	PathPositions []frontend.Variable `gnark:",secret"`
	// The user's secret key we check the leaf value against.
	SecretKey frontend.Variable `gnark:",secret"`
}

func (c *MerkleCircuit) Define(api frontend.API) error {
	hasher, err := mimc.NewMiMC(api)
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't instantiate mimc hash. err : %s", err))
		panic(err)
	}
	// Check if secret key matches the user commitment!
	hasher.Write(c.SecretKey)
	api.AssertIsEqual(hasher.Sum(), c.LeafValue)

	current := c.LeafValue
	for i := 0; i < len(c.MerklePath); i++ {
		sibling := c.MerklePath[i]
		isRight := c.PathPositions[i]

		// Always obtain the final hash input by concatenating the children from left to right.
		right := api.Select(isRight, current, sibling)
		left := api.Select(isRight, sibling, current)
		api.Println("left is ", left)
		api.Println("right is ", right)
		hasher.Reset()
		hasher.Write(left, right)
		current = hasher.Sum()
	}
	api.AssertIsEqual(current, c.MerkleRoot)
	return nil
}

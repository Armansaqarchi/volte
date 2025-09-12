package constraintsys

import (
	"flag"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"log/slog"
	circuits2 "volte/backend/crypto/circuits"
)

var treeDepth = flag.Int("tree_depth", 10, "Merkle trees depth.")

type volteCircuit struct {
	Merkle   frontend.Circuit
	Poseidon frontend.Circuit
	Vote     frontend.Circuit
}

func (c *volteCircuit) Define(api frontend.API) error {
	if err := c.Merkle.Define(api); err != nil {
		slog.Error(fmt.Sprintf("Failed to define merkle circuit. err : %s", err))
	}
	if err := c.Poseidon.Define(api); err != nil {
		slog.Error(fmt.Sprintf("Failed to define poseidon circuit. err : %s", err))
	}
	if err := c.Vote.Define(api); err != nil {
		slog.Error(fmt.Sprintf("Failed to define vote curcuit. err : %s", err))
	}
	return nil
}

type VolteR1CS interface {
	Compile() constraint.ConstraintSystem
}

type volteR1CS struct {
	volteCircuit volteCircuit
	field        ecc.ID
}

func NewVolteBLS12377R1CS() VolteR1CS {
	return &volteR1CS{
		volteCircuit: volteCircuit{
			Merkle: &circuits2.MerkleCircuit{
				PathPositions: make([]frontend.Variable, *treeDepth),
				MerklePath:    make([]frontend.Variable, *treeDepth),
			},
			Poseidon: &circuits2.PoseidonCircuit{},
			Vote:     &circuits2.VoteCircuit{},
		},
		field: ecc.BLS12_377,
	}
}

func (v volteR1CS) Compile() constraint.ConstraintSystem {
	volteR1CS, err := frontend.Compile(ecc.BLS12_377.ScalarField(), r1cs.NewBuilder, &v.volteCircuit)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to compile volte R1CS. err : %s", err))
		panic(err)
	}
	return volteR1CS
}

// generate R1CS for other circuits as well.

package groth16

import (
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"log/slog"
	"volte/backend/zkproofs/groth16/circuits"
)

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

type VolteR1CS struct {
	volteCircuit volteCircuit
	field        ecc.ID
}

func NewVolteBLS12377R1CS(nullifierTreeDepth int) VolteR1CS {
	return VolteR1CS{
		volteCircuit: volteCircuit{
			Merkle:   &circuits.MerkleCircuit{},
			Poseidon: &circuits.PoseidonCircuit{},
			Vote:     &circuits.VoteCircuit{},
		},
		field: ecc.BLS12_377,
	}
}

func (v VolteR1CS) Compile() constraint.ConstraintSystem {
	volteR1CS, err := frontend.Compile(ecc.BLS12_377.ScalarField(), r1cs.NewBuilder, &v.volteCircuit)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to compile volte R1CS. err : %s", err))
		panic(err)
	}
	return volteR1CS
}

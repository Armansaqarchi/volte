package constraintsys

import (
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"log/slog"
)

type R1CS interface {
	GetConstraintSystem() constraint.ConstraintSystem
	Compile() constraint.ConstraintSystem
}

type volteR1CS struct {
	circuit          frontend.Circuit
	constraintSystem constraint.ConstraintSystem
	field            ecc.ID
}

func NewVolteBLS12377R1CS(circuit frontend.Circuit) R1CS {
	return &volteR1CS{
		circuit: circuit,
		field:   ecc.BLS12_377,
	}
}

func (v *volteR1CS) GetConstraintSystem() constraint.ConstraintSystem {
	return v.constraintSystem
}

func (v *volteR1CS) Compile() constraint.ConstraintSystem {
	css, err := frontend.Compile(ecc.BLS12_377.ScalarField(), r1cs.NewBuilder, v.circuit)
	v.constraintSystem = css
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to compile volte R1CS. err : %s", err))
		panic(err)
	}
	return css
}

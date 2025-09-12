package zkproofs

import (
	"fmt"
	"github.com/consensys/gnark/backend/groth16"
	"log/slog"
	"volte/backend/crypto/constraintsys"
)

// Groth16 is a base Groth16 wrapper that corresponding to an R1CS.
type Groth16 struct {
	r1cs         constraintsys.VolteR1CS
	provingKey   groth16.ProvingKey
	verifyingKey groth16.VerifyingKey
}

func (g *Groth16) GetProvingKey() groth16.ProvingKey {
	return g.provingKey
}

func (g *Groth16) GetVerifyingKey() groth16.VerifyingKey {
	return g.verifyingKey
}

func SetupNewGroth16(constraintSystem constraintsys.VolteR1CS) *Groth16 {
	g := new(Groth16)
	g.r1cs = constraintSystem
	cs := g.r1cs.Compile()
	slog.Info("Successfully compiled the volte circuit.")
	slog.Info(fmt.Sprintf("Number of public variables in the circuit : %d", cs.GetNbPublicVariables()))
	slog.Info(fmt.Sprintf("Number of secret variables in the circuit : %d", cs.GetNbSecretVariables()))
	slog.Info(fmt.Sprintf("Number of internal variables in the circuit : %d", cs.GetNbInternalVariables()))
	provingKey, verifyingKey, err := groth16.Setup(cs)
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't setup groth16 zkproof: %v", err))
		panic(err)
	}
	slog.Info("Successfully build groth16 zkproof")
	g.provingKey = provingKey
	g.verifyingKey = verifyingKey

	return g
}

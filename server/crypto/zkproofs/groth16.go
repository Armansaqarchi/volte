package zkproofs

import (
	"fmt"
	"github.com/consensys/gnark/backend/groth16"
	"log/slog"
	"volte/backend/crypto/constraintsys"
)

type VolteGroth16 struct {
	r1cs         constraintsys.VolteR1CS
	provingKey   groth16.ProvingKey
	verifyingKey groth16.VerifyingKey
}

func (g *VolteGroth16) Setup() {
	g.r1cs = constraintsys.NewVolteBLS12377R1CS()
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
}

func (g *VolteGroth16) GetProvingKey() groth16.ProvingKey {
	return g.provingKey
}

func (g *VolteGroth16) GetVerifyingKey() groth16.VerifyingKey {
	return g.verifyingKey
}

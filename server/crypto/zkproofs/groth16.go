package zkproofs

import (
	"fmt"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"log/slog"
	"volte/backend/crypto/circuits"
	"volte/backend/crypto/constraintsys"
)

// Groth16 is a base Groth16 wrapper that corresponding to an R1CS.
type Groth16 struct {
	r1cs         constraintsys.R1CS
	provingKey   groth16.ProvingKey
	verifyingKey groth16.VerifyingKey
}

func (g *Groth16) GetProvingKey() groth16.ProvingKey {
	return g.provingKey
}

func (g *Groth16) GetVerifyingKey() groth16.VerifyingKey {
	return g.verifyingKey
}

func SetupNewGroth16(constraintSystem constraintsys.R1CS) *Groth16 {
	g := new(Groth16)
	g.r1cs = constraintSystem
	cs := g.r1cs.Compile()

	fmt.Println(cs.GetNbPublicVariables())

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

func NewBallotGroth16() *Groth16 {
	return SetupNewGroth16(constraintsys.NewVolteBN254R1CS(new(circuits.BallotCircuit)))
}

func NewNullifierGroth16() *Groth16 {
	return SetupNewGroth16(constraintsys.NewVolteBN254R1CS(new(circuits.NullifierCircuit)))
}

func NewMembershipGroth16(len int) *Groth16 {
	// Length of arrays for this circuit are dynamic, so proving key and verifying key varies between other events.
	return SetupNewGroth16(constraintsys.NewVolteBN254R1CS(&circuits.MerkleCircuit{
		MerklePath:    make([]frontend.Variable, len),
		PathPositions: make([]frontend.Variable, len),
	}))
}

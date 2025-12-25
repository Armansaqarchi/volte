package zkproofs

import (
	"fmt"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/minio/sha256-simd"
	"io"
	"log/slog"
	"math/big"
	"volte/backend/crypto/circuits"
	"volte/backend/crypto/constraintsys"
	"volte/backend/crypto/utils"
)

func CreateProof(assignment frontend.Circuit, g16 *Groth16) ([]*big.Int, any) {
	fullWitness, err := frontend.NewWitness(assignment, ecc.BN254.ScalarField())
	if err != nil {
		slog.Error("new witness failed: %v", err)
	}
	slog.Info("Created witness for cs circuit.")
	pubWitness, err := fullWitness.Public()
	if err != nil {
		slog.Error(fmt.Sprintf("failed to extract public inputs from witness, err : %s", err))
	}
	proof, err := groth16.Prove(g16.GetR1CSSystem().GetConstraintSystem(), g16.GetProvingKey(), fullWitness, func(config *backend.ProverConfig) error {
		config.HashToFieldFn = sha256.New()
		return nil
	})
	if err != nil {
		slog.Error(fmt.Sprintf("failed to generate proof: %v", err))
		panic(err)
	}
	slog.Info("Proof created.")
	proofParts, err := utils.ExtractProof(proof)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to extract public inputs from witness, err : %s", err))
	}
	return proofParts, pubWitness.Vector()
}

// Groth16 is a base Groth16 wrapper that corresponds to an R1CS.
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

func (g *Groth16) GetR1CSSystem() constraintsys.R1CS {
	return g.r1cs
}

func SetupNewGroth16(constraintSystem constraintsys.R1CS) *Groth16 {
	g := new(Groth16)
	g.r1cs = constraintSystem
	cs := g.r1cs.Compile()

	fmt.Println(cs.GetNbPublicVariables())

	slog.Info("Successfully compiled the circuit.")
	slog.Info(fmt.Sprintf("Number of public variables in the circuit : %d", cs.GetNbPublicVariables()))
	slog.Info(fmt.Sprintf("Number of secret variables in the circuit : X%d", cs.GetNbSecretVariables()))
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

func SetupNewGroth16FromKeys(constraintSystem constraintsys.R1CS, provingKey io.Reader,
	VerifyingKey io.Reader) *Groth16 {
	g := new(Groth16)
	g.r1cs = constraintSystem
	cs := g.r1cs.Compile()

	slog.Info("Successfully compiled the circuit.")
	slog.Info(fmt.Sprintf("Number of public variables in the circuit : %d", cs.GetNbPublicVariables()))
	slog.Info(fmt.Sprintf("Number of secret variables in the circuit : %d", cs.GetNbSecretVariables()))
	slog.Info(fmt.Sprintf("Number of internal variables in the circuit : %d", cs.GetNbInternalVariables()))

	pk := groth16.NewProvingKey(g.r1cs.GetCurve())
	slog.Info("Reading proving key...")
	if _, err := pk.ReadFrom(provingKey); err != nil {
		slog.Error(fmt.Sprintf("failed to read circuit proving key: %v", err))
	}
	g.provingKey = pk
	slog.Info("Finished reading proving key.")
	vk := groth16.NewVerifyingKey(g.r1cs.GetCurve())
	slog.Info("Reading verifying key...")
	if _, err := vk.ReadFrom(VerifyingKey); err != nil {
		slog.Error(fmt.Sprintf("failed to read circuit verifying key: %v", err))
	}
	g.verifyingKey = vk
	slog.Info("Successfully build groth16 zkproof")

	return g
}

func NewBallotGroth16() *Groth16 {
	return SetupNewGroth16(constraintsys.NewVolteBN254R1CS(new(circuits.BallotCircuit)))
}

func NewBallotGroth16FromExistingKeys(verifyingKey io.Reader, provingKey io.Reader) *Groth16 {
	slog.Info("Instantiating the groth16 ballot from existing keys.")
	return SetupNewGroth16FromKeys(
		constraintsys.NewVolteBN254R1CS(new(circuits.BallotCircuit)), provingKey, verifyingKey,
	)
}

func NewNullifierGroth16() *Groth16 {
	return SetupNewGroth16(constraintsys.NewVolteBN254R1CS(new(circuits.NullifierCircuit)))
}

func NewNullifierGroth16FromExistingKeys(verifyingKey io.Reader, provingKey io.Reader) *Groth16 {
	return SetupNewGroth16FromKeys(
		constraintsys.NewVolteBN254R1CS(new(circuits.NullifierCircuit)), provingKey, verifyingKey,
	)
}

func NewMembershipGroth16(len int) *Groth16 {
	// Length of arrays for this circuit are dynamic, so proving key and verifying key varies between other events.
	return SetupNewGroth16(constraintsys.NewVolteBN254R1CS(&circuits.MerkleCircuit{
		MerklePath:    make([]frontend.Variable, len),
		PathPositions: make([]frontend.Variable, len),
	}))
}

func NewMembershipGroth16FromExistingKeys(len int, verifyingKey io.Reader, provingKey io.Reader) *Groth16 {
	// Length of arrays for this circuit are dynamic, so proving key and verifying key varies between other events.
	return SetupNewGroth16FromKeys(constraintsys.NewVolteBN254R1CS(&circuits.MerkleCircuit{
		MerklePath:    make([]frontend.Variable, len),
		PathPositions: make([]frontend.Variable, len),
	}), provingKey, verifyingKey)
}

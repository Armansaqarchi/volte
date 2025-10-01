package zkproofs

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"testing"
)

func TestBallotGroth16(t *testing.T) {
	g16 := NewBallotGroth16()

	witness, err := frontend.NewWitness(&assignment, ecc.BLS12_377.ScalarField())
	if err != nil {
		t.Fatal(err)
	}
	publicWitness, _ := witness.Public()

	proof, err := groth16.Prove(g16.r1cs.GetConstraintSystem(), g16.provingKey, witness)
	if err != nil {
		t.Fatal(err)
	}
	if err := groth16.Verify(proof, g16.verifyingKey, publicWitness); err != nil {
		t.Fatal(err)
	}
}

func TestNullifierGroth16(t *testing.T) {

}

func TestMembershipGroth16(t *testing.T) {

}

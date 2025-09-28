package zkproofs

import (
	"crypto/rand"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark-crypto/ecc/bls12-377/fr"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"math/big"
	"testing"
	"volte/backend/crypto/circuits"
)

func TestBallotGroth16(t *testing.T) {
	g16 := NewBallotGroth16()
	secretBytes := make([]byte, 32)
	secretKeyBigInt := new(big.Int).SetBytes(make([]byte, 32))
	_, _ = rand.Read(secretBytes)
	EncryptionKey := new(big.Int).SetInt64(1)
	var skEl fr.Element
	skEl.SetBigInt(secretKeyBigInt)
	var encryptionKeyEl fr.Element
	encryptionKeyEl.SetBigInt(EncryptionKey)
	assignment := circuits.BallotCircuit{
		Vote:          1,
		EncryptedVote: 2,
		EncryptionKey: encryptionKeyEl,
	}
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

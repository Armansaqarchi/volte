package circuits

import (
	"flag"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"math/big"
	"testing"
	"volte/backend/crypto/utils"
)

func TestBallotCircuit_Define(t *testing.T) {

	flag.Set("Gx", "Fake Gx")
	flag.Set("Gy", "Fake Gy")
	flag.Set("Yx", "Fake Yx")
	flag.Set("Yy", "Fake Yy")
	flag.Parse()

	G := utils.GenerateBaseECC()
	x := big.NewInt(30)
	Y := utils.G1MulAffine(&G, x)
	k := big.NewInt(10)
	m := big.NewInt(1)
	mG := utils.G1MulAffine(&G, m)
	kY := utils.G1MulAffine(&Y, k)
	C1 := utils.G1MulAffine(&G, k)
	C2 := utils.G1AddAffine(&mG, &kY)

	ballotCircuitMeta = &BallotCircuitMeta{
		g: utils.ECCToAffinePoint(G),
		y: utils.ECCToAffinePoint(Y),
	}
	assignment := &BallotCircuit{
		C1: utils.ECCToAffinePoint(C1),
		C2: utils.ECCToAffinePoint(C2),
		M:  utils.ValFr(m),
		k:  utils.ValFr(k),
	}

	var circuit BallotCircuit
	cs, err := frontend.Compile(ecc.BLS12_377.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		t.Fatalf("compile failed: %v", err)
	}

	fullWitness, err := frontend.NewWitness(assignment, ecc.BLS12_377.ScalarField())
	if err != nil {
		t.Fatalf("new witness failed: %v", err)
	}
	//publicWitness, err := fullWitness.Public()
	//if err != nil {
	//	t.Fatalf("public witness failed: %v", err)
	//}
	if err := cs.IsSolved(fullWitness); err != nil {
		t.Fatalf("constraint system not satisfied: %v", err)
	}
}

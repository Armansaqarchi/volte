package zkproofs

import (
	"flag"
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/consensys/gnark/std/math/emulated"
	"github.com/consensys/gnark/test"
	"log/slog"
	"math/big"
	"testing"
	"volte/backend/crypto/circuits"
	"volte/backend/crypto/utils"
)

func TestBallotCircuitProof(t *testing.T) {

	G := utils.GenerateBaseECC()
	x := big.NewInt(30)
	Y := utils.G1MulAffine(&G, x)

	flag.Set("Gx", G.X.String())
	flag.Set("Gy", G.Y.String())
	flag.Set("Yx", Y.X.String())
	flag.Set("Yy", Y.Y.String())
	flag.Parse()

	k := big.NewInt(13220)
	m := big.NewInt(99)
	mG := utils.G1MulAffine(&G, m)
	kY := utils.G1MulAffine(&Y, k)
	C1 := utils.G1MulAffine(&G, k)
	C2 := utils.G1AddAffine(&mG, &kY)

	assignment := &circuits.BallotCircuit{
		C1: utils.ECCToAffinePoint(C1),
		C2: utils.ECCToAffinePoint(C2),
		M:  emulated.ValueOf[emulated.BN254Fr](m),
		K:  emulated.ValueOf[emulated.BN254Fr](k),
	}

	cs, err := frontend.Compile(ecc.BN254.ScalarField(), r1cs.NewBuilder, &circuits.BallotCircuit{})
	if err != nil {
		t.Fatalf("compile failed: %v", err)
	}

	fullWitness, err := frontend.NewWitness(assignment, ecc.BN254.ScalarField())
	if err != nil {
		t.Fatalf("new witness failed: %v", err)
	}
	ballotGroth16 := NewBallotGroth16()
	publicWitness, err := fullWitness.Public()
	proof, _ := groth16.Prove(cs, ballotGroth16.GetProvingKey(), fullWitness)
	if err := groth16.Verify(proof, ballotGroth16.GetVerifyingKey(), publicWitness); err != nil {
		slog.Error(err.Error())
		t.Fail()
	}
	if err := cs.IsSolved(fullWitness); err != nil {
		panic(err)
	}
}

func TestNullifierProof(t *testing.T) {

	assert := test.NewAssert(t)

	assignment := circuits.NullifierCircuit{
		SecretKey: "5483851157728431092195247496931641754059230332466588572732387297395134828078",
		EventID:   "3523196653250260958887739657950671762678466692388251624290163732010351636053",
		Nullifier: "7193063174895994881691285216265347816684682062856292682827508574744185066741",
	}

	assert.SolvingSucceeded(&circuits.NullifierCircuit{}, &assignment, test.WithCurves(ecc.BLS12_377))
	cs, err := frontend.Compile(ecc.BLS12_377.ScalarField(), r1cs.NewBuilder, &circuits.NullifierCircuit{})
	if err != nil {
		t.Fatalf("compile failed: %v", err)
	}

	fullWitness, err := frontend.NewWitness(&assignment, ecc.BLS12_377.ScalarField())
	if err != nil {
		t.Fatalf("new witness failed: %v", err)
	}
	nullifierGroth16 := NewNullifierGroth16()
	proof, err := groth16.Prove(cs, nullifierGroth16.GetProvingKey(), fullWitness)
	publicWitness, err := fullWitness.Public()
	if err != nil {
		slog.Error(err.Error())
		t.Fail()
	}

	if err := groth16.Verify(proof, nullifierGroth16.GetVerifyingKey(), publicWitness); err != nil {
		slog.Error(err.Error())
		t.Fail()
	}
}

func TestMerklePathProof(t *testing.T) {
	assert := test.NewAssert(t)

	assignment := circuits.MerkleCircuit{
		MerkleRoot:    "7356920758325201059908902373941132312982459912789273619201826660050410247863",
		LeafValue:     "4137760094704180852789719500758563423980885922685717827383305955441808899436",
		PathPositions: []frontend.Variable{0, 0, 1, 0, 1, 1, 1, 0},
		MerklePath: []frontend.Variable{
			"2201405337893316815918273758681089615279726740723671154548377099132895554651",
			"424860605035859383704348310352427067128543432054056537987894636310336670904",
			"1733733094908890202308645489071232861946837868060311922544482735299421656899",
			"4738184262270190554129232471504937848386298074999300056607424141837299375993",
			"2891247542101801713889729682482526880106482743800010573169732230883085419928",
			"2234929124811023109737676314622609840962882856595367442516400778278767910233",
			"4814974093518456386896925216260606687605659407293043447225930833688370931203",
			"3310810034820363093334307352403815617830879961547922857199997655893198882998",
		},
	}
	circuit := &circuits.MerkleCircuit{
		MerklePath:    make([]frontend.Variable, 8),
		PathPositions: make([]frontend.Variable, 8),
	}
	assert.SolvingSucceeded(circuit, &assignment, test.WithCurves(ecc.BLS12_377))
	cs, err := frontend.Compile(ecc.BLS12_377.ScalarField(), r1cs.NewBuilder, circuit)
	if err != nil {
		t.Fatalf("compile failed: %v", err)
	}
	fullWitness, err := frontend.NewWitness(&assignment, ecc.BLS12_377.ScalarField())
	if err != nil {
		t.Fatalf("new witness failed: %v", err)
	}
	membershipGroth16 := NewMembershipGroth16(8)
	proof, err := groth16.Prove(cs, membershipGroth16.GetProvingKey(), fullWitness)
	if err != nil {
		t.Fatalf("%v", err)
	}
	publicWitness, err := fullWitness.Public()
	if err != nil {
		t.Fatalf("%v", err)
	}

	if err := groth16.Verify(proof, membershipGroth16.GetVerifyingKey(), publicWitness); err != nil {
		t.Fatalf("%v", err)
	}
}

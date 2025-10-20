package circuits

import (
	"flag"
	"fmt"
	"log/slog"
	"volte/backend/crypto/utils"

	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/emulated/sw_emulated"
	"github.com/consensys/gnark/std/math/emulated"
)

var (
	maxVoteValues = flag.Int("max_vote_values", 100, "Maximum possible values for vote.")
	curveParams   = sw_emulated.GetCurveParams[emulated.BN254Fp]()
	// Gx and Gy are coordinates of generator point G in BN254Fp.
	Gx = flag.String("Gx", "", "Specifies the x-coordinate of point G.")
	Gy = flag.String("Gy", "", "Specifies the Y-coordinate of point G.")
	// Yx and Yy specify coordinates of ecc point Y = x[G] where x is the common secret key in elgamal encryption.
	Yx = flag.String("Yx", "", "Specifies the x-coordinate of point Y.")
	Yy = flag.String("Yy", "", "Specifies the Y-coordinate of point Y.")
)

type BallotCircuitMeta struct {
	G sw_emulated.AffinePoint[emulated.BN254Fp]
	Y sw_emulated.AffinePoint[emulated.BN254Fp]
}

var ballotCircuitMeta *BallotCircuitMeta

func GetBallotCircuitMeta() *BallotCircuitMeta {
	if ballotCircuitMeta != nil {
		return ballotCircuitMeta
	}

	ballotCircuitMeta = &BallotCircuitMeta{
		G: sw_emulated.AffinePoint[emulated.BN254Fp]{
			X: utils.StringToElement[emulated.BN254Fp](*Gx),
			Y: utils.StringToElement[emulated.BN254Fp](*Gy),
		},
		Y: sw_emulated.AffinePoint[emulated.BN254Fp]{
			X: utils.StringToElement[emulated.BN254Fp](*Yx),
			Y: utils.StringToElement[emulated.BN254Fp](*Yy),
		},
	}
	return ballotCircuitMeta
}

type BallotCircuit struct {
	C1 sw_emulated.AffinePoint[emulated.BN254Fp] `gnark:",public"`
	C2 sw_emulated.AffinePoint[emulated.BN254Fp] `gnark:",public"`
	M  emulated.Element[emulated.BN254Fr]
	K  emulated.Element[emulated.BN254Fr]
}

func NewBallotCircuit(
	C1, C2 sw_emulated.AffinePoint[emulated.BN254Fp],
	M, k emulated.Element[emulated.BN254Fr],
) *BallotCircuit {

	return &BallotCircuit{C1: C1, C2: C2, M: M, K: k}
}

// Define creates the circuit corresponding to BallotCircuit for validating the ciphertext!
// For this goal, we need to make sure the following requirements are met:
// 1: ScalarMul(G,k) == C1.
// 2: Add(ScalarMul(Y, k), ScalarMul(M . G)) == C2.
// 3: M(M-1)(M-2)....(M - maxVoteValues) = 0.
func (c *BallotCircuit) Define(api frontend.API) error {
	// Creating an instance of sw_emulator used to perform ecc operations.
	sw, err := sw_emulated.New[emulated.BN254Fp, emulated.BN254Fr](api, curveParams)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to initialize the sw_emulator instance, err = %s", err.Error()))
		panic(err)
	}
	fr, err := emulated.NewField[emulated.BN254Fr](api)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to initialize the fr emulator instance, err = %s", err.Error()))
		panic(err)
	}
	fp, err := emulated.NewField[emulated.BN254Fp](api)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to initialize the fp emulator instance, err = %s", err.Error()))
		panic(err)
	}
	// Use circuit's meta values G and Y for validation.
	meta := GetBallotCircuitMeta()
	c1 := sw.ScalarMul(&meta.G, &c.K)
	fp.AssertIsEqual(&c1.X, &c.C1.X)
	fp.AssertIsEqual(&c1.Y, &c.C1.Y)
	Ky := sw.ScalarMul(&meta.Y, &c.K)
	Mg := sw.ScalarMul(&meta.G, &c.M)

	c2 := sw.Add(Ky, Mg)
	fp.AssertIsEqual(&c2.X, &c.C2.X)
	fp.AssertIsEqual(&c2.Y, &c.C2.Y)

	prod := fr.NewElement(1)
	for i := 0; i < *maxVoteValues; i++ {
		prod = fr.Mul(prod, fr.Sub(&c.M, fr.NewElement(i)))
	}
	fr.AssertIsEqual(prod, fr.NewElement(0))
	return nil
}

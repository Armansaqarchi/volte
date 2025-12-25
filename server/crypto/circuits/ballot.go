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
	maxVoteValues = flag.Int("max_vote_values", 2, "Maximum possible values for vote.")
	curveParams   = sw_emulated.GetCurveParams[emulated.BN254Fp]()
	// Gx and Gy are coordinates of generator point G in BN254Fp.
	Gx = flag.String("Gx", "", "Specifies the x-coordinate of EC point G.")
	Gy = flag.String("Gy", "", "Specifies the Y-coordinate of EC point G.")
	// Yx and Yy specify coordinates of ecc point Y = x[G] where x is the common secret key in elgamal encryption.
	Yx = flag.String("Yx", "", "Specifies the x-coordinate of EC point Y.")
	Yy = flag.String("Yy", "", "Specifies the Y-coordinate of EC point Y.")
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

	meta := GetBallotCircuitMeta()

	// 1) C1 == kG
	c1 := sw.ScalarMul(&meta.G, &c.K)
	fp.AssertIsEqual(&c1.X, &c.C1.X)
	fp.AssertIsEqual(&c1.Y, &c.C1.Y)

	// 2) Ky == kY
	Ky := sw.ScalarMul(&meta.Y, &c.K)

	// 3) Enforce M is boolean: M*(M-1)=0
	one := fr.NewElement(1)
	zero := fr.NewElement(0)
	fr.AssertIsEqual(fr.Mul(&c.M, fr.Sub(&c.M, one)), zero)

	// Build the two valid C2 options without ever computing M*G:
	// if M=0 => C2 = Ky
	// if M=1 => C2 = Ky + G
	c2IfOne := sw.Add(Ky, &meta.G)

	// Select based on the LSB of M (safe because M is constrained boolean above)
	// Note: ToBits returns bits as frontend.Variables (native field booleans)
	mBits := fr.ToBits(&c.M)
	mBit := mBits[0]

	expX := fp.Select(mBit, &c2IfOne.X, &Ky.X) // if mBit==1 pick (Ky+G).X else Ky.X
	expY := fp.Select(mBit, &c2IfOne.Y, &Ky.Y) // if mBit==1 pick (Ky+G).Y else Ky.Y

	fp.AssertIsEqual(expX, &c.C2.X)
	fp.AssertIsEqual(expY, &c.C2.Y)

	return nil
}

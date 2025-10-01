package circuits

import (
	"flag"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/emulated/sw_emulated"
	"github.com/consensys/gnark/std/math/emulated"
	"log/slog"
	"volte/backend/crypto/utils"
)

var (
	maxVoteValues = flag.Int("max_vote_values", 100, "Maximum possible values for vote.")
	curveParams   = sw_emulated.GetCurveParams[emulated.BLS12377Fp]()
	// Gx and Gy are coordinates of generator point G in BLS12377
	Gx = flag.String("Yx", "", "Specifies the x-coordinate of point G.")
	Gy = flag.String("Yy", "", "Specifies the y-coordinate of point G.")
	// Yx and Yy specify coordinates of ecc point Y = x[G] where x is the common secret key in elgamal encryption.
	Yx = flag.String("Yx", "", "Specifies the x-coordinate of point Y.")
	Yy = flag.String("Yy", "", "Specifies the y-coordinate of point Y.")
)

type BallotCircuitMeta struct {
	g *sw_emulated.AffinePoint[emulated.BLS12377Fp]
	y *sw_emulated.AffinePoint[emulated.BLS12377Fp]
}

var ballotCircuitMeta *BallotCircuitMeta

func GetBallotCircuitMeta() *BallotCircuitMeta {
	if ballotCircuitMeta != nil {
		return ballotCircuitMeta
	}
	ballotCircuitMeta = &BallotCircuitMeta{
		g: &sw_emulated.AffinePoint[emulated.BLS12377Fp]{
			X: utils.StringToBLS12377Element(*Gx),
			Y: utils.StringToBLS12377Element(*Gy),
		},
		y: &sw_emulated.AffinePoint[emulated.BLS12377Fp]{
			X: utils.StringToBLS12377Element(*Yx),
			Y: utils.StringToBLS12377Element(*Yy),
		},
	}
	return ballotCircuitMeta
}

type BallotCircuit struct {
	C1 *sw_emulated.AffinePoint[emulated.BLS12377Fp] `gnark:"public,"`
	C2 *sw_emulated.AffinePoint[emulated.BLS12377Fp] `gnark:"public,"`
	M  *emulated.Element[emulated.BLS12377Fp]
	k  *emulated.Element[emulated.BLS12377Fp]
}

func NewBallotCircuit(
	C1, C2 *sw_emulated.AffinePoint[emulated.BLS12377Fp],
	M, k *emulated.Element[emulated.BLS12377Fp],
) *BallotCircuit {

	return &BallotCircuit{C1: C1, C2: C2, M: M, k: k}
}

// Define creates the circuit corresponding to BallotCircuit for validating the ciphertext!
// For this goal, we need to make sure the following requirements meets:
// 1: ScalarMul(G,k) == C1.
// 2: Add(ScalarMul(Y, k), ScalarMul(M . G)) == C2.
// 3: M(M-1)(M-2)....(M - maxVoteValues) = 0.
func (c *BallotCircuit) Define(api frontend.API) error {
	// Creating an instance of sw_emulator used to perform ecc operations.
	sw, err := sw_emulated.New[emulated.BLS12377Fp, emulated.BLS12377Fp](api, curveParams)
	if err != nil {
		slog.Error("Failed to initialize the sw_emulator instance, err = %s", err.Error())
		panic(err)
	}
	// Use circuit's meta values G and Y for validation.
	meta := GetBallotCircuitMeta()
	api.AssertIsEqual(sw.ScalarMul(meta.g, c.k).Y, c.C1)
	api.AssertIsEqual(sw.Add(sw.ScalarMul(meta.y, c.k), sw.ScalarMul(meta.g, c.M)).Y, c.C2)
	em, err := emulated.NewField[emulated.BLS12377Fp](api)
	if err != nil {
		slog.Error("Failed to initialize the emulator instance, err = %s", err.Error())
		panic(err)
	}
	prod := em.NewElement(1)
	for i := 0; i < *maxVoteValues; i++ {
		prod = em.Mul(prod, em.Sub(c.M, em.NewElement(i)))
	}
	api.AssertIsEqual(prod, em.NewElement(0))
	return nil
}

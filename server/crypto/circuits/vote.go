package circuits

import (
	"flag"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/emulated/sw_emulated"
	"github.com/consensys/gnark/std/math/emulated"
	"log/slog"
)

var (
	maxVoteValues = flag.Int("max_vote_values", 100, "Maximum possible values for vote.")
	curveParams   = sw_emulated.GetCurveParams[emulated.BLS12377Fp]()
)

type BallotCircuit struct {
	C1 *sw_emulated.AffinePoint[emulated.BLS12377Fp] `gnark:"public,"`
	C2 *sw_emulated.AffinePoint[emulated.BLS12377Fp] `gnark:"public,"`
	G  *sw_emulated.AffinePoint[emulated.BLS12377Fp] `gnark:"public,"`
	Y  *sw_emulated.AffinePoint[emulated.BLS12377Fp] `gnark:"public,"`
	M  *emulated.Element[emulated.BLS12377Fp]
	k  *emulated.Element[emulated.BLS12377Fp]
}

// Define creates the circuit corresponding to BallotCircuit for validating the ciphertext!
// For this goal, we need to make sure the following requirements meets:
// 1: ScalarMul(G,k) == C1.
// 2: Add(ScalarMul(Y, k), ScalarMul(M . G)) == C2.
// 3: M(M-1)(M-2)....(M-100) = 0.
func (c *BallotCircuit) Define(api frontend.API) error {
	// Creating an instance of sw_emulator used to perform ecc operations.
	sw, err := sw_emulated.New[emulated.BLS12377Fp, emulated.BLS12377Fp](api, curveParams)
	if err != nil {
		slog.Error("Failed to initialize the sw_emulator instance, err = %s", err.Error())
		panic(err)
	}
	api.AssertIsEqual(sw.ScalarMul(c.G, c.k).Y, c.C1)
	api.AssertIsEqual(sw.Add(sw.ScalarMul(c.Y, c.k), sw.ScalarMul(c.G, c.M)).Y, c.C2)
	em, err := emulated.NewField[emulated.BLS12377Fp](api)
	if err != nil {
		slog.Error("Failed to initialize the emulator instance, err = %s", err.Error())
		panic(err)
	}
	em.AssertIsInRange(c.M)
	return nil
}

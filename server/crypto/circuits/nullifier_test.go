package circuits

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/test"
	"testing"
)

func TestNullifier(t *testing.T) {

	assert := test.NewAssert(t)

	circuit := NullifierCircuit{
		SecretKey: "5483851157728431092195247496931641754059230332466588572732387297395134828078",
		EventID:   "3523196653250260958887739657950671762678466692388251624290163732010351636053",
		Nullifier: "7193063174895994881691285216265347816684682062856292682827508574744185066741",
	}

	assert.SolvingSucceeded(&circuit, &circuit, test.WithCurves(ecc.BLS12_377))
}

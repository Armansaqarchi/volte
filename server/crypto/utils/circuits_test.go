package utils

import (
	"fmt"
	"github.com/iden3/go-iden3-crypto/v2/babyjub"
	"math"
	"math/big"
	"testing"
)

func TestFindM(t *testing.T) {
	C1 := babyjub.NewPoint()
	C2 := babyjub.NewPoint()
	C1.X = NewBigIntFromString("14143394589216999491167590241056903469890491202580299340012053369632212315960")
	C1.Y = NewBigIntFromString("11857660852706073722557758331793886782915607013317321971646634272150373979221")

	C2.X = NewBigIntFromString("10666958084136858807587443784291324259844140813722403635252333953295545417950")
	C2.Y = NewBigIntFromString("3446437679768586016390761797560054434279946790686033375868979553073753644823")

	m, err := DecryptM_BSGS(C1, C2, big.NewInt(20), uint64(math.Pow(2, 32)))
	if err != nil {
		t.Error(err)
	}

	fmt.Println(m)
}

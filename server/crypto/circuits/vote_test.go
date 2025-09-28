package circuits

import (
	"math/big"
	"testing"

	"github.com/consensys/gnark/test"
)

func TestVote(t *testing.T) {
	assert := test.NewAssert(t)
	circuit := BallotCircuit{
		Vote: big.NewInt(40),
	}
	assert.SolvingSucceeded(new(BallotCircuit), &circuit)
}

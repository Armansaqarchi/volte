package proof

import (
	"crypto/rand"
	"flag"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark/std/math/emulated"
	"log/slog"
	"math/big"
	"strconv"
	"volte/backend/chain/contracts"
	"volte/backend/cmd/proof/keys"
	"volte/backend/crypto/circuits"
	"volte/backend/crypto/utils"
	"volte/backend/crypto/zkproofs"
)

var (
	Msg = flag.String("msg", "", "the message one hot encoded value to store")
)

func RunBallotProof() *contracts.VolteContractBallotProof {

	slog.Info("Generating random proof for Ballot verification.")

	// Make sure G and Y correspond otherwise verification would fail.
	if *circuits.Gx == "" || *circuits.Gy == "" || *circuits.Yx == "" || *circuits.Yy == "" {
		panic(fmt.Errorf("g and y affine points are not defined in flags"))
	}
	G, err := utils.MakeG1Affine(*circuits.Gx, *circuits.Gy)
	if err != nil {
		panic(err)
	}
	Y, err := utils.MakeG1Affine(*circuits.Yx, *circuits.Yy)
	if err != nil {
		panic(err)
	}
	k, err := rand.Int(rand.Reader, emulated.BN254Fr{}.Modulus())
	if err != nil {
	}
	slog.Info(fmt.Sprintf("Setting msg to %s", *Msg))
	m, err := strconv.Atoi(*Msg)
	if err != nil {
		panic(fmt.Sprintf("failed to convert the msg string to integer, err : %s", err))
	}
	if m != 0 && m != 1 {
		panic(fmt.Sprintf("Message must be boolean, got invalid value : %d", m))
	}
	C1 := utils.G1MulAffine(&G, k)
	// kY = k*Y
	kY := utils.G1MulAffine(&Y, k)
	// C2 = M + kY   (point addition)
	var C2 bn254.G1Affine
	if m == 0 {
		C2 = kY
	} else {
		C2 = utils.G1AddAffine(&G, &kY)
	}

	assignment := circuits.BallotCircuit{
		C1: utils.ECCToAffinePoint(C1),
		C2: utils.ECCToAffinePoint(C2),
		M:  emulated.ValueOf[emulated.BN254Fr](m),
		K:  emulated.ValueOf[emulated.BN254Fr](k),
	}
	vkFile, pkFile := keys.GetBallotKeys()
	slog.Info("Successfully read ballot provingKey and verifyingKey from files.")
	g16 := zkproofs.NewBallotGroth16FromExistingKeys(vkFile, pkFile)
	slog.Info("Created ballot groth16.")
	proofParts, publicWitness := zkproofs.CreateProof(&assignment, g16)
	var publicInputs [16]*big.Int
	for i, input := range publicWitness.(fr.Vector) {
		publicInput, _ := big.NewInt(0).SetString(input.String(), 10)
		publicInputs[i] = publicInput
	}
	return &contracts.VolteContractBallotProof{
		Proof: contracts.VolteContractProof{
			Arx:  proofParts[0],
			Ary:  proofParts[1],
			Brx0: proofParts[3],
			Brx1: proofParts[2],
			Bry0: proofParts[5],
			Bry1: proofParts[4],
			Cx:   proofParts[6],
			Cy:   proofParts[7],
		},
	}
}

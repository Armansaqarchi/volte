package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"github.com/consensys/gnark/std/math/emulated"
	"log/slog"
	"math/big"
	"volte/backend/cmd/proof/keys"
	"volte/backend/crypto/circuits"
	"volte/backend/crypto/utils"
	"volte/backend/crypto/zkproofs"
)

var (
	Msg = flag.String("m", "", "the message one hot encoded value to store")
)

func runRandomBallotProof() {

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
	flag.Parse()
	k, err := rand.Int(rand.Reader, emulated.BN254Fr{}.Modulus())
	if err != nil {
	}
	m, ok := big.NewInt(0).SetString(*Msg, 10)
	if !ok {
		panic(fmt.Errorf("failed to set big.Int to message"))
	}

	M := utils.G1MulAffine(&G, m)
	// C1 = k*G
	C1 := utils.G1MulAffine(&G, k)
	// kY = k*Y
	kY := utils.G1MulAffine(&Y, k)
	// C2 = M + kY   (point addition)
	C2 := utils.G1AddAffine(&M, &kY)

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
	zkproofs.CreateProof(&assignment, g16)
}

func main() {
	flag.Parse()
	runRandomBallotProof()
}

package main

import (
	"flag"
	"github.com/consensys/gnark/frontend"
	"log/slog"
	"strings"
	"volte/backend/cmd/proof/keys"
	"volte/backend/crypto/circuits"
	"volte/backend/crypto/utils"
	"volte/backend/crypto/zkproofs"
)

func runRandomMerklePathProof() {
	slog.Info("Generating random proof for membership verification.")

	var (
		root          = flag.String("membership_merkle_root", "", "membership merkle tree root")
		secretKey     = flag.String("membership_secret_key", "", "membership secret key")
		pathPositions = flag.String("membership_path_positions", "", "membership path positions")
		merklePath    = flag.String("membership merkle path", "", "membership flag path")
	)

	var MerklePath []frontend.Variable
	for _, path := range strings.Split(*merklePath, ",") {
		MerklePath = append(MerklePath, path)
	}

	var PathPositions []frontend.Variable
	for _, path := range strings.Split(*pathPositions, ",") {
		PathPositions = append(PathPositions, path)
	}

	leafValue, err := utils.MimcHash(*secretKey)
	if err != nil {
		panic(err)
	}
	assignment := circuits.MerkleCircuit{
		MerkleRoot:    *root,
		SecretKey:     *secretKey,
		LeafValue:     leafValue,
		PathPositions: PathPositions,
		MerklePath:    MerklePath,
	}

	vkFile, pkFile := keys.GetMembershipKeys()
	slog.Info("Successfully read membership provingKey and verifyingKey from files.")
	g16 := zkproofs.NewMembershipGroth16FromExistingKeys(8, vkFile, pkFile)
	slog.Info("Created membership groth16.")
	zkproofs.CreateProof(&assignment, g16)
}

func main() {
	flag.Parse()
	runRandomMerklePathProof()
}

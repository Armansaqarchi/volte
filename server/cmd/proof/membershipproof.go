package proof

import (
	"flag"
	"fmt"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/consensys/gnark/frontend"
	"log/slog"
	"math/big"
	"strings"
	"volte/backend/chain/contracts"
	"volte/backend/cmd/proof/keys"
	"volte/backend/crypto/circuits"
	"volte/backend/crypto/utils"
	"volte/backend/crypto/zkproofs"
)

var (
	root          = flag.String("membership_merkle_root", "", "membership merkle tree root")
	pathPositions = flag.String("membership_path_positions", "", "membership path positions")
	merklePath    = flag.String("membership_merkle_path", "", "membership flag path")
	secretKey     = flag.String("user_secret_key", "", "membership secret key")
)

func RunMerklePathProof() *contracts.VolteContractMembershipProof {
	slog.Info("Generating random proof for membership verification.")

	var MerklePath []frontend.Variable
	for _, path := range strings.Split(*merklePath, ",") {
		MerklePath = append(MerklePath, path)
	}

	var PathPositions []frontend.Variable
	for _, path := range strings.Split(*pathPositions, ",") {
		PathPositions = append(PathPositions, path)
	}

	leafValue, err := utils.MimcHash([]byte(*secretKey))
	fmt.Println("root : ", *root)
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
	proofParts, publicWitness := zkproofs.CreateProof(&assignment, g16)
	var publicInputs [2]*big.Int
	for i, input := range publicWitness.(fr.Vector) {
		publicInput, _ := big.NewInt(0).SetString(input.String(), 10)
		publicInputs[i] = publicInput
	}
	return &contracts.VolteContractMembershipProof{
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
		Input: publicInputs,
	}
}

package proof

import (
	"flag"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"log/slog"
	"math/big"
	"volte/backend/chain/contracts"
	"volte/backend/cmd/proof/keys"
	"volte/backend/crypto/circuits"
	"volte/backend/crypto/utils"
	"volte/backend/crypto/zkproofs"
)

var (
	eventId = flag.String("event_id", "", "event id")
)

func RunNullifierProof() *contracts.VolteContractNullifierProof {

	nullifier, err := utils.MimcHash([]byte(*secretKey), []byte(*eventId))
	if err != nil {
		panic(err)
	}
	commitment, err := utils.MimcHash([]byte(*secretKey))
	if err != nil {
		panic(err)
	}

	assignment := circuits.NullifierCircuit{
		Commitment: commitment,
		SecretKey:  *secretKey,
		EventID:    *eventId,
		Nullifier:  nullifier,
	}

	vkFile, pkFile := keys.GetNullifierKeys()
	slog.Info("Successfully read nullifier provingKey and verifyingKey from files.")
	g16 := zkproofs.NewNullifierGroth16FromExistingKeys(vkFile, pkFile)
	slog.Info("Created nullifier groth16.")
	proofParts, publicWitness := zkproofs.CreateProof(&assignment, g16)
	var publicInputs [3]*big.Int
	for i, input := range publicWitness.(fr.Vector) {
		publicInput, _ := big.NewInt(0).SetString(input.String(), 10)
		publicInputs[i] = publicInput
	}
	return &contracts.VolteContractNullifierProof{
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

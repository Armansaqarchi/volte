package main

import (
	"flag"
	"log/slog"
	"volte/backend/cmd/proof/keys"
	"volte/backend/crypto/circuits"
	"volte/backend/crypto/utils"
	"volte/backend/crypto/zkproofs"
)

var (
	secretKey = flag.String("user_secret_key", "", "user secret key")
	eventId   = flag.String("event_id", "", "event id")
)

func runRandomNullifierProof() {
	nullifier, err := utils.MimcHash(*secretKey, *eventId)
	if err != nil {
		panic(err)
	}
	commitment, err := utils.MimcHash(*secretKey)
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
	zkproofs.CreateProof(&assignment, g16)
}

func main() {
	flag.Parse()
	runRandomNullifierProof()
}

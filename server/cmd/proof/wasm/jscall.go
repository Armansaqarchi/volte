//go:build js

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"syscall/js"
	"volte/backend/chain/contracts"
	"volte/backend/cmd/proof"
	"volte/backend/crypto/utils"
)

func runProof(this js.Value, args []js.Value) any {
	flag.Parse()
	jsonData, err := json.Marshal(contracts.VolteContractProofs{
		Ballot:     *proof.RunBallotProof(),
		Membership: *proof.RunMerklePathProof(),
		Nullifier:  *proof.RunNullifierProof(),
	})
	if err != nil {
		panic(err)
	}
	return string(jsonData)
}

func getMIMCHash(this js.Value, args []js.Value) any {
	secretKey := flag.String("secret", "", "secret key")
	flag.Parse()
	hash, err := utils.MimcHash([]byte(*secretKey))
	if err != nil {
		panic(fmt.Sprintf("Failed to calculate hash, err : %s", err.Error()))
	}
	return hash
}

func main() {
	js.Global().Set("runProof", js.FuncOf(runProof))
	js.Global().Set("getMIMCHash", js.FuncOf(getMIMCHash))
	// keep runtime alive
	select {}
}

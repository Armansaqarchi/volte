package cmd

import (
	"crypto/sha256"
	"flag"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/backend/solidity"
	"github.com/spf13/cobra"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"volte/backend/crypto/utils"
	"volte/backend/crypto/zkproofs"
)

var exportSolidityCmd = &cobra.Command{
	Use:   "export-solidity",
	Short: "Generates solidity contract from the verifying key of circuits.",
	Run:   runExportSolidity,
}

func init() {
	rootCmd.AddCommand(exportSolidityCmd)
}

func runExportSolidity(_ *cobra.Command, _ []string) {

	G := utils.GenerateBaseECC()
	x := big.NewInt(30)
	Y := utils.G1MulAffine(&G, x)

	flag.Set("Gx", G.X.String())
	flag.Set("Gy", G.Y.String())
	flag.Set("Yx", Y.X.String())
	flag.Set("Yy", Y.Y.String())

	flag.Parse()

	vkFile, pkFile := utils.GetCircuitKeys("../keys/groth16/ballot")
	ballotG16 := zkproofs.NewBallotGroth16FromExistingKeys(vkFile, pkFile)
	vkFile, pkFile = utils.GetCircuitKeys("../keys/groth16/membership")
	membershipG16 := zkproofs.NewMembershipGroth16FromExistingKeys(8, vkFile, pkFile)
	vkFile, pkFile = utils.GetCircuitKeys("../keys/groth16/nullifier")
	nullifierG16 := zkproofs.NewNullifierGroth16FromExistingKeys(vkFile, pkFile)

	exportSolidityFile("../contracts/groth16/ballot.sol", ballotG16.GetVerifyingKey())
	exportSolidityFile("../contracts/groth16/nullifier.sol", nullifierG16.GetVerifyingKey())
	exportSolidityFile("../contracts/groth16/membership.sol", membershipG16.GetVerifyingKey())

}

func exportSolidityFile(path string, key groth16.VerifyingKey) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("failed to create directories for %s: %v", path, err)
	}
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	if err := key.ExportSolidity(file, func(config *solidity.ExportConfig) error {
		config.HashToFieldFn = sha256.New()
		return nil
	}); err != nil {
		panic(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

package cmd

import (
	"github.com/consensys/gnark/backend/groth16"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
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
	ballotG16 := zkproofs.NewBallotGroth16()
	//nullifierG16 := zkproofs.NewNullifierGroth16()
	//membershipG16 := zkproofs.NewMembershipGroth16(32)

	exportSolidityFile("../contracts/groth16/ballot.sol", ballotG16.GetVerifyingKey())
	//exportSolidityFile("../contracts/groth16/nullifier.sol", nullifierG16.GetVerifyingKey())
	//exportSolidityFile("../contracts/groth16/membership.sol", membershipG16.GetVerifyingKey())
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
	if err := key.ExportSolidity(file); err != nil {
		panic(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

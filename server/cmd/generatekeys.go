package cmd

import (
	"flag"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"io"
	"log"
	"os"
	"path/filepath"
	"volte/backend/crypto/zkproofs"
)

var generateKeysCmd = &cobra.Command{
	Use:   "generate-keys",
	Short: "Generate verifying and proving keys.",
	Run:   runGenerateKeys,
}

func init() {
	rootCmd.AddCommand(generateKeysCmd)
}

func runGenerateKeys(_ *cobra.Command, _ []string) {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Parse()
	ballotG16 := zkproofs.NewBallotGroth16()
	nullifierG16 := zkproofs.NewNullifierGroth16()
	membershipG16 := zkproofs.NewMembershipGroth16(32)
	createFileAndWrite("../keys/groth16/ballot/verifyingKey", ballotG16.GetVerifyingKey())
	createFileAndWrite("../keys/groth16/ballot/provingKey", ballotG16.GetProvingKey())
	createFileAndWrite("../keys/groth16/nullifier/verifyingKey", nullifierG16.GetVerifyingKey())
	createFileAndWrite("../keys/groth16/nullifier/provingKey", nullifierG16.GetProvingKey())
	createFileAndWrite("../keys/groth16/membership/verifyingKey", membershipG16.GetVerifyingKey())
	createFileAndWrite("../keys/groth16/membership/provingKey", membershipG16.GetProvingKey())
}

func createFileAndWrite(path string, content io.WriterTo) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("failed to create directories for %s: %v", path, err)
	}
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := content.WriteTo(file); err != nil {
		log.Fatal(err)
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}

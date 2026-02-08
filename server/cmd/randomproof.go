package cmd

import (
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"log/slog"
	"time"
	"volte/backend/cmd/proof"
)

var randomCmd = &cobra.Command{
	Use:   "random-proof",
	Short: "Generate verifying and proving keys.",
	Run:   runRandomProof,
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

func runRandomProof(_ *cobra.Command, _ []string) {
	flag.Parse()
	slog.Info("Start generating proof.")
	start := time.Now()
	proof.RunBallotProof()
	proof.RunMerklePathProof()
	proof.RunNullifierProof()
	duration := time.Since(start)
	slog.Info("End generating proof.")
	fmt.Println("Proof generation total time : ", duration)
}

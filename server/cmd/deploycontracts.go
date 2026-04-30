package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log/slog"
	"volte/backend/chain"
	"volte/backend/chain/contracts"
	"volte/backend/chain/contracts/ballot"
	"volte/backend/chain/contracts/membership"
	"volte/backend/chain/contracts/nullifier"
)

var deployContractCmd = &cobra.Command{
	Use:   "deploy-contract",
	Short: "Deploy contract",
	Run:   runDeployContract,
}

func init() {
	rootCmd.AddCommand(deployContractCmd)
}

func runDeployContract(_ *cobra.Command, _ []string) {
	connectionHandler := chain.NewConnectionHandler()
	slog.Info(fmt.Sprintf("Successfully connected to chain using address : %s", connectionHandler.FromAddress))
	// after connecting to the server, deploy the contracts.

	nullifierAddr, _, _, err := nullifier.DeployVolte(connectionHandler.TransactionOpts, connectionHandler.Client)
	if err != nil {
		fmt.Println("Err : ", err.Error())
		panic(err)
	}
	membershipAddr, _, _, err := membership.DeployVolte(connectionHandler.TransactionOpts, connectionHandler.Client)
	if err != nil {
		fmt.Println("Err : ", err)
		panic(err)
	}
	ballotAddr, _, _, err := ballot.DeployVolte(connectionHandler.TransactionOpts, connectionHandler.Client)
	if err != nil {
		panic(err)
	}
	volteAddr, txn, _, err := contracts.DeployVolte(
		connectionHandler.TransactionOpts,
		connectionHandler.Client,
		ballotAddr,
		membershipAddr,
		nullifierAddr,
	)
	if err != nil {
		panic(err)
	}
	slog.Info(fmt.Sprintf("volte address : %s", volteAddr.Hex()))
	slog.Info(fmt.Sprintf("Transaction : %s", txn.Hash().Hex()))
}

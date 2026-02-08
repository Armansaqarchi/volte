package cmd

import (
	"flag"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate verifying and proving keys.",
}

func Execute() {
	arg := os.Args[1]
	for _, command := range rootCmd.Commands() {
		if arg == command.Use {
			if err := rootCmd.Execute(); err != nil {
				os.Exit(1)
			} else {
				os.Exit(0)
			}
		}
	}
}

func init() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Parse()
}

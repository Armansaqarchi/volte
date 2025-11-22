package main

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

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Parse()
}

func main() {
	Execute()
}

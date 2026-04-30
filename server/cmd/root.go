package cmd

import (
	"flag"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate verifying and proving keys.",
}

func Execute() {
	panicOnErr(rootCmd.Execute())
}

func init() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Parse()
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

package cmd

import (
	"flag"

	"github.com/spf13/cobra"
)

var cmdPattern *cobra.Command

func runPattern(cmd *cobra.Command, args []string) error {

	patName := flag.Arg(1)
	files := flag.Arg(2)

	gclib.grepPattern(patName, files)
	return nil
}

func init() {
	cmdPattern = &cobra.Command{
		Use:   "pattern",
		Short: "Grep for a pattern type",
		RunE:  runPattern,
	}

	rootCmd.AddCommand(cmdPattern)
}

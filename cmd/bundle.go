package cmd

import (
	"flag"

	"github.com/spf13/cobra"
)

var cmdBundle *cobra.Command

func runBundle(cmd *cobra.Command, args []string) error {
	// flag declaration

	bundleName := flag.Arg(1)
	files := flag.Arg(2)

	gclib.grepBundle(bundleName, files)
	return nil
}

// nolint:gochecknoinits
func init() {
	cmdBundle = &cobra.Command{
		Use:   "bundle",
		Short: "Grep for multiple patterns type",
		RunE:  runBundle,
	}

	rootCmd.AddCommand(cmdBundle)
}

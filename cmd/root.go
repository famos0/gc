package cmd

import (
	"github.com/famos0/gc/gclib"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "gc",
	SilenceUsage: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func getOptions() (*gclib.Options, error) {
	opts := gclib.NewOptions()

	opts.Quiet, _ = rootCmd.Flags().GetBool("quiet")
	opts.Testless, _ = rootCmd.Flags().GetBool("testless")
	opts.Stdin, _ = rootCmd.Flags().GetBool("stdin")
	return opts, nil
}

func init() {
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "Don't print Patterns and Bundles name and comment")
	rootCmd.PersistentFlags().BoolP("testless", "t", false, "Don't grep test/mock code")
	rootCmd.PersistentFlags().BoolP("stdin", "s", false, "Display generated grep command instead of running them")
}

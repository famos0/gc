package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/famos0/gc/gclib"
	"github.com/spf13/cobra"
)

var cmdBundle *cobra.Command

func runBundle(cmd *cobra.Command, args []string) error {

	listmode, _ := cmdBundle.Flags().GetBool("list")
	if listmode {
		bundlesdir, _ := gclib.GetBundleDir()

		filepath.Walk(bundlesdir,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if file, err := os.Stat(path); file.Mode().IsRegular() && err == nil {
					fmt.Println(strings.Replace(path[:len(path)-5], bundlesdir+"/", "", 1))
				}
				return nil
			})
	} else {
		if len(args) < 2 {
			fmt.Println("Usage : gc bundle [OPTIONS] [PATTERN] [TARGET]")
			fmt.Println("To list options, use gc bundle -h")
		} else {
			options, _ := getOptions()
			bundleName := args[0]
			files := args[1]

			gclib.GrepBundle(bundleName, files, options)
		}
	}
	return nil
}

// nolint:gochecknoinits
func init() {
	cmdBundle = &cobra.Command{
		Use:   "bundle",
		Short: "Grep for multiple patterns type",
		RunE:  runBundle,
	}
	cmdBundle.Flags().BoolP("list", "l", false, "Display available bundles")
	rootCmd.AddCommand(cmdBundle)
}

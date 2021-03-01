package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/famos0/gc/gclib"
	"github.com/spf13/cobra"
)

var cmdPattern *cobra.Command

func runPattern(cmd *cobra.Command, args []string) error {

	listmode, _ := cmdPattern.Flags().GetBool("list")
	if listmode {
		patternsdir, _ := gclib.GetPatternDir()
		filepath.Walk(patternsdir,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if file, err := os.Stat(path); file.Mode().IsRegular() && err == nil {
					fmt.Println(strings.Replace(path[:len(path)-5], patternsdir+"/", "", 1))
				}
				return nil
			})
	} else {
		if len(args) < 2 {
			fmt.Println("Usage : gc pattern [OPTIONS] [PATTERN] [TARGET]")
			fmt.Println("To list options, use gc pattern -h")
		} else {
			options, _ := getOptions()
			patName := args[0]
			files := args[1]

			gclib.GrepPattern(patName, files, options)
		}
	}
	return nil
}

func init() {
	cmdPattern = &cobra.Command{
		Use:   "pattern",
		Short: "Grep for a pattern type",
		RunE:  runPattern,
	}
	cmdPattern.Flags().BoolP("list", "l", false, "Display available patterns")
	rootCmd.AddCommand(cmdPattern)
}

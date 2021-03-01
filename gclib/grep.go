package gclib

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

type pattern struct {
	Flags    string   `json:"flags,omitempty"`
	Pattern  string   `json:"pattern,omitempty"`
	Patterns []string `json:"patterns,omitempty"`
	Comment  string   `json:"comment,omitempty"`
}

type bundle struct {
	PatternsPath []string `json:"patternspath,omitempty"`
	Bundles      []string `json:"bundles,omitempty"`
	Comment      string   `json:"comment,omitempty"`
}

// GrepPattern throw a gc pattern
func GrepPattern(patName, files string, options *Options) {

	if files == "" {
		files = "."
	}

	patDir, err := GetPatternDir()
	if err != nil {
		if !options.Quiet {
			fmt.Fprintln(os.Stderr, "unable to open user's pattern directory")
		}

		return
	}

	filename := filepath.Join(patDir, patName+".json")
	f, err := os.Open(filename)
	if err != nil {
		if !options.Quiet {
			fmt.Fprintln(os.Stderr, "no such pattern")
		}
		return
	}
	defer f.Close()

	pat := pattern{}
	dec := json.NewDecoder(f)
	dec.Decode(&pat)

	if !options.Quiet {
		printComment("pattern", patName, pat.Comment)
	}

	if pat.Pattern == "" {

		if len(pat.Patterns) == 0 {
			if !options.Quiet {
				fmt.Fprintf(os.Stderr, "pattern file '%s' contains no pattern(s)\n", filename)
			}
			return
		}

		pat.Pattern = "(" + strings.Join(pat.Patterns, "|") + ")"
	}

	if options.Testless {

		var c1 *exec.Cmd
		var c2 *exec.Cmd

		c1 = exec.Command("grep", pat.Flags, pat.Pattern, files)
		c2 = exec.Command("grep", "-vi", "(test\\|mock)")

		if options.Stdin {
			fmt.Printf("%s | %s\n", c1.String(), c2.String())
		} else {
			c1.Stdin = os.Stdin
			c2.Stdin, _ = c1.StdoutPipe()
			c2.Stdout = os.Stdout
			c2.Stderr = os.Stderr

			c2.Start()
			c1.Run()
			c2.Wait()
		}
	} else {
		var cmd *exec.Cmd
		cmd = exec.Command("grep", pat.Flags, pat.Pattern, files)

		if options.Stdin {
			fmt.Println(cmd.String())
		} else {
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		}
	}
}

// GrepBundle throw a gc Bundle
func GrepBundle(bundleName, files string, options *Options) {

	bundleDir, err := GetBundleDir()
	if err != nil {
		if !options.Quiet {
			fmt.Fprintln(os.Stderr, "unable to open user's pattern directory")
		}
		return
	}

	filename := filepath.Join(bundleDir, bundleName+".json")
	f, err := os.Open(filename)
	if err != nil {
		if !options.Quiet {
			fmt.Fprintln(os.Stderr, "no such pattern")
		}
		return
	}
	defer f.Close()

	bundle := bundle{}
	dec := json.NewDecoder(f)
	dec.Decode(&bundle)

	if !options.Quiet {
		printComment("bundle", bundleName, bundle.Comment)
	}

	if len(bundle.Bundles) > 0 {
		for _, bundle := range bundle.Bundles {
			GrepBundle(bundle, files, options)
		}
	}

	for _, pattern := range bundle.PatternsPath {
		GrepPattern(pattern, files, options)
	}
}

// GetPatternDir return pattern directory
func GetPatternDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	path := filepath.Join(usr.HomeDir, ".config/gc")
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		// .config/gc exists
		return path, nil
	}
	return "templates/patterns", nil
	//return filepath.Join(usr.HomeDir, ".gc"), nil
}

// GetBundleDir return bundle directory
func GetBundleDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	path := filepath.Join(usr.HomeDir, ".config/gc")
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		// .config/gc exists
		return path, nil
	}
	return "templates/bundles", nil
	//return filepath.Join(usr.HomeDir, ".gc"), nil
}

func printComment(filetype, name, comment string) {
	fmt.Printf("\n%s: %s\n", strings.Title(filetype), name)
	fmt.Printf("%s\n\n", comment)
}

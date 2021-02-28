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
	Flags    string   `json:"flags:omitempty"`
	Pattern  string   `json:"pattern:omitempty"`
	Patterns []string `json:"patterns:omitempty"`
	Comment  string   `json:"comments:omitempty"`
}

type bundle struct {
	PatternsPath []string `json:"patternspath:omitempty"`
	Bundles      []string `json:"bundles:omitempty`
	Comment      string   `json:"comments:omitempty"`
}

func grepPattern(patName, files string) {

	if files == "" {
		files = "."
	}

	patDir, err := getPatternDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to open user's pattern directory")
		return
	}

	filename := filepath.Join(patDir, patName+".json")
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "no such pattern")
		return
	}
	defer f.Close()

	pat := pattern{}
	dec := json.NewDecoder(f)
	err = dec.Decode(&pat)

	if pat.Pattern == "" {
		// check for multiple patterns
		if len(pat.Patterns) == 0 {
			fmt.Fprintf(os.Stderr, "pattern file '%s' contains no pattern(s)\n", filename)
			return
		}

		pat.Pattern = "(" + strings.Join(pat.Patterns, "|") + ")"
	}

	var cmd *exec.Cmd

	cmd = exec.Command("grep", pat.Flags, pat.Pattern, files)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func grepBundle(bundleName, files string) {

	bundleDir, err := getPatternBundle()
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to open user's pattern directory")
		return
	}

	filename := filepath.Join(bundleDir, bundleName+".json")
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "no such pattern")
		return
	}
	defer f.Close()

	bundle := bundle{}
	dec := json.NewDecoder(f)
	err = dec.Decode(&bundle)

	if len(bundle.Bundles) > 0 {
		for _, bundle := range bundle.Bundles {
			grepBundle(bundle, files)
		}
	}

	for _, pattern := range bundle.Patterns {
		grepPattern(pattern, files)
	}
}

func getPatternDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	path := filepath.Join(usr.HomeDir, ".config/gc")
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		// .config/gf exists
		return path, nil
	}
	return "../templates/patterns", nil
	//return filepath.Join(usr.HomeDir, ".gc"), nil
}

func getBundleDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	path := filepath.Join(usr.HomeDir, ".config/gc")
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		// .config/gf exists
		return path, nil
	}
	return "../templates/bundles", nil
	//return filepath.Join(usr.HomeDir, ".gc"), nil
}

func printComment(bundlename, comment string) {
	fmt.Println("Bundle: %s", bundlename)
	fmt.Println("%s\n", comment)
}

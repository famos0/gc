# gc
Grep wrapper for source code review

# Description
This tool is largely inspired from [gf](https://github.com/tomnomnom/gf), a wonderfull idea and tool from Tomnomnom.
gc aims to be an improvment of gf, specifically built for source code review.
You can use some **Patterns** (just as gf) to simplify some grep commands, and also use **Bundles**, which are a combination of Patterns (and Bundles)

For example, if you have a pattern to hunt for concatenation and another pattern to hunt for dangerous function, you can bundle them up and run them in one command.

```
Usage:
  gc [command]

Available Commands:
  bundle      Grep for multiple patterns type
  help        Help about any command
  pattern     Grep for a pattern type

Flags:
  -h, --help       help for gc
  -q, --quiet      Don't print Patterns and Bundles name and comment
  -t, --testless   Don't grep test/mock code

Use "gc [command] --help" for more information about a command.
```

To use a pattern, just use:
```
gc pattern <pattern_name> <dir>
```

To use a bundle, just use:
```
gc bundle <bundle_name> <dir>
```

Some templates are available in **templates** folder

# Install

If you've installed [golang](https://golang.org/), you can install gc with:
```
go get -u github.com/famos0/gc
```

To add bash autocompletion, just use
```
source ./gc-completion.bash
```
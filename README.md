# gc
Grep wrapper for source code review

# Description
This tool is largely inspired from [gf](https://github.com/tomnomnom/gf), a wonderfull idea and tool from Tomnomnom.
gc aims to be an improvment of gf, specifically built for source code review.
You can use some **Patterns** (just as gf) to simplify some grep commands, and also use **Bundles**, which are a combination of Patterns (and Bundles).

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
  -s, --stdin      Display generated grep command instead of running them
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

Some templates are available in **templates** folder.

# Install

If you've installed [golang](https://golang.org/), you can install gc with:
```
go get -u github.com/famos0/gc
```

To add bash autocompletion, just use:
```
source ./gc-completion.bash
```

# Create bundles and patterns

## Pattern

A **Pattern** is json object where you declare **flags** that will be used in the grep command, and **patterns** that will be used in the grep command. The **comment** is optional.
```
$ cat fruits/leftovers.json 
{
  "flags": "-HnriE",
  "patterns": [
    "todo",
    "to do",
    "fixme",
    "fix me",
    "tofix",
    "to fix",
    "comment",
    "debug"
  ],
  "comment":"Hunt for leftovers from developers"
}
$ gc pattern fruits/leftovers .

# THE COMMAND LAUNCHED WILL BE 
grep -HnriE "(todo|to do|fixme|fix me|tofix|to fix|comment|debug)"
```

## Bundle
A **Bundle** is a combination of **patternspath** and **bundles**. As previously, **comment** is optional.

```
$ cat java.json 
{
    "patternspath":[
        "java/dangerous_functions",
        "java/heap_inspection",
        "java/xml_parser"
    ],
    "bundles": [
        "find-fruits"
    ],
    "comment":"Java classic bundle"
}
```
The bundle will launch grep commands for every patterns in **patternspath**. It will also launch bundles in **bundles**.
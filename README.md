# GoTea
Dead Simple CLI Tea Timer written in Go

# Quick start

In a terminal :
```bash
go get github.com/PPACI/gotea
```

Then just start the timer you want :
```bash
gotea black
```
or maybe, for 15 minutes timer (why not ?)
```bash
gotea custom 15
```

# Usage
```bash
Easily create timer for your tea, using a CLI !

Usage:
  gotea [command]

Available Commands:
  black       4 minutes
  custom      custom timer
  darjeeling  3 minutes
  green       3 minutes
  help        Help about any command
  oolong      4 minutes
  tisane      6 minutes
  white       2 minutes

Flags:
  -h, --help   help for gotea

Use "gotea [command] --help" for more information about a command.
```

# Warning
* As gotea will display a dialog to warn you about the end of the timer, it will not work on ssh or non-desktop server. 
* gotea use https://github.com/sqweek/dialog as dialog library, which use https://github.com/AllenDang/w32 on windows... Which can be picky to install due tot the fact it need gcc and complete golang build chain.
  * gotea may use another method for windows in the future as my goal is to provide a simple to use, __and to install__ cli timer.

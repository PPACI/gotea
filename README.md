# GoTea
Dead Simple CLI Tea Timer written in Go

Start a Tea timer and display a notification on your desktop ! 
**Work with Windows, MacOS, Linux !**

# Quick start
## go get
In a terminal :
```bash
go get github.com/PPACI/gotea
```
## compiled
Download [gotea](https://github.com/PPACI/gotea/releases) and put the binary file on your path.
## create a timer
Just start the timer you want :
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

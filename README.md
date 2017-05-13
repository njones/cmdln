[![Build Status](https://travis-ci.org/njones/cmdln.svg?branch=master)](https://travis-ci.org/njones/cmdln) [![GoDoc](https://godoc.org/github.com/njones/cmdln?status.svg)](https://godoc.org/github.com/njones/cmdln) [![Go Report Card](https://goreportcard.com/badge/github.com/njones/cmdln)](https://goreportcard.com/report/github.com/njones/cmdln)

# CmdLn

CmdLn is a simple library to split a command line string into a command and slice of arguments. 

# Documentation
Documentation can be found at [Godoc](https://godoc.org/github.com/njones/cmdln)

# Example of how to use cmdln
```Go
package main

import (
    "fmt"
    "os/exec"

	"github.com/njones/cmdln"
)

func main() {
    // here is a somewhat complex command line to execute
	fullCommand := `echo -e 'Starting LS\n===========' && ls -la && echo -e "===========\nI'm Done."`

    // split, so we can run this on the command line...
    cmmd, args := cmdln.Split(fullCommand)

    // you should be checking errors, even though I don't here.
    // run it!
    o, _ := os.Command(cmmd, args...).Output()
    
    // output the current directory
    fmt.Println(string(o))
}
```

# LICENSE

CmdLn is available under the [MIT License](https://opensource.org/licenses/MIT).

Copyright (c) 2017 Nika Jones <copyright@nikajon.es> All Rights Reserved.
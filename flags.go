package main

import (
	"flag"
	"os"

	"github.com/ppltools/cmsg"
)

var (
	v bool // enable debug
	s bool // short path, directly under github
	u bool // force update
)

func init() {
	flag.BoolVar(&v, "v", false, "verbose")
	flag.BoolVar(&s, "s", false, "short path")
	flag.BoolVar(&u, "u", false, "update")
}

func flagParse() []string {
	flag.Parse()

	args := os.Args[1:]

	if v {
		args = args[1:]
		cmsg.Default.IsDebugging = true
	}

	if s {
		args = args[1:]
	}

	if u {
		args = args[1:]
	}

	return args
}

package main

import (
	"flag"
	"log"
	"os"
)

var (
	l bool // long mode, create full path, otherwise, create in github.com
)

const (
	FMT   = "\033[%sm%s\033[m"
	RED   = "0;31"
	ERROR = "[ERROR]"
)

func init() {
	flag.BoolVar(&l, "l", false, "get source code with full path")

	log.SetFlags(log.LstdFlags)
}

func main() {
	flag.Parse()

	args := os.Args[1:]
	if l {
		args = args[1:]
	}

	for _, arg := range args {
		if !validate(arg) {
			log.Fatalf(FMT+" invalid name of repository: %s\n", RED, ERROR, arg)
		}
	}
}

func validate(repo string) bool {
	return false
}

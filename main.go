package main

import (
	"github.com/ppltools/cmsg"
)

func main() {
	args := flagParse()

	if len(args) == 0 {
		cmsg.Die("usage: gitrepo [-v] [-s] [-u] [-l c|go] repo1 repo2 repo3 ...")
	}

	downloader := NewDownloader()

	for _, arg := range args {
		gitrepo, group, module, ok := validate(arg)
		if !ok {
			cmsg.Die("invalid name of repository: %s\n", arg)
		}
		downloader.Download(gitrepo, arg, group, module)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/ppltools/cmsg"
)

func main() {
	args := flagParse()

	if len(args) == 0 {
		cmsg.Die("usage: gitrepo [-v] [-s] [-u] repo1 repo2 repo3 ...")
	}

	gopath := GetGOPATH()

	for _, arg := range args {
		gitrepo, group, module, ok := validate(arg)
		if !ok {
			cmsg.Die("invalid name of repository: %s\n", arg)
		}
		gitDownload(gopath, gitrepo, arg, group, module)
	}
}

func gitDownload(gopath, gitrepo, repo, group, module string) {
	cmsg.Debug("%s %s %s %s\n", gitrepo, repo, group, module)

	groupPath := CreatePath(gopath, gitrepo, group)
	ChangePath(groupPath)

	if _, err := os.Stat(module); err == nil {
		cmsg.Info("module %s is already exists", fmt.Sprintf("%s/%s", groupPath, module))

		if u {
			ChangePath(module)

			RunCmd("git", []string{"pull"}, cmsg.Die, "git pull %s failed: %s", repo)
		}
		return
	}

	RunCmd("git", []string{"clone", fmt.Sprintf("git@%s:%s/%s.git", gitrepo, group, module)},
		cmsg.Die, "git clone repo %s failed: %s", repo)

	cmsg.Info("git clone %s successfully", repo)
}

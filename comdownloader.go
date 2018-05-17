package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ppltools/cmsg"
)

type ComDownload struct {
	LangEnv  string
	ShortFmt string
	LongFmt  string
}

func (com *ComDownload) GetPath() string {
	path := os.Getenv(com.LangEnv)
	if pos := strings.IndexByte(path, ':'); pos > 0 {
		path = path[:pos-1]
	}
	if path == "" {
		cmsg.Die(com.LangEnv + " is not defined\n")
	}

	return path
}

func (com *ComDownload) CreatePath(gopath, gitrepo, group string) string {
	var groupPath string
	if s {
		groupPath = fmt.Sprintf(com.ShortFmt, gopath)
	} else {
		groupPath = fmt.Sprintf(com.LongFmt, gopath, gitrepo, group)
	}

	err := os.MkdirAll(groupPath, 0755)
	if err != nil {
		cmsg.Die("create dir %s failed: %s", groupPath, err)
	}

	return groupPath
}

func (com *ComDownload) Download(gitrepo, repo, group, module string) {
	cmsg.Debug("%s %s %s %s\n", gitrepo, repo, group, module)

	groupPath := com.CreatePath(com.GetPath(), gitrepo, group)
	ChangePath(groupPath)

	if _, err := os.Stat(module); err == nil {
		cmsg.Warn("module %s is already exists", fmt.Sprintf("%s/%s", groupPath, module))

		if u {
			ChangePath(module)
			RunCmd("git", []string{"pull"}, cmsg.Die, "git pull %s failed: %s", repo)
			return
		}

		cmsg.Info("use -u option for updating")
		return
	}

	RunCmd("git", []string{"clone", fmt.Sprintf("git@%s:%s/%s.git", gitrepo, group, module)},
		cmsg.Die, "git clone repo %s failed: %s", repo)

	cmsg.Info("git clone %s successfully", repo)
}

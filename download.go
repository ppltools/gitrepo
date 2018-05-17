package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ppltools/cmsg"
)

type Download struct {
	LangEnv  string // name of env path
	ShortFmt string // short format: path of the source code, eg: PATH/code
	LongFmt  string // long format: path of the source code, eg: PATH/gitrepo/group/code
}

func NewDownload(env, sf, lf string) *Download {
	return &Download{
		LangEnv:  env,
		ShortFmt: sf,
		LongFmt:  lf,
	}
}

func (d *Download) GetPath() string {
	path := os.Getenv(d.LangEnv)
	if pos := strings.IndexByte(path, ':'); pos > 0 {
		path = path[:pos-1]
	}
	if path == "" {
		cmsg.Die(d.LangEnv + " is not defined\n")
	}

	return path
}

func (d *Download) CreatePath(gopath, gitrepo, group string) string {
	var groupPath string
	if s {
		groupPath = fmt.Sprintf(d.ShortFmt, gopath)
	} else {
		groupPath = fmt.Sprintf(d.LongFmt, gopath, gitrepo, group)
	}

	err := os.MkdirAll(groupPath, 0755)
	if err != nil {
		cmsg.Die("create dir %s failed: %s", groupPath, err)
	}

	return groupPath
}

func (d *Download) Download(gr GitRepo, repo string) {
	cmsg.Debug("%s %s %s %s\n", gr.gitrepo, repo, gr.group, gr.module)

	groupPath := d.CreatePath(d.GetPath(), gr.gitrepo, gr.group)
	ChangePath(groupPath)

	if _, err := os.Stat(gr.module); err == nil {
		cmsg.Warn("module %s is already exists", fmt.Sprintf("%s/%s", groupPath, gr.module))

		if u {
			ChangePath(gr.module)
			RunCmd("git", []string{"pull"}, cmsg.Die, "git pull %s failed: %s", repo)
			return
		}

		cmsg.Info("use -u option for updating")
		return
	}

	RunCmd("git", []string{"clone", fmt.Sprintf("git@%s:%s/%s.git", gr.gitrepo, gr.group, gr.module)},
		cmsg.Die, "git clone repo %s failed: %s", repo)

	cmsg.Info("git clone %s successfully", repo)
}

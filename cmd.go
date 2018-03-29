package main

import (
	"os"
	"os/exec"
	"strings"

	"fmt"

	"github.com/ppltools/cmsg"
)

func GetGOPATH() string {
	gopath := os.Getenv("GOPATH")
	if pos := strings.IndexByte(gopath, ':'); pos > 0 {
		gopath = gopath[:pos-1]
	}
	if gopath == "" {
		cmsg.Die("GOPATH is not defined\n")
	}

	return gopath
}

func RunCmd(cmd string, args []string, msgFunc cmsg.MsgFunc, msgInfo string, msgArgs ...string) {
	cmsg.Info("run command: %s %s", cmd, strings.Join(args, " "))

	c := exec.Command(cmd, args...)
	err := c.Run()
	if err != nil {
		msgFunc(msgInfo, msgArgs, err)
	}
}

func CreatePath(gopath string, gitrepo string, group string) string {
	var groupPath string
	if s {
		groupPath = fmt.Sprintf("%s/src", gopath)
	} else {
		groupPath = fmt.Sprintf("%s/src/%s/%s", gopath, gitrepo, group)
	}

	err := os.MkdirAll(groupPath, 0755)
	if err != nil {
		cmsg.Die("create dir %s failed: %s", groupPath, err)
	}

	return groupPath
}

func ChangePath(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		cmsg.Die("switch to dir %s failed: %s", dir, err)
	}
}

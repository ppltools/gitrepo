package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/ppltools/cmsg"
)

func RunCmd(cmd string, args []string, msgFunc cmsg.MsgFunc, msgInfo string, msgArgs ...string) error {
	cmsg.Info("run command: %s %s", cmd, strings.Join(args, " "))

	c := exec.Command(cmd, args...)
	err := c.Run()
	if err != nil {
		msgFunc(msgInfo, msgArgs, err)
	}

	return err
}

func ChangePath(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		cmsg.Die("switch to dir %s failed: %s", dir, err)
	}
}

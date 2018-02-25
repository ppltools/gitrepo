package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/ppltools/cmsg"
)

var (
	v bool // enable debug
	s bool // short path, directly under github
)

var (
	pats []string
	exps []*regexp.Regexp
)

const (
	HTTPSPAT  = `^https\://([\w-]+(\.[\w-]+){1,2})/([\w-]+)/([\w-]+)\.git$`
	HTTPSPAT2 = `^https\://([\w-]+(\.[\w-]+){1,2})/([\w-]+)/([\w-]+)$`
	GITSPAT   = `^git\@([\w-]+(\.[\w-]+){1,2})\:([\w-]+)/([\w-]+).git$`
	GITSPAT2  = `^git\@([\w-]+(\.[\w-]+){1,2})\:([\w-]+)/([\w-]+)$`
	PAT       = `^([\w-]+(\.[\w-]+){1,2})/([\w-]+)/([\w-]+)$`
)

func init() {
	flag.BoolVar(&v, "v", false, "verbose")
	flag.BoolVar(&s, "s", false, "short path")

	pats = []string{HTTPSPAT, HTTPSPAT2, GITSPAT, GITSPAT2, PAT}
	for _, pat := range pats {
		exp := regexp.MustCompile(pat)
		exps = append(exps, exp)
	}
}

func main() {
	flag.Parse()

	args := os.Args[1:]

	if v {
		args = args[1:]
	}

	if s {
		args = args[1:]
	}

	if len(args) == 0 {
		cmsg.Die("usage: gitrepo [-v] [-s] repo1 repo2 repo3 ...")
	}

	// get gopath
	gopath := os.Getenv("GOPATH")
	if pos := strings.IndexByte(gopath, ':'); pos > 0 {
		gopath = gopath[:pos-1]
	}
	if gopath == "" {
		cmsg.Die("GOPATH is not defined\n")
	}

	for _, arg := range args {
		gitrepo, group, module, ok := validate(arg)
		if !ok {
			cmsg.Die("invalid name of repository: %s\n", arg)
		}
		gitPull(gopath, gitrepo, arg, group, module)
	}
}

func gitPull(gopath, gitrepo, repo, group, module string) {
	cmsg.Info("%s %s %s %s\n", gitrepo, repo, group, module)
	var groupPath string
	if s {
		groupPath = fmt.Sprintf("%s/src/%s", gopath, gitrepo)
	} else {
		groupPath = fmt.Sprintf("%s/src/%s/%s", gopath, gitrepo, group)
	}
	err := os.MkdirAll(groupPath, 0755)
	if err != nil {
		cmsg.Die("create dir %s failed: %s", groupPath, err)
	}
	err = os.Chdir(groupPath)
	if err != nil {
		cmsg.Die("switch to dir %s failed: %s", groupPath, err)
	}

	if _, err = os.Stat(module); err == nil {
		cmsg.Info("module %s is already exists", fmt.Sprintf("%s/%s", groupPath, module))
		return
	}

	cmd := exec.Command("git", "clone", fmt.Sprintf("git@%s:%s/%s.git", gitrepo, group, module))
	err = cmd.Run()
	if err != nil {
		cmsg.Die("git clone repo %s failed: %s", repo, err)
	}

	cmsg.Info("git clone %s successfully", repo)
}

func validate(repo string) (string, string, string, bool) {
	for _, exp := range exps {
		res := exp.MatchString(repo)
		if res {
			splits := exp.FindStringSubmatch(repo)
			if v {
				cmsg.Info("repo %s matched, splits: %v", repo, splits)
			}
			length := len(splits)
			if v {
				cmsg.Info("repo %s matched, repository: %s, group: %s, module: %s",
					repo, splits[length-4], splits[length-2], splits[length-1])
			}
			return splits[length-4], splits[length-2], splits[length-1], true
		}
	}
	return "", "", "", false
}

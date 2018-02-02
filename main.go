package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var (
	d bool // enable debug
	s bool // short path, directly under github
)

const (
	FMT    = "\033[%sm%s\033[m"
	RED    = "0;31"
	ERROR  = "[ERROR]"
	YELLOW = "0;33"
	WARN   = "[WARN]"
	GREEN  = "0;32"
	INFO   = "[INFO]"
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
	flag.BoolVar(&d, "d", false, "verbose")
	flag.BoolVar(&s, "s", false, "short path")

	log.SetFlags(log.LstdFlags)

	pats = []string{HTTPSPAT, HTTPSPAT2, GITSPAT, GITSPAT2, PAT}
	for _, pat := range pats {
		exp := regexp.MustCompile(pat)
		exps = append(exps, exp)
	}
}

func main() {
	flag.Parse()

	args := os.Args[1:]

	if d {
		args = args[1:]
	}

	if s {
		args = args[1:]
	}

	if len(args) == 0 {
		log.Fatalf(FMT+" usage: gitrepo repo1 repo2 repo3 ...", RED, ERROR)
	}

	// get gopath
	gopath := os.Getenv("GOPATH")
	if pos := strings.IndexByte(gopath, ':'); pos > 0 {
		gopath = gopath[:pos-1]
	}
	if gopath == "" {
		log.Fatalf(FMT+" GOPATH is not defined\n", RED, ERROR)
	}

	for _, arg := range args {
		gitrepo, group, module, ok := validate(arg)
		if !ok {
			log.Fatalf(FMT+" invalid name of repository: %s\n", RED, ERROR, arg)
		}
		gitPull(gopath, gitrepo, arg, group, module)
	}
}

func gitPull(gopath, gitrepo, repo, group, module string) {
	var groupPath string
	if s {
		groupPath = fmt.Sprintf("%s/src/%s", gopath, gitrepo)
	} else {
		groupPath = fmt.Sprintf("%s/src/%s/%s", gopath, gitrepo, group)
	}
	err := os.MkdirAll(groupPath, 0755)
	if err != nil {
		log.Fatalf(FMT+" create dir %s failed: %s", RED, ERROR, groupPath, err)
	}
	err = os.Chdir(groupPath)
	if err != nil {
		log.Fatalf(FMT+" switch to dir %s failed: %s", RED, ERROR, groupPath, err)
	}

	if _, err = os.Stat(module); err == nil {
		log.Printf(FMT+" module %s is already exists", YELLOW, WARN, fmt.Sprintf("%s/%s", groupPath, module))
		return
	}

	cmd := exec.Command("git", "clone", repo)
	err = cmd.Run()
	if err != nil {
		log.Fatalf(FMT+" git clone repo %s failed: %s", RED, ERROR, repo, err)
	}

	log.Printf(FMT+" git clone %s successfully", GREEN, INFO, repo)
}

func validate(repo string) (string, string, string, bool) {
	for _, exp := range exps {
		res := exp.MatchString(repo)
		if res {
			splits := exp.FindStringSubmatch(repo)
			if d {
				log.Printf(FMT+" repo %s matched, splits: %v", GREEN, INFO, repo, splits)
			}
			length := len(splits)
			if d {
				log.Printf(FMT+" repo %s matched, repository: %s, group: %s, module: %s",
					GREEN, INFO, repo, splits[length-4], splits[length-2], splits[length-1])
			}
			return splits[length-4], splits[length-2], splits[length-1], true
		}
	}
	return "", "", "", false
}

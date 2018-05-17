package main

import (
	"regexp"

	"github.com/ppltools/cmsg"
)

var (
	pats []string
	exps []*regexp.Regexp
)

type GitRepo struct {
	gitrepo string
	group   string
	module  string
}

const (
	HTTPSPAT  = `^https\://([\w-]+(\.[\w-]+){1,2})/([\w-]+)/([\w-]+)\.git$`
	HTTPSPAT2 = `^https\://([\w-]+(\.[\w-]+){1,2})/([\w-]+)/([\w-]+)$`
	GITSPAT   = `^git\@([\w-]+(\.[\w-]+){1,2})\:([\w-]+)/([\w-]+).git$`
	GITSPAT2  = `^git\@([\w-]+(\.[\w-]+){1,2})\:([\w-]+)/([\w-]+)$`
	PAT       = `^([\w-]+(\.[\w-]+){1,2})/([\w-]+)/([\w-]+)$`
)

func init() {
	pats = []string{HTTPSPAT, HTTPSPAT2, GITSPAT, GITSPAT2, PAT}
	for _, pat := range pats {
		exp := regexp.MustCompile(pat)
		exps = append(exps, exp)
	}
}

func validate(repo string) (GitRepo, bool) {
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
			return GitRepo{splits[length-4], splits[length-2], splits[length-1]}, true
		}
	}
	return GitRepo{"", "", ""}, false
}

package main

import (
	"github.com/ppltools/cmsg"
)

const (
	CPATH  = "CPATH"
	GOPATH = "GOPATH"

	CFmtShort  = "%s"
	CFmtLong   = "%s/%s/%s"
	GoFmtShort = "%s/src"
	GoFmtLong  = "%s/src/%s/%s"
)

type Downloader interface {
	GetPath() string
	CreatePath(root string, gitrepo string, group string) string

	Download(gitrepo, repo, group, module string)
}

func NewDownloader() Downloader {
	switch l {
	case "go":
		return NewDownload(GOPATH, GoFmtShort, GoFmtLong)
	case "c":
		return NewDownload(CPATH, CFmtShort, CFmtLong)
	default:
		cmsg.Die("Unsupport language: %s", l)
	}
	return nil
}

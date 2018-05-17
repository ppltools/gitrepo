package main

import (
	"github.com/ppltools/cmsg"
)

type Downloader interface {
	GetPath() string
	CreatePath(root string, gitrepo string, group string) string

	Download(gitrepo, repo, group, module string)
}

func NewDownloader() Downloader {
	switch l {
	case "go":
		return NewGoDown()
	case "c":
		return NewCDown()
	default:
		cmsg.Die("Unsupport language: %s", l)
	}
	return nil
}

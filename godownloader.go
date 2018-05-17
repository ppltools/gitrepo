package main

const (
	GOPATH = "GOPATH"

	GoFmtShort = "%s/src"
	GoFmtLong  = "%s/src/%s/%s"
)

type GoDown struct {
	d ComDownload
}

func NewGoDown() *GoDown {
	return &GoDown{d: ComDownload{
		LangEnv:  GOPATH,
		ShortFmt: GoFmtShort,
		LongFmt:  GoFmtLong,
	}}
}

func (g *GoDown) GetPath() string {
	return g.d.GetPath()
}

func (g *GoDown) CreatePath(root string, gitrepo string, group string) string {
	return g.d.CreatePath(root, gitrepo, group)
}

func (g *GoDown) Download(gitrepo, repo, group, module string) {
	g.d.Download(gitrepo, repo, group, module)
}

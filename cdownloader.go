package main

const (
	CPATH = "CPATH"

	CFmtShort = "%s"
	CFmtLong  = "%s/%s/%s"
)

type CDown struct {
	d ComDownload
}

func NewCDown() *CDown {
	return &CDown{d: ComDownload{
		LangEnv:  CPATH,
		ShortFmt: CFmtShort,
		LongFmt:  CFmtLong,
	}}
}

func (c *CDown) GetPath() string {
	return c.d.GetPath()
}

func (c *CDown) CreatePath(root string, gitrepo string, group string) string {
	return c.d.CreatePath(root, gitrepo, group)
}

func (c *CDown) Download(gitrepo, repo, group, module string) {
	c.d.Download(gitrepo, repo, group, module)
}

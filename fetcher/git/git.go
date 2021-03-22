package git

import (
	"github.com/bukowa/pkgo/fetcher"
	"github.com/bukowa/pkgo/src"
	log "github.com/sirupsen/logrus"
	"os"
)

type Fetcher struct {

}

func (f Fetcher) Fetch(p src.Package) (string, error) {
	p.WithFields().Infof("fetching")
	dir, err := os.MkdirTemp("", "")
	if err != nil {
		return "", err
	}
	cmd := fetcher.NewCommandWithArgs("git", dir, "-C", dir)
	args := [][]string{
		{"clone", p.Source, "."},
		{"fetch", "--recurse-submodules", "--depth=1", "origin", p.Version},
		{"checkout", "FETCH_HEAD"},
		{"submodule", "sync", "--recursive"},
		{"submodule", "update", "--init", "--recursive", "--remote", "--depth=1", "--no-single-branch"},
	}
	return dir, gits(args, cmd)
}

func gits(args[][]string, f func(a ...string) ([]byte, error)) error {
	for _, arg := range args {
		if err := git(f(arg...)); err != nil {
			return err
		}
	}
	return nil
}

func git(b []byte, err error) error {
	if err != nil {
		log.Print(string(b))
	}
	return err
}
package git

import (
	"github.com/bukowa/pkgo/fetcher"
	"github.com/bukowa/pkgo/src"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type Fetcher struct {

}

func (f Fetcher) Fetch(p src.Package) (string, error) {
	loc, err := filepath.Abs(p.Location)
	if err != nil {
		return "", err
	}
	p.WithFields().Infof("fetching into %s", loc)

	err = os.MkdirAll(loc, 0777)
	if err != nil {
		return "", nil
	}

	cmd := fetcher.NewCommandWithArgs("git", loc, "-C", loc)
	args := [][]string{
		{"clone", p.Source, "."},
		{"fetch", "--recurse-submodules", "--depth=1", "origin", p.Version},
		{"checkout", "FETCH_HEAD"},
		{"submodule", "sync", "--recursive"},
		{"submodule", "update", "--init", "--recursive", "--remote", "--depth=1", "--no-single-branch"},
	}
	return loc, gits(args, cmd)
}

func gits(args[][]string, f func(a ...string) (string, []byte, error)) error {
	for _, arg := range args {
		if s, b, err := f(arg...); err != nil {
			log.WithFields(log.Fields{
				"cmd": s,
				"out": string(b),
			}).Error(err)
			return err
		}
	}
	return nil
}

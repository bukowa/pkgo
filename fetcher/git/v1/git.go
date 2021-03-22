package v1git

import (
	"github.com/bukowa/pkgo/src"
	log "github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"path/filepath"
)

type Fetcher struct {

}

func (f Fetcher) Fetch(p src.Package) (string, error) {
	loc, err := filepath.Abs(p.Destination)
	if err != nil {
		return "", err
	}
	p.WithFields().Infof("fetching into %s", loc)

	err = os.MkdirAll(loc, 0777)
	if err != nil {
		return "", nil
	}

	cmd := newCommandWithArgs("git", loc, "-C", loc)
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

func newCommandWithArgs(name, workDir string, args ...string) func(a ...string) (string, []byte, error) {
	return func(a ...string) (string, []byte, error) {
		a = append(args, a...)
		cmd := exec.Command(name, a...)
		cmd.Dir = workDir
		cmd.Env = append(cmd.Env, "GIT_TERMINAL_PROMPT=0")
		b, err := cmd.CombinedOutput()
		return cmd.String(), b, err
	}
}

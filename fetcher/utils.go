package fetcher

import "os/exec"

func NewCommandWithArgs(name, workDir string, args ...string) func(a ...string) ([]byte, error) {
	return func(a ...string) ([]byte, error) {
		a = append(args, a...)
		cmd := exec.Command(name, a...)
		cmd.Dir = workDir
		return cmd.CombinedOutput()
	}
}

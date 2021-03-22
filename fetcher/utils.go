package fetcher

import "os/exec"

func NewCommandWithArgs(name, workDir string, args ...string) func(a ...string) (string, []byte, error) {
	return func(a ...string) (string, []byte, error) {
		a = append(args, a...)
		cmd := exec.Command(name, a...)
		cmd.Dir = workDir
		cmd.Env = append(cmd.Env, "GIT_TERMINAL_PROMPT=0")
		b, err := cmd.CombinedOutput()
		return cmd.String(), b, err
	}
}

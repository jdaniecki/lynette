package runner

import (
	"context"
	"os"
	"os/exec"
	"syscall"
)

type runner struct {
	attributes *syscall.SysProcAttr
	binary     string
}

func New(binary string) runner {
	r := runner{}
	r.binary = binary
	r.attributes = &syscall.SysProcAttr{
		//Cloneflags: syscall.CLONE_NEWUTS,
	}
	return r
}

func (r *runner) Run(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, r.binary)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = []string{"LYNETTE=true"}
	cmd.SysProcAttr = r.attributes
	return cmd.Run()
}

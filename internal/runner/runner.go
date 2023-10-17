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
	args       []string
}

type opt func(r *runner)

func New(binary string, opts ...opt) runner {
	r := runner{}
	r.binary = binary
	r.attributes = &syscall.SysProcAttr{}

	for _, o := range opts {
		o(&r)
	}

	return r
}

func WithArgs(args ...string) opt {
	return func(r *runner) {
		r.args = args
	}
}

func WithNewUts() opt {
	return func(r *runner) {
		r.attributes.Cloneflags |= syscall.CLONE_NEWUTS
	}
}

func (r *runner) Run(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, r.binary, r.args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = []string{"LYNETTE=true"}
	cmd.SysProcAttr = r.attributes
	return cmd.Run()
}

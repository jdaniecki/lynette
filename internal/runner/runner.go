// Package runner provides facilities to execute command on host in an isolated way
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

// New is a configurable runner constructor
func New(binary string, opts ...opt) *runner {
	r := &runner{}
	r.binary = binary
	r.attributes = &syscall.SysProcAttr{}

	for _, o := range opts {
		o(r)
	}

	return r
}

// Runner option which configures executable binary arguments
func WithArgs(args ...string) opt {
	return func(r *runner) {
		r.args = args
	}
}

// Runner option which configures command to be executed in new UTS namespace
func WithNewNamespaces() opt {
	return func(r *runner) {
		r.attributes.Cloneflags = syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNET | syscall.CLONE_NEWUSER
		r.attributes.UidMappings = append(r.attributes.UidMappings, syscall.SysProcIDMap{
			ContainerID: 0,
			HostID:      os.Getuid(),
			Size:        1,
		})
		r.attributes.GidMappings = append(r.attributes.GidMappings, syscall.SysProcIDMap{
			ContainerID: 0,
			HostID:      os.Getgid(),
			Size:        1,
		})
	}
}

// Run executes binary and waits until its completion
func (r *runner) Run(ctx context.Context) error {
	cmd := exec.CommandContext(ctx, r.binary, r.args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = []string{"LYNETTE=true"}
	cmd.SysProcAttr = r.attributes
	return cmd.Run()
}

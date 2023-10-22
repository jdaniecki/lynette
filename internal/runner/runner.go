// Package runner provides facilities to execute command on host in an isolated way
package runner

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type runner struct {
	attributes *syscall.SysProcAttr
	binary     string
	args       []string
	rootfs     string
}

// The runner constructor
func New(binary string, args ...string) *runner {
	r := &runner{}
	r.binary = binary
	r.args = args
	r.attributes = &syscall.SysProcAttr{
		Cloneflags:  syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNET | syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{{ContainerID: 0, HostID: os.Getuid(), Size: 1}},
		GidMappings: []syscall.SysProcIDMap{{ContainerID: 0, HostID: os.Getgid(), Size: 1}},
	}

	r.rootfs = os.Getenv("ROOTFS")
	return r
}

// Run executes binary and waits until its completion
func (r *runner) Run(ctx context.Context) error {
	if os.Args[0] == "/proc/self/exe" { // in container ?
		// setup container hostname
		err := syscall.Sethostname([]byte("container"))
		if err != nil {
			return fmt.Errorf("container hostname setup failed: %v", err)
		}

		// setup rootfs
		err = syscall.Chroot(r.rootfs)
		if err != nil {
			return fmt.Errorf("rootfs setup failed: %v", err)
		}

		err = os.Chdir("/")
		if err != nil {
			return fmt.Errorf("rootfs setup failed: %v", err)
		}

		// execute target process
		cmd := exec.CommandContext(ctx, r.binary, r.args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	// setup container process isolation
	cmd := exec.CommandContext(ctx, "/proc/self/exe", os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = r.attributes
	return cmd.Run()
}

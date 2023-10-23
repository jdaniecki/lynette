// Package runner provides facilities to execute command on host in an isolated way
package runner

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

type runner struct {
	attributes *syscall.SysProcAttr
	binary     string
	args       []string
	rootfs     string
}

// The runner constructor
func New(rootfs string, binary string, args ...string) *runner {
	r := &runner{}
	r.binary = binary
	r.args = args
	r.attributes = &syscall.SysProcAttr{
		Cloneflags:  syscall.CLONE_NEWNS | syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWNET | syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{{ContainerID: 0, HostID: os.Getuid(), Size: 1}},
		GidMappings: []syscall.SysProcIDMap{{ContainerID: 0, HostID: os.Getgid(), Size: 1}},
	}

	r.rootfs = rootfs
	return r
}

// Run executes binary and waits until its completion
func (r *runner) Run(ctx context.Context) error {
	if os.Args[0] == "/proc/self/exe" { // in container ?
		if err := setupHostname("container"); err != nil {
			return fmt.Errorf("container hostname setup failed: %v", err)
		}

		if err := setupRootfs(r.rootfs); err != nil {
			return fmt.Errorf("rootfs setup failed: %v", err)
		}

		// execute target process
		cmd := exec.CommandContext(ctx, r.binary, r.args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		slog.Debug("Executing command", "command", cmd)
		return cmd.Run()
	}

	// setup container process isolation
	cmd := exec.CommandContext(ctx, "/proc/self/exe", os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = r.attributes
	slog.Debug("Executing command", "command", cmd)
	return cmd.Run()
}

func setupHostname(hostname string) error {
	slog.Debug("Setting up hostname", "hostname", hostname)
	return syscall.Sethostname([]byte(hostname))
}

func setupRootfs(rootfsPath string) error {
	slog.Debug("Setting up rootfs...")

	slog.Debug("Changing root", "root", rootfsPath)
	if err := syscall.Chroot(rootfsPath); err != nil {
		return err
	}

	if err := os.Chdir("/"); err != nil {
		return err
	}

	proc := filepath.Join("/", "proc")
	slog.Debug("Mounting proc", "proc", proc)
	if err := os.MkdirAll(proc, 0755); err != nil {
		return fmt.Errorf("dir creation failed: %v", err)
	}

	if err := syscall.Mount("proc", proc, "proc", 0, ""); err != nil {
		return fmt.Errorf("mount failed: %v", err)
	}

	return nil
}

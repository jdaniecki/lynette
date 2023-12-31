package main_test

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var lynetteBinary = os.Getenv("LYNETTE_BINARY_PATH")
var rootfs = os.Getenv("ROOTFS")

func init() {
	fmt.Printf("Executing tests on %q\n", lynetteBinary)
}

func TestRootCmdFailure(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, lynetteBinary, "non-existend-command")
	require.Error(t, cmd.Run())
}

func TestRunSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, lynetteBinary, "run", rootfs, "true")
	require.NoError(t, cmd.Run())
}

func TestRunFailure(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, lynetteBinary, "run", rootfs, "false")
	require.Error(t, cmd.Run())
}

func TestRunTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, lynetteBinary, "run", rootfs, "sleep", "10")
	require.Error(t, cmd.Run())
}

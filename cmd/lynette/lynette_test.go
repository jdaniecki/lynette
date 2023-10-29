package main_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var lynetteBinary = os.Getenv("LYNETTE_BINARY_PATH")
var rootfs = os.Getenv("ROOTFS")

func init() {
	fmt.Printf("Executing tests on %q\n", lynetteBinary)

	fmt.Printf("Create host bridge... %q\n", lynetteBinary)
	cmd := strings.Split("sudo ip link add lynette0 type bridge", " ")
	err := exec.Command(cmd[0], cmd[1:]...).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func TestRootCmdFailure(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, lynetteBinary, "non-existend-command")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	require.Error(t, cmd.Run())
}

func TestRunSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, lynetteBinary, "run", rootfs, "true")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	require.NoError(t, cmd.Run())
}

func TestRunFailure(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, lynetteBinary, "run", rootfs, "false")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	require.Error(t, cmd.Run())
}

func TestRunTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, lynetteBinary, "run", rootfs, "sleep", "10")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	require.Error(t, cmd.Run())
}

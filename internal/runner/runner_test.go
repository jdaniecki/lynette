package runner_test

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

func init() {
	fmt.Printf("Executing tests on %q\n", lynetteBinary)
}

func TestRunSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, lynetteBinary, "run", "true")
	require.NoError(t, cmd.Run())
}

func TestRunFailure(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, lynetteBinary, "run", "false")
	require.Error(t, cmd.Run())
}

func TestRunTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, lynetteBinary, "run", "sleep", "10")
	require.Error(t, cmd.Run())
}

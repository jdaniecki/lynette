package runner_test

import (
	"context"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var lynetteBinary = os.Getenv("LYNETTE_BINARY_PATH")

func TestRunSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, lynetteBinary, "run", "true")
	assert.NoError(t, cmd.Run())
}

func TestRunFailure(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, lynetteBinary, "run", "false")
	assert.Error(t, cmd.Run())
}

func TestRunTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, lynetteBinary, "run", "sleep", "10")
	assert.Error(t, cmd.Run())
}

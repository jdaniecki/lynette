package runner_test

import (
	"context"
	"testing"
	"time"

	"github.com/jdaniecki/lynette/internal/runner"
	"github.com/stretchr/testify/assert"
)

func TestRunSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	runner := runner.New("true")
	assert.NoError(t, runner.Run(ctx))
}

func TestRunFailure(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	runner := runner.New("false")
	assert.Error(t, runner.Run(ctx))
}

func TestRunTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	runner := runner.New("sleep", runner.WithArgs("10"))
	assert.Error(t, runner.Run(ctx))
}

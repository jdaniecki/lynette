package runner_test

import (
	"context"
	"testing"
	"time"

	_ "github.com/jdaniecki/lynette/internal/run"
	runner "github.com/jdaniecki/lynette/internal/run"
	"github.com/stretchr/testify/assert"
)

func TestRunSuccess(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	runner := runner.New("true")
	assert.NoError(t, runner.Run(ctx))
}

func TestRunFailure(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	runner := runner.New("false")
	assert.Error(t, runner.Run(ctx))
}

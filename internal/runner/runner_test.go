package runner_test

// TODO: fix tests to work with isolation
/*
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
*/

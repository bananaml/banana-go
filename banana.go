package banana

import (
	"context"
	"fmt"
	"strings"
	"time"

	uuid "github.com/google/uuid"
)

// Run will call the inference pipeline on custom models with the use of a model key.
// It is a syncronous wrapper around the async Start and Check functions.
func Run(apiKey string, modelKey string, inputs []byte) (Result, error) {
	return RunWithContext(context.Background(), apiKey, modelKey, inputs)
}

// RunWithContext is the same as Run, but with a non-default context.
// This allows for in-progress cancellation.
func RunWithContext(
	ctx context.Context,
	apiKey string,
	modelKey string,
	inputs []byte,
) (Result, error) {
	// Start the task
	startOnly := false
	res, err := subStart(ctx, apiKey, modelKey, inputs, startOnly)
	if err != nil {
		return Result{}, err
	}

	// catch the case where the start endpoint actually returned results
	if res.Finished {
		out := Result{
			ID:           res.ID,
			Message:      res.Message,
			Created:      res.Created,
			APIVersion:   res.APIVersion,
			ModelOutputs: res.ModelOutputs,
		}
		return out, nil
	}

	// Else if long running, poll check until done
	out := Result{}
	for {
		out, err = CheckWithContext(ctx, apiKey, res.CallID)
		if err != nil {
			return Result{}, err
		}

		if strings.ToLower(out.Message) == "success" {
			break
		}
	}

	return out, nil
}

// Start will start an async inference task and return a task ID.
func Start(apiKey string, modelKey string, inputs []byte) (callID string, err error) {
	return StartWithContext(context.Background(), apiKey, modelKey, inputs)
}

// StartWithContext is the same as Start, but with a non-default context.
// This allows for in-progress cancellation.
func StartWithContext(
	ctx context.Context,
	apiKey string,
	modelKey string,
	inputs []byte,
) (callID string, err error) {
	startOnly := true
	re, err := subStart(ctx, apiKey, modelKey, inputs, startOnly)
	if err != nil {
		return "", err
	}
	return re.CallID, err
}

// subStart is a start call wrapper returning the whole payload, for use by Start and Run
func subStart(
	ctx context.Context,
	apiKey string,
	modelKey string,
	inputs []byte,
	startOnly bool,
) (outStartV3, error) {
	p := inStartV3{
		ID:          uuid.New().String(),
		Created:     time.Now().Unix(),
		APIKey:      apiKey,
		ModelKey:    modelKey,
		ModelInputs: inputs,
		StartOnly:   startOnly,
	}

	re := outStartV3{}

	url := endpoint + "start/v3/"

	err := post(ctx, url, &p, &re)
	if err != nil {
		return re, err
	}

	lower := strings.ToLower(re.Message)
	if strings.Contains(lower, "error") {
		return re, fmt.Errorf(re.Message)
	}

	return re, nil
}

// Check will check the status of an existing async inference task.
func Check(apiKey string, callID string) (Result, error) {
	return CheckWithContext(context.Background(), apiKey, callID)
}

// CheckWithContext is the same as Check, but with a non-default context.
// This allows for in-progress cancellation.
func CheckWithContext(ctx context.Context, apiKey string, callID string) (Result, error) {
	p := inCheckV3{
		ID:       uuid.New().String(),
		Created:  time.Now().Unix(),
		APIKey:   apiKey,
		CallID:   callID,
		LongPoll: true,
	}

	re := Result{}

	url := endpoint + "check/v3/"

	err := post(ctx, url, &p, &re)
	if err != nil {
		return Result{}, err
	}

	lower := strings.ToLower(re.Message)
	if strings.Contains(lower, "error") {
		return Result{}, fmt.Errorf(re.Message)
	}

	return re, nil
}

package banana

import (
	"fmt"
	"strings"
	"time"

	uuid "github.com/google/uuid"
)

// Run will call the inference pipeline on custom models with the use of a model key.
// It is a syncronous wrapper around the async Start and Check functions.
func Run(apiKey string, modelKey string, inputs []byte) (outCheckV2, error) {

	// Start the task
	callID, err := Start(apiKey, modelKey, inputs)
	if err != nil {
		return outCheckV2{}, err
	}

	re := outCheckV2{}

	// Poll check until done
	for {
		re, err = Check(apiKey, callID)
		if err != nil {
			return outCheckV2{}, err
		}

		if strings.ToLower(re.Message) == "success" {
			break
		}
	}

	return re, nil
}

// Start will start an async inference task and return a task ID.
func Start(apiKey string, modelKey string, inputs []byte) (callID string, err error) {

	p := inStartV2{
		ID:          uuid.New().String(),
		Created:     time.Now().Unix(),
		APIKey:      apiKey,
		ModelKey:    modelKey,
		ModelInputs: inputs,
	}

	re := outStartV2{}

	url := endpoint + "start/v2/"

	err = post(url, &p, &re)
	if err != nil {
		return "", err
	}

	lower := strings.ToLower(re.Message)
	if strings.Contains(lower, "error") {
		return "", fmt.Errorf(re.Message)
	}

	return re.CallID, nil
}

// Check will check the status of an existing async inference task. If the task has finished, the task's return values will be marshalled into payloadOut.
// The "done" boolean return value indicates if the requested async inference task has finished (true) or is still running (false).
func Check(apiKey string, callID string) (outCheckV2, error) {

	p := inCheckV2{
		ID:       uuid.New().String(),
		Created:  time.Now().Unix(),
		APIKey:   apiKey,
		CallID:   callID,
		LongPoll: true,
	}

	re := outCheckV2{}

	url := endpoint + "check/v2/"

	err := post(url, &p, &re)
	if err != nil {
		return outCheckV2{}, err
	}

	lower := strings.ToLower(re.Message)
	if strings.Contains(lower, "error") {
		return outCheckV2{}, fmt.Errorf(re.Message)
	}

	return re, nil
}

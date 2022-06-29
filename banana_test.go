package banana

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"
)

const (
	testingAPIKey         = "demo"
	testingModelKey       = "carrot"
	testingCarrotImageURL = "https://demo-images-banana.s3.us-west-1.amazonaws.com/image2.jpg"
)

func TestRunCarrot(t *testing.T) {

	type input struct {
		ImageURL string `json:"imageURL"`
	}

	in := input{
		ImageURL: testingCarrotImageURL,
	}

	bytesIn, _ := json.Marshal(in)

	bananaOut, err := Run(testingAPIKey, testingModelKey, bytesIn)
	if err != nil {
		panic(err)
	}

	type output []struct {
		Caption string `json:"caption"`
	}

	out := output{}
	json.Unmarshal(bananaOut.ModelOutputs, &out)
	fmt.Println(out[0].Caption)

}

func TestContext(t *testing.T) {
	buf, err := json.Marshal(map[string]string{
		"imageURL": testingCarrotImageURL,
	})
	if err != nil {
		t.Fatal(err)
	}

	// We put a stupidly low timeout, to make sure we get a timeout error.
	// The laws of physics *should* prohibit a value of 1ms from ever being
	// a valid timeout... we'd hope at least.
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	_, err = RunWithContext(ctx, testingAPIKey, testingModelKey, buf)

	// Ensure that the context was cancelled.
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected context.DeadlineExceeded error, got %v", err)
	}
}

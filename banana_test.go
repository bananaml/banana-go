package banana

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRunCarrot(t *testing.T) {

	apiKey := "demo"
	modelKey := "carrot"

	type input struct {
		ImageURL string `json:"imageURL"`
	}

	in := input{
		ImageURL: "https://demo-images-banana.s3.us-west-1.amazonaws.com/image2.jpg",
	}

	bytesIn, _ := json.Marshal(in)

	bananaOut, err := Run(apiKey, modelKey, bytesIn)
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

package main

import (
	"encoding/json"
	"fmt"
	banana "github.com/bananaml/banana-go"
)

func main() {
	apiKey := "demo"
	modelKey := "carrot"

	// a model input struct specific to the json intake of your model
	type input struct {
		ImageURL string `json:"imageURL"`
	}

	in := input{
		ImageURL: "https://demo-images-banana.s3.us-west-1.amazonaws.com/image2.jpg",
	}

	bytesIn, _ := json.Marshal(in)

	bananaOut, err := banana.Run(apiKey, modelKey, bytesIn)
	if err != nil {
		panic(err)
	}

	// a model output struct specific to the json intake of your model
	type output []struct {
		Caption string `json:"caption"`
	}

	out := output{}
	json.Unmarshal(bananaOut.ModelOutputs, &out)
	fmt.Println(out[0].Caption)
}
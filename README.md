# Banana Go SDK

### Getting Started

Install via go get
`go get github.com/bananaml/banana-go`

Get your API Key
- [Sign in / log in here](https://app.banana.dev)

Run:
```go
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

```

bananaOut type:
```javascript
{
    "id": "12345678-1234-1234-1234-123456789012", 
    "message": "success", 
    "created": 1649712752, 
    "apiVersion": "26 Nov 2021", 
    "modelOutputs": [
        {
            // a json specific to your model. In this example, the caption of the image
            "caption": "a baseball player throwing a ball", 
        }
    ]
}
```

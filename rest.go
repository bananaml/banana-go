package banana

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type errJSON struct {
	Message string
}

func post(ctx context.Context, url string, p interface{}, re interface{}) error {

	// Marshal the payload
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Post it
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	// Catch non-200 status codes
	if resp.StatusCode != http.StatusOK {

		if resp.StatusCode == 500 {
			defer resp.Body.Close()
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			errStruct := errJSON{}
			err = json.Unmarshal(bodyBytes, &errStruct)
			if err != nil {
				return err
			}
			return fmt.Errorf(
				"banana returned status code %v with message:%s",
				resp.StatusCode,
				errStruct.Message,
			)
		}

		return fmt.Errorf("banana returned status code %v", resp.StatusCode)
	}

	// Read body into bytes
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse returned bytes into the struct
	err = json.Unmarshal(bodyBytes, re)
	if err != nil {
		return err
	}
	return nil
}

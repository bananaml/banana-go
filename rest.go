package banana

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type errJSON struct {
	Message string
}

func post(url string, p interface{}, re interface{}) error {

	// Marshal the payload
	jsonBytes, _ := json.Marshal(p)

	// Post it
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
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
			return fmt.Errorf("banana returned status code %v with message:%s", resp.StatusCode, errStruct.Message)
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

package banana

import "encoding/json"

// This package defines the entirety of what data can be accepted on the inbound Post requests

// InStartV3 Defines the acceptable payload into the task/start/v3 endpoint
type inStartV3 struct {
	ID          string          `json:"id" xml:"id" form:"id"`
	Created     int64           `json:"created" xml:"created" form:"created"`
	APIKey      string          `json:"apiKey" xml:"apiKey" form:"apiKey"`
	ModelKey    string          `json:"modelKey" xml:"modelKey" form:"modelKey"`
	StartOnly   bool            `json:"startOnly" xml:"startOnly" form:"startOnly"`
	ModelInputs json.RawMessage `json:"modelInputs" xml:"modelInputs" form:"modelInputs"` // Keeps the dynamic substruct as raw bytes
}

// OutStartV3 defines the payload to be returned by the task/start/v3 endpoint
type outStartV3 struct {
	ID           string          `json:"id" xml:"id" form:"id"`
	Message      string          `json:"message" xml:"message" form:"message"`
	Created      int64           `json:"created" xml:"created" form:"created"`
	APIVersion   string          `json:"apiVersion" xml:"apiVersion" form:"apiVersion"`
	CallID       string          `json:"callID" xml:"callID" form:"callID"`
	Finished     bool            `json:"finished" xml:"finished" form:"finished"`
	ModelOutputs json.RawMessage `json:"modelOutputs" xml:"modelOutputs" form:"modelOutputs"`
}

// InCheckV3 Defines the acceptable payload into the task/check/v3 endpoint
type inCheckV3 struct {
	ID       string `json:"id" xml:"id" form:"id"`
	Created  int64  `json:"created" xml:"created" form:"created"`
	APIKey   string `json:"apiKey" xml:"apiKey" form:"apiKey"`
	LongPoll bool   `json:"longPoll" xml:"longPoll" form:"longPoll"`
	CallID   string `json:"callID" xml:"callID" form:"callID"`
}

// OutCheckV3 defines the payload to be returned by the task/check/v3 endpoint
type outCheckV3 struct {
	ID           string          `json:"id" xml:"id" form:"id"`
	Message      string          `json:"message" xml:"message" form:"message"`
	Created      int64           `json:"created" xml:"created" form:"created"`
	APIVersion   string          `json:"apiVersion" xml:"apiVersion" form:"apiVersion"`
	ModelOutputs json.RawMessage `json:"modelOutputs" xml:"modelOutputs" form:"modelOutputs"`
}

// Export Result, aliased from outCheckV3, which is the return of banana.Run and banana.Check
type Result outCheckV3

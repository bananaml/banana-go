package banana

import "encoding/json"

// This package defines the entirety of what data can be accepted on the inbound Post requests

// InStartV2 Defines the acceptable payload into the task/start/v2 endpoint
type inStartV2 struct {
	ID          string          `json:"id" xml:"id" form:"id"`
	Created     int64           `json:"created" xml:"created" form:"created"`
	APIKey      string          `json:"apiKey" xml:"apiKey" form:"apiKey"`
	ModelKey    string          `json:"modelKey" xml:"modelKey" form:"modelKey"`
	ModelInputs json.RawMessage `json:"modelInputs" xml:"modelInputs" form:"modelInputs"` // Keeps the dynamic substruct as raw bytes
	Config      struct {
		MachineID string `json:"machineID" xml:"machineID" form:"machineID"`
	} `json:"config" xml:"config" form:"config"`
}

// OutStartV1 defines the payload to be returned by the task/start/v2 endpoint
type outStartV2 struct {
	ID         string `json:"id" xml:"id" form:"id"`
	Message    string `json:"message" xml:"message" form:"message"`
	Created    int64  `json:"created" xml:"created" form:"created"`
	APIVersion string `json:"apiVersion" xml:"apiVersion" form:"apiVersion"`
	CallID     string `json:"callID" xml:"callID" form:"callID"`
}

// InCheckV1 Defines the acceptable payload into the task/check/v2 endpoint
type inCheckV2 struct {
	ID       string `json:"id" xml:"id" form:"id"`
	Created  int64  `json:"created" xml:"created" form:"created"`
	APIKey   string `json:"apiKey" xml:"apiKey" form:"apiKey"`
	LongPoll bool   `json:"longPoll" xml:"longPoll" form:"longPoll"`
	CallID   string `json:"callID" xml:"callID" form:"callID"`
}

// OutCheckV1 defines the payload to be returned by the task/check/v2 endpoint
type outCheckV2 struct {
	ID           string          `json:"id" xml:"id" form:"id"`
	Message      string          `json:"message" xml:"message" form:"message"`
	Created      int64           `json:"created" xml:"created" form:"created"`
	APIVersion   string          `json:"apiVersion" xml:"apiVersion" form:"apiVersion"`
	ModelOutputs json.RawMessage `json:"modelOutputs" xml:"modelOutputs" form:"modelOutputs"`
}

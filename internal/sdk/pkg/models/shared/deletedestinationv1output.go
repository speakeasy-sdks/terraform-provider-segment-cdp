// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DeleteDestinationV1OutputStatus - The status of the Warehouse deletion operation.
type DeleteDestinationV1OutputStatus string

const (
	DeleteDestinationV1OutputStatusSuccess DeleteDestinationV1OutputStatus = "SUCCESS"
)

func (e DeleteDestinationV1OutputStatus) ToPointer() *DeleteDestinationV1OutputStatus {
	return &e
}

func (e *DeleteDestinationV1OutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		*e = DeleteDestinationV1OutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteDestinationV1OutputStatus: %v", v)
	}
}

// DeleteDestinationV1Output - Returns the status of a Destination deletion.
type DeleteDestinationV1Output struct {
	// The status of the Warehouse deletion operation.
	Status DeleteDestinationV1OutputStatus `json:"status"`
}

func (o *DeleteDestinationV1Output) GetStatus() DeleteDestinationV1OutputStatus {
	if o == nil {
		return DeleteDestinationV1OutputStatus("")
	}
	return o.Status
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DeleteUsersV1OutputStatus - A flag that indicates the status of a successful deletion operation.
type DeleteUsersV1OutputStatus string

const (
	DeleteUsersV1OutputStatusSuccess DeleteUsersV1OutputStatus = "SUCCESS"
)

func (e DeleteUsersV1OutputStatus) ToPointer() *DeleteUsersV1OutputStatus {
	return &e
}

func (e *DeleteUsersV1OutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		*e = DeleteUsersV1OutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteUsersV1OutputStatus: %v", v)
	}
}

// DeleteUsersV1Output - Returns the status of the removal operation.
type DeleteUsersV1Output struct {
	// A flag that indicates the status of a successful deletion operation.
	Status DeleteUsersV1OutputStatus `json:"status"`
}

func (o *DeleteUsersV1Output) GetStatus() DeleteUsersV1OutputStatus {
	if o == nil {
		return DeleteUsersV1OutputStatus("")
	}
	return o.Status
}

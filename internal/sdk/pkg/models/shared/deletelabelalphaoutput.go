// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DeleteLabelAlphaOutputStatus - The status of the label deletion operation.
type DeleteLabelAlphaOutputStatus string

const (
	DeleteLabelAlphaOutputStatusSuccess DeleteLabelAlphaOutputStatus = "SUCCESS"
)

func (e DeleteLabelAlphaOutputStatus) ToPointer() *DeleteLabelAlphaOutputStatus {
	return &e
}

func (e *DeleteLabelAlphaOutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		*e = DeleteLabelAlphaOutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteLabelAlphaOutputStatus: %v", v)
	}
}

// DeleteLabelAlphaOutput - Returns the status of a label deletion.
type DeleteLabelAlphaOutput struct {
	// The status of the label deletion operation.
	Status DeleteLabelAlphaOutputStatus `json:"status"`
}
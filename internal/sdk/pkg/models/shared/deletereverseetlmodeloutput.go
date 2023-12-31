// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DeleteReverseEtlModelOutputStatus - The result of the deletion.
type DeleteReverseEtlModelOutputStatus string

const (
	DeleteReverseEtlModelOutputStatusSuccess DeleteReverseEtlModelOutputStatus = "SUCCESS"
)

func (e DeleteReverseEtlModelOutputStatus) ToPointer() *DeleteReverseEtlModelOutputStatus {
	return &e
}

func (e *DeleteReverseEtlModelOutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SUCCESS":
		*e = DeleteReverseEtlModelOutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeleteReverseEtlModelOutputStatus: %v", v)
	}
}

// DeleteReverseEtlModelOutput - Defines the result of getting a Model.
type DeleteReverseEtlModelOutput struct {
	// The result of the deletion.
	Status DeleteReverseEtlModelOutputStatus `json:"status"`
}

func (o *DeleteReverseEtlModelOutput) GetStatus() DeleteReverseEtlModelOutputStatus {
	if o == nil {
		return DeleteReverseEtlModelOutputStatus("")
	}
	return o.Status
}

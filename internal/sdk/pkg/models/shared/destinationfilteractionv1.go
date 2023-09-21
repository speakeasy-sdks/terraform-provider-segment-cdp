// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DestinationFilterActionV1Type - The kind of Transformation to apply to any matched properties.
type DestinationFilterActionV1Type string

const (
	DestinationFilterActionV1TypeAllowProperties DestinationFilterActionV1Type = "ALLOW_PROPERTIES"
	DestinationFilterActionV1TypeDrop            DestinationFilterActionV1Type = "DROP"
	DestinationFilterActionV1TypeDropProperties  DestinationFilterActionV1Type = "DROP_PROPERTIES"
	DestinationFilterActionV1TypeSample          DestinationFilterActionV1Type = "SAMPLE"
)

func (e DestinationFilterActionV1Type) ToPointer() *DestinationFilterActionV1Type {
	return &e
}

func (e *DestinationFilterActionV1Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ALLOW_PROPERTIES":
		fallthrough
	case "DROP":
		fallthrough
	case "DROP_PROPERTIES":
		fallthrough
	case "SAMPLE":
		*e = DestinationFilterActionV1Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationFilterActionV1Type: %v", v)
	}
}

// DestinationFilterActionV1 - Represents a Destination filter action.
type DestinationFilterActionV1 struct {
	// A dictionary of paths to object keys that this filter applies to.
	//   The literal string '' represents the top level of the object.
	Fields map[string]interface{} `json:"fields,omitempty"`
	// The JSON path to a property within a payload object from which Segment generates a deterministic
	// sampling rate.
	Path *string `json:"path,omitempty"`
	// A decimal between 0 and 1 used for 'sample' type events and
	// influences the likelihood of sampling to occur.
	Percent *float64 `json:"percent,omitempty"`
	// The kind of Transformation to apply to any matched properties.
	Type DestinationFilterActionV1Type `json:"type"`
}

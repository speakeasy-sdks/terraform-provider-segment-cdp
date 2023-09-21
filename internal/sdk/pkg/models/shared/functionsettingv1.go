// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// FunctionSettingV1Type - The type of this Function Setting.
type FunctionSettingV1Type string

const (
	FunctionSettingV1TypeArray   FunctionSettingV1Type = "ARRAY"
	FunctionSettingV1TypeBoolean FunctionSettingV1Type = "BOOLEAN"
	FunctionSettingV1TypeString  FunctionSettingV1Type = "STRING"
	FunctionSettingV1TypeTextMap FunctionSettingV1Type = "TEXT_MAP"
)

func (e FunctionSettingV1Type) ToPointer() *FunctionSettingV1Type {
	return &e
}

func (e *FunctionSettingV1Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ARRAY":
		fallthrough
	case "BOOLEAN":
		fallthrough
	case "STRING":
		fallthrough
	case "TEXT_MAP":
		*e = FunctionSettingV1Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for FunctionSettingV1Type: %v", v)
	}
}

type FunctionSettingV1 struct {
	// A description of this Function Setting.
	Description string `json:"description"`
	// The label for this Function Setting.
	Label string `json:"label"`
	// The name of this Function Setting.
	Name string `json:"name"`
	// Whether this Function Setting is required.
	Required bool `json:"required"`
	// Whether this Function Setting contains sensitive information.
	Sensitive bool `json:"sensitive"`
	// The type of this Function Setting.
	Type FunctionSettingV1Type `json:"type"`
}

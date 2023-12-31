// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// RuleV1Type - The type for this Tracking Plan rule.
type RuleV1Type string

const (
	RuleV1TypeCommon   RuleV1Type = "COMMON"
	RuleV1TypeGroup    RuleV1Type = "GROUP"
	RuleV1TypeIdentify RuleV1Type = "IDENTIFY"
	RuleV1TypePage     RuleV1Type = "PAGE"
	RuleV1TypeScreen   RuleV1Type = "SCREEN"
	RuleV1TypeTrack    RuleV1Type = "TRACK"
)

func (e RuleV1Type) ToPointer() *RuleV1Type {
	return &e
}

func (e *RuleV1Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "COMMON":
		fallthrough
	case "GROUP":
		fallthrough
	case "IDENTIFY":
		fallthrough
	case "PAGE":
		fallthrough
	case "SCREEN":
		fallthrough
	case "TRACK":
		*e = RuleV1Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RuleV1Type: %v", v)
	}
}

// RuleV1 - Represents a rule from a Tracking Plan.
type RuleV1 struct {
	// The timestamp of this rule's creation.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The timestamp of this rule's deprecation.
	DeprecatedAt *string `json:"deprecatedAt,omitempty"`
	// JSON Schema of this rule.
	JSONSchema interface{} `json:"jsonSchema"`
	// Key to this rule (free-form string like 'Button clicked').
	Key *string `json:"key,omitempty"`
	// The type for this Tracking Plan rule.
	Type RuleV1Type `json:"type"`
	// The timestamp of this rule's last change.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// Version of this rule.
	Version float64 `json:"version"`
}

func (o *RuleV1) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RuleV1) GetDeprecatedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeprecatedAt
}

func (o *RuleV1) GetJSONSchema() interface{} {
	if o == nil {
		return nil
	}
	return o.JSONSchema
}

func (o *RuleV1) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *RuleV1) GetType() RuleV1Type {
	if o == nil {
		return RuleV1Type("")
	}
	return o.Type
}

func (o *RuleV1) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *RuleV1) GetVersion() float64 {
	if o == nil {
		return 0.0
	}
	return o.Version
}

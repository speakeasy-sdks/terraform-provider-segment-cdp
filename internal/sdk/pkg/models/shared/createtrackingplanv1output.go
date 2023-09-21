// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreateTrackingPlanV1OutputTrackingPlanV1Type - The Tracking Plan's type.
type CreateTrackingPlanV1OutputTrackingPlanV1Type string

const (
	CreateTrackingPlanV1OutputTrackingPlanV1TypeEngage          CreateTrackingPlanV1OutputTrackingPlanV1Type = "ENGAGE"
	CreateTrackingPlanV1OutputTrackingPlanV1TypeLive            CreateTrackingPlanV1OutputTrackingPlanV1Type = "LIVE"
	CreateTrackingPlanV1OutputTrackingPlanV1TypePropertyLibrary CreateTrackingPlanV1OutputTrackingPlanV1Type = "PROPERTY_LIBRARY"
	CreateTrackingPlanV1OutputTrackingPlanV1TypeRuleLibrary     CreateTrackingPlanV1OutputTrackingPlanV1Type = "RULE_LIBRARY"
	CreateTrackingPlanV1OutputTrackingPlanV1TypeTemplate        CreateTrackingPlanV1OutputTrackingPlanV1Type = "TEMPLATE"
)

func (e CreateTrackingPlanV1OutputTrackingPlanV1Type) ToPointer() *CreateTrackingPlanV1OutputTrackingPlanV1Type {
	return &e
}

func (e *CreateTrackingPlanV1OutputTrackingPlanV1Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ENGAGE":
		fallthrough
	case "LIVE":
		fallthrough
	case "PROPERTY_LIBRARY":
		fallthrough
	case "RULE_LIBRARY":
		fallthrough
	case "TEMPLATE":
		*e = CreateTrackingPlanV1OutputTrackingPlanV1Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTrackingPlanV1OutputTrackingPlanV1Type: %v", v)
	}
}

// CreateTrackingPlanV1OutputTrackingPlanV1 - Represents a Tracking Plan.
type CreateTrackingPlanV1OutputTrackingPlanV1 struct {
	// The timestamp of this Tracking Plan's creation.
	//
	// Config API note: equal to `createTime`.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The Tracking Plan's description.
	Description *string `json:"description,omitempty"`
	// The Tracking Plan's identifier.
	//
	// Config API note: analogous to `name`.
	ID string `json:"id"`
	// The Tracking Plan's name.
	//
	// Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	// URL-friendly slug of this Tracking Plan.
	//
	// Config API note: equal to `name`.
	Slug *string `json:"slug,omitempty"`
	// The Tracking Plan's type.
	Type CreateTrackingPlanV1OutputTrackingPlanV1Type `json:"type"`
	// The timestamp of the last change to the Tracking Plan.
	//
	// Config API note: equal to `updateTime`.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// CreateTrackingPlanV1Output - Result of a CreateTrackingPlan call.
type CreateTrackingPlanV1Output struct {
	// The created Tracking Plan.
	TrackingPlan CreateTrackingPlanV1OutputTrackingPlanV1 `json:"trackingPlan"`
}

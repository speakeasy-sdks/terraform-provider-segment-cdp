// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ReplaceRulesInTrackingPlanRequest struct {
	ReplaceRulesInTrackingPlanV1Input shared.ReplaceRulesInTrackingPlanV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	TrackingPlanID                    string                                   `pathParam:"style=simple,explode=false,name=trackingPlanId"`
}

// ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSON - OK
type ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSON struct {
	// Replaces Tracking Plan rules.
	Data *shared.ReplaceRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

// ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON - OK
type ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Replaces Tracking Plan rules.
	Data *shared.ReplaceRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

// ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSON - OK
type ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSON struct {
	// Replaces Tracking Plan rules.
	Data *shared.ReplaceRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

// ReplaceRulesInTrackingPlan200ApplicationJSON - OK
type ReplaceRulesInTrackingPlan200ApplicationJSON struct {
	// Replaces Tracking Plan rules.
	Data *shared.ReplaceRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

type ReplaceRulesInTrackingPlanResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	ReplaceRulesInTrackingPlan200ApplicationJSONObject *ReplaceRulesInTrackingPlan200ApplicationJSON
	// OK
	ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSONObject *ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSON
	// OK
	ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject *ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject *ReplaceRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
}

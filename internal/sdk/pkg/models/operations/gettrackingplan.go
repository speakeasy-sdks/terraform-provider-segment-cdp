// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetTrackingPlanRequest struct {
	TrackingPlanID string `pathParam:"style=simple,explode=false,name=trackingPlanId"`
}

// GetTrackingPlan200ApplicationVndSegmentV1betaPlusJSON - OK
type GetTrackingPlan200ApplicationVndSegmentV1betaPlusJSON struct {
	// Gets a single Tracking Plan.
	Data *shared.GetTrackingPlanV1Output `json:"data,omitempty"`
}

// GetTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Gets a single Tracking Plan.
	Data *shared.GetTrackingPlanV1Output `json:"data,omitempty"`
}

// GetTrackingPlan200ApplicationVndSegmentV1PlusJSON - OK
type GetTrackingPlan200ApplicationVndSegmentV1PlusJSON struct {
	// Gets a single Tracking Plan.
	Data *shared.GetTrackingPlanV1Output `json:"data,omitempty"`
}

// GetTrackingPlan200ApplicationJSON - OK
type GetTrackingPlan200ApplicationJSON struct {
	// Gets a single Tracking Plan.
	Data *shared.GetTrackingPlanV1Output `json:"data,omitempty"`
}

type GetTrackingPlanResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	GetTrackingPlan200ApplicationJSONObject *GetTrackingPlan200ApplicationJSON
	// OK
	GetTrackingPlan200ApplicationVndSegmentV1PlusJSONObject *GetTrackingPlan200ApplicationVndSegmentV1PlusJSON
	// OK
	GetTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject *GetTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	GetTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject *GetTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
}
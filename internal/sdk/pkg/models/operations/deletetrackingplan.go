// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type DeleteTrackingPlanRequest struct {
	TrackingPlanID string `pathParam:"style=simple,explode=false,name=trackingPlanId"`
}

// DeleteTrackingPlan200ApplicationVndSegmentV1betaPlusJSON - OK
type DeleteTrackingPlan200ApplicationVndSegmentV1betaPlusJSON struct {
	// Result of a DeleteTrackingPlan call.
	Data *shared.DeleteTrackingPlanV1Output `json:"data,omitempty"`
}

// DeleteTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON - OK
type DeleteTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Result of a DeleteTrackingPlan call.
	Data *shared.DeleteTrackingPlanV1Output `json:"data,omitempty"`
}

// DeleteTrackingPlan200ApplicationVndSegmentV1PlusJSON - OK
type DeleteTrackingPlan200ApplicationVndSegmentV1PlusJSON struct {
	// Result of a DeleteTrackingPlan call.
	Data *shared.DeleteTrackingPlanV1Output `json:"data,omitempty"`
}

// DeleteTrackingPlan200ApplicationJSON - OK
type DeleteTrackingPlan200ApplicationJSON struct {
	// Result of a DeleteTrackingPlan call.
	Data *shared.DeleteTrackingPlanV1Output `json:"data,omitempty"`
}

type DeleteTrackingPlanResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DeleteTrackingPlan200ApplicationJSONObject *DeleteTrackingPlan200ApplicationJSON
	// OK
	DeleteTrackingPlan200ApplicationVndSegmentV1PlusJSONObject *DeleteTrackingPlan200ApplicationVndSegmentV1PlusJSON
	// OK
	DeleteTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject *DeleteTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	DeleteTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject *DeleteTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
}

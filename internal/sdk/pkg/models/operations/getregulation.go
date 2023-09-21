// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetRegulationRequest struct {
	RegulateID string `pathParam:"style=simple,explode=false,name=regulateId"`
}

// GetRegulation200ApplicationVndSegmentV1betaPlusJSON - OK
type GetRegulation200ApplicationVndSegmentV1betaPlusJSON struct {
	// The regulate request returned.
	Data *shared.GetRegulationV1Output `json:"data,omitempty"`
}

// GetRegulation200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetRegulation200ApplicationVndSegmentV1alphaPlusJSON struct {
	// The regulate request returned.
	Data *shared.GetRegulationV1Output `json:"data,omitempty"`
}

// GetRegulation200ApplicationVndSegmentV1PlusJSON - OK
type GetRegulation200ApplicationVndSegmentV1PlusJSON struct {
	// The regulate request returned.
	Data *shared.GetRegulationV1Output `json:"data,omitempty"`
}

// GetRegulation200ApplicationJSON - OK
type GetRegulation200ApplicationJSON struct {
	// The regulate request returned.
	Data *shared.GetRegulationV1Output `json:"data,omitempty"`
}

type GetRegulationResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	GetRegulation200ApplicationJSONObject *GetRegulation200ApplicationJSON
	// OK
	GetRegulation200ApplicationVndSegmentV1PlusJSONObject *GetRegulation200ApplicationVndSegmentV1PlusJSON
	// OK
	GetRegulation200ApplicationVndSegmentV1alphaPlusJSONObject *GetRegulation200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	GetRegulation200ApplicationVndSegmentV1betaPlusJSONObject *GetRegulation200ApplicationVndSegmentV1betaPlusJSON
}

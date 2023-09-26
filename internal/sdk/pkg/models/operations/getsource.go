// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetSourceRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

// GetSource200ApplicationVndSegmentV1betaPlusJSON - OK
type GetSource200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a Source.
	Data *shared.GetSourceV1Output `json:"data,omitempty"`
}

// GetSource200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetSource200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a Source.
	Data *shared.GetSourceAlphaOutput `json:"data,omitempty"`
}

// GetSource200ApplicationVndSegmentV1PlusJSON - OK
type GetSource200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a Source.
	Data *shared.GetSourceV1Output `json:"data,omitempty"`
}

// GetSource200ApplicationJSON - OK
type GetSource200ApplicationJSON struct {
	// Returns a Source.
	Data *shared.GetSourceV1Output `json:"data,omitempty"`
}

type GetSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	GetSource200ApplicationJSONObject *GetSource200ApplicationJSON
	// OK
	GetSource200ApplicationVndSegmentV1PlusJSONObject *GetSource200ApplicationVndSegmentV1PlusJSON
	// OK
	GetSource200ApplicationVndSegmentV1alphaPlusJSONObject *GetSource200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	GetSource200ApplicationVndSegmentV1betaPlusJSONObject *GetSource200ApplicationVndSegmentV1betaPlusJSON
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetDestinationRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
}

// GetDestination200ApplicationVndSegmentV1betaPlusJSON - OK
type GetDestination200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a single Destination by its id.
	Data *shared.GetDestinationV1Output `json:"data,omitempty"`
}

// GetDestination200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetDestination200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a single Destination by its id.
	Data *shared.GetDestinationV1Output `json:"data,omitempty"`
}

// GetDestination200ApplicationVndSegmentV1PlusJSON - OK
type GetDestination200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a single Destination by its id.
	Data *shared.GetDestinationV1Output `json:"data,omitempty"`
}

// GetDestination200ApplicationJSON - OK
type GetDestination200ApplicationJSON struct {
	// Returns a single Destination by its id.
	Data *shared.GetDestinationV1Output `json:"data,omitempty"`
}

type GetDestinationResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	GetDestination200ApplicationJSONObject *GetDestination200ApplicationJSON
	// OK
	GetDestination200ApplicationVndSegmentV1PlusJSONObject *GetDestination200ApplicationVndSegmentV1PlusJSON
	// OK
	GetDestination200ApplicationVndSegmentV1alphaPlusJSONObject *GetDestination200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	GetDestination200ApplicationVndSegmentV1betaPlusJSONObject *GetDestination200ApplicationVndSegmentV1betaPlusJSON
}

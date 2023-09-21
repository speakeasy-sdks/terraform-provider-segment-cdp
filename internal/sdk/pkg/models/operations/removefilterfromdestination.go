// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type RemoveFilterFromDestinationRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
	FilterID      string `pathParam:"style=simple,explode=false,name=filterId"`
}

// RemoveFilterFromDestination200ApplicationVndSegmentV1betaPlusJSON - OK
type RemoveFilterFromDestination200ApplicationVndSegmentV1betaPlusJSON struct {
	// Output for DeleteDestinationFilterV1.
	Data *shared.RemoveFilterFromDestinationV1Output `json:"data,omitempty"`
}

// RemoveFilterFromDestination200ApplicationVndSegmentV1alphaPlusJSON - OK
type RemoveFilterFromDestination200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Output for DeleteDestinationFilterV1.
	Data *shared.RemoveFilterFromDestinationV1Output `json:"data,omitempty"`
}

// RemoveFilterFromDestination200ApplicationVndSegmentV1PlusJSON - OK
type RemoveFilterFromDestination200ApplicationVndSegmentV1PlusJSON struct {
	// Output for DeleteDestinationFilterV1.
	Data *shared.RemoveFilterFromDestinationV1Output `json:"data,omitempty"`
}

// RemoveFilterFromDestination200ApplicationJSON - OK
type RemoveFilterFromDestination200ApplicationJSON struct {
	// Output for DeleteDestinationFilterV1.
	Data *shared.RemoveFilterFromDestinationV1Output `json:"data,omitempty"`
}

type RemoveFilterFromDestinationResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	RemoveFilterFromDestination200ApplicationJSONObject *RemoveFilterFromDestination200ApplicationJSON
	// OK
	RemoveFilterFromDestination200ApplicationVndSegmentV1PlusJSONObject *RemoveFilterFromDestination200ApplicationVndSegmentV1PlusJSON
	// OK
	RemoveFilterFromDestination200ApplicationVndSegmentV1alphaPlusJSONObject *RemoveFilterFromDestination200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	RemoveFilterFromDestination200ApplicationVndSegmentV1betaPlusJSONObject *RemoveFilterFromDestination200ApplicationVndSegmentV1betaPlusJSON
}

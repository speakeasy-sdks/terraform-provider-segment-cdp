// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type DeleteDestinationRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
}

// DeleteDestination200ApplicationVndSegmentV1betaPlusJSON - OK
type DeleteDestination200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the status of a Destination deletion.
	Data *shared.DeleteDestinationV1Output `json:"data,omitempty"`
}

// DeleteDestination200ApplicationVndSegmentV1alphaPlusJSON - OK
type DeleteDestination200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the status of a Destination deletion.
	Data *shared.DeleteDestinationV1Output `json:"data,omitempty"`
}

// DeleteDestination200ApplicationVndSegmentV1PlusJSON - OK
type DeleteDestination200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the status of a Destination deletion.
	Data *shared.DeleteDestinationV1Output `json:"data,omitempty"`
}

// DeleteDestination200ApplicationJSON - OK
type DeleteDestination200ApplicationJSON struct {
	// Returns the status of a Destination deletion.
	Data *shared.DeleteDestinationV1Output `json:"data,omitempty"`
}

type DeleteDestinationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DeleteDestination200ApplicationJSONObject *DeleteDestination200ApplicationJSON
	// OK
	DeleteDestination200ApplicationVndSegmentV1PlusJSONObject *DeleteDestination200ApplicationVndSegmentV1PlusJSON
	// OK
	DeleteDestination200ApplicationVndSegmentV1alphaPlusJSONObject *DeleteDestination200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	DeleteDestination200ApplicationVndSegmentV1betaPlusJSONObject *DeleteDestination200ApplicationVndSegmentV1betaPlusJSON
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type UpdateFilterForDestinationRequest struct {
	UpdateFilterForDestinationV1Input shared.UpdateFilterForDestinationV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	DestinationID                     string                                   `pathParam:"style=simple,explode=false,name=destinationId"`
	FilterID                          string                                   `pathParam:"style=simple,explode=false,name=filterId"`
}

// UpdateFilterForDestination200ApplicationVndSegmentV1betaPlusJSON - OK
type UpdateFilterForDestination200ApplicationVndSegmentV1betaPlusJSON struct {
	// Output for UpdateDestinationFilterV1.
	Data *shared.UpdateFilterForDestinationV1Output `json:"data,omitempty"`
}

// UpdateFilterForDestination200ApplicationVndSegmentV1alphaPlusJSON - OK
type UpdateFilterForDestination200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Output for UpdateDestinationFilterV1.
	Data *shared.UpdateFilterForDestinationV1Output `json:"data,omitempty"`
}

// UpdateFilterForDestination200ApplicationVndSegmentV1PlusJSON - OK
type UpdateFilterForDestination200ApplicationVndSegmentV1PlusJSON struct {
	// Output for UpdateDestinationFilterV1.
	Data *shared.UpdateFilterForDestinationV1Output `json:"data,omitempty"`
}

// UpdateFilterForDestination200ApplicationJSON - OK
type UpdateFilterForDestination200ApplicationJSON struct {
	// Output for UpdateDestinationFilterV1.
	Data *shared.UpdateFilterForDestinationV1Output `json:"data,omitempty"`
}

type UpdateFilterForDestinationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UpdateFilterForDestination200ApplicationJSONObject *UpdateFilterForDestination200ApplicationJSON
	// OK
	UpdateFilterForDestination200ApplicationVndSegmentV1PlusJSONObject *UpdateFilterForDestination200ApplicationVndSegmentV1PlusJSON
	// OK
	UpdateFilterForDestination200ApplicationVndSegmentV1alphaPlusJSONObject *UpdateFilterForDestination200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	UpdateFilterForDestination200ApplicationVndSegmentV1betaPlusJSONObject *UpdateFilterForDestination200ApplicationVndSegmentV1betaPlusJSON
}
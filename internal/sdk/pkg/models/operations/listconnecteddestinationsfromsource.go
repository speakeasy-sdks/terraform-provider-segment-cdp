// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListConnectedDestinationsFromSourceRequest struct {
	// Required pagination params for the request.
	//
	// This parameter exists in alpha.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	SourceID   string                 `pathParam:"style=simple,explode=false,name=sourceId"`
}

// ListConnectedDestinationsFromSource200ApplicationVndSegmentV1betaPlusJSON - OK
type ListConnectedDestinationsFromSource200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a list of Destinations connected to a Source.
	Data *shared.ListConnectedDestinationsFromSourceV1Output `json:"data,omitempty"`
}

// ListConnectedDestinationsFromSource200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListConnectedDestinationsFromSource200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a list of Destinations connected to a Source.
	Data *shared.ListConnectedDestinationsFromSourceAlphaOutput `json:"data,omitempty"`
}

// ListConnectedDestinationsFromSource200ApplicationVndSegmentV1PlusJSON - OK
type ListConnectedDestinationsFromSource200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a list of Destinations connected to a Source.
	Data *shared.ListConnectedDestinationsFromSourceV1Output `json:"data,omitempty"`
}

// ListConnectedDestinationsFromSource200ApplicationJSON - OK
type ListConnectedDestinationsFromSource200ApplicationJSON struct {
	// Returns a list of Destinations connected to a Source.
	Data *shared.ListConnectedDestinationsFromSourceV1Output `json:"data,omitempty"`
}

type ListConnectedDestinationsFromSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ListConnectedDestinationsFromSource200ApplicationJSONObject *ListConnectedDestinationsFromSource200ApplicationJSON
	// OK
	ListConnectedDestinationsFromSource200ApplicationVndSegmentV1PlusJSONObject *ListConnectedDestinationsFromSource200ApplicationVndSegmentV1PlusJSON
	// OK
	ListConnectedDestinationsFromSource200ApplicationVndSegmentV1alphaPlusJSONObject *ListConnectedDestinationsFromSource200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListConnectedDestinationsFromSource200ApplicationVndSegmentV1betaPlusJSONObject *ListConnectedDestinationsFromSource200ApplicationVndSegmentV1betaPlusJSON
}
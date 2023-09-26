// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

// ListLabels200ApplicationVndSegmentV1betaPlusJSON - OK
type ListLabels200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns all available labels for the current Workspace.
	Data *shared.ListLabelsV1Output `json:"data,omitempty"`
}

// ListLabels200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListLabels200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns all available labels for the current Workspace.
	Data *shared.ListLabelsAlphaOutput `json:"data,omitempty"`
}

// ListLabels200ApplicationVndSegmentV1PlusJSON - OK
type ListLabels200ApplicationVndSegmentV1PlusJSON struct {
	// Returns all available labels for the current Workspace.
	Data *shared.ListLabelsV1Output `json:"data,omitempty"`
}

// ListLabels200ApplicationJSON - OK
type ListLabels200ApplicationJSON struct {
	// Returns all available labels for the current Workspace.
	Data *shared.ListLabelsV1Output `json:"data,omitempty"`
}

type ListLabelsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ListLabels200ApplicationJSONObject *ListLabels200ApplicationJSON
	// OK
	ListLabels200ApplicationVndSegmentV1PlusJSONObject *ListLabels200ApplicationVndSegmentV1PlusJSON
	// OK
	ListLabels200ApplicationVndSegmentV1alphaPlusJSONObject *ListLabels200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListLabels200ApplicationVndSegmentV1betaPlusJSONObject *ListLabels200ApplicationVndSegmentV1betaPlusJSON
}

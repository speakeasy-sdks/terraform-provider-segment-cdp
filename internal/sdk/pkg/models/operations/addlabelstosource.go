// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type AddLabelsToSourceRequest struct {
	AddLabelsToSourceV1Input shared.AddLabelsToSourceV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	SourceID                 string                          `pathParam:"style=simple,explode=false,name=sourceId"`
}

// AddLabelsToSource200ApplicationVndSegmentV1betaPlusJSON - OK
type AddLabelsToSource200ApplicationVndSegmentV1betaPlusJSON struct {
	// Applies an existing label to an existing Source.
	Data *shared.AddLabelsToSourceV1Output `json:"data,omitempty"`
}

// AddLabelsToSource200ApplicationVndSegmentV1alphaPlusJSON - OK
type AddLabelsToSource200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Applies an existing label to an existing Source.
	Data *shared.AddLabelsToSourceAlphaOutput `json:"data,omitempty"`
}

// AddLabelsToSource200ApplicationVndSegmentV1PlusJSON - OK
type AddLabelsToSource200ApplicationVndSegmentV1PlusJSON struct {
	// Applies an existing label to an existing Source.
	Data *shared.AddLabelsToSourceV1Output `json:"data,omitempty"`
}

// AddLabelsToSource200ApplicationJSON - OK
type AddLabelsToSource200ApplicationJSON struct {
	// Applies an existing label to an existing Source.
	Data *shared.AddLabelsToSourceV1Output `json:"data,omitempty"`
}

type AddLabelsToSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	AddLabelsToSource200ApplicationJSONObject *AddLabelsToSource200ApplicationJSON
	// OK
	AddLabelsToSource200ApplicationVndSegmentV1PlusJSONObject *AddLabelsToSource200ApplicationVndSegmentV1PlusJSON
	// OK
	AddLabelsToSource200ApplicationVndSegmentV1alphaPlusJSONObject *AddLabelsToSource200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	AddLabelsToSource200ApplicationVndSegmentV1betaPlusJSONObject *AddLabelsToSource200ApplicationVndSegmentV1betaPlusJSON
}

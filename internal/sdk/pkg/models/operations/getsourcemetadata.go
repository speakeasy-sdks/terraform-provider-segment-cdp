// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetSourceMetadataRequest struct {
	SourceMetadataID string `pathParam:"style=simple,explode=false,name=sourceMetadataId"`
}

// GetSourceMetadata200ApplicationVndSegmentV1betaPlusJSON - OK
type GetSourceMetadata200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the Source catalog item looked up by id.
	Data *shared.GetSourceMetadataV1Output `json:"data,omitempty"`
}

// GetSourceMetadata200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetSourceMetadata200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the Source catalog item looked up by id.
	Data *shared.GetSourceMetadataV1Output `json:"data,omitempty"`
}

// GetSourceMetadata200ApplicationVndSegmentV1PlusJSON - OK
type GetSourceMetadata200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the Source catalog item looked up by id.
	Data *shared.GetSourceMetadataV1Output `json:"data,omitempty"`
}

// GetSourceMetadata200ApplicationJSON - OK
type GetSourceMetadata200ApplicationJSON struct {
	// Returns the Source catalog item looked up by id.
	Data *shared.GetSourceMetadataV1Output `json:"data,omitempty"`
}

type GetSourceMetadataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	GetSourceMetadata200ApplicationJSONObject *GetSourceMetadata200ApplicationJSON
	// OK
	GetSourceMetadata200ApplicationVndSegmentV1PlusJSONObject *GetSourceMetadata200ApplicationVndSegmentV1PlusJSON
	// OK
	GetSourceMetadata200ApplicationVndSegmentV1alphaPlusJSONObject *GetSourceMetadata200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	GetSourceMetadata200ApplicationVndSegmentV1betaPlusJSONObject *GetSourceMetadata200ApplicationVndSegmentV1betaPlusJSON
}
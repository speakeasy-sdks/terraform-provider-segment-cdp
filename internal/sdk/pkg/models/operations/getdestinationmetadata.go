// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetDestinationMetadataRequest struct {
	DestinationMetadataID string `pathParam:"style=simple,explode=false,name=destinationMetadataId"`
}

// GetDestinationMetadata200ApplicationVndSegmentV1betaPlusJSON - OK
type GetDestinationMetadata200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a Destination catalog item.
	Data *shared.GetDestinationMetadataV1Output `json:"data,omitempty"`
}

// GetDestinationMetadata200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetDestinationMetadata200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a Destination catalog item.
	Data *shared.GetDestinationMetadataV1Output `json:"data,omitempty"`
}

// GetDestinationMetadata200ApplicationVndSegmentV1PlusJSON - OK
type GetDestinationMetadata200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a Destination catalog item.
	Data *shared.GetDestinationMetadataV1Output `json:"data,omitempty"`
}

// GetDestinationMetadata200ApplicationJSON - OK
type GetDestinationMetadata200ApplicationJSON struct {
	// Returns a Destination catalog item.
	Data *shared.GetDestinationMetadataV1Output `json:"data,omitempty"`
}

type GetDestinationMetadataResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	GetDestinationMetadata200ApplicationJSONObject *GetDestinationMetadata200ApplicationJSON
	// OK
	GetDestinationMetadata200ApplicationVndSegmentV1PlusJSONObject *GetDestinationMetadata200ApplicationVndSegmentV1PlusJSON
	// OK
	GetDestinationMetadata200ApplicationVndSegmentV1alphaPlusJSONObject *GetDestinationMetadata200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	GetDestinationMetadata200ApplicationVndSegmentV1betaPlusJSONObject *GetDestinationMetadata200ApplicationVndSegmentV1betaPlusJSON
}
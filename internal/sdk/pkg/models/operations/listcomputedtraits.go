// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListComputedTraitsRequest struct {
	// Information about the pagination of this response.
	//
	// This parameter exists in alpha.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	SpaceID    string                 `pathParam:"style=simple,explode=false,name=spaceId"`
}

// ListComputedTraits200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListComputedTraits200ApplicationVndSegmentV1alphaPlusJSON struct {
	// List computed traits endpoint output.
	Data *shared.ListComputedTraitsAlphaOutput `json:"data,omitempty"`
}

type ListComputedTraitsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	ListComputedTraits200ApplicationVndSegmentV1alphaPlusJSONObject *ListComputedTraits200ApplicationVndSegmentV1alphaPlusJSON
}

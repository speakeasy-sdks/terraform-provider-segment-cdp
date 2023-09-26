// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListAudiencesRequest struct {
	// Information about the pagination of this response.
	//
	// This parameter exists in alpha.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	SpaceID    string                 `pathParam:"style=simple,explode=false,name=spaceId"`
}

// ListAudiences200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListAudiences200ApplicationVndSegmentV1alphaPlusJSON struct {
	// List audiences endpoint output.
	Data *shared.ListAudiencesAlphaOutput `json:"data,omitempty"`
}

type ListAudiencesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	ListAudiences200ApplicationVndSegmentV1alphaPlusJSONObject *ListAudiences200ApplicationVndSegmentV1alphaPlusJSON
}

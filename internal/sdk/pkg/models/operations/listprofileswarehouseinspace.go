// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListProfilesWarehouseInSpaceRequest struct {
	// Defines the pagination parameters.
	//
	// This parameter exists in alpha.
	Pagination *shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	SpaceID    string                  `pathParam:"style=simple,explode=false,name=spaceId"`
}

// ListProfilesWarehouseInSpace200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListProfilesWarehouseInSpace200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns all Profiles Warehouse based on spaceID and pagination.
	Data *shared.ListProfilesWarehouseInSpaceAlphaOutput `json:"data,omitempty"`
}

type ListProfilesWarehouseInSpaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	ListProfilesWarehouseInSpace200ApplicationVndSegmentV1alphaPlusJSONObject *ListProfilesWarehouseInSpace200ApplicationVndSegmentV1alphaPlusJSON
}
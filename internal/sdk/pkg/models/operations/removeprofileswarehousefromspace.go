// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type RemoveProfilesWarehouseFromSpaceRequest struct {
	SpaceID     string `pathParam:"style=simple,explode=false,name=spaceId"`
	WarehouseID string `pathParam:"style=simple,explode=false,name=warehouseId"`
}

// RemoveProfilesWarehouseFromSpace200ApplicationVndSegmentV1alphaPlusJSON - OK
type RemoveProfilesWarehouseFromSpace200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the status of a Warehouse deletion.
	Data *shared.RemoveProfilesWarehouseFromSpaceAlphaOutput `json:"data,omitempty"`
}

type RemoveProfilesWarehouseFromSpaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	RemoveProfilesWarehouseFromSpace200ApplicationVndSegmentV1alphaPlusJSONObject *RemoveProfilesWarehouseFromSpace200ApplicationVndSegmentV1alphaPlusJSON
}
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetSpaceRequest struct {
	SpaceID string `pathParam:"style=simple,explode=false,name=spaceId"`
}

// GetSpace200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetSpace200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Response for the getSpaceById endpoint.
	Data *shared.GetSpaceAlphaOutput `json:"data,omitempty"`
}

type GetSpaceResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	GetSpace200ApplicationVndSegmentV1alphaPlusJSONObject *GetSpace200ApplicationVndSegmentV1alphaPlusJSON
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetLatestFromEdgeFunctionsRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

// GetLatestFromEdgeFunctions200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetLatestFromEdgeFunctions200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Output for GetLatestFromEdgeFunctions.
	Data *shared.GetLatestFromEdgeFunctionsAlphaOutput `json:"data,omitempty"`
}

type GetLatestFromEdgeFunctionsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	GetLatestFromEdgeFunctions200ApplicationVndSegmentV1alphaPlusJSONObject *GetLatestFromEdgeFunctions200ApplicationVndSegmentV1alphaPlusJSON
}

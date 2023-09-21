// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListSubscriptionsFromDestinationRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
	// Pagination options.
	//
	// This parameter exists in alpha.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

// ListSubscriptionsFromDestination200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListSubscriptionsFromDestination200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Output for ListDestinationSubscriptionsAlpha.
	Data *shared.ListSubscriptionsFromDestinationAlphaOutput `json:"data,omitempty"`
}

type ListSubscriptionsFromDestinationResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	ListSubscriptionsFromDestination200ApplicationVndSegmentV1alphaPlusJSONObject *ListSubscriptionsFromDestination200ApplicationVndSegmentV1alphaPlusJSON
}

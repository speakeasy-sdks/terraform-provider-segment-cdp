// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type CreateDestinationSubscriptionRequest struct {
	CreateDestinationSubscriptionAlphaInput shared.CreateDestinationSubscriptionAlphaInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	DestinationID                           string                                         `pathParam:"style=simple,explode=false,name=destinationId"`
}

// CreateDestinationSubscription200ApplicationVndSegmentV1alphaPlusJSON - OK
type CreateDestinationSubscription200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a newly created Destination subscription.
	Data *shared.CreateDestinationSubscriptionAlphaOutput `json:"data,omitempty"`
}

type CreateDestinationSubscriptionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	CreateDestinationSubscription200ApplicationVndSegmentV1alphaPlusJSONObject *CreateDestinationSubscription200ApplicationVndSegmentV1alphaPlusJSON
}

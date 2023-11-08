// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateDestinationSubscriptionRequest struct {
	CreateDestinationSubscriptionAlphaInput shared.CreateDestinationSubscriptionAlphaInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	DestinationID                           string                                         `pathParam:"style=simple,explode=false,name=destinationId"`
}

func (o *CreateDestinationSubscriptionRequest) GetCreateDestinationSubscriptionAlphaInput() shared.CreateDestinationSubscriptionAlphaInput {
	if o == nil {
		return shared.CreateDestinationSubscriptionAlphaInput{}
	}
	return o.CreateDestinationSubscriptionAlphaInput
}

func (o *CreateDestinationSubscriptionRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

// CreateDestinationSubscriptionResponseBody - OK
type CreateDestinationSubscriptionResponseBody struct {
	// Returns a newly created Destination subscription.
	Data *shared.CreateDestinationSubscriptionAlphaOutput `json:"data,omitempty"`
}

func (o *CreateDestinationSubscriptionResponseBody) GetData() *shared.CreateDestinationSubscriptionAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
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
	Object *CreateDestinationSubscriptionResponseBody
}

func (o *CreateDestinationSubscriptionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateDestinationSubscriptionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateDestinationSubscriptionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateDestinationSubscriptionResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *CreateDestinationSubscriptionResponse) GetObject() *CreateDestinationSubscriptionResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

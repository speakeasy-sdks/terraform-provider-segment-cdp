// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"net/http"
)

type RemoveSubscriptionFromDestinationRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
	ID            string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *RemoveSubscriptionFromDestinationRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *RemoveSubscriptionFromDestinationRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// RemoveSubscriptionFromDestinationResponseBody - OK
type RemoveSubscriptionFromDestinationResponseBody struct {
	// Returns a Destination deletion flag.
	Data *shared.RemoveSubscriptionFromDestinationAlphaOutput `json:"data,omitempty"`
}

func (o *RemoveSubscriptionFromDestinationResponseBody) GetData() *shared.RemoveSubscriptionFromDestinationAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type RemoveSubscriptionFromDestinationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *RemoveSubscriptionFromDestinationResponseBody
}

func (o *RemoveSubscriptionFromDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveSubscriptionFromDestinationResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *RemoveSubscriptionFromDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveSubscriptionFromDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *RemoveSubscriptionFromDestinationResponse) GetObject() *RemoveSubscriptionFromDestinationResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

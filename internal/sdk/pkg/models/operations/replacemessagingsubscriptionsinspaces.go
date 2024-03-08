// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"net/http"
)

type ReplaceMessagingSubscriptionsInSpacesRequest struct {
	ReplaceMessagingSubscriptionsInSpacesAlphaInput shared.ReplaceMessagingSubscriptionsInSpacesAlphaInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	SpaceID                                         string                                                 `pathParam:"style=simple,explode=false,name=spaceId"`
}

func (o *ReplaceMessagingSubscriptionsInSpacesRequest) GetReplaceMessagingSubscriptionsInSpacesAlphaInput() shared.ReplaceMessagingSubscriptionsInSpacesAlphaInput {
	if o == nil {
		return shared.ReplaceMessagingSubscriptionsInSpacesAlphaInput{}
	}
	return o.ReplaceMessagingSubscriptionsInSpacesAlphaInput
}

func (o *ReplaceMessagingSubscriptionsInSpacesRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

// ReplaceMessagingSubscriptionsInSpacesResponseBody - OK
type ReplaceMessagingSubscriptionsInSpacesResponseBody struct {
	// Output for the endpoint.
	Data *shared.ReplaceMessagingSubscriptionsInSpacesAlphaOutput `json:"data,omitempty"`
}

func (o *ReplaceMessagingSubscriptionsInSpacesResponseBody) GetData() *shared.ReplaceMessagingSubscriptionsInSpacesAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type ReplaceMessagingSubscriptionsInSpacesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *ReplaceMessagingSubscriptionsInSpacesResponseBody
}

func (o *ReplaceMessagingSubscriptionsInSpacesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReplaceMessagingSubscriptionsInSpacesResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ReplaceMessagingSubscriptionsInSpacesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReplaceMessagingSubscriptionsInSpacesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ReplaceMessagingSubscriptionsInSpacesResponse) GetObject() *ReplaceMessagingSubscriptionsInSpacesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/models/shared"
	"net/http"
)

type GetSpaceRequest struct {
	SpaceID string `pathParam:"style=simple,explode=false,name=spaceId"`
}

func (o *GetSpaceRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

// GetSpaceResponseBody - OK
type GetSpaceResponseBody struct {
	// Response for the getSpaceById endpoint.
	Data *shared.GetSpaceAlphaOutput `json:"data,omitempty"`
}

func (o *GetSpaceResponseBody) GetData() *shared.GetSpaceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetSpaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *GetSpaceResponseBody
}

func (o *GetSpaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSpaceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GetSpaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSpaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSpaceResponse) GetObject() *GetSpaceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
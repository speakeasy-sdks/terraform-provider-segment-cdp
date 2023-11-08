// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetComputedTraitRequest struct {
	ID      string `pathParam:"style=simple,explode=false,name=id"`
	SpaceID string `pathParam:"style=simple,explode=false,name=spaceId"`
}

func (o *GetComputedTraitRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetComputedTraitRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

// GetComputedTraitResponseBody - OK
type GetComputedTraitResponseBody struct {
	// Computed Trait output for get and update.
	Data *shared.GetComputedTraitAlphaOutput `json:"data,omitempty"`
}

func (o *GetComputedTraitResponseBody) GetData() *shared.GetComputedTraitAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetComputedTraitResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	Object *GetComputedTraitResponseBody
}

func (o *GetComputedTraitResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetComputedTraitResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetComputedTraitResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetComputedTraitResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GetComputedTraitResponse) GetObject() *GetComputedTraitResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

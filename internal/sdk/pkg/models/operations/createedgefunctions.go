// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateEdgeFunctionsRequest struct {
	CreateEdgeFunctionsAlphaInput shared.CreateEdgeFunctionsAlphaInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	SourceID                      string                               `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *CreateEdgeFunctionsRequest) GetCreateEdgeFunctionsAlphaInput() shared.CreateEdgeFunctionsAlphaInput {
	if o == nil {
		return shared.CreateEdgeFunctionsAlphaInput{}
	}
	return o.CreateEdgeFunctionsAlphaInput
}

func (o *CreateEdgeFunctionsRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

// CreateEdgeFunctionsResponseBody - OK
type CreateEdgeFunctionsResponseBody struct {
	// Output for CreateEdgeFunctions.
	Data *shared.CreateEdgeFunctionsAlphaOutput `json:"data,omitempty"`
}

func (o *CreateEdgeFunctionsResponseBody) GetData() *shared.CreateEdgeFunctionsAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateEdgeFunctionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	Object *CreateEdgeFunctionsResponseBody
}

func (o *CreateEdgeFunctionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateEdgeFunctionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateEdgeFunctionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateEdgeFunctionsResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *CreateEdgeFunctionsResponse) GetObject() *CreateEdgeFunctionsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

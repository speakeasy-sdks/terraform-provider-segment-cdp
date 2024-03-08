// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"net/http"
)

type DeleteReverseEtlModelRequest struct {
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

func (o *DeleteReverseEtlModelRequest) GetModelID() string {
	if o == nil {
		return ""
	}
	return o.ModelID
}

// DeleteReverseEtlModelResponseBody - OK
type DeleteReverseEtlModelResponseBody struct {
	// Defines the result of getting a Model.
	Data *shared.DeleteReverseEtlModelOutput `json:"data,omitempty"`
}

func (o *DeleteReverseEtlModelResponseBody) GetData() *shared.DeleteReverseEtlModelOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type DeleteReverseEtlModelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *DeleteReverseEtlModelResponseBody
}

func (o *DeleteReverseEtlModelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteReverseEtlModelResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *DeleteReverseEtlModelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteReverseEtlModelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteReverseEtlModelResponse) GetObject() *DeleteReverseEtlModelResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

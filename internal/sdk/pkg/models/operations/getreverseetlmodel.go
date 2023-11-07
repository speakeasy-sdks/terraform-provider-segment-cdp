// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetReverseEtlModelRequest struct {
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

func (o *GetReverseEtlModelRequest) GetModelID() string {
	if o == nil {
		return ""
	}
	return o.ModelID
}

// GetReverseEtlModelResponseBody - OK
type GetReverseEtlModelResponseBody struct {
	// Defines the result of getting a Model.
	Data *shared.GetReverseEtlModelOutput `json:"data,omitempty"`
}

func (o *GetReverseEtlModelResponseBody) GetData() *shared.GetReverseEtlModelOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetReverseEtlModelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	Object *GetReverseEtlModelResponseBody
}

func (o *GetReverseEtlModelResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetReverseEtlModelResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetReverseEtlModelResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetReverseEtlModelResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GetReverseEtlModelResponse) GetObject() *GetReverseEtlModelResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

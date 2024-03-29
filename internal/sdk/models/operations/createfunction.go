// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/models/shared"
	"net/http"
)

// CreateFunctionFunctionsResponse200ResponseBody - OK
type CreateFunctionFunctionsResponse200ResponseBody struct {
	// Create a Function.
	Data *shared.CreateFunctionV1Output `json:"data,omitempty"`
}

func (o *CreateFunctionFunctionsResponse200ResponseBody) GetData() *shared.CreateFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateFunctionFunctionsResponseResponseBody - OK
type CreateFunctionFunctionsResponseResponseBody struct {
	// Create a Function.
	Data *shared.CreateFunctionV1Output `json:"data,omitempty"`
}

func (o *CreateFunctionFunctionsResponseResponseBody) GetData() *shared.CreateFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateFunctionFunctionsResponseBody - OK
type CreateFunctionFunctionsResponseBody struct {
	// Create a Function.
	Data *shared.CreateFunctionV1Output `json:"data,omitempty"`
}

func (o *CreateFunctionFunctionsResponseBody) GetData() *shared.CreateFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateFunctionResponseBody - OK
type CreateFunctionResponseBody struct {
	// Create a Function.
	Data *shared.CreateFunctionV1Output `json:"data,omitempty"`
}

func (o *CreateFunctionResponseBody) GetData() *shared.CreateFunctionV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateFunctionResponse struct {
	// OK
	TwoHundredApplicationJSONObject *CreateFunctionResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *CreateFunctionFunctionsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *CreateFunctionFunctionsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *CreateFunctionFunctionsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateFunctionResponse) GetTwoHundredApplicationJSONObject() *CreateFunctionResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *CreateFunctionResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *CreateFunctionFunctionsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *CreateFunctionResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *CreateFunctionFunctionsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *CreateFunctionResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *CreateFunctionFunctionsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *CreateFunctionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateFunctionResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *CreateFunctionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateFunctionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

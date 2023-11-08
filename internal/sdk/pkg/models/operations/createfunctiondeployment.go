// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateFunctionDeploymentRequest struct {
	FunctionID string `pathParam:"style=simple,explode=false,name=functionId"`
}

func (o *CreateFunctionDeploymentRequest) GetFunctionID() string {
	if o == nil {
		return ""
	}
	return o.FunctionID
}

// CreateFunctionDeploymentFunctionsResponse200ResponseBody - OK
type CreateFunctionDeploymentFunctionsResponse200ResponseBody struct {
	// Updates the deployment for a Source Function instance.
	Data *shared.CreateFunctionDeploymentV1Output `json:"data,omitempty"`
}

func (o *CreateFunctionDeploymentFunctionsResponse200ResponseBody) GetData() *shared.CreateFunctionDeploymentV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateFunctionDeploymentFunctionsResponseResponseBody - OK
type CreateFunctionDeploymentFunctionsResponseResponseBody struct {
	// Updates the deployment for a Source Function instance.
	Data *shared.CreateFunctionDeploymentV1Output `json:"data,omitempty"`
}

func (o *CreateFunctionDeploymentFunctionsResponseResponseBody) GetData() *shared.CreateFunctionDeploymentV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateFunctionDeploymentFunctionsResponseBody - OK
type CreateFunctionDeploymentFunctionsResponseBody struct {
	// Updates the deployment for a Source Function instance.
	Data *shared.CreateFunctionDeploymentV1Output `json:"data,omitempty"`
}

func (o *CreateFunctionDeploymentFunctionsResponseBody) GetData() *shared.CreateFunctionDeploymentV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateFunctionDeploymentResponseBody - OK
type CreateFunctionDeploymentResponseBody struct {
	// Updates the deployment for a Source Function instance.
	Data *shared.CreateFunctionDeploymentV1Output `json:"data,omitempty"`
}

func (o *CreateFunctionDeploymentResponseBody) GetData() *shared.CreateFunctionDeploymentV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateFunctionDeploymentResponse struct {
	// OK
	TwoHundredApplicationJSONObject *CreateFunctionDeploymentResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *CreateFunctionDeploymentFunctionsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *CreateFunctionDeploymentFunctionsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *CreateFunctionDeploymentFunctionsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateFunctionDeploymentResponse) GetTwoHundredApplicationJSONObject() *CreateFunctionDeploymentResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *CreateFunctionDeploymentResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *CreateFunctionDeploymentFunctionsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *CreateFunctionDeploymentResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *CreateFunctionDeploymentFunctionsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *CreateFunctionDeploymentResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *CreateFunctionDeploymentFunctionsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *CreateFunctionDeploymentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateFunctionDeploymentResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *CreateFunctionDeploymentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateFunctionDeploymentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

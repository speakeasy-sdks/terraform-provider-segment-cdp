// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetFunctionVersionRequest struct {
	FunctionID string `pathParam:"style=simple,explode=false,name=functionId"`
	VersionID  string `pathParam:"style=simple,explode=false,name=versionId"`
}

func (o *GetFunctionVersionRequest) GetFunctionID() string {
	if o == nil {
		return ""
	}
	return o.FunctionID
}

func (o *GetFunctionVersionRequest) GetVersionID() string {
	if o == nil {
		return ""
	}
	return o.VersionID
}

// GetFunctionVersionResponseBody - OK
type GetFunctionVersionResponseBody struct {
	// Get Function version output.
	Data *shared.GetFunctionVersionAlphaOutput `json:"data,omitempty"`
}

func (o *GetFunctionVersionResponseBody) GetData() *shared.GetFunctionVersionAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetFunctionVersionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	Object *GetFunctionVersionResponseBody
}

func (o *GetFunctionVersionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFunctionVersionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFunctionVersionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetFunctionVersionResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GetFunctionVersionResponse) GetObject() *GetFunctionVersionResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

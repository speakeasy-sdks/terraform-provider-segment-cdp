// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/models/shared"
	"net/http"
)

type ListFunctionVersionsRequest struct {
	FunctionID string `pathParam:"style=simple,explode=false,name=functionId"`
	// Pagination parameters.
	//
	// This parameter exists in alpha.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

func (o *ListFunctionVersionsRequest) GetFunctionID() string {
	if o == nil {
		return ""
	}
	return o.FunctionID
}

func (o *ListFunctionVersionsRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

// ListFunctionVersionsResponseBody - OK
type ListFunctionVersionsResponseBody struct {
	// Lists Versions of a Function.
	Data *shared.ListFunctionVersionsAlphaOutput `json:"data,omitempty"`
}

func (o *ListFunctionVersionsResponseBody) GetData() *shared.ListFunctionVersionsAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListFunctionVersionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *ListFunctionVersionsResponseBody
}

func (o *ListFunctionVersionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListFunctionVersionsResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListFunctionVersionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListFunctionVersionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListFunctionVersionsResponse) GetObject() *ListFunctionVersionsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

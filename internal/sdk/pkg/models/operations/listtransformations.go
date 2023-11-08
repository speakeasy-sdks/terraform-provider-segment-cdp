// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListTransformationsRequest struct {
	// Pagination options.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

func (o *ListTransformationsRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

// ListTransformationsTransformationsResponse200ResponseBody - OK
type ListTransformationsTransformationsResponse200ResponseBody struct {
	// Lists the Transformations associated with the current Workspace.
	Data *shared.ListTransformationsV1Output `json:"data,omitempty"`
}

func (o *ListTransformationsTransformationsResponse200ResponseBody) GetData() *shared.ListTransformationsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListTransformationsTransformationsResponseResponseBody - OK
type ListTransformationsTransformationsResponseResponseBody struct {
	// Lists the Transformations associated with the current Workspace.
	Data *shared.ListTransformationsV1Output `json:"data,omitempty"`
}

func (o *ListTransformationsTransformationsResponseResponseBody) GetData() *shared.ListTransformationsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListTransformationsTransformationsResponseBody - OK
type ListTransformationsTransformationsResponseBody struct {
	// Lists the Transformations associated with the current Workspace.
	Data *shared.ListTransformationsV1Output `json:"data,omitempty"`
}

func (o *ListTransformationsTransformationsResponseBody) GetData() *shared.ListTransformationsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListTransformationsResponseBody - OK
type ListTransformationsResponseBody struct {
	// Lists the Transformations associated with the current Workspace.
	Data *shared.ListTransformationsV1Output `json:"data,omitempty"`
}

func (o *ListTransformationsResponseBody) GetData() *shared.ListTransformationsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListTransformationsResponse struct {
	// OK
	TwoHundredApplicationJSONObject *ListTransformationsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *ListTransformationsTransformationsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *ListTransformationsTransformationsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *ListTransformationsTransformationsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListTransformationsResponse) GetTwoHundredApplicationJSONObject() *ListTransformationsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *ListTransformationsResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *ListTransformationsTransformationsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *ListTransformationsResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *ListTransformationsTransformationsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *ListTransformationsResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *ListTransformationsTransformationsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *ListTransformationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListTransformationsResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListTransformationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListTransformationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

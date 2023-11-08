// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"net/http"
)

type ListDestinationsRequest struct {
	// Required pagination params for the request.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

func (o *ListDestinationsRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

// ListDestinationsDestinationsResponse200ResponseBody - OK
type ListDestinationsDestinationsResponse200ResponseBody struct {
	// Returns all Destinations connected to the current Workspace.
	Data *shared.ListDestinationsV1Output `json:"data,omitempty"`
}

func (o *ListDestinationsDestinationsResponse200ResponseBody) GetData() *shared.ListDestinationsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListDestinationsDestinationsResponseResponseBody - OK
type ListDestinationsDestinationsResponseResponseBody struct {
	// Returns all Destinations connected to the current Workspace.
	Data *shared.ListDestinationsV1Output `json:"data,omitempty"`
}

func (o *ListDestinationsDestinationsResponseResponseBody) GetData() *shared.ListDestinationsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListDestinationsDestinationsResponseBody - OK
type ListDestinationsDestinationsResponseBody struct {
	// Returns all Destinations connected to the current Workspace.
	Data *shared.ListDestinationsV1Output `json:"data,omitempty"`
}

func (o *ListDestinationsDestinationsResponseBody) GetData() *shared.ListDestinationsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListDestinationsResponseBody - OK
type ListDestinationsResponseBody struct {
	// Returns all Destinations connected to the current Workspace.
	Data *shared.ListDestinationsV1Output `json:"data,omitempty"`
}

func (o *ListDestinationsResponseBody) GetData() *shared.ListDestinationsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListDestinationsResponse struct {
	// OK
	TwoHundredApplicationJSONObject *ListDestinationsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *ListDestinationsDestinationsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *ListDestinationsDestinationsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *ListDestinationsDestinationsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListDestinationsResponse) GetTwoHundredApplicationJSONObject() *ListDestinationsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *ListDestinationsResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *ListDestinationsDestinationsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *ListDestinationsResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *ListDestinationsDestinationsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *ListDestinationsResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *ListDestinationsDestinationsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *ListDestinationsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListDestinationsResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListDestinationsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListDestinationsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

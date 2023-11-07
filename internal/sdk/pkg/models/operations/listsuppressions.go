// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListSuppressionsRequest struct {
	// Pagination parameters.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

func (o *ListSuppressionsRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

// ListSuppressionsDeletionAndSuppressionResponse200ResponseBody - OK
type ListSuppressionsDeletionAndSuppressionResponse200ResponseBody struct {
	// The output of a list suppressed call for a Workspace.
	Data *shared.ListSuppressionsV1Output `json:"data,omitempty"`
}

func (o *ListSuppressionsDeletionAndSuppressionResponse200ResponseBody) GetData() *shared.ListSuppressionsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListSuppressionsDeletionAndSuppressionResponseResponseBody - OK
type ListSuppressionsDeletionAndSuppressionResponseResponseBody struct {
	// The output of a list suppressed call for a Workspace.
	Data *shared.ListSuppressionsV1Output `json:"data,omitempty"`
}

func (o *ListSuppressionsDeletionAndSuppressionResponseResponseBody) GetData() *shared.ListSuppressionsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListSuppressionsDeletionAndSuppressionResponseBody - OK
type ListSuppressionsDeletionAndSuppressionResponseBody struct {
	// The output of a list suppressed call for a Workspace.
	Data *shared.ListSuppressionsV1Output `json:"data,omitempty"`
}

func (o *ListSuppressionsDeletionAndSuppressionResponseBody) GetData() *shared.ListSuppressionsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListSuppressionsResponseBody - OK
type ListSuppressionsResponseBody struct {
	// The output of a list suppressed call for a Workspace.
	Data *shared.ListSuppressionsV1Output `json:"data,omitempty"`
}

func (o *ListSuppressionsResponseBody) GetData() *shared.ListSuppressionsV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListSuppressionsResponse struct {
	// OK
	TwoHundredApplicationJSONObject *ListSuppressionsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *ListSuppressionsDeletionAndSuppressionResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *ListSuppressionsDeletionAndSuppressionResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *ListSuppressionsDeletionAndSuppressionResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListSuppressionsResponse) GetTwoHundredApplicationJSONObject() *ListSuppressionsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *ListSuppressionsResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *ListSuppressionsDeletionAndSuppressionResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *ListSuppressionsResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *ListSuppressionsDeletionAndSuppressionResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *ListSuppressionsResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *ListSuppressionsDeletionAndSuppressionResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *ListSuppressionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListSuppressionsResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListSuppressionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListSuppressionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/models/shared"
	"net/http"
)

type GetSourcesCatalogRequest struct {
	// Defines the pagination parameters.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

func (o *GetSourcesCatalogRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

// GetSourcesCatalogCatalogResponse200ResponseBody - OK
type GetSourcesCatalogCatalogResponse200ResponseBody struct {
	// Returns a list of all Source catalog items contained within a given page.
	Data *shared.GetSourcesCatalogV1Output `json:"data,omitempty"`
}

func (o *GetSourcesCatalogCatalogResponse200ResponseBody) GetData() *shared.GetSourcesCatalogV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetSourcesCatalogCatalogResponseResponseBody - OK
type GetSourcesCatalogCatalogResponseResponseBody struct {
	// Returns a list of all Source catalog items contained within a given page.
	Data *shared.GetSourcesCatalogV1Output `json:"data,omitempty"`
}

func (o *GetSourcesCatalogCatalogResponseResponseBody) GetData() *shared.GetSourcesCatalogV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetSourcesCatalogCatalogResponseBody - OK
type GetSourcesCatalogCatalogResponseBody struct {
	// Returns a list of all Source catalog items contained within a given page.
	Data *shared.GetSourcesCatalogV1Output `json:"data,omitempty"`
}

func (o *GetSourcesCatalogCatalogResponseBody) GetData() *shared.GetSourcesCatalogV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetSourcesCatalogResponseBody - OK
type GetSourcesCatalogResponseBody struct {
	// Returns a list of all Source catalog items contained within a given page.
	Data *shared.GetSourcesCatalogV1Output `json:"data,omitempty"`
}

func (o *GetSourcesCatalogResponseBody) GetData() *shared.GetSourcesCatalogV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetSourcesCatalogResponse struct {
	// OK
	TwoHundredApplicationJSONObject *GetSourcesCatalogResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *GetSourcesCatalogCatalogResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *GetSourcesCatalogCatalogResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *GetSourcesCatalogCatalogResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetSourcesCatalogResponse) GetTwoHundredApplicationJSONObject() *GetSourcesCatalogResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetSourcesCatalogResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *GetSourcesCatalogCatalogResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *GetSourcesCatalogResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *GetSourcesCatalogCatalogResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *GetSourcesCatalogResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *GetSourcesCatalogCatalogResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *GetSourcesCatalogResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSourcesCatalogResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GetSourcesCatalogResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSourcesCatalogResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

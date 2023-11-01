// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListConnectedWarehousesFromSourceRequest struct {
	// Required pagination params for the request.
	//
	// This parameter exists in alpha.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	SourceID   string                 `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *ListConnectedWarehousesFromSourceRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

func (o *ListConnectedWarehousesFromSourceRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

// ListConnectedWarehousesFromSource200ApplicationVndSegmentV1betaPlusJSON - OK
type ListConnectedWarehousesFromSource200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a list of Warehouses connected to a Source.
	Data *shared.ListConnectedWarehousesFromSourceV1Output `json:"data,omitempty"`
}

func (o *ListConnectedWarehousesFromSource200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.ListConnectedWarehousesFromSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListConnectedWarehousesFromSource200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListConnectedWarehousesFromSource200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a list of Warehouses connected to a Source.
	Data *shared.ListConnectedWarehousesFromSourceAlphaOutput `json:"data,omitempty"`
}

func (o *ListConnectedWarehousesFromSource200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.ListConnectedWarehousesFromSourceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListConnectedWarehousesFromSource200ApplicationVndSegmentV1PlusJSON - OK
type ListConnectedWarehousesFromSource200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a list of Warehouses connected to a Source.
	Data *shared.ListConnectedWarehousesFromSourceV1Output `json:"data,omitempty"`
}

func (o *ListConnectedWarehousesFromSource200ApplicationVndSegmentV1PlusJSON) GetData() *shared.ListConnectedWarehousesFromSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListConnectedWarehousesFromSource200ApplicationJSON - OK
type ListConnectedWarehousesFromSource200ApplicationJSON struct {
	// Returns a list of Warehouses connected to a Source.
	Data *shared.ListConnectedWarehousesFromSourceV1Output `json:"data,omitempty"`
}

func (o *ListConnectedWarehousesFromSource200ApplicationJSON) GetData() *shared.ListConnectedWarehousesFromSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListConnectedWarehousesFromSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ListConnectedWarehousesFromSource200ApplicationJSONObject *ListConnectedWarehousesFromSource200ApplicationJSON
	// OK
	ListConnectedWarehousesFromSource200ApplicationVndSegmentV1PlusJSONObject *ListConnectedWarehousesFromSource200ApplicationVndSegmentV1PlusJSON
	// OK
	ListConnectedWarehousesFromSource200ApplicationVndSegmentV1alphaPlusJSONObject *ListConnectedWarehousesFromSource200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListConnectedWarehousesFromSource200ApplicationVndSegmentV1betaPlusJSONObject *ListConnectedWarehousesFromSource200ApplicationVndSegmentV1betaPlusJSON
}

func (o *ListConnectedWarehousesFromSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListConnectedWarehousesFromSourceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListConnectedWarehousesFromSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListConnectedWarehousesFromSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListConnectedWarehousesFromSourceResponse) GetListConnectedWarehousesFromSource200ApplicationJSONObject() *ListConnectedWarehousesFromSource200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ListConnectedWarehousesFromSource200ApplicationJSONObject
}

func (o *ListConnectedWarehousesFromSourceResponse) GetListConnectedWarehousesFromSource200ApplicationVndSegmentV1PlusJSONObject() *ListConnectedWarehousesFromSource200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.ListConnectedWarehousesFromSource200ApplicationVndSegmentV1PlusJSONObject
}

func (o *ListConnectedWarehousesFromSourceResponse) GetListConnectedWarehousesFromSource200ApplicationVndSegmentV1alphaPlusJSONObject() *ListConnectedWarehousesFromSource200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.ListConnectedWarehousesFromSource200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *ListConnectedWarehousesFromSourceResponse) GetListConnectedWarehousesFromSource200ApplicationVndSegmentV1betaPlusJSONObject() *ListConnectedWarehousesFromSource200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.ListConnectedWarehousesFromSource200ApplicationVndSegmentV1betaPlusJSONObject
}

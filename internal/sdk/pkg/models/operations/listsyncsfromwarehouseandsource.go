// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListSyncsFromWarehouseAndSourceRequest struct {
	// Defines the pagination parameters.
	//
	// This parameter exists in v1.
	Pagination  shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	SourceID    string                 `pathParam:"style=simple,explode=false,name=sourceId"`
	WarehouseID string                 `pathParam:"style=simple,explode=false,name=warehouseId"`
}

// ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1betaPlusJSON - OK
type ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a list that contains the most recent syncs for a Warehouse-source pair, filtered and constrained by a given
	// set of pagination parameters.
	Data *shared.ListSyncsFromWarehouseAndSourceV1Output `json:"data,omitempty"`
}

// ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a list that contains the most recent syncs for a Warehouse-source pair, filtered and constrained by a given
	// set of pagination parameters.
	Data *shared.ListSyncsFromWarehouseAndSourceV1Output `json:"data,omitempty"`
}

// ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1PlusJSON - OK
type ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a list that contains the most recent syncs for a Warehouse-source pair, filtered and constrained by a given
	// set of pagination parameters.
	Data *shared.ListSyncsFromWarehouseAndSourceV1Output `json:"data,omitempty"`
}

// ListSyncsFromWarehouseAndSource200ApplicationJSON - OK
type ListSyncsFromWarehouseAndSource200ApplicationJSON struct {
	// Returns a list that contains the most recent syncs for a Warehouse-source pair, filtered and constrained by a given
	// set of pagination parameters.
	Data *shared.ListSyncsFromWarehouseAndSourceV1Output `json:"data,omitempty"`
}

type ListSyncsFromWarehouseAndSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ListSyncsFromWarehouseAndSource200ApplicationJSONObject *ListSyncsFromWarehouseAndSource200ApplicationJSON
	// OK
	ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1PlusJSONObject *ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1PlusJSON
	// OK
	ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1alphaPlusJSONObject *ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1betaPlusJSONObject *ListSyncsFromWarehouseAndSource200ApplicationVndSegmentV1betaPlusJSON
}

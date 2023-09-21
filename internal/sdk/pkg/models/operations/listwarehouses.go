// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListWarehousesRequest struct {
	// Defines the pagination parameters.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

// ListWarehouses200ApplicationVndSegmentV1betaPlusJSON - OK
type ListWarehouses200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a list of Warehouses that belong to a Workspace.
	Data *shared.ListWarehousesV1Output `json:"data,omitempty"`
}

// ListWarehouses200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListWarehouses200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a list of Warehouses that belong to a Workspace.
	Data *shared.ListWarehousesV1Output `json:"data,omitempty"`
}

// ListWarehouses200ApplicationVndSegmentV1PlusJSON - OK
type ListWarehouses200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a list of Warehouses that belong to a Workspace.
	Data *shared.ListWarehousesV1Output `json:"data,omitempty"`
}

// ListWarehouses200ApplicationJSON - OK
type ListWarehouses200ApplicationJSON struct {
	// Returns a list of Warehouses that belong to a Workspace.
	Data *shared.ListWarehousesV1Output `json:"data,omitempty"`
}

type ListWarehousesResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	ListWarehouses200ApplicationJSONObject *ListWarehouses200ApplicationJSON
	// OK
	ListWarehouses200ApplicationVndSegmentV1PlusJSONObject *ListWarehouses200ApplicationVndSegmentV1PlusJSON
	// OK
	ListWarehouses200ApplicationVndSegmentV1alphaPlusJSONObject *ListWarehouses200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListWarehouses200ApplicationVndSegmentV1betaPlusJSONObject *ListWarehouses200ApplicationVndSegmentV1betaPlusJSON
}

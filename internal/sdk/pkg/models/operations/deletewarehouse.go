// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type DeleteWarehouseRequest struct {
	WarehouseID string `pathParam:"style=simple,explode=false,name=warehouseId"`
}

// DeleteWarehouse200ApplicationVndSegmentV1betaPlusJSON - OK
type DeleteWarehouse200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the status of a Warehouse deletion.
	Data *shared.DeleteWarehouseV1Output `json:"data,omitempty"`
}

// DeleteWarehouse200ApplicationVndSegmentV1alphaPlusJSON - OK
type DeleteWarehouse200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the status of a Warehouse deletion.
	Data *shared.DeleteWarehouseV1Output `json:"data,omitempty"`
}

// DeleteWarehouse200ApplicationVndSegmentV1PlusJSON - OK
type DeleteWarehouse200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the status of a Warehouse deletion.
	Data *shared.DeleteWarehouseV1Output `json:"data,omitempty"`
}

// DeleteWarehouse200ApplicationJSON - OK
type DeleteWarehouse200ApplicationJSON struct {
	// Returns the status of a Warehouse deletion.
	Data *shared.DeleteWarehouseV1Output `json:"data,omitempty"`
}

type DeleteWarehouseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	DeleteWarehouse200ApplicationJSONObject *DeleteWarehouse200ApplicationJSON
	// OK
	DeleteWarehouse200ApplicationVndSegmentV1PlusJSONObject *DeleteWarehouse200ApplicationVndSegmentV1PlusJSON
	// OK
	DeleteWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject *DeleteWarehouse200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	DeleteWarehouse200ApplicationVndSegmentV1betaPlusJSONObject *DeleteWarehouse200ApplicationVndSegmentV1betaPlusJSON
}
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type UpdateSelectiveSyncForWarehouseRequest struct {
	UpdateSelectiveSyncForWarehouseV1Input shared.UpdateSelectiveSyncForWarehouseV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	WarehouseID                            string                                        `pathParam:"style=simple,explode=false,name=warehouseId"`
}

// UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1betaPlusJSON - OK
type UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1betaPlusJSON struct {
	// Results from updating the schema for a Warehouse/source pair.
	Data *shared.UpdateSelectiveSyncForWarehouseV1Output `json:"data,omitempty"`
}

// UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1alphaPlusJSON - OK
type UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Results from updating the schema for a Warehouse/source pair.
	Data *shared.UpdateSelectiveSyncForWarehouseV1Output `json:"data,omitempty"`
}

// UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1PlusJSON - OK
type UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1PlusJSON struct {
	// Results from updating the schema for a Warehouse/source pair.
	Data *shared.UpdateSelectiveSyncForWarehouseV1Output `json:"data,omitempty"`
}

// UpdateSelectiveSyncForWarehouse200ApplicationJSON - OK
type UpdateSelectiveSyncForWarehouse200ApplicationJSON struct {
	// Results from updating the schema for a Warehouse/source pair.
	Data *shared.UpdateSelectiveSyncForWarehouseV1Output `json:"data,omitempty"`
}

type UpdateSelectiveSyncForWarehouseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UpdateSelectiveSyncForWarehouse200ApplicationJSONObject *UpdateSelectiveSyncForWarehouse200ApplicationJSON
	// OK
	UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1PlusJSONObject *UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1PlusJSON
	// OK
	UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject *UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1betaPlusJSONObject *UpdateSelectiveSyncForWarehouse200ApplicationVndSegmentV1betaPlusJSON
}
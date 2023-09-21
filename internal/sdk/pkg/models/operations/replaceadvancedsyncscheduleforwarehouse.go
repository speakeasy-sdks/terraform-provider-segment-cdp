// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ReplaceAdvancedSyncScheduleForWarehouseRequest struct {
	ReplaceAdvancedSyncScheduleForWarehouseV1Input shared.ReplaceAdvancedSyncScheduleForWarehouseV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	WarehouseID                                    string                                                `pathParam:"style=simple,explode=false,name=warehouseId"`
}

// ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1betaPlusJSON - OK
type ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the advanced sync schedule for a Warehouse.
	Data *shared.ReplaceAdvancedSyncScheduleForWarehouseV1Output `json:"data,omitempty"`
}

// ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1alphaPlusJSON - OK
type ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the advanced sync schedule for a Warehouse.
	Data *shared.ReplaceAdvancedSyncScheduleForWarehouseV1Output `json:"data,omitempty"`
}

// ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1PlusJSON - OK
type ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the advanced sync schedule for a Warehouse.
	Data *shared.ReplaceAdvancedSyncScheduleForWarehouseV1Output `json:"data,omitempty"`
}

// ReplaceAdvancedSyncScheduleForWarehouse200ApplicationJSON - OK
type ReplaceAdvancedSyncScheduleForWarehouse200ApplicationJSON struct {
	// Returns the advanced sync schedule for a Warehouse.
	Data *shared.ReplaceAdvancedSyncScheduleForWarehouseV1Output `json:"data,omitempty"`
}

type ReplaceAdvancedSyncScheduleForWarehouseResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	ReplaceAdvancedSyncScheduleForWarehouse200ApplicationJSONObject *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationJSON
	// OK
	ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1PlusJSONObject *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1PlusJSON
	// OK
	ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1betaPlusJSONObject *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1betaPlusJSON
}

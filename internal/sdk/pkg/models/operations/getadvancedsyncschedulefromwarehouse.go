// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetAdvancedSyncScheduleFromWarehouseRequest struct {
	WarehouseID string `pathParam:"style=simple,explode=false,name=warehouseId"`
}

// GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1betaPlusJSON - OK
type GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the advanced sync schedule for a Warehouse.
	Data *shared.GetAdvancedSyncScheduleFromWarehouseV1Output `json:"data,omitempty"`
}

// GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the advanced sync schedule for a Warehouse.
	Data *shared.GetAdvancedSyncScheduleFromWarehouseV1Output `json:"data,omitempty"`
}

// GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1PlusJSON - OK
type GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the advanced sync schedule for a Warehouse.
	Data *shared.GetAdvancedSyncScheduleFromWarehouseV1Output `json:"data,omitempty"`
}

// GetAdvancedSyncScheduleFromWarehouse200ApplicationJSON - OK
type GetAdvancedSyncScheduleFromWarehouse200ApplicationJSON struct {
	// Returns the advanced sync schedule for a Warehouse.
	Data *shared.GetAdvancedSyncScheduleFromWarehouseV1Output `json:"data,omitempty"`
}

type GetAdvancedSyncScheduleFromWarehouseResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	GetAdvancedSyncScheduleFromWarehouse200ApplicationJSONObject *GetAdvancedSyncScheduleFromWarehouse200ApplicationJSON
	// OK
	GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1PlusJSONObject *GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1PlusJSON
	// OK
	GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject *GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1betaPlusJSONObject *GetAdvancedSyncScheduleFromWarehouse200ApplicationVndSegmentV1betaPlusJSON
}

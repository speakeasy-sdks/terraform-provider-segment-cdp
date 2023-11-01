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

func (o *ReplaceAdvancedSyncScheduleForWarehouseRequest) GetReplaceAdvancedSyncScheduleForWarehouseV1Input() shared.ReplaceAdvancedSyncScheduleForWarehouseV1Input {
	if o == nil {
		return shared.ReplaceAdvancedSyncScheduleForWarehouseV1Input{}
	}
	return o.ReplaceAdvancedSyncScheduleForWarehouseV1Input
}

func (o *ReplaceAdvancedSyncScheduleForWarehouseRequest) GetWarehouseID() string {
	if o == nil {
		return ""
	}
	return o.WarehouseID
}

// ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1betaPlusJSON - OK
type ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the advanced sync schedule for a Warehouse.
	Data *shared.ReplaceAdvancedSyncScheduleForWarehouseV1Output `json:"data,omitempty"`
}

func (o *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.ReplaceAdvancedSyncScheduleForWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1alphaPlusJSON - OK
type ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the advanced sync schedule for a Warehouse.
	Data *shared.ReplaceAdvancedSyncScheduleForWarehouseV1Output `json:"data,omitempty"`
}

func (o *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.ReplaceAdvancedSyncScheduleForWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1PlusJSON - OK
type ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the advanced sync schedule for a Warehouse.
	Data *shared.ReplaceAdvancedSyncScheduleForWarehouseV1Output `json:"data,omitempty"`
}

func (o *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1PlusJSON) GetData() *shared.ReplaceAdvancedSyncScheduleForWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ReplaceAdvancedSyncScheduleForWarehouse200ApplicationJSON - OK
type ReplaceAdvancedSyncScheduleForWarehouse200ApplicationJSON struct {
	// Returns the advanced sync schedule for a Warehouse.
	Data *shared.ReplaceAdvancedSyncScheduleForWarehouseV1Output `json:"data,omitempty"`
}

func (o *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationJSON) GetData() *shared.ReplaceAdvancedSyncScheduleForWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type ReplaceAdvancedSyncScheduleForWarehouseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ReplaceAdvancedSyncScheduleForWarehouse200ApplicationJSONObject *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationJSON
	// OK
	ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1PlusJSONObject *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1PlusJSON
	// OK
	ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1betaPlusJSONObject *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1betaPlusJSON
}

func (o *ReplaceAdvancedSyncScheduleForWarehouseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReplaceAdvancedSyncScheduleForWarehouseResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ReplaceAdvancedSyncScheduleForWarehouseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReplaceAdvancedSyncScheduleForWarehouseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ReplaceAdvancedSyncScheduleForWarehouseResponse) GetReplaceAdvancedSyncScheduleForWarehouse200ApplicationJSONObject() *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.ReplaceAdvancedSyncScheduleForWarehouse200ApplicationJSONObject
}

func (o *ReplaceAdvancedSyncScheduleForWarehouseResponse) GetReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1PlusJSONObject() *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1PlusJSONObject
}

func (o *ReplaceAdvancedSyncScheduleForWarehouseResponse) GetReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject() *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *ReplaceAdvancedSyncScheduleForWarehouseResponse) GetReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1betaPlusJSONObject() *ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.ReplaceAdvancedSyncScheduleForWarehouse200ApplicationVndSegmentV1betaPlusJSONObject
}

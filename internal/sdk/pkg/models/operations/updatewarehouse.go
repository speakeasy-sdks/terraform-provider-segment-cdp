// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type UpdateWarehouseRequest struct {
	UpdateWarehouseV1Input shared.UpdateWarehouseV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	WarehouseID            string                        `pathParam:"style=simple,explode=false,name=warehouseId"`
}

func (o *UpdateWarehouseRequest) GetUpdateWarehouseV1Input() shared.UpdateWarehouseV1Input {
	if o == nil {
		return shared.UpdateWarehouseV1Input{}
	}
	return o.UpdateWarehouseV1Input
}

func (o *UpdateWarehouseRequest) GetWarehouseID() string {
	if o == nil {
		return ""
	}
	return o.WarehouseID
}

// UpdateWarehouse200ApplicationVndSegmentV1betaPlusJSON - OK
type UpdateWarehouse200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the updated Warehouse.
	Data *shared.UpdateWarehouseV1Output `json:"data,omitempty"`
}

func (o *UpdateWarehouse200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.UpdateWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateWarehouse200ApplicationVndSegmentV1alphaPlusJSON - OK
type UpdateWarehouse200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the updated Warehouse.
	Data *shared.UpdateWarehouseV1Output `json:"data,omitempty"`
}

func (o *UpdateWarehouse200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.UpdateWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateWarehouse200ApplicationVndSegmentV1PlusJSON - OK
type UpdateWarehouse200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the updated Warehouse.
	Data *shared.UpdateWarehouseV1Output `json:"data,omitempty"`
}

func (o *UpdateWarehouse200ApplicationVndSegmentV1PlusJSON) GetData() *shared.UpdateWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateWarehouse200ApplicationJSON - OK
type UpdateWarehouse200ApplicationJSON struct {
	// Returns the updated Warehouse.
	Data *shared.UpdateWarehouseV1Output `json:"data,omitempty"`
}

func (o *UpdateWarehouse200ApplicationJSON) GetData() *shared.UpdateWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateWarehouseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UpdateWarehouse200ApplicationJSONObject *UpdateWarehouse200ApplicationJSON
	// OK
	UpdateWarehouse200ApplicationVndSegmentV1PlusJSONObject *UpdateWarehouse200ApplicationVndSegmentV1PlusJSON
	// OK
	UpdateWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject *UpdateWarehouse200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	UpdateWarehouse200ApplicationVndSegmentV1betaPlusJSONObject *UpdateWarehouse200ApplicationVndSegmentV1betaPlusJSON
}

func (o *UpdateWarehouseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateWarehouseResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *UpdateWarehouseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateWarehouseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateWarehouseResponse) GetUpdateWarehouse200ApplicationJSONObject() *UpdateWarehouse200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateWarehouse200ApplicationJSONObject
}

func (o *UpdateWarehouseResponse) GetUpdateWarehouse200ApplicationVndSegmentV1PlusJSONObject() *UpdateWarehouse200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.UpdateWarehouse200ApplicationVndSegmentV1PlusJSONObject
}

func (o *UpdateWarehouseResponse) GetUpdateWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject() *UpdateWarehouse200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.UpdateWarehouse200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *UpdateWarehouseResponse) GetUpdateWarehouse200ApplicationVndSegmentV1betaPlusJSONObject() *UpdateWarehouse200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.UpdateWarehouse200ApplicationVndSegmentV1betaPlusJSONObject
}

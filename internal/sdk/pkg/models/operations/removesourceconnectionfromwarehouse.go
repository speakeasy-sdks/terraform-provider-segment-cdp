// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type RemoveSourceConnectionFromWarehouseRequest struct {
	SourceID    string `pathParam:"style=simple,explode=false,name=sourceId"`
	WarehouseID string `pathParam:"style=simple,explode=false,name=warehouseId"`
}

func (o *RemoveSourceConnectionFromWarehouseRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *RemoveSourceConnectionFromWarehouseRequest) GetWarehouseID() string {
	if o == nil {
		return ""
	}
	return o.WarehouseID
}

// RemoveSourceConnectionFromWarehouseWarehousesResponse200ResponseBody - OK
type RemoveSourceConnectionFromWarehouseWarehousesResponse200ResponseBody struct {
	// Response indicating whether the disconnection was successful.
	Data *shared.RemoveSourceConnectionFromWarehouseV1Output `json:"data,omitempty"`
}

func (o *RemoveSourceConnectionFromWarehouseWarehousesResponse200ResponseBody) GetData() *shared.RemoveSourceConnectionFromWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// RemoveSourceConnectionFromWarehouseWarehousesResponseResponseBody - OK
type RemoveSourceConnectionFromWarehouseWarehousesResponseResponseBody struct {
	// Response indicating whether the disconnection was successful.
	Data *shared.RemoveSourceConnectionFromWarehouseV1Output `json:"data,omitempty"`
}

func (o *RemoveSourceConnectionFromWarehouseWarehousesResponseResponseBody) GetData() *shared.RemoveSourceConnectionFromWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// RemoveSourceConnectionFromWarehouseWarehousesResponseBody - OK
type RemoveSourceConnectionFromWarehouseWarehousesResponseBody struct {
	// Response indicating whether the disconnection was successful.
	Data *shared.RemoveSourceConnectionFromWarehouseV1Output `json:"data,omitempty"`
}

func (o *RemoveSourceConnectionFromWarehouseWarehousesResponseBody) GetData() *shared.RemoveSourceConnectionFromWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// RemoveSourceConnectionFromWarehouseResponseBody - OK
type RemoveSourceConnectionFromWarehouseResponseBody struct {
	// Response indicating whether the disconnection was successful.
	Data *shared.RemoveSourceConnectionFromWarehouseV1Output `json:"data,omitempty"`
}

func (o *RemoveSourceConnectionFromWarehouseResponseBody) GetData() *shared.RemoveSourceConnectionFromWarehouseV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type RemoveSourceConnectionFromWarehouseResponse struct {
	// OK
	TwoHundredApplicationJSONObject *RemoveSourceConnectionFromWarehouseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *RemoveSourceConnectionFromWarehouseWarehousesResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *RemoveSourceConnectionFromWarehouseWarehousesResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *RemoveSourceConnectionFromWarehouseWarehousesResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RemoveSourceConnectionFromWarehouseResponse) GetTwoHundredApplicationJSONObject() *RemoveSourceConnectionFromWarehouseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *RemoveSourceConnectionFromWarehouseResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *RemoveSourceConnectionFromWarehouseWarehousesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *RemoveSourceConnectionFromWarehouseResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *RemoveSourceConnectionFromWarehouseWarehousesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *RemoveSourceConnectionFromWarehouseResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *RemoveSourceConnectionFromWarehouseWarehousesResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *RemoveSourceConnectionFromWarehouseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveSourceConnectionFromWarehouseResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *RemoveSourceConnectionFromWarehouseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveSourceConnectionFromWarehouseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

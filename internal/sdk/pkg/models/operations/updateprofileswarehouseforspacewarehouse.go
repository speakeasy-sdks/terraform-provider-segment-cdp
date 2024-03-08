// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateProfilesWarehouseForSpaceWarehouseRequest struct {
	UpdateProfilesWarehouseForSpaceWarehouseAlphaInput shared.UpdateProfilesWarehouseForSpaceWarehouseAlphaInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	SpaceID                                            string                                                    `pathParam:"style=simple,explode=false,name=spaceId"`
	WarehouseID                                        string                                                    `pathParam:"style=simple,explode=false,name=warehouseId"`
}

func (o *UpdateProfilesWarehouseForSpaceWarehouseRequest) GetUpdateProfilesWarehouseForSpaceWarehouseAlphaInput() shared.UpdateProfilesWarehouseForSpaceWarehouseAlphaInput {
	if o == nil {
		return shared.UpdateProfilesWarehouseForSpaceWarehouseAlphaInput{}
	}
	return o.UpdateProfilesWarehouseForSpaceWarehouseAlphaInput
}

func (o *UpdateProfilesWarehouseForSpaceWarehouseRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

func (o *UpdateProfilesWarehouseForSpaceWarehouseRequest) GetWarehouseID() string {
	if o == nil {
		return ""
	}
	return o.WarehouseID
}

// UpdateProfilesWarehouseForSpaceWarehouseResponseBody - OK
type UpdateProfilesWarehouseForSpaceWarehouseResponseBody struct {
	// Returns the updated Warehouse.
	Data *shared.UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput `json:"data,omitempty"`
}

func (o *UpdateProfilesWarehouseForSpaceWarehouseResponseBody) GetData() *shared.UpdateProfilesWarehouseForSpaceWarehouseAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateProfilesWarehouseForSpaceWarehouseResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Object *UpdateProfilesWarehouseForSpaceWarehouseResponseBody
}

func (o *UpdateProfilesWarehouseForSpaceWarehouseResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateProfilesWarehouseForSpaceWarehouseResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *UpdateProfilesWarehouseForSpaceWarehouseResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateProfilesWarehouseForSpaceWarehouseResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateProfilesWarehouseForSpaceWarehouseResponse) GetObject() *UpdateProfilesWarehouseForSpaceWarehouseResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetRegulationRequest struct {
	RegulateID string `pathParam:"style=simple,explode=false,name=regulateId"`
}

func (o *GetRegulationRequest) GetRegulateID() string {
	if o == nil {
		return ""
	}
	return o.RegulateID
}

// GetRegulation200ApplicationVndSegmentV1betaPlusJSON - OK
type GetRegulation200ApplicationVndSegmentV1betaPlusJSON struct {
	// The regulate request returned.
	Data *shared.GetRegulationV1Output `json:"data,omitempty"`
}

func (o *GetRegulation200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.GetRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetRegulation200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetRegulation200ApplicationVndSegmentV1alphaPlusJSON struct {
	// The regulate request returned.
	Data *shared.GetRegulationV1Output `json:"data,omitempty"`
}

func (o *GetRegulation200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.GetRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetRegulation200ApplicationVndSegmentV1PlusJSON - OK
type GetRegulation200ApplicationVndSegmentV1PlusJSON struct {
	// The regulate request returned.
	Data *shared.GetRegulationV1Output `json:"data,omitempty"`
}

func (o *GetRegulation200ApplicationVndSegmentV1PlusJSON) GetData() *shared.GetRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetRegulation200ApplicationJSON - OK
type GetRegulation200ApplicationJSON struct {
	// The regulate request returned.
	Data *shared.GetRegulationV1Output `json:"data,omitempty"`
}

func (o *GetRegulation200ApplicationJSON) GetData() *shared.GetRegulationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetRegulationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	GetRegulation200ApplicationJSONObject *GetRegulation200ApplicationJSON
	// OK
	GetRegulation200ApplicationVndSegmentV1PlusJSONObject *GetRegulation200ApplicationVndSegmentV1PlusJSON
	// OK
	GetRegulation200ApplicationVndSegmentV1alphaPlusJSONObject *GetRegulation200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	GetRegulation200ApplicationVndSegmentV1betaPlusJSONObject *GetRegulation200ApplicationVndSegmentV1betaPlusJSON
}

func (o *GetRegulationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRegulationResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GetRegulationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRegulationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetRegulationResponse) GetGetRegulation200ApplicationJSONObject() *GetRegulation200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetRegulation200ApplicationJSONObject
}

func (o *GetRegulationResponse) GetGetRegulation200ApplicationVndSegmentV1PlusJSONObject() *GetRegulation200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.GetRegulation200ApplicationVndSegmentV1PlusJSONObject
}

func (o *GetRegulationResponse) GetGetRegulation200ApplicationVndSegmentV1alphaPlusJSONObject() *GetRegulation200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.GetRegulation200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *GetRegulationResponse) GetGetRegulation200ApplicationVndSegmentV1betaPlusJSONObject() *GetRegulation200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.GetRegulation200ApplicationVndSegmentV1betaPlusJSONObject
}

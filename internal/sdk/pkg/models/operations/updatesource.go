// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type UpdateSourceRequest struct {
	UpdateSourceV1Input shared.UpdateSourceV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	SourceID            string                     `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *UpdateSourceRequest) GetUpdateSourceV1Input() shared.UpdateSourceV1Input {
	if o == nil {
		return shared.UpdateSourceV1Input{}
	}
	return o.UpdateSourceV1Input
}

func (o *UpdateSourceRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

// UpdateSource200ApplicationVndSegmentV1betaPlusJSON - OK
type UpdateSource200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the updated Source.
	Data *shared.UpdateSourceV1Output `json:"data,omitempty"`
}

func (o *UpdateSource200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.UpdateSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateSource200ApplicationVndSegmentV1alphaPlusJSON - OK
type UpdateSource200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the updated Source.
	Data *shared.UpdateSourceAlphaOutput `json:"data,omitempty"`
}

func (o *UpdateSource200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.UpdateSourceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateSource200ApplicationVndSegmentV1PlusJSON - OK
type UpdateSource200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the updated Source.
	Data *shared.UpdateSourceV1Output `json:"data,omitempty"`
}

func (o *UpdateSource200ApplicationVndSegmentV1PlusJSON) GetData() *shared.UpdateSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateSource200ApplicationJSON - OK
type UpdateSource200ApplicationJSON struct {
	// Returns the updated Source.
	Data *shared.UpdateSourceV1Output `json:"data,omitempty"`
}

func (o *UpdateSource200ApplicationJSON) GetData() *shared.UpdateSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateSourceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UpdateSource200ApplicationJSONObject *UpdateSource200ApplicationJSON
	// OK
	UpdateSource200ApplicationVndSegmentV1PlusJSONObject *UpdateSource200ApplicationVndSegmentV1PlusJSON
	// OK
	UpdateSource200ApplicationVndSegmentV1alphaPlusJSONObject *UpdateSource200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	UpdateSource200ApplicationVndSegmentV1betaPlusJSONObject *UpdateSource200ApplicationVndSegmentV1betaPlusJSON
}

func (o *UpdateSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateSourceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *UpdateSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateSourceResponse) GetUpdateSource200ApplicationJSONObject() *UpdateSource200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateSource200ApplicationJSONObject
}

func (o *UpdateSourceResponse) GetUpdateSource200ApplicationVndSegmentV1PlusJSONObject() *UpdateSource200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.UpdateSource200ApplicationVndSegmentV1PlusJSONObject
}

func (o *UpdateSourceResponse) GetUpdateSource200ApplicationVndSegmentV1alphaPlusJSONObject() *UpdateSource200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.UpdateSource200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *UpdateSourceResponse) GetUpdateSource200ApplicationVndSegmentV1betaPlusJSONObject() *UpdateSource200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.UpdateSource200ApplicationVndSegmentV1betaPlusJSONObject
}

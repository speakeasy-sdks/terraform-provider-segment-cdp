// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

// PreviewDestinationFilter200ApplicationVndSegmentV1betaPlusJSON - OK
type PreviewDestinationFilter200ApplicationVndSegmentV1betaPlusJSON struct {
	// Preview output from applying the Destination filter.
	// Segment modifies or nullifies payloads depending on the provided filter actions.
	Data *shared.PreviewDestinationFilterV1Output `json:"data,omitempty"`
}

func (o *PreviewDestinationFilter200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.PreviewDestinationFilterV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// PreviewDestinationFilter200ApplicationVndSegmentV1alphaPlusJSON - OK
type PreviewDestinationFilter200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Preview output from applying the Destination filter.
	// Segment modifies or nullifies payloads depending on the provided filter actions.
	Data *shared.PreviewDestinationFilterV1Output `json:"data,omitempty"`
}

func (o *PreviewDestinationFilter200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.PreviewDestinationFilterV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// PreviewDestinationFilter200ApplicationVndSegmentV1PlusJSON - OK
type PreviewDestinationFilter200ApplicationVndSegmentV1PlusJSON struct {
	// Preview output from applying the Destination filter.
	// Segment modifies or nullifies payloads depending on the provided filter actions.
	Data *shared.PreviewDestinationFilterV1Output `json:"data,omitempty"`
}

func (o *PreviewDestinationFilter200ApplicationVndSegmentV1PlusJSON) GetData() *shared.PreviewDestinationFilterV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// PreviewDestinationFilter200ApplicationJSON - OK
type PreviewDestinationFilter200ApplicationJSON struct {
	// Preview output from applying the Destination filter.
	// Segment modifies or nullifies payloads depending on the provided filter actions.
	Data *shared.PreviewDestinationFilterV1Output `json:"data,omitempty"`
}

func (o *PreviewDestinationFilter200ApplicationJSON) GetData() *shared.PreviewDestinationFilterV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type PreviewDestinationFilterResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	PreviewDestinationFilter200ApplicationJSONObject *PreviewDestinationFilter200ApplicationJSON
	// OK
	PreviewDestinationFilter200ApplicationVndSegmentV1PlusJSONObject *PreviewDestinationFilter200ApplicationVndSegmentV1PlusJSON
	// OK
	PreviewDestinationFilter200ApplicationVndSegmentV1alphaPlusJSONObject *PreviewDestinationFilter200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	PreviewDestinationFilter200ApplicationVndSegmentV1betaPlusJSONObject *PreviewDestinationFilter200ApplicationVndSegmentV1betaPlusJSON
}

func (o *PreviewDestinationFilterResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PreviewDestinationFilterResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *PreviewDestinationFilterResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PreviewDestinationFilterResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PreviewDestinationFilterResponse) GetPreviewDestinationFilter200ApplicationJSONObject() *PreviewDestinationFilter200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.PreviewDestinationFilter200ApplicationJSONObject
}

func (o *PreviewDestinationFilterResponse) GetPreviewDestinationFilter200ApplicationVndSegmentV1PlusJSONObject() *PreviewDestinationFilter200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.PreviewDestinationFilter200ApplicationVndSegmentV1PlusJSONObject
}

func (o *PreviewDestinationFilterResponse) GetPreviewDestinationFilter200ApplicationVndSegmentV1alphaPlusJSONObject() *PreviewDestinationFilter200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.PreviewDestinationFilter200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *PreviewDestinationFilterResponse) GetPreviewDestinationFilter200ApplicationVndSegmentV1betaPlusJSONObject() *PreviewDestinationFilter200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.PreviewDestinationFilter200ApplicationVndSegmentV1betaPlusJSONObject
}

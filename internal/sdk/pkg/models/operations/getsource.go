// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetSourceRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

func (o *GetSourceRequest) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

// GetSourceSourcesResponse200ResponseBody - OK
type GetSourceSourcesResponse200ResponseBody struct {
	// Returns a Source.
	Data *shared.GetSourceV1Output `json:"data,omitempty"`
}

func (o *GetSourceSourcesResponse200ResponseBody) GetData() *shared.GetSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetSourceSourcesResponseResponseBody - OK
type GetSourceSourcesResponseResponseBody struct {
	// Returns a Source.
	Data *shared.GetSourceAlphaOutput `json:"data,omitempty"`
}

func (o *GetSourceSourcesResponseResponseBody) GetData() *shared.GetSourceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetSourceSourcesResponseBody - OK
type GetSourceSourcesResponseBody struct {
	// Returns a Source.
	Data *shared.GetSourceV1Output `json:"data,omitempty"`
}

func (o *GetSourceSourcesResponseBody) GetData() *shared.GetSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetSourceResponseBody - OK
type GetSourceResponseBody struct {
	// Returns a Source.
	Data *shared.GetSourceV1Output `json:"data,omitempty"`
}

func (o *GetSourceResponseBody) GetData() *shared.GetSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetSourceResponse struct {
	// OK
	TwoHundredApplicationJSONObject *GetSourceResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *GetSourceSourcesResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *GetSourceSourcesResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *GetSourceSourcesResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetSourceResponse) GetTwoHundredApplicationJSONObject() *GetSourceResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetSourceResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *GetSourceSourcesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *GetSourceResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *GetSourceSourcesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *GetSourceResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *GetSourceSourcesResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *GetSourceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSourceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GetSourceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSourceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

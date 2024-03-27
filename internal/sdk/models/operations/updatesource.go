// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/models/shared"
	"net/http"
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

// UpdateSourceSourcesResponse200ResponseBody - OK
type UpdateSourceSourcesResponse200ResponseBody struct {
	// Returns the updated Source.
	Data *shared.UpdateSourceV1Output `json:"data,omitempty"`
}

func (o *UpdateSourceSourcesResponse200ResponseBody) GetData() *shared.UpdateSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateSourceSourcesResponseResponseBody - OK
type UpdateSourceSourcesResponseResponseBody struct {
	// Returns the updated Source.
	Data *shared.UpdateSourceAlphaOutput `json:"data,omitempty"`
}

func (o *UpdateSourceSourcesResponseResponseBody) GetData() *shared.UpdateSourceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateSourceSourcesResponseBody - OK
type UpdateSourceSourcesResponseBody struct {
	// Returns the updated Source.
	Data *shared.UpdateSourceV1Output `json:"data,omitempty"`
}

func (o *UpdateSourceSourcesResponseBody) GetData() *shared.UpdateSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateSourceResponseBody - OK
type UpdateSourceResponseBody struct {
	// Returns the updated Source.
	Data *shared.UpdateSourceV1Output `json:"data,omitempty"`
}

func (o *UpdateSourceResponseBody) GetData() *shared.UpdateSourceV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateSourceResponse struct {
	// OK
	TwoHundredApplicationJSONObject *UpdateSourceResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *UpdateSourceSourcesResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *UpdateSourceSourcesResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *UpdateSourceSourcesResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateSourceResponse) GetTwoHundredApplicationJSONObject() *UpdateSourceResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *UpdateSourceResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *UpdateSourceSourcesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *UpdateSourceResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *UpdateSourceSourcesResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *UpdateSourceResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *UpdateSourceSourcesResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
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
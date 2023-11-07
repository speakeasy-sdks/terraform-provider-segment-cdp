// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetTransformationRequest struct {
	TransformationID string `pathParam:"style=simple,explode=false,name=transformationId"`
}

func (o *GetTransformationRequest) GetTransformationID() string {
	if o == nil {
		return ""
	}
	return o.TransformationID
}

// GetTransformationTransformationsResponse200ResponseBody - OK
type GetTransformationTransformationsResponse200ResponseBody struct {
	// The output of Transformation retrieval.
	Data *shared.GetTransformationV1Output `json:"data,omitempty"`
}

func (o *GetTransformationTransformationsResponse200ResponseBody) GetData() *shared.GetTransformationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetTransformationTransformationsResponseResponseBody - OK
type GetTransformationTransformationsResponseResponseBody struct {
	// The output of Transformation retrieval.
	Data *shared.GetTransformationV1Output `json:"data,omitempty"`
}

func (o *GetTransformationTransformationsResponseResponseBody) GetData() *shared.GetTransformationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetTransformationTransformationsResponseBody - OK
type GetTransformationTransformationsResponseBody struct {
	// The output of Transformation retrieval.
	Data *shared.GetTransformationV1Output `json:"data,omitempty"`
}

func (o *GetTransformationTransformationsResponseBody) GetData() *shared.GetTransformationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// GetTransformationResponseBody - OK
type GetTransformationResponseBody struct {
	// The output of Transformation retrieval.
	Data *shared.GetTransformationV1Output `json:"data,omitempty"`
}

func (o *GetTransformationResponseBody) GetData() *shared.GetTransformationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetTransformationResponse struct {
	// OK
	TwoHundredApplicationJSONObject *GetTransformationResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *GetTransformationTransformationsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *GetTransformationTransformationsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *GetTransformationTransformationsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetTransformationResponse) GetTwoHundredApplicationJSONObject() *GetTransformationResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetTransformationResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *GetTransformationTransformationsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *GetTransformationResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *GetTransformationTransformationsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *GetTransformationResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *GetTransformationTransformationsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *GetTransformationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTransformationResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *GetTransformationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTransformationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

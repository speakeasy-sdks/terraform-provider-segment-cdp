// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/models/shared"
	"net/http"
)

type UpdateTransformationRequest struct {
	UpdateTransformationV1Input shared.UpdateTransformationV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	TransformationID            string                             `pathParam:"style=simple,explode=false,name=transformationId"`
}

func (o *UpdateTransformationRequest) GetUpdateTransformationV1Input() shared.UpdateTransformationV1Input {
	if o == nil {
		return shared.UpdateTransformationV1Input{}
	}
	return o.UpdateTransformationV1Input
}

func (o *UpdateTransformationRequest) GetTransformationID() string {
	if o == nil {
		return ""
	}
	return o.TransformationID
}

// UpdateTransformationTransformationsResponse200ResponseBody - OK
type UpdateTransformationTransformationsResponse200ResponseBody struct {
	// The output of an updated Transformation.
	Data *shared.UpdateTransformationV1Output `json:"data,omitempty"`
}

func (o *UpdateTransformationTransformationsResponse200ResponseBody) GetData() *shared.UpdateTransformationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateTransformationTransformationsResponseResponseBody - OK
type UpdateTransformationTransformationsResponseResponseBody struct {
	// The output of an updated Transformation.
	Data *shared.UpdateTransformationV1Output `json:"data,omitempty"`
}

func (o *UpdateTransformationTransformationsResponseResponseBody) GetData() *shared.UpdateTransformationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateTransformationTransformationsResponseBody - OK
type UpdateTransformationTransformationsResponseBody struct {
	// The output of an updated Transformation.
	Data *shared.UpdateTransformationV1Output `json:"data,omitempty"`
}

func (o *UpdateTransformationTransformationsResponseBody) GetData() *shared.UpdateTransformationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateTransformationResponseBody - OK
type UpdateTransformationResponseBody struct {
	// The output of an updated Transformation.
	Data *shared.UpdateTransformationV1Output `json:"data,omitempty"`
}

func (o *UpdateTransformationResponseBody) GetData() *shared.UpdateTransformationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateTransformationResponse struct {
	// OK
	TwoHundredApplicationJSONObject *UpdateTransformationResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *UpdateTransformationTransformationsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *UpdateTransformationTransformationsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *UpdateTransformationTransformationsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateTransformationResponse) GetTwoHundredApplicationJSONObject() *UpdateTransformationResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *UpdateTransformationResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *UpdateTransformationTransformationsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *UpdateTransformationResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *UpdateTransformationTransformationsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *UpdateTransformationResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *UpdateTransformationTransformationsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *UpdateTransformationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTransformationResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *UpdateTransformationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTransformationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

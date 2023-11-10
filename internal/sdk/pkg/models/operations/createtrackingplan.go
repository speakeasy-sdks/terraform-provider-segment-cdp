// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"net/http"
)

// CreateTrackingPlanTrackingPlansResponse200ResponseBody - OK
type CreateTrackingPlanTrackingPlansResponse200ResponseBody struct {
	// Result of a CreateTrackingPlan call.
	Data *shared.CreateTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *CreateTrackingPlanTrackingPlansResponse200ResponseBody) GetData() *shared.CreateTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateTrackingPlanTrackingPlansResponseResponseBody - OK
type CreateTrackingPlanTrackingPlansResponseResponseBody struct {
	// Result of a CreateTrackingPlan call.
	Data *shared.CreateTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *CreateTrackingPlanTrackingPlansResponseResponseBody) GetData() *shared.CreateTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateTrackingPlanTrackingPlansResponseBody - OK
type CreateTrackingPlanTrackingPlansResponseBody struct {
	// Result of a CreateTrackingPlan call.
	Data *shared.CreateTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *CreateTrackingPlanTrackingPlansResponseBody) GetData() *shared.CreateTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateTrackingPlanResponseBody - OK
type CreateTrackingPlanResponseBody struct {
	// Result of a CreateTrackingPlan call.
	Data *shared.CreateTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *CreateTrackingPlanResponseBody) GetData() *shared.CreateTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateTrackingPlanResponse struct {
	// OK
	TwoHundredApplicationJSONObject *CreateTrackingPlanResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *CreateTrackingPlanTrackingPlansResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *CreateTrackingPlanTrackingPlansResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *CreateTrackingPlanTrackingPlansResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateTrackingPlanResponse) GetTwoHundredApplicationJSONObject() *CreateTrackingPlanResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *CreateTrackingPlanResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *CreateTrackingPlanTrackingPlansResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *CreateTrackingPlanResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *CreateTrackingPlanTrackingPlansResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *CreateTrackingPlanResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *CreateTrackingPlanTrackingPlansResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *CreateTrackingPlanResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTrackingPlanResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *CreateTrackingPlanResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTrackingPlanResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/models/shared"
	"net/http"
)

type UpdateRulesInTrackingPlanRequest struct {
	UpdateRulesInTrackingPlanV1Input shared.UpdateRulesInTrackingPlanV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	TrackingPlanID                   string                                  `pathParam:"style=simple,explode=false,name=trackingPlanId"`
}

func (o *UpdateRulesInTrackingPlanRequest) GetUpdateRulesInTrackingPlanV1Input() shared.UpdateRulesInTrackingPlanV1Input {
	if o == nil {
		return shared.UpdateRulesInTrackingPlanV1Input{}
	}
	return o.UpdateRulesInTrackingPlanV1Input
}

func (o *UpdateRulesInTrackingPlanRequest) GetTrackingPlanID() string {
	if o == nil {
		return ""
	}
	return o.TrackingPlanID
}

// UpdateRulesInTrackingPlanTrackingPlansResponse200ResponseBody - OK
type UpdateRulesInTrackingPlanTrackingPlansResponse200ResponseBody struct {
	// Updates Tracking Plan rules. Non-existent rules are added.
	Data *shared.UpdateRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *UpdateRulesInTrackingPlanTrackingPlansResponse200ResponseBody) GetData() *shared.UpdateRulesInTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateRulesInTrackingPlanTrackingPlansResponseResponseBody - OK
type UpdateRulesInTrackingPlanTrackingPlansResponseResponseBody struct {
	// Updates Tracking Plan rules. Non-existent rules are added.
	Data *shared.UpdateRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *UpdateRulesInTrackingPlanTrackingPlansResponseResponseBody) GetData() *shared.UpdateRulesInTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateRulesInTrackingPlanTrackingPlansResponseBody - OK
type UpdateRulesInTrackingPlanTrackingPlansResponseBody struct {
	// Updates Tracking Plan rules. Non-existent rules are added.
	Data *shared.UpdateRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *UpdateRulesInTrackingPlanTrackingPlansResponseBody) GetData() *shared.UpdateRulesInTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateRulesInTrackingPlanResponseBody - OK
type UpdateRulesInTrackingPlanResponseBody struct {
	// Updates Tracking Plan rules. Non-existent rules are added.
	Data *shared.UpdateRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *UpdateRulesInTrackingPlanResponseBody) GetData() *shared.UpdateRulesInTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateRulesInTrackingPlanResponse struct {
	// OK
	TwoHundredApplicationJSONObject *UpdateRulesInTrackingPlanResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *UpdateRulesInTrackingPlanTrackingPlansResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *UpdateRulesInTrackingPlanTrackingPlansResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *UpdateRulesInTrackingPlanTrackingPlansResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateRulesInTrackingPlanResponse) GetTwoHundredApplicationJSONObject() *UpdateRulesInTrackingPlanResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *UpdateRulesInTrackingPlanResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *UpdateRulesInTrackingPlanTrackingPlansResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *UpdateRulesInTrackingPlanResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *UpdateRulesInTrackingPlanTrackingPlansResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *UpdateRulesInTrackingPlanResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *UpdateRulesInTrackingPlanTrackingPlansResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *UpdateRulesInTrackingPlanResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateRulesInTrackingPlanResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *UpdateRulesInTrackingPlanResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateRulesInTrackingPlanResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
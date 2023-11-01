// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
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

// UpdateRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSON - OK
type UpdateRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSON struct {
	// Updates Tracking Plan rules. Non-existent rules are added.
	Data *shared.UpdateRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *UpdateRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.UpdateRulesInTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON - OK
type UpdateRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Updates Tracking Plan rules. Non-existent rules are added.
	Data *shared.UpdateRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *UpdateRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.UpdateRulesInTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSON - OK
type UpdateRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSON struct {
	// Updates Tracking Plan rules. Non-existent rules are added.
	Data *shared.UpdateRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *UpdateRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSON) GetData() *shared.UpdateRulesInTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateRulesInTrackingPlan200ApplicationJSON - OK
type UpdateRulesInTrackingPlan200ApplicationJSON struct {
	// Updates Tracking Plan rules. Non-existent rules are added.
	Data *shared.UpdateRulesInTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *UpdateRulesInTrackingPlan200ApplicationJSON) GetData() *shared.UpdateRulesInTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateRulesInTrackingPlanResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UpdateRulesInTrackingPlan200ApplicationJSONObject *UpdateRulesInTrackingPlan200ApplicationJSON
	// OK
	UpdateRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSONObject *UpdateRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSON
	// OK
	UpdateRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject *UpdateRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	UpdateRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject *UpdateRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSON
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

func (o *UpdateRulesInTrackingPlanResponse) GetUpdateRulesInTrackingPlan200ApplicationJSONObject() *UpdateRulesInTrackingPlan200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateRulesInTrackingPlan200ApplicationJSONObject
}

func (o *UpdateRulesInTrackingPlanResponse) GetUpdateRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSONObject() *UpdateRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.UpdateRulesInTrackingPlan200ApplicationVndSegmentV1PlusJSONObject
}

func (o *UpdateRulesInTrackingPlanResponse) GetUpdateRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject() *UpdateRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.UpdateRulesInTrackingPlan200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *UpdateRulesInTrackingPlanResponse) GetUpdateRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject() *UpdateRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.UpdateRulesInTrackingPlan200ApplicationVndSegmentV1betaPlusJSONObject
}

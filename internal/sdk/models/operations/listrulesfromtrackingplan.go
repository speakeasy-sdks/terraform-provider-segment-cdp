// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/models/shared"
	"net/http"
)

type ListRulesFromTrackingPlanRequest struct {
	// Pagination options.
	//
	// This parameter exists in v1.
	Pagination     shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	TrackingPlanID string                 `pathParam:"style=simple,explode=false,name=trackingPlanId"`
}

func (o *ListRulesFromTrackingPlanRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

func (o *ListRulesFromTrackingPlanRequest) GetTrackingPlanID() string {
	if o == nil {
		return ""
	}
	return o.TrackingPlanID
}

// ListRulesFromTrackingPlanTrackingPlansResponse200ResponseBody - OK
type ListRulesFromTrackingPlanTrackingPlansResponse200ResponseBody struct {
	// Lists a Tracking Plan's rules.
	Data *shared.ListRulesFromTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *ListRulesFromTrackingPlanTrackingPlansResponse200ResponseBody) GetData() *shared.ListRulesFromTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListRulesFromTrackingPlanTrackingPlansResponseResponseBody - OK
type ListRulesFromTrackingPlanTrackingPlansResponseResponseBody struct {
	// Lists a Tracking Plan's rules.
	Data *shared.ListRulesFromTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *ListRulesFromTrackingPlanTrackingPlansResponseResponseBody) GetData() *shared.ListRulesFromTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListRulesFromTrackingPlanTrackingPlansResponseBody - OK
type ListRulesFromTrackingPlanTrackingPlansResponseBody struct {
	// Lists a Tracking Plan's rules.
	Data *shared.ListRulesFromTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *ListRulesFromTrackingPlanTrackingPlansResponseBody) GetData() *shared.ListRulesFromTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListRulesFromTrackingPlanResponseBody - OK
type ListRulesFromTrackingPlanResponseBody struct {
	// Lists a Tracking Plan's rules.
	Data *shared.ListRulesFromTrackingPlanV1Output `json:"data,omitempty"`
}

func (o *ListRulesFromTrackingPlanResponseBody) GetData() *shared.ListRulesFromTrackingPlanV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListRulesFromTrackingPlanResponse struct {
	// OK
	TwoHundredApplicationJSONObject *ListRulesFromTrackingPlanResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *ListRulesFromTrackingPlanTrackingPlansResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *ListRulesFromTrackingPlanTrackingPlansResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *ListRulesFromTrackingPlanTrackingPlansResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListRulesFromTrackingPlanResponse) GetTwoHundredApplicationJSONObject() *ListRulesFromTrackingPlanResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *ListRulesFromTrackingPlanResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *ListRulesFromTrackingPlanTrackingPlansResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *ListRulesFromTrackingPlanResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *ListRulesFromTrackingPlanTrackingPlansResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *ListRulesFromTrackingPlanResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *ListRulesFromTrackingPlanTrackingPlansResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *ListRulesFromTrackingPlanResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListRulesFromTrackingPlanResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListRulesFromTrackingPlanResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListRulesFromTrackingPlanResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

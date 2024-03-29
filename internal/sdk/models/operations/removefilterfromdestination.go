// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/models/shared"
	"net/http"
)

type RemoveFilterFromDestinationRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
	FilterID      string `pathParam:"style=simple,explode=false,name=filterId"`
}

func (o *RemoveFilterFromDestinationRequest) GetDestinationID() string {
	if o == nil {
		return ""
	}
	return o.DestinationID
}

func (o *RemoveFilterFromDestinationRequest) GetFilterID() string {
	if o == nil {
		return ""
	}
	return o.FilterID
}

// RemoveFilterFromDestinationDestinationFiltersResponse200ResponseBody - OK
type RemoveFilterFromDestinationDestinationFiltersResponse200ResponseBody struct {
	// Output for DeleteDestinationFilterV1.
	Data *shared.RemoveFilterFromDestinationV1Output `json:"data,omitempty"`
}

func (o *RemoveFilterFromDestinationDestinationFiltersResponse200ResponseBody) GetData() *shared.RemoveFilterFromDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// RemoveFilterFromDestinationDestinationFiltersResponseResponseBody - OK
type RemoveFilterFromDestinationDestinationFiltersResponseResponseBody struct {
	// Output for DeleteDestinationFilterV1.
	Data *shared.RemoveFilterFromDestinationV1Output `json:"data,omitempty"`
}

func (o *RemoveFilterFromDestinationDestinationFiltersResponseResponseBody) GetData() *shared.RemoveFilterFromDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// RemoveFilterFromDestinationDestinationFiltersResponseBody - OK
type RemoveFilterFromDestinationDestinationFiltersResponseBody struct {
	// Output for DeleteDestinationFilterV1.
	Data *shared.RemoveFilterFromDestinationV1Output `json:"data,omitempty"`
}

func (o *RemoveFilterFromDestinationDestinationFiltersResponseBody) GetData() *shared.RemoveFilterFromDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// RemoveFilterFromDestinationResponseBody - OK
type RemoveFilterFromDestinationResponseBody struct {
	// Output for DeleteDestinationFilterV1.
	Data *shared.RemoveFilterFromDestinationV1Output `json:"data,omitempty"`
}

func (o *RemoveFilterFromDestinationResponseBody) GetData() *shared.RemoveFilterFromDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type RemoveFilterFromDestinationResponse struct {
	// OK
	TwoHundredApplicationJSONObject *RemoveFilterFromDestinationResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *RemoveFilterFromDestinationDestinationFiltersResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *RemoveFilterFromDestinationDestinationFiltersResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *RemoveFilterFromDestinationDestinationFiltersResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RemoveFilterFromDestinationResponse) GetTwoHundredApplicationJSONObject() *RemoveFilterFromDestinationResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *RemoveFilterFromDestinationResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *RemoveFilterFromDestinationDestinationFiltersResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *RemoveFilterFromDestinationResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *RemoveFilterFromDestinationDestinationFiltersResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *RemoveFilterFromDestinationResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *RemoveFilterFromDestinationDestinationFiltersResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *RemoveFilterFromDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RemoveFilterFromDestinationResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *RemoveFilterFromDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RemoveFilterFromDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

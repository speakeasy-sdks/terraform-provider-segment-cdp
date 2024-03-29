// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/models/shared"
	"net/http"
)

// CreateDestinationDestinationsResponse200ResponseBody - OK
type CreateDestinationDestinationsResponse200ResponseBody struct {
	// Creates a new Destination.
	Data *shared.CreateDestinationV1Output `json:"data,omitempty"`
}

func (o *CreateDestinationDestinationsResponse200ResponseBody) GetData() *shared.CreateDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateDestinationDestinationsResponseResponseBody - OK
type CreateDestinationDestinationsResponseResponseBody struct {
	// Creates a new Destination.
	Data *shared.CreateDestinationV1Output `json:"data,omitempty"`
}

func (o *CreateDestinationDestinationsResponseResponseBody) GetData() *shared.CreateDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateDestinationDestinationsResponseBody - OK
type CreateDestinationDestinationsResponseBody struct {
	// Creates a new Destination.
	Data *shared.CreateDestinationV1Output `json:"data,omitempty"`
}

func (o *CreateDestinationDestinationsResponseBody) GetData() *shared.CreateDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateDestinationResponseBody - OK
type CreateDestinationResponseBody struct {
	// Creates a new Destination.
	Data *shared.CreateDestinationV1Output `json:"data,omitempty"`
}

func (o *CreateDestinationResponseBody) GetData() *shared.CreateDestinationV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateDestinationResponse struct {
	// OK
	TwoHundredApplicationJSONObject *CreateDestinationResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *CreateDestinationDestinationsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *CreateDestinationDestinationsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *CreateDestinationDestinationsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateDestinationResponse) GetTwoHundredApplicationJSONObject() *CreateDestinationResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *CreateDestinationResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *CreateDestinationDestinationsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *CreateDestinationResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *CreateDestinationDestinationsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *CreateDestinationResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *CreateDestinationDestinationsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *CreateDestinationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateDestinationResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *CreateDestinationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateDestinationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

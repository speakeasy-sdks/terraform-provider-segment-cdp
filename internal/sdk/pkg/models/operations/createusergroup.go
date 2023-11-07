// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

// CreateUserGroupIAMGroupsResponse200ResponseBody - OK
type CreateUserGroupIAMGroupsResponse200ResponseBody struct {
	// Returns the newly created user group.
	Data *shared.CreateUserGroupV1Output `json:"data,omitempty"`
}

func (o *CreateUserGroupIAMGroupsResponse200ResponseBody) GetData() *shared.CreateUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateUserGroupIAMGroupsResponseResponseBody - OK
type CreateUserGroupIAMGroupsResponseResponseBody struct {
	// Returns the newly created user group.
	Data *shared.CreateUserGroupV1Output `json:"data,omitempty"`
}

func (o *CreateUserGroupIAMGroupsResponseResponseBody) GetData() *shared.CreateUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateUserGroupIAMGroupsResponseBody - OK
type CreateUserGroupIAMGroupsResponseBody struct {
	// Returns the newly created user group.
	Data *shared.CreateUserGroupV1Output `json:"data,omitempty"`
}

func (o *CreateUserGroupIAMGroupsResponseBody) GetData() *shared.CreateUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateUserGroupResponseBody - OK
type CreateUserGroupResponseBody struct {
	// Returns the newly created user group.
	Data *shared.CreateUserGroupV1Output `json:"data,omitempty"`
}

func (o *CreateUserGroupResponseBody) GetData() *shared.CreateUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateUserGroupResponse struct {
	// OK
	TwoHundredApplicationJSONObject *CreateUserGroupResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *CreateUserGroupIAMGroupsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *CreateUserGroupIAMGroupsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *CreateUserGroupIAMGroupsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateUserGroupResponse) GetTwoHundredApplicationJSONObject() *CreateUserGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *CreateUserGroupResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *CreateUserGroupIAMGroupsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *CreateUserGroupResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *CreateUserGroupIAMGroupsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *CreateUserGroupResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *CreateUserGroupIAMGroupsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *CreateUserGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateUserGroupResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *CreateUserGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateUserGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

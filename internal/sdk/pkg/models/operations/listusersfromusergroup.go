// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListUsersFromUserGroupRequest struct {
	// Pagination for members of a group.
	//
	// This parameter exists in v1.
	Pagination  shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	UserGroupID string                 `pathParam:"style=simple,explode=false,name=userGroupId"`
}

func (o *ListUsersFromUserGroupRequest) GetPagination() shared.PaginationInput {
	if o == nil {
		return shared.PaginationInput{}
	}
	return o.Pagination
}

func (o *ListUsersFromUserGroupRequest) GetUserGroupID() string {
	if o == nil {
		return ""
	}
	return o.UserGroupID
}

// ListUsersFromUserGroupIAMGroupsResponse200ResponseBody - OK
type ListUsersFromUserGroupIAMGroupsResponse200ResponseBody struct {
	// Returns the members of a user group with the given group id.
	Data *shared.ListUsersFromUserGroupV1Output `json:"data,omitempty"`
}

func (o *ListUsersFromUserGroupIAMGroupsResponse200ResponseBody) GetData() *shared.ListUsersFromUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListUsersFromUserGroupIAMGroupsResponseResponseBody - OK
type ListUsersFromUserGroupIAMGroupsResponseResponseBody struct {
	// Returns the members of a user group with the given group id.
	Data *shared.ListUsersFromUserGroupV1Output `json:"data,omitempty"`
}

func (o *ListUsersFromUserGroupIAMGroupsResponseResponseBody) GetData() *shared.ListUsersFromUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListUsersFromUserGroupIAMGroupsResponseBody - OK
type ListUsersFromUserGroupIAMGroupsResponseBody struct {
	// Returns the members of a user group with the given group id.
	Data *shared.ListUsersFromUserGroupV1Output `json:"data,omitempty"`
}

func (o *ListUsersFromUserGroupIAMGroupsResponseBody) GetData() *shared.ListUsersFromUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ListUsersFromUserGroupResponseBody - OK
type ListUsersFromUserGroupResponseBody struct {
	// Returns the members of a user group with the given group id.
	Data *shared.ListUsersFromUserGroupV1Output `json:"data,omitempty"`
}

func (o *ListUsersFromUserGroupResponseBody) GetData() *shared.ListUsersFromUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type ListUsersFromUserGroupResponse struct {
	// OK
	TwoHundredApplicationJSONObject *ListUsersFromUserGroupResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *ListUsersFromUserGroupIAMGroupsResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *ListUsersFromUserGroupIAMGroupsResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *ListUsersFromUserGroupIAMGroupsResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListUsersFromUserGroupResponse) GetTwoHundredApplicationJSONObject() *ListUsersFromUserGroupResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *ListUsersFromUserGroupResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *ListUsersFromUserGroupIAMGroupsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *ListUsersFromUserGroupResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *ListUsersFromUserGroupIAMGroupsResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *ListUsersFromUserGroupResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *ListUsersFromUserGroupIAMGroupsResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *ListUsersFromUserGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListUsersFromUserGroupResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ListUsersFromUserGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListUsersFromUserGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListUserGroupsRequest struct {
	// Pagination for user groups.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

// ListUserGroups200ApplicationVndSegmentV1betaPlusJSON - OK
type ListUserGroups200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a list of user groups with the given Workspace id.
	Data *shared.ListUserGroupsV1Output `json:"data,omitempty"`
}

// ListUserGroups200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListUserGroups200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a list of user groups with the given Workspace id.
	Data *shared.ListUserGroupsV1Output `json:"data,omitempty"`
}

// ListUserGroups200ApplicationVndSegmentV1PlusJSON - OK
type ListUserGroups200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a list of user groups with the given Workspace id.
	Data *shared.ListUserGroupsV1Output `json:"data,omitempty"`
}

// ListUserGroups200ApplicationJSON - OK
type ListUserGroups200ApplicationJSON struct {
	// Returns a list of user groups with the given Workspace id.
	Data *shared.ListUserGroupsV1Output `json:"data,omitempty"`
}

type ListUserGroupsResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	ListUserGroups200ApplicationJSONObject *ListUserGroups200ApplicationJSON
	// OK
	ListUserGroups200ApplicationVndSegmentV1PlusJSONObject *ListUserGroups200ApplicationVndSegmentV1PlusJSON
	// OK
	ListUserGroups200ApplicationVndSegmentV1alphaPlusJSONObject *ListUserGroups200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListUserGroups200ApplicationVndSegmentV1betaPlusJSONObject *ListUserGroups200ApplicationVndSegmentV1betaPlusJSON
}

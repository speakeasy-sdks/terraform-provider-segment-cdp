// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListUserGroupsFromUserRequest struct {
	// Pagination for groups.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	UserID     string                 `pathParam:"style=simple,explode=false,name=userId"`
}

// ListUserGroupsFromUser200ApplicationVndSegmentV1betaPlusJSON - OK
type ListUserGroupsFromUser200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns all user groups the user belongs to.
	Data *shared.ListUserGroupsFromUserV1Output `json:"data,omitempty"`
}

// ListUserGroupsFromUser200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListUserGroupsFromUser200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns all user groups the user belongs to.
	Data *shared.ListUserGroupsFromUserV1Output `json:"data,omitempty"`
}

// ListUserGroupsFromUser200ApplicationVndSegmentV1PlusJSON - OK
type ListUserGroupsFromUser200ApplicationVndSegmentV1PlusJSON struct {
	// Returns all user groups the user belongs to.
	Data *shared.ListUserGroupsFromUserV1Output `json:"data,omitempty"`
}

// ListUserGroupsFromUser200ApplicationJSON - OK
type ListUserGroupsFromUser200ApplicationJSON struct {
	// Returns all user groups the user belongs to.
	Data *shared.ListUserGroupsFromUserV1Output `json:"data,omitempty"`
}

type ListUserGroupsFromUserResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	ListUserGroupsFromUser200ApplicationJSONObject *ListUserGroupsFromUser200ApplicationJSON
	// OK
	ListUserGroupsFromUser200ApplicationVndSegmentV1PlusJSONObject *ListUserGroupsFromUser200ApplicationVndSegmentV1PlusJSON
	// OK
	ListUserGroupsFromUser200ApplicationVndSegmentV1alphaPlusJSONObject *ListUserGroupsFromUser200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListUserGroupsFromUser200ApplicationVndSegmentV1betaPlusJSONObject *ListUserGroupsFromUser200ApplicationVndSegmentV1betaPlusJSON
}

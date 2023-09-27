// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListInvitesFromUserGroupRequest struct {
	// Pagination for invites to the group.
	//
	// This parameter exists in v1.
	Pagination  shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
	UserGroupID string                 `pathParam:"style=simple,explode=false,name=userGroupId"`
}

// ListInvitesFromUserGroup200ApplicationVndSegmentV1betaPlusJSON - OK
type ListInvitesFromUserGroup200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the emails of invitees to a user group with the given group id.
	Data *shared.ListInvitesFromUserGroupV1Output `json:"data,omitempty"`
}

// ListInvitesFromUserGroup200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListInvitesFromUserGroup200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the emails of invitees to a user group with the given group id.
	Data *shared.ListInvitesFromUserGroupV1Output `json:"data,omitempty"`
}

// ListInvitesFromUserGroup200ApplicationVndSegmentV1PlusJSON - OK
type ListInvitesFromUserGroup200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the emails of invitees to a user group with the given group id.
	Data *shared.ListInvitesFromUserGroupV1Output `json:"data,omitempty"`
}

// ListInvitesFromUserGroup200ApplicationJSON - OK
type ListInvitesFromUserGroup200ApplicationJSON struct {
	// Returns the emails of invitees to a user group with the given group id.
	Data *shared.ListInvitesFromUserGroupV1Output `json:"data,omitempty"`
}

type ListInvitesFromUserGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ListInvitesFromUserGroup200ApplicationJSONObject *ListInvitesFromUserGroup200ApplicationJSON
	// OK
	ListInvitesFromUserGroup200ApplicationVndSegmentV1PlusJSONObject *ListInvitesFromUserGroup200ApplicationVndSegmentV1PlusJSON
	// OK
	ListInvitesFromUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject *ListInvitesFromUserGroup200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListInvitesFromUserGroup200ApplicationVndSegmentV1betaPlusJSONObject *ListInvitesFromUserGroup200ApplicationVndSegmentV1betaPlusJSON
}
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ReplacePermissionsForUserGroupRequest struct {
	ReplacePermissionsForUserGroupV1Input shared.ReplacePermissionsForUserGroupV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	UserGroupID                           string                                       `pathParam:"style=simple,explode=false,name=userGroupId"`
}

// ReplacePermissionsForUserGroup200ApplicationVndSegmentV1betaPlusJSON - OK
type ReplacePermissionsForUserGroup200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the user group's permissions, including the updated permissions.
	Data *shared.ReplacePermissionsForUserGroupV1Output `json:"data,omitempty"`
}

// ReplacePermissionsForUserGroup200ApplicationVndSegmentV1alphaPlusJSON - OK
type ReplacePermissionsForUserGroup200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the user group's permissions, including the updated permissions.
	Data *shared.ReplacePermissionsForUserGroupV1Output `json:"data,omitempty"`
}

// ReplacePermissionsForUserGroup200ApplicationVndSegmentV1PlusJSON - OK
type ReplacePermissionsForUserGroup200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the user group's permissions, including the updated permissions.
	Data *shared.ReplacePermissionsForUserGroupV1Output `json:"data,omitempty"`
}

// ReplacePermissionsForUserGroup200ApplicationJSON - OK
type ReplacePermissionsForUserGroup200ApplicationJSON struct {
	// Returns the user group's permissions, including the updated permissions.
	Data *shared.ReplacePermissionsForUserGroupV1Output `json:"data,omitempty"`
}

type ReplacePermissionsForUserGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ReplacePermissionsForUserGroup200ApplicationJSONObject *ReplacePermissionsForUserGroup200ApplicationJSON
	// OK
	ReplacePermissionsForUserGroup200ApplicationVndSegmentV1PlusJSONObject *ReplacePermissionsForUserGroup200ApplicationVndSegmentV1PlusJSON
	// OK
	ReplacePermissionsForUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject *ReplacePermissionsForUserGroup200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ReplacePermissionsForUserGroup200ApplicationVndSegmentV1betaPlusJSONObject *ReplacePermissionsForUserGroup200ApplicationVndSegmentV1betaPlusJSON
}

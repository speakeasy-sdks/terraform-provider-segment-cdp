// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetUserGroupRequest struct {
	UserGroupID string `pathParam:"style=simple,explode=false,name=userGroupId"`
}

// GetUserGroup200ApplicationVndSegmentV1betaPlusJSON - OK
type GetUserGroup200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a user group with the given id.
	Data *shared.GetUserGroupV1Output `json:"data,omitempty"`
}

// GetUserGroup200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetUserGroup200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a user group with the given id.
	Data *shared.GetUserGroupV1Output `json:"data,omitempty"`
}

// GetUserGroup200ApplicationVndSegmentV1PlusJSON - OK
type GetUserGroup200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a user group with the given id.
	Data *shared.GetUserGroupV1Output `json:"data,omitempty"`
}

// GetUserGroup200ApplicationJSON - OK
type GetUserGroup200ApplicationJSON struct {
	// Returns a user group with the given id.
	Data *shared.GetUserGroupV1Output `json:"data,omitempty"`
}

type GetUserGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	GetUserGroup200ApplicationJSONObject *GetUserGroup200ApplicationJSON
	// OK
	GetUserGroup200ApplicationVndSegmentV1PlusJSONObject *GetUserGroup200ApplicationVndSegmentV1PlusJSON
	// OK
	GetUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject *GetUserGroup200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	GetUserGroup200ApplicationVndSegmentV1betaPlusJSONObject *GetUserGroup200ApplicationVndSegmentV1betaPlusJSON
}

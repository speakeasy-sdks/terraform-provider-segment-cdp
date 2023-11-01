// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type UpdateUserGroupRequest struct {
	UpdateUserGroupV1Input shared.UpdateUserGroupV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	UserGroupID            string                        `pathParam:"style=simple,explode=false,name=userGroupId"`
}

func (o *UpdateUserGroupRequest) GetUpdateUserGroupV1Input() shared.UpdateUserGroupV1Input {
	if o == nil {
		return shared.UpdateUserGroupV1Input{}
	}
	return o.UpdateUserGroupV1Input
}

func (o *UpdateUserGroupRequest) GetUserGroupID() string {
	if o == nil {
		return ""
	}
	return o.UserGroupID
}

// UpdateUserGroup200ApplicationVndSegmentV1betaPlusJSON - OK
type UpdateUserGroup200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the updated user group.
	Data *shared.UpdateUserGroupV1Output `json:"data,omitempty"`
}

func (o *UpdateUserGroup200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.UpdateUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateUserGroup200ApplicationVndSegmentV1alphaPlusJSON - OK
type UpdateUserGroup200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the updated user group.
	Data *shared.UpdateUserGroupV1Output `json:"data,omitempty"`
}

func (o *UpdateUserGroup200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.UpdateUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateUserGroup200ApplicationVndSegmentV1PlusJSON - OK
type UpdateUserGroup200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the updated user group.
	Data *shared.UpdateUserGroupV1Output `json:"data,omitempty"`
}

func (o *UpdateUserGroup200ApplicationVndSegmentV1PlusJSON) GetData() *shared.UpdateUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// UpdateUserGroup200ApplicationJSON - OK
type UpdateUserGroup200ApplicationJSON struct {
	// Returns the updated user group.
	Data *shared.UpdateUserGroupV1Output `json:"data,omitempty"`
}

func (o *UpdateUserGroup200ApplicationJSON) GetData() *shared.UpdateUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateUserGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UpdateUserGroup200ApplicationJSONObject *UpdateUserGroup200ApplicationJSON
	// OK
	UpdateUserGroup200ApplicationVndSegmentV1PlusJSONObject *UpdateUserGroup200ApplicationVndSegmentV1PlusJSON
	// OK
	UpdateUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject *UpdateUserGroup200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	UpdateUserGroup200ApplicationVndSegmentV1betaPlusJSONObject *UpdateUserGroup200ApplicationVndSegmentV1betaPlusJSON
}

func (o *UpdateUserGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateUserGroupResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *UpdateUserGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateUserGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateUserGroupResponse) GetUpdateUserGroup200ApplicationJSONObject() *UpdateUserGroup200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateUserGroup200ApplicationJSONObject
}

func (o *UpdateUserGroupResponse) GetUpdateUserGroup200ApplicationVndSegmentV1PlusJSONObject() *UpdateUserGroup200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.UpdateUserGroup200ApplicationVndSegmentV1PlusJSONObject
}

func (o *UpdateUserGroupResponse) GetUpdateUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject() *UpdateUserGroup200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.UpdateUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *UpdateUserGroupResponse) GetUpdateUserGroup200ApplicationVndSegmentV1betaPlusJSONObject() *UpdateUserGroup200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.UpdateUserGroup200ApplicationVndSegmentV1betaPlusJSONObject
}

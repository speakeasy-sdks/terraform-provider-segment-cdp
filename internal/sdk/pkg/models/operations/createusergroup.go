// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

// CreateUserGroup200ApplicationVndSegmentV1betaPlusJSON - OK
type CreateUserGroup200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns the newly created user group.
	Data *shared.CreateUserGroupV1Output `json:"data,omitempty"`
}

func (o *CreateUserGroup200ApplicationVndSegmentV1betaPlusJSON) GetData() *shared.CreateUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateUserGroup200ApplicationVndSegmentV1alphaPlusJSON - OK
type CreateUserGroup200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns the newly created user group.
	Data *shared.CreateUserGroupV1Output `json:"data,omitempty"`
}

func (o *CreateUserGroup200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.CreateUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateUserGroup200ApplicationVndSegmentV1PlusJSON - OK
type CreateUserGroup200ApplicationVndSegmentV1PlusJSON struct {
	// Returns the newly created user group.
	Data *shared.CreateUserGroupV1Output `json:"data,omitempty"`
}

func (o *CreateUserGroup200ApplicationVndSegmentV1PlusJSON) GetData() *shared.CreateUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// CreateUserGroup200ApplicationJSON - OK
type CreateUserGroup200ApplicationJSON struct {
	// Returns the newly created user group.
	Data *shared.CreateUserGroupV1Output `json:"data,omitempty"`
}

func (o *CreateUserGroup200ApplicationJSON) GetData() *shared.CreateUserGroupV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateUserGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	CreateUserGroup200ApplicationJSONObject *CreateUserGroup200ApplicationJSON
	// OK
	CreateUserGroup200ApplicationVndSegmentV1PlusJSONObject *CreateUserGroup200ApplicationVndSegmentV1PlusJSON
	// OK
	CreateUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject *CreateUserGroup200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	CreateUserGroup200ApplicationVndSegmentV1betaPlusJSONObject *CreateUserGroup200ApplicationVndSegmentV1betaPlusJSON
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

func (o *CreateUserGroupResponse) GetCreateUserGroup200ApplicationJSONObject() *CreateUserGroup200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateUserGroup200ApplicationJSONObject
}

func (o *CreateUserGroupResponse) GetCreateUserGroup200ApplicationVndSegmentV1PlusJSONObject() *CreateUserGroup200ApplicationVndSegmentV1PlusJSON {
	if o == nil {
		return nil
	}
	return o.CreateUserGroup200ApplicationVndSegmentV1PlusJSONObject
}

func (o *CreateUserGroupResponse) GetCreateUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject() *CreateUserGroup200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.CreateUserGroup200ApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *CreateUserGroupResponse) GetCreateUserGroup200ApplicationVndSegmentV1betaPlusJSONObject() *CreateUserGroup200ApplicationVndSegmentV1betaPlusJSON {
	if o == nil {
		return nil
	}
	return o.CreateUserGroup200ApplicationVndSegmentV1betaPlusJSONObject
}

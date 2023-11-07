// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ReplacePermissionsForUserRequest struct {
	ReplacePermissionsForUserV1Input shared.ReplacePermissionsForUserV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	UserID                           string                                  `pathParam:"style=simple,explode=false,name=userId"`
}

func (o *ReplacePermissionsForUserRequest) GetReplacePermissionsForUserV1Input() shared.ReplacePermissionsForUserV1Input {
	if o == nil {
		return shared.ReplacePermissionsForUserV1Input{}
	}
	return o.ReplacePermissionsForUserV1Input
}

func (o *ReplacePermissionsForUserRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

// ReplacePermissionsForUserIAMUsersResponse200ResponseBody - OK
type ReplacePermissionsForUserIAMUsersResponse200ResponseBody struct {
	// Returns the user's permissions, including the updated permissions.
	Data *shared.ReplacePermissionsForUserV1Output `json:"data,omitempty"`
}

func (o *ReplacePermissionsForUserIAMUsersResponse200ResponseBody) GetData() *shared.ReplacePermissionsForUserV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ReplacePermissionsForUserIAMUsersResponseResponseBody - OK
type ReplacePermissionsForUserIAMUsersResponseResponseBody struct {
	// Returns the user's permissions, including the updated permissions.
	Data *shared.ReplacePermissionsForUserV1Output `json:"data,omitempty"`
}

func (o *ReplacePermissionsForUserIAMUsersResponseResponseBody) GetData() *shared.ReplacePermissionsForUserV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ReplacePermissionsForUserIAMUsersResponseBody - OK
type ReplacePermissionsForUserIAMUsersResponseBody struct {
	// Returns the user's permissions, including the updated permissions.
	Data *shared.ReplacePermissionsForUserV1Output `json:"data,omitempty"`
}

func (o *ReplacePermissionsForUserIAMUsersResponseBody) GetData() *shared.ReplacePermissionsForUserV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// ReplacePermissionsForUserResponseBody - OK
type ReplacePermissionsForUserResponseBody struct {
	// Returns the user's permissions, including the updated permissions.
	Data *shared.ReplacePermissionsForUserV1Output `json:"data,omitempty"`
}

func (o *ReplacePermissionsForUserResponseBody) GetData() *shared.ReplacePermissionsForUserV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type ReplacePermissionsForUserResponse struct {
	// OK
	TwoHundredApplicationJSONObject *ReplacePermissionsForUserResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *ReplacePermissionsForUserIAMUsersResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *ReplacePermissionsForUserIAMUsersResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *ReplacePermissionsForUserIAMUsersResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ReplacePermissionsForUserResponse) GetTwoHundredApplicationJSONObject() *ReplacePermissionsForUserResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *ReplacePermissionsForUserResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *ReplacePermissionsForUserIAMUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *ReplacePermissionsForUserResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *ReplacePermissionsForUserIAMUsersResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *ReplacePermissionsForUserResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *ReplacePermissionsForUserIAMUsersResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *ReplacePermissionsForUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ReplacePermissionsForUserResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *ReplacePermissionsForUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ReplacePermissionsForUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

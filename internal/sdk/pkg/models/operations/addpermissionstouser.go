// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"net/http"
)

type AddPermissionsToUserRequest struct {
	AddPermissionsToUserV1Input shared.AddPermissionsToUserV1Input `request:"mediaType=application/vnd.segment.v1beta+json"`
	UserID                      string                             `pathParam:"style=simple,explode=false,name=userId"`
}

func (o *AddPermissionsToUserRequest) GetAddPermissionsToUserV1Input() shared.AddPermissionsToUserV1Input {
	if o == nil {
		return shared.AddPermissionsToUserV1Input{}
	}
	return o.AddPermissionsToUserV1Input
}

func (o *AddPermissionsToUserRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}

// AddPermissionsToUserIAMUsersResponse200ResponseBody - OK
type AddPermissionsToUserIAMUsersResponse200ResponseBody struct {
	// Returns the user's permissions, including the added permissions.
	Data *shared.AddPermissionsToUserV1Output `json:"data,omitempty"`
}

func (o *AddPermissionsToUserIAMUsersResponse200ResponseBody) GetData() *shared.AddPermissionsToUserV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// AddPermissionsToUserIAMUsersResponseResponseBody - OK
type AddPermissionsToUserIAMUsersResponseResponseBody struct {
	// Returns the user's permissions, including the added permissions.
	Data *shared.AddPermissionsToUserV1Output `json:"data,omitempty"`
}

func (o *AddPermissionsToUserIAMUsersResponseResponseBody) GetData() *shared.AddPermissionsToUserV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// AddPermissionsToUserIAMUsersResponseBody - OK
type AddPermissionsToUserIAMUsersResponseBody struct {
	// Returns the user's permissions, including the added permissions.
	Data *shared.AddPermissionsToUserV1Output `json:"data,omitempty"`
}

func (o *AddPermissionsToUserIAMUsersResponseBody) GetData() *shared.AddPermissionsToUserV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// AddPermissionsToUserResponseBody - OK
type AddPermissionsToUserResponseBody struct {
	// Returns the user's permissions, including the added permissions.
	Data *shared.AddPermissionsToUserV1Output `json:"data,omitempty"`
}

func (o *AddPermissionsToUserResponseBody) GetData() *shared.AddPermissionsToUserV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type AddPermissionsToUserResponse struct {
	// OK
	TwoHundredApplicationJSONObject *AddPermissionsToUserResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *AddPermissionsToUserIAMUsersResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *AddPermissionsToUserIAMUsersResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *AddPermissionsToUserIAMUsersResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AddPermissionsToUserResponse) GetTwoHundredApplicationJSONObject() *AddPermissionsToUserResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *AddPermissionsToUserResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *AddPermissionsToUserIAMUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *AddPermissionsToUserResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *AddPermissionsToUserIAMUsersResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *AddPermissionsToUserResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *AddPermissionsToUserIAMUsersResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *AddPermissionsToUserResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AddPermissionsToUserResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *AddPermissionsToUserResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AddPermissionsToUserResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

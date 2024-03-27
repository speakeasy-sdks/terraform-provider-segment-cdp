// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/models/shared"
	"net/http"
)

type DeleteInvitesRequest struct {
	// The list of emails to delete invites for.
	//
	// This parameter exists in v1.
	Emails []string `queryParam:"style=deepObject,explode=true,name=emails"`
}

func (o *DeleteInvitesRequest) GetEmails() []string {
	if o == nil {
		return []string{}
	}
	return o.Emails
}

// DeleteInvitesIAMUsersResponse200ResponseBody - OK
type DeleteInvitesIAMUsersResponse200ResponseBody struct {
	// Returns the status of the removal operation.
	Data *shared.DeleteInvitesV1Output `json:"data,omitempty"`
}

func (o *DeleteInvitesIAMUsersResponse200ResponseBody) GetData() *shared.DeleteInvitesV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteInvitesIAMUsersResponseResponseBody - OK
type DeleteInvitesIAMUsersResponseResponseBody struct {
	// Returns the status of the removal operation.
	Data *shared.DeleteInvitesV1Output `json:"data,omitempty"`
}

func (o *DeleteInvitesIAMUsersResponseResponseBody) GetData() *shared.DeleteInvitesV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteInvitesIAMUsersResponseBody - OK
type DeleteInvitesIAMUsersResponseBody struct {
	// Returns the status of the removal operation.
	Data *shared.DeleteInvitesV1Output `json:"data,omitempty"`
}

func (o *DeleteInvitesIAMUsersResponseBody) GetData() *shared.DeleteInvitesV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

// DeleteInvitesResponseBody - OK
type DeleteInvitesResponseBody struct {
	// Returns the status of the removal operation.
	Data *shared.DeleteInvitesV1Output `json:"data,omitempty"`
}

func (o *DeleteInvitesResponseBody) GetData() *shared.DeleteInvitesV1Output {
	if o == nil {
		return nil
	}
	return o.Data
}

type DeleteInvitesResponse struct {
	// OK
	TwoHundredApplicationJSONObject *DeleteInvitesResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1PlusJSONObject *DeleteInvitesIAMUsersResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1alphaPlusJSONObject *DeleteInvitesIAMUsersResponseResponseBody
	// OK
	TwoHundredApplicationVndSegmentV1betaPlusJSONObject *DeleteInvitesIAMUsersResponse200ResponseBody
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteInvitesResponse) GetTwoHundredApplicationJSONObject() *DeleteInvitesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *DeleteInvitesResponse) GetTwoHundredApplicationVndSegmentV1PlusJSONObject() *DeleteInvitesIAMUsersResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1PlusJSONObject
}

func (o *DeleteInvitesResponse) GetTwoHundredApplicationVndSegmentV1alphaPlusJSONObject() *DeleteInvitesIAMUsersResponseResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject
}

func (o *DeleteInvitesResponse) GetTwoHundredApplicationVndSegmentV1betaPlusJSONObject() *DeleteInvitesIAMUsersResponse200ResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationVndSegmentV1betaPlusJSONObject
}

func (o *DeleteInvitesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteInvitesResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *DeleteInvitesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteInvitesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
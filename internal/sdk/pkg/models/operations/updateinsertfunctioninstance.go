// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"net/http"
)

type UpdateInsertFunctionInstanceRequest struct {
	UpdateInsertFunctionInstanceAlphaInput shared.UpdateInsertFunctionInstanceAlphaInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	InstanceID                             string                                        `pathParam:"style=simple,explode=false,name=instanceId"`
}

func (o *UpdateInsertFunctionInstanceRequest) GetUpdateInsertFunctionInstanceAlphaInput() shared.UpdateInsertFunctionInstanceAlphaInput {
	if o == nil {
		return shared.UpdateInsertFunctionInstanceAlphaInput{}
	}
	return o.UpdateInsertFunctionInstanceAlphaInput
}

func (o *UpdateInsertFunctionInstanceRequest) GetInstanceID() string {
	if o == nil {
		return ""
	}
	return o.InstanceID
}

// UpdateInsertFunctionInstanceResponseBody - OK
type UpdateInsertFunctionInstanceResponseBody struct {
	// Returns the updated insert Function instance.
	Data *shared.UpdateInsertFunctionInstanceAlphaOutput `json:"data,omitempty"`
}

func (o *UpdateInsertFunctionInstanceResponseBody) GetData() *shared.UpdateInsertFunctionInstanceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateInsertFunctionInstanceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	Object *UpdateInsertFunctionInstanceResponseBody
}

func (o *UpdateInsertFunctionInstanceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateInsertFunctionInstanceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateInsertFunctionInstanceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateInsertFunctionInstanceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *UpdateInsertFunctionInstanceResponse) GetObject() *UpdateInsertFunctionInstanceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

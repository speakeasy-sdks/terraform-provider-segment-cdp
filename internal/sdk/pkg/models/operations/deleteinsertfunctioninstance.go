// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type DeleteInsertFunctionInstanceRequest struct {
	InstanceID string `pathParam:"style=simple,explode=false,name=instanceId"`
}

func (o *DeleteInsertFunctionInstanceRequest) GetInstanceID() string {
	if o == nil {
		return ""
	}
	return o.InstanceID
}

// DeleteInsertFunctionInstanceResponseBody - OK
type DeleteInsertFunctionInstanceResponseBody struct {
	// Delete an insert Function instance.
	Data *shared.DeleteInsertFunctionInstanceAlphaOutput `json:"data,omitempty"`
}

func (o *DeleteInsertFunctionInstanceResponseBody) GetData() *shared.DeleteInsertFunctionInstanceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type DeleteInsertFunctionInstanceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	Object *DeleteInsertFunctionInstanceResponseBody
}

func (o *DeleteInsertFunctionInstanceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteInsertFunctionInstanceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteInsertFunctionInstanceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteInsertFunctionInstanceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *DeleteInsertFunctionInstanceResponse) GetObject() *DeleteInsertFunctionInstanceResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}

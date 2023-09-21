// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

// CreateInsertFunctionInstance200ApplicationVndSegmentV1alphaPlusJSON - OK
type CreateInsertFunctionInstance200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Creates an insert Function instance.
	Data *shared.CreateInsertFunctionInstanceAlphaOutput `json:"data,omitempty"`
}

type CreateInsertFunctionInstanceResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	CreateInsertFunctionInstance200ApplicationVndSegmentV1alphaPlusJSONObject *CreateInsertFunctionInstance200ApplicationVndSegmentV1alphaPlusJSON
}

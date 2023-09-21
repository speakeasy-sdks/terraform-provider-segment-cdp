// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetFunctionRequest struct {
	FunctionID string `pathParam:"style=simple,explode=false,name=functionId"`
}

// GetFunction200ApplicationVndSegmentV1betaPlusJSON - OK
type GetFunction200ApplicationVndSegmentV1betaPlusJSON struct {
	// Gets a single Function.
	Data *shared.GetFunctionV1Output `json:"data,omitempty"`
}

// GetFunction200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetFunction200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Gets a single Function.
	Data *shared.GetFunctionV1Output `json:"data,omitempty"`
}

// GetFunction200ApplicationVndSegmentV1PlusJSON - OK
type GetFunction200ApplicationVndSegmentV1PlusJSON struct {
	// Gets a single Function.
	Data *shared.GetFunctionV1Output `json:"data,omitempty"`
}

// GetFunction200ApplicationJSON - OK
type GetFunction200ApplicationJSON struct {
	// Gets a single Function.
	Data *shared.GetFunctionV1Output `json:"data,omitempty"`
}

type GetFunctionResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	GetFunction200ApplicationJSONObject *GetFunction200ApplicationJSON
	// OK
	GetFunction200ApplicationVndSegmentV1PlusJSONObject *GetFunction200ApplicationVndSegmentV1PlusJSON
	// OK
	GetFunction200ApplicationVndSegmentV1alphaPlusJSONObject *GetFunction200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	GetFunction200ApplicationVndSegmentV1betaPlusJSONObject *GetFunction200ApplicationVndSegmentV1betaPlusJSON
}

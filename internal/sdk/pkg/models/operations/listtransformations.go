// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListTransformationsRequest struct {
	// Pagination options.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

// ListTransformations200ApplicationVndSegmentV1betaPlusJSON - OK
type ListTransformations200ApplicationVndSegmentV1betaPlusJSON struct {
	// Lists the Transformations associated with the current Workspace.
	Data *shared.ListTransformationsV1Output `json:"data,omitempty"`
}

// ListTransformations200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListTransformations200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Lists the Transformations associated with the current Workspace.
	Data *shared.ListTransformationsV1Output `json:"data,omitempty"`
}

// ListTransformations200ApplicationVndSegmentV1PlusJSON - OK
type ListTransformations200ApplicationVndSegmentV1PlusJSON struct {
	// Lists the Transformations associated with the current Workspace.
	Data *shared.ListTransformationsV1Output `json:"data,omitempty"`
}

// ListTransformations200ApplicationJSON - OK
type ListTransformations200ApplicationJSON struct {
	// Lists the Transformations associated with the current Workspace.
	Data *shared.ListTransformationsV1Output `json:"data,omitempty"`
}

type ListTransformationsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	ListTransformations200ApplicationJSONObject *ListTransformations200ApplicationJSON
	// OK
	ListTransformations200ApplicationVndSegmentV1PlusJSONObject *ListTransformations200ApplicationVndSegmentV1PlusJSON
	// OK
	ListTransformations200ApplicationVndSegmentV1alphaPlusJSONObject *ListTransformations200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListTransformations200ApplicationVndSegmentV1betaPlusJSONObject *ListTransformations200ApplicationVndSegmentV1betaPlusJSON
}

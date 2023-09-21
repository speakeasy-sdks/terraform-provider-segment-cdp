// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type ListSuppressionsRequest struct {
	// Pagination parameters.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

// ListSuppressions200ApplicationVndSegmentV1betaPlusJSON - OK
type ListSuppressions200ApplicationVndSegmentV1betaPlusJSON struct {
	// The output of a list suppressed call for a Workspace.
	Data *shared.ListSuppressionsV1Output `json:"data,omitempty"`
}

// ListSuppressions200ApplicationVndSegmentV1alphaPlusJSON - OK
type ListSuppressions200ApplicationVndSegmentV1alphaPlusJSON struct {
	// The output of a list suppressed call for a Workspace.
	Data *shared.ListSuppressionsV1Output `json:"data,omitempty"`
}

// ListSuppressions200ApplicationVndSegmentV1PlusJSON - OK
type ListSuppressions200ApplicationVndSegmentV1PlusJSON struct {
	// The output of a list suppressed call for a Workspace.
	Data *shared.ListSuppressionsV1Output `json:"data,omitempty"`
}

// ListSuppressions200ApplicationJSON - OK
type ListSuppressions200ApplicationJSON struct {
	// The output of a list suppressed call for a Workspace.
	Data *shared.ListSuppressionsV1Output `json:"data,omitempty"`
}

type ListSuppressionsResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	ListSuppressions200ApplicationJSONObject *ListSuppressions200ApplicationJSON
	// OK
	ListSuppressions200ApplicationVndSegmentV1PlusJSONObject *ListSuppressions200ApplicationVndSegmentV1PlusJSON
	// OK
	ListSuppressions200ApplicationVndSegmentV1alphaPlusJSONObject *ListSuppressions200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	ListSuppressions200ApplicationVndSegmentV1betaPlusJSONObject *ListSuppressions200ApplicationVndSegmentV1betaPlusJSON
}

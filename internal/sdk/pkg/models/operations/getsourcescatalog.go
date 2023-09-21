// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type GetSourcesCatalogRequest struct {
	// Defines the pagination parameters.
	//
	// This parameter exists in v1.
	Pagination shared.PaginationInput `queryParam:"style=deepObject,explode=true,name=pagination"`
}

// GetSourcesCatalog200ApplicationVndSegmentV1betaPlusJSON - OK
type GetSourcesCatalog200ApplicationVndSegmentV1betaPlusJSON struct {
	// Returns a list of all Source catalog items contained within a given page.
	Data *shared.GetSourcesCatalogV1Output `json:"data,omitempty"`
}

// GetSourcesCatalog200ApplicationVndSegmentV1alphaPlusJSON - OK
type GetSourcesCatalog200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Returns a list of all Source catalog items contained within a given page.
	Data *shared.GetSourcesCatalogV1Output `json:"data,omitempty"`
}

// GetSourcesCatalog200ApplicationVndSegmentV1PlusJSON - OK
type GetSourcesCatalog200ApplicationVndSegmentV1PlusJSON struct {
	// Returns a list of all Source catalog items contained within a given page.
	Data *shared.GetSourcesCatalogV1Output `json:"data,omitempty"`
}

// GetSourcesCatalog200ApplicationJSON - OK
type GetSourcesCatalog200ApplicationJSON struct {
	// Returns a list of all Source catalog items contained within a given page.
	Data *shared.GetSourcesCatalogV1Output `json:"data,omitempty"`
}

type GetSourcesCatalogResponse struct {
	ContentType string
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	StatusCode           int
	RawResponse          *http.Response
	// OK
	GetSourcesCatalog200ApplicationJSONObject *GetSourcesCatalog200ApplicationJSON
	// OK
	GetSourcesCatalog200ApplicationVndSegmentV1PlusJSONObject *GetSourcesCatalog200ApplicationVndSegmentV1PlusJSON
	// OK
	GetSourcesCatalog200ApplicationVndSegmentV1alphaPlusJSONObject *GetSourcesCatalog200ApplicationVndSegmentV1alphaPlusJSON
	// OK
	GetSourcesCatalog200ApplicationVndSegmentV1betaPlusJSONObject *GetSourcesCatalog200ApplicationVndSegmentV1betaPlusJSON
}

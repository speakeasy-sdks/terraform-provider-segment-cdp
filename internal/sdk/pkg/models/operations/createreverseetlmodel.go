// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

// CreateReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSON - OK
type CreateReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Defines the results of creating a Model.
	Data *shared.CreateReverseEtlModelOutput `json:"data,omitempty"`
}

type CreateReverseEtlModelResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	CreateReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSONObject *CreateReverseEtlModel200ApplicationVndSegmentV1alphaPlusJSON
}

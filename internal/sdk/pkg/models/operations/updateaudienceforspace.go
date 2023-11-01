// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
)

type UpdateAudienceForSpaceRequest struct {
	UpdateAudienceForSpaceInput shared.UpdateAudienceForSpaceInput `request:"mediaType=application/vnd.segment.v1alpha+json"`
	ID                          string                             `pathParam:"style=simple,explode=false,name=id"`
	SpaceID                     string                             `pathParam:"style=simple,explode=false,name=spaceId"`
}

func (o *UpdateAudienceForSpaceRequest) GetUpdateAudienceForSpaceInput() shared.UpdateAudienceForSpaceInput {
	if o == nil {
		return shared.UpdateAudienceForSpaceInput{}
	}
	return o.UpdateAudienceForSpaceInput
}

func (o *UpdateAudienceForSpaceRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateAudienceForSpaceRequest) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

// UpdateAudienceForSpace200ApplicationVndSegmentV1alphaPlusJSON - OK
type UpdateAudienceForSpace200ApplicationVndSegmentV1alphaPlusJSON struct {
	// Audience output for get.
	Data *shared.UpdateAudienceForSpaceAlphaOutput `json:"data,omitempty"`
}

func (o *UpdateAudienceForSpace200ApplicationVndSegmentV1alphaPlusJSON) GetData() *shared.UpdateAudienceForSpaceAlphaOutput {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateAudienceForSpaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Resource not found
	RequestErrorEnvelope *shared.RequestErrorEnvelope
	// OK
	UpdateAudienceForSpace200ApplicationVndSegmentV1alphaPlusJSONObject *UpdateAudienceForSpace200ApplicationVndSegmentV1alphaPlusJSON
}

func (o *UpdateAudienceForSpaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateAudienceForSpaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateAudienceForSpaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateAudienceForSpaceResponse) GetRequestErrorEnvelope() *shared.RequestErrorEnvelope {
	if o == nil {
		return nil
	}
	return o.RequestErrorEnvelope
}

func (o *UpdateAudienceForSpaceResponse) GetUpdateAudienceForSpace200ApplicationVndSegmentV1alphaPlusJSONObject() *UpdateAudienceForSpace200ApplicationVndSegmentV1alphaPlusJSON {
	if o == nil {
		return nil
	}
	return o.UpdateAudienceForSpace200ApplicationVndSegmentV1alphaPlusJSONObject
}

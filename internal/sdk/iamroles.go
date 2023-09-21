// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/operations"
	"segment_public_api/internal/sdk/pkg/models/shared"
	"segment_public_api/internal/sdk/pkg/utils"
	"strings"
)

// iamRoles - A role gives a user access to resources within a Workspace. Roles are additive, and can combine to configure a custom policy for a Team Member or a Group. A policy is at least one role plus one resource applied to an individual user or group.
type iamRoles struct {
	sdkConfiguration sdkConfiguration
}

func newIAMRoles(sdkConfig sdkConfiguration) *iamRoles {
	return &iamRoles{
		sdkConfiguration: sdkConfig,
	}
}

// ListRoles - List Roles
// Returns a list of Roles available to apply to permissions for users and/or groups.
func (s *iamRoles) ListRoles(ctx context.Context, request operations.ListRolesRequest) (*operations.ListRolesResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/roles"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s %s", s.sdkConfiguration.Language, s.sdkConfiguration.SDKVersion, s.sdkConfiguration.GenVersion, s.sdkConfiguration.OpenAPIDocVersion))

	if err := utils.PopulateQueryParams(ctx, req, request, nil); err != nil {
		return nil, fmt.Errorf("error populating query params: %w", err)
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListRolesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListRoles200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListRoles200ApplicationJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out *operations.ListRoles200ApplicationVndSegmentV1PlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListRoles200ApplicationVndSegmentV1PlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out *operations.ListRoles200ApplicationVndSegmentV1alphaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListRoles200ApplicationVndSegmentV1alphaPlusJSONObject = out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out *operations.ListRoles200ApplicationVndSegmentV1betaPlusJSON
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.ListRoles200ApplicationVndSegmentV1betaPlusJSONObject = out
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out); err != nil {
				return res, err
			}

			res.RequestErrorEnvelope = out
		}
	}

	return res, nil
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/internal/hooks"
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/internal/utils"
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/models/errors"
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/models/operations"
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/models/shared"
	"io"
	"net/http"
	"net/url"
)

// Workspaces - A Workspace is a group of Sources that you administer and Segment bills together. Workspaces help companies manage access for different users and data Sources and let you collaborate with team members, add permissions, and share Sources across your team using a shared billing account.
//
// When you first log in to your Segment account, you can create a new Workspace, or choose to log into an existing Workspace if your account is part of an existing organization.
//
// As the Segment Public API scopes tokens to a Workspace, all operations within the API are also limited to the Workspace to which that token belongs.
//
// ## Migrate from the Config API
//
// Like the Segment Public API, the Config API has one endpoint to retrieve details about a Workspace. The [getWorkspace endpoint](https://reference.segmentapis.com/#7ed2968b-c4a5-4cfb-b4bf-7d28c7b38bd2) returns the following fields:
//
// | Config API     | Public API                           |
// | -------------- | ------------------------------------ |
// | `create_time`  | Not returned                         |
// | `display_name` | `name`                               |
// | `id`           | `id`                                 |
// | `name`         | `slug` (`workspace/` prefix removed) |
//
// To migrate, replace any use of the Config API endpoints with the Segment Public API counterparts, using the field mappings in the table above.
type Workspaces struct {
	sdkConfiguration sdkConfiguration
}

func newWorkspaces(sdkConfig sdkConfiguration) *Workspaces {
	return &Workspaces{
		sdkConfiguration: sdkConfig,
	}
}

// GetWorkspace - Get Workspace
// Returns the Workspace associated with the token used to access this resource.
func (s *Workspaces) GetWorkspace(ctx context.Context, opts ...operations.Option) (*operations.GetWorkspaceResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "getWorkspace",
		SecuritySource: s.sdkConfiguration.Security,
	}

	o := operations.Options{}
	supportedOptions := []string{
		operations.SupportedOptionAcceptHeaderOverride,
	}

	for _, opt := range opts {
		if err := opt(&o, supportedOptions...); err != nil {
			return nil, fmt.Errorf("error applying option: %w", err)
		}
	}
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
	}

	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.GetWorkspaceResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: httpRes.Header.Get("Content-Type"),
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out operations.GetWorkspaceResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationJSONObject = &out
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/vnd.segment.v1+json`):
			var out operations.GetWorkspaceWorkspacesResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1PlusJSONObject = &out
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/vnd.segment.v1alpha+json`):
			var out operations.GetWorkspaceWorkspacesResponseResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject = &out
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/vnd.segment.v1beta+json`):
			var out operations.GetWorkspaceWorkspacesResponse200ResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1betaPlusJSONObject = &out
		default:
			return nil, errors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.RequestErrorEnvelope = &out
		default:
			return nil, errors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	default:
		return nil, errors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}

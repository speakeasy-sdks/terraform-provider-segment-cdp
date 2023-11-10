// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/operations"
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/sdkerrors"
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/models/shared"
	"github.com/ds-terraform/terraform-provider-segment_public_api/internal/sdk/pkg/utils"
	"io"
	"net/http"
	"strings"
)

// ReverseETL - Reverse ETL allows the use of a database (aka: Segment Warehouse) as a source of
// data to be connected and sent to supported Segment Destinations. Previously, it
// was only possible to use a Segment Warehouse as a destination.
//
// ## Sync schedule JSON configuration
//
// The structure of `scheduleConfig` varies with the value in `scheduleStrategy`.
// All strategies will employ a key/value object with _minimum_ depth of 1.
//
// ### Strategy: "manual"
//
// When `scheduleStrategy` is "manual", this field can be excluded altogether. With
// this strategy, syncs will never be triggered automatically, so they must be
// requested manually.
//
// ### Strategy: "periodic"
//
// When `scheduleStrategy` is "periodic", the only supported key is "interval",
// which will be a string that is in a format accepted by Go's
// [time.ParseDuration][go-time-parse-duration] function. For example, "3h" will
// sync every 3 hours, while "30m" will sync every 30 minutes. The interval must be
// greater than 10 minutes ("10m") and shorter than 1 week ("168h").
//
// An example config would look like this:
//
// ```json
//
//	{
//	  "interval": "1h"
//	}
//
// ```
//
// ### Strategy: "specific_days"
//
// When `scheduleStrategy` is "specific_days", there are only 3 supported keys:
// "days", "hours" and "timezone".
//
// The "days" field must be an array of numbers, corresponding to the desired days
// of the week that syncs should run. Each item must be a number from 0 (Sunday) up
// to 6 (Saturday). There needs to be at least 1 item, but also no more than 7.
//
// For example: `[0, 6]` would correspond to the weekend (Saturday and Sunday)
// while `[1, 2, 3, 4, 5]` would correspond to the weekdays (Monday through
// Friday).
//
// The "hours" field must be an array of numbers, corresponding to the desired
// hours of the day that syncs should run. Each item must be a number from 0 (12am
// or 0:00) up to 23 (11pm or 23:00). There needs to be at least 1 item, but also
// no more than 24.
//
// For example: `[0, 12]` would correspond to the running at midnight (12:00am /
// 0:00) and noon (12:00pm / 12:00).
//
// The "timezone" field must be a valid value from the IANA database used by
// Golang, which should match the list found [here][iana-timezones]. This allows
// the "hours" above to reflect your own time zone which may be easier to
// configure, but you can also choose "UTC" if you want to avoid things like DST
// (daylight savings).
//
// An example config would look like:
//
// ```json
//
//	{
//	  "days": [0, 6],
//	  "hours": [0, 6, 12, 18],
//	  "timezone": "America/Los_Angeles"
//	}
//
// ```
//
// This would run syncs 4 times per day (midnight, 6am, noon, 6pm) but only on
// weekends (Saturday and Sunday).
//
// [go-time-parse-duration]: https://pkg.go.dev/time#ParseDuration
// [iana-timezones]: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
type ReverseETL struct {
	sdkConfiguration sdkConfiguration
}

func newReverseETL(sdkConfig sdkConfiguration) *ReverseETL {
	return &ReverseETL{
		sdkConfiguration: sdkConfig,
	}
}

// CreateReverseEtlModel - Create Reverse Etl Model
// Creates a new Reverse ETL Model.
//
// • When called, this endpoint may generate the `Model Created` event in the [audit trail](/tag/Audit-Trail).
func (s *ReverseETL) CreateReverseEtlModel(ctx context.Context, request shared.CreateReverseEtlModelInput, opts ...operations.Option) (*operations.CreateReverseEtlModelResponse, error) {
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
	url := strings.TrimSuffix(baseURL, "/") + "/reverse-etl-models"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "Request", "json", `request:"mediaType=application/vnd.segment.v1alpha+json"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "POST", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1alpha+json;q=0")
	}

	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

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
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateReverseEtlModelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.CreateReverseEtlModelResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.RequestErrorEnvelope = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// DeleteReverseEtlModel - Delete Reverse Etl Model
// Deletes an existing Model.
//
// • When called, this endpoint may generate the `Model Deleted` event in the [audit trail](/tag/Audit-Trail).
func (s *ReverseETL) DeleteReverseEtlModel(ctx context.Context, request operations.DeleteReverseEtlModelRequest, opts ...operations.Option) (*operations.DeleteReverseEtlModelResponse, error) {
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
	url, err := utils.GenerateURL(ctx, baseURL, "/reverse-etl-models/{modelId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1alpha+json;q=0")
	}

	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.DeleteReverseEtlModelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.DeleteReverseEtlModelResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.RequestErrorEnvelope = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// GetReverseEtlModel - Get Reverse Etl Model
// Returns a Reverse ETL Model by its id.
func (s *ReverseETL) GetReverseEtlModel(ctx context.Context, request operations.GetReverseEtlModelRequest, opts ...operations.Option) (*operations.GetReverseEtlModelResponse, error) {
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
	url, err := utils.GenerateURL(ctx, baseURL, "/reverse-etl-models/{modelId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1alpha+json;q=0")
	}

	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.GetReverseEtlModelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.GetReverseEtlModelResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.RequestErrorEnvelope = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// ListReverseEtlModels - List Reverse Etl Models
// Returns a list of Reverse ETL Models.
func (s *ReverseETL) ListReverseEtlModels(ctx context.Context, request operations.ListReverseEtlModelsRequest, opts ...operations.Option) (*operations.ListReverseEtlModelsResponse, error) {
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
	url := strings.TrimSuffix(baseURL, "/") + "/reverse-etl-models"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1alpha+json;q=0")
	}

	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

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

	res := &operations.ListReverseEtlModelsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.ListReverseEtlModelsResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.RequestErrorEnvelope = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

// UpdateReverseEtlModel - Update Reverse Etl Model
// Updates an existing Reverse ETL Model.
//
// • When called, this endpoint may generate one or more of the following [audit trail](/tag/Audit-Trail) events:* Model Settings Saved
// * Model State Change Toggled
func (s *ReverseETL) UpdateReverseEtlModel(ctx context.Context, request operations.UpdateReverseEtlModelRequest, opts ...operations.Option) (*operations.UpdateReverseEtlModelResponse, error) {
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
	url, err := utils.GenerateURL(ctx, baseURL, "/reverse-etl-models/{modelId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "UpdateReverseEtlModelInput", "json", `request:"mediaType=application/vnd.segment.v1alpha+json"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	debugBody := bytes.NewBuffer([]byte{})
	debugReader := io.TeeReader(bodyReader, debugBody)

	req, err := http.NewRequestWithContext(ctx, "PATCH", url, debugReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	if o.AcceptHeaderOverride != nil {
		req.Header.Set("Accept", string(*o.AcceptHeaderOverride))
	} else {
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1alpha+json;q=0")
	}

	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

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
	httpRes.Request.Body = io.NopCloser(debugBody)
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateReverseEtlModelResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.UpdateReverseEtlModelResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.Object = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.RequestErrorEnvelope
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.RequestErrorEnvelope = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}

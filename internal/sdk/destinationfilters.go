// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/operations"
	"segment_public_api/internal/sdk/pkg/models/sdkerrors"
	"segment_public_api/internal/sdk/pkg/models/shared"
	"segment_public_api/internal/sdk/pkg/utils"
	"strings"
)

// DestinationFilters - The Destination Filters API provides fine-grained controls that allow you to conditionally prevent data delivery to specific destinations. You can filter entire events (for example, selectively drop them) or block/allow individual fields in events before you send them. You can conditionally apply filters to each event based on the contents of that event’s payload. For example, you could apply a filter to Track events with a `plan` property equal to `Professional`, or you could apply a filter to events with a user email address that does not match `*@example.com`.
//
// ## Use cases
//
// Use Destination Filters to:
//
// - Reduce the delivery volume of events to a Destination to save on costs
// - Filter out Personally Identifying Information (PII) from the events sent to a Destination that shouldn't receive or store PII
// - Prevent internal user activity from reaching an analytics tool
// - Send the events that you care about to an custom webhook
//
// ## Availability
//
// Destination Filters are available to all Business customers and support cloud-based (server-side) and web device-mode destinations. Mobile device-mode destinations are in beta and only support [Swift](https://segment.com/docs/connections/sources/catalog/libraries/mobile/swift-ios/), [Kotlin](https://segment.com/docs/connections/sources/catalog/libraries/mobile/kotlin-android/), and [React Native 2.0](https://segment.com/docs/connections/sources/catalog/libraries/mobile/react-native/) libraries. S3 and Warehouse Destinations are not yet supported.
//
// ## Types of filters
//
// There are three filters that can be applied to events:
// | Type | Action |
// | ---------------- | ---------------------------------------- |
// | drop | Do not send the event to the destination |
// | sample | Send only a percentage of events through to the destination. You can optionally sample based on a field. For example, you can sample 10% of users by sampling on userId instead of a random 10% of all events, which is the default. |
// | allow_properties | Allow properties in the event by specifying nested JSON objects (for example, context, properties.productDetails, etc.) and a list of fields that should be retained in that nested object, with all others dropped before the event is sent to the destination.|
// | drop_properties | Drop properties in the event by specifying nested JSON objects (for example, context, properties.productDetails, etc.) and a list of fields to drop from those nested objects before the event is sent to the destination. |
//
// The `allow_properties` and `drop_properties` filters may only filter fields inside the following top-level fields of Segment events:
//
// - `properties`
// - `context`
// - `traits`
//
// ## Filter order
//
// Filters are applied in the following order for each event:
//
// 1. `drop`
// 2. `sample`
// 3. `allow_properties`
// 4. `drop_properties`
//
// ## Conditional filters
//
// Filters can optionally be applied to an event by writing an "if" statement using Filter Query Language ("FQL"), Segment’s simple query language for streaming JSON. FQL statements evaluate to true or false based on the contents of each event, allowing you to only apply filters to specific events.
//
// For example, given the following event payload:
//
// ```
//
//	{
//	  "event": "Button Clicked",
//	  "type": "track",
//	  "context": {
//	    "library": {
//	      "name": "analytics.js",
//	      "version": "1.0",
//	    }
//	  }
//	}
//
// ```
//
// The following FQL statements will evaluate to "true", causing the filter to be applied:
//
// - `event = 'Button Clicked'`
// - `event = 'Button Clicked' and type = 'track'`
// - `match(context.library.version, '1.*')`
// - `type = 'identify' or type = 'track'`
//
// And the following FQL statements will evaluate to "false", causing the filter to be skipped:
//
// - `event = 'Screen Tapped'`
// - `context.path.path = '/login'`
// - `match (context.library.version, '2.*')`
//
// For more detailed documentation including a list of all operators and functions, see the [FQL Documentation](https://segment.com/docs/api/public-api/fql).
//
// ## Universal filters
//
// To apply a filter to all events going to a destination, you can supply an "if" statement of "all". All actions in that filter will be applied to events before being delivered to that destination.
//
// ## Limits
//
// Destination Filters currently have the following limits:
//
// - Each Destination may have no more than 10 filters.
// - FQL "if" statements may be a maximum of 3Kb in size.
// - Filters may have a maximum of 4 actions.
// - Allow / Drop properties filters may target a maximum of 8 objects and each object may have a maximum of 64 fields in the allow or drop.
// - The Preview API "input" payload may be a maximum of 128Kb.
//
// If you would like any of these limits lifted, please reach out to Segment at friends@segment.com.
//
// ## API models
//
// ### Filter
//
// | Attribute   | Type     | Description                                                                                                                                       |           |
// | ----------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | --------- |
// | name        | string   | The URL path of this filter.                                                                                                                      | read-only |
// | enabled     | boolean  | Whether or not this filter should be active.                                                                                                      |           |
// | title       | string   | A human-readable title for this filter.                                                                                                           |           |
// | description | string   | A longer human-readable description of this filter.                                                                                               |           |
// | if          | string   | A FQL statement that causes this filter’s action to be applied if it evaluates to true. "all" will cause the filter to be applied to all events   | required  |
// | actions     | []action | The filtering action to take on events that match the "if" statement. Must be one of: "drop", "sample", "allow_properties", or "drop_properties". | required  |
//
// ## Action
//
// ### Drop
//
// The `drop` action will cause the event to be dropped and not sent to the destination if the "if" statement evaluates to true.
//
// | Attribute | Type   | Description                                                |          |
// | --------- | ------ | ---------------------------------------------------------- | -------- |
// | type      | string | The action name. For the drop action, this must be "drop". | required |
//
// ### Sample
//
// The `sample` action will allow only a percentage of events through. It can sample randomly or, if given a path attribute, it can sample a percentage of events based on the contents of a field. This is useful for sampling all events for a percentage of users rather than a percentage of all events for all users.
//
// | Attribute | Type   | Description                                                                                                                                                                            |          |
// | --------- | ------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------- |
// | type      | string | The action name. For the sample action, this must be "sample".                                                                                                                         | required |
// | percent   | float  | A percentage in the range [0.0, 1.0] that determines the percent of events to allow through. 0.0 will allow no events and 1.0 will allow all events. The default is 0.0.               | required |
// | path      | string | If non-empty, events will be sampled based on the value at this path. For example, if path is userId, a percentage of users will have their events allowed through to the destination. |          |
//
// ### Allow properties
//
// The allow_properties action takes a list of nested objects (for example, `context`, `properties.orderDetails`) and a list of fields for each object that should be allowed, with all other fields in those objects dropped.
// | Attribute | Type | Description | |
// | --------- | ------ | ---------------------------------------------------------------------------------------------------------- | -------- |
// | type | string | The action name. For the allow properties action, this must be "allow_properties". | required |
// | fields | map of string → list of string | A map of nested JSON objects that should have their properties filtered. The map key is the path to the nested JSON object (for example, context.ip, properties, etc.) and the value is a list of string that specifies which fields within the object to allow. Nested JSON objects not specified in this map will be untouched by this filter | required |
//
// ### Drop properties
//
// The drop_properties action takes a list of nested objects (for example, `context`, `properties.orderDetails`) and a list of fields for each object that should be dropped, with all other fields in those objects untouched.
// | Attribute | Type | Description | |
// | --------- | ------------------------------ | ---------------------------------------------------------------------------------- | -------- |
// | type | string | The action name. For the drop properties action, this must be "drop_properties". | required |
// | fields | map of string → list of string | A map of nested JSON objects that should have their properties filtered. The map key is the path to the nested JSON object (for example, context.ip, properties, etc.) and the value is a list of string that specifies which fields within the object to drop. Nested JSON objects not specified in this map will be untouched by this filter | required |
//
// ## Migrate from the Config API
//
// ## Destination filter model
//
// | Config API field | Public API counterpart     |
// | ---------------- | -------------------------- |
// | name             | Use the `id` field instead |
// | enabled          | `enabled`                  |
// | title            | `title`                    |
// | description      | `description`              |
// | if               | `if`                       |
// | actions          | `actions`                  |
//
// ## Action model differences
//
// ### Type property
//
// | Config API       | Public API counterpart |
// | ---------------- | ---------------------- |
// | drop_event       | `drop`                 |
// | sample_event     | `sample`               |
// | whitelist_fields | `allow_properties`     |
// | blacklist_fields | `drop_properties`      |
//
// ### Fields property
//
// The type of `fields` property has been changed:
//
// | Config API                             | Public API counterpart         |
// | -------------------------------------- | ------------------------------ |
// | map of string → Field List (see below) | map of string → list of string |
//
// #### Field list
//
// | Attribute | Type     | Description                                                                                                |
// | --------- | -------- | ---------------------------------------------------------------------------------------------------------- |
// | fields    | []string | One or more JSON object field names. Nested fields (that is, dot-separated field names) are not supported. |
//
// To migrate, replace any usages of the Config API endpoints with the Segment Public API counterparts, using the field mappings in the table above.
type DestinationFilters struct {
	sdkConfiguration sdkConfiguration
}

func newDestinationFilters(sdkConfig sdkConfiguration) *DestinationFilters {
	return &DestinationFilters{
		sdkConfiguration: sdkConfig,
	}
}

// CreateFilterForDestination - Create Filter for Destination
// Creates a filter in a Destination.
//
// • When called, this endpoint may generate the `Destination Filter Created` event in the [audit trail](/tag/Audit-Trail).
func (s *DestinationFilters) CreateFilterForDestination(ctx context.Context, request operations.CreateFilterForDestinationRequest, opts ...operations.Option) (*operations.CreateFilterForDestinationResponse, error) {
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
	url, err := utils.GenerateURL(ctx, baseURL, "/destination/{destinationId}/filters", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "CreateFilterForDestinationV1Input", "json", `request:"mediaType=application/vnd.segment.v1beta+json"`)
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
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
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

	res := &operations.CreateFilterForDestinationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.CreateFilterForDestinationResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out operations.CreateFilterForDestinationDestinationFiltersResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1PlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.CreateFilterForDestinationDestinationFiltersResponseResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out operations.CreateFilterForDestinationDestinationFiltersResponse200ResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1betaPlusJSONObject = &out
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

// GetFilterInDestination - Get Filter in Destination
// Gets a Destination filter by id.
func (s *DestinationFilters) GetFilterInDestination(ctx context.Context, request operations.GetFilterInDestinationRequest, opts ...operations.Option) (*operations.GetFilterInDestinationResponse, error) {
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
	url, err := utils.GenerateURL(ctx, baseURL, "/destination/{destinationId}/filters/{filterId}", request, nil)
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
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
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

	res := &operations.GetFilterInDestinationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.GetFilterInDestinationResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out operations.GetFilterInDestinationDestinationFiltersResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1PlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.GetFilterInDestinationDestinationFiltersResponseResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out operations.GetFilterInDestinationDestinationFiltersResponse200ResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1betaPlusJSONObject = &out
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

// ListFiltersFromDestination - List Filters from Destination
// Lists filters for a Destination.
func (s *DestinationFilters) ListFiltersFromDestination(ctx context.Context, request operations.ListFiltersFromDestinationRequest, opts ...operations.Option) (*operations.ListFiltersFromDestinationResponse, error) {
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
	url, err := utils.GenerateURL(ctx, baseURL, "/destination/{destinationId}/filters", request, nil)
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
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
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

	res := &operations.ListFiltersFromDestinationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.ListFiltersFromDestinationResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out operations.ListFiltersFromDestinationDestinationFiltersResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1PlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.ListFiltersFromDestinationDestinationFiltersResponseResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out operations.ListFiltersFromDestinationDestinationFiltersResponse200ResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1betaPlusJSONObject = &out
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

// PreviewDestinationFilter - Preview Destination Filter
// Simulates the application of a Destination filter to a provided JSON payload.
func (s *DestinationFilters) PreviewDestinationFilter(ctx context.Context, request shared.PreviewDestinationFilterV1Input, opts ...operations.Option) (*operations.PreviewDestinationFilterResponse, error) {
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
	url := strings.TrimSuffix(baseURL, "/") + "/destination/filters/preview"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "Request", "json", `request:"mediaType=application/vnd.segment.v1beta+json"`)
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
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
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

	res := &operations.PreviewDestinationFilterResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.PreviewDestinationFilterResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out operations.PreviewDestinationFilterDestinationFiltersResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1PlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.PreviewDestinationFilterDestinationFiltersResponseResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out operations.PreviewDestinationFilterDestinationFiltersResponse200ResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1betaPlusJSONObject = &out
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

// RemoveFilterFromDestination - Remove Filter from Destination
// Deletes a Destination filter.
//
// • When called, this endpoint may generate the `Destination Filter Deleted` event in the [audit trail](/tag/Audit-Trail).
func (s *DestinationFilters) RemoveFilterFromDestination(ctx context.Context, request operations.RemoveFilterFromDestinationRequest, opts ...operations.Option) (*operations.RemoveFilterFromDestinationResponse, error) {
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
	url, err := utils.GenerateURL(ctx, baseURL, "/destination/{destinationId}/filters/{filterId}", request, nil)
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
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
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

	res := &operations.RemoveFilterFromDestinationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.RemoveFilterFromDestinationResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out operations.RemoveFilterFromDestinationDestinationFiltersResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1PlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.RemoveFilterFromDestinationDestinationFiltersResponseResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out operations.RemoveFilterFromDestinationDestinationFiltersResponse200ResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1betaPlusJSONObject = &out
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

// UpdateFilterForDestination - Update Filter for Destination
// Updates a filter in a Destination.
//
// • When called, this endpoint may generate one or more of the following [audit trail](/tag/Audit-Trail) events:* Destination Filter Enabled
// * Destination Filter Disabled
func (s *DestinationFilters) UpdateFilterForDestination(ctx context.Context, request operations.UpdateFilterForDestinationRequest, opts ...operations.Option) (*operations.UpdateFilterForDestinationResponse, error) {
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
	url, err := utils.GenerateURL(ctx, baseURL, "/destination/{destinationId}/filters/{filterId}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "UpdateFilterForDestinationV1Input", "json", `request:"mediaType=application/vnd.segment.v1beta+json"`)
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
		req.Header.Set("Accept", "application/json;q=1, application/vnd.segment.v1+json;q=0.8, application/vnd.segment.v1alpha+json;q=0.5, application/vnd.segment.v1beta+json;q=0")
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

	res := &operations.UpdateFilterForDestinationResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out operations.UpdateFilterForDestinationResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1+json`):
			var out operations.UpdateFilterForDestinationDestinationFiltersResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1PlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1alpha+json`):
			var out operations.UpdateFilterForDestinationDestinationFiltersResponseResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1alphaPlusJSONObject = &out
		case utils.MatchContentType(contentType, `application/vnd.segment.v1beta+json`):
			var out operations.UpdateFilterForDestinationDestinationFiltersResponse200ResponseBody
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.TwoHundredApplicationVndSegmentV1betaPlusJSONObject = &out
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

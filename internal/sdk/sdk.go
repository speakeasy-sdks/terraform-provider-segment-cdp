// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"net/http"
	"segment_public_api/internal/sdk/pkg/models/shared"
	"segment_public_api/internal/sdk/pkg/utils"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// Production
	"https://api.segmentapis.com",
	// Staging
	"https://api.segmentapis.build",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// SegmentPublicAPI - Segment Public API: The Segment Public API helps you manage your Segment Workspaces and its resources. You can use the API to perform CRUD (create, read, update, delete) operations at no extra charge. This includes working with resources such as Sources, Destinations, Warehouses, Tracking Plans, and the Segment Destinations and Sources Catalogs.
//
// All CRUD endpoints in the API follow REST conventions and use standard HTTP methods. Different URL endpoints represent different resources in a Workspace.
//
// See the next sections for more information on how to use the Segment Public API.
type SegmentPublicAPI struct {
	// A Workspace is a group of Sources that you administer and Segment bills together. Workspaces help companies manage access for different users and data Sources and let you collaborate with team members, add permissions, and share Sources across your team using a shared billing account.
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
	//
	Workspaces *Workspaces
	// These endpoints let you list all the Sources, Destinations, and Destination settings that Segment supports.
	//
	// ## Migrate from the Config API
	//
	// Like the Segment Public API, the Config API has endpoints to navigate the catalog of Integrations supported by Segment. See the sections below for some key differences:
	//
	// ### Source catalog items
	//
	// | Config API  | Public API                                      |
	// | ----------- | ----------------------------------------------- |
	// | name        | See note on names vs IDs in the migration guide |
	// | displayName | `name`                                          |
	// | description | `description`                                   |
	// | categories  | `categories`                                    |
	// | logos       | `logos`                                         |
	// | type        | Removed                                         |
	//
	// ### Destination catalog items
	//
	// | Config API                 | Public API                                      |
	// | -------------------------- | ----------------------------------------------- |
	// | browserUnbundlingPublic    | `supportedFeatures.browserUnbundlingPublic`     |
	// | browserUnbundlingSupported | `supportedFeatures.browserUnbundling`           |
	// | categories                 | `categories`                                    |
	// | components                 | `components`                                    |
	// | description                | `description`                                   |
	// | displayName                | `name`,                                         |
	// | id                         | `id`                                            |
	// | logos                      | `logos`                                         |
	// | methods                    | `supportedMethods`                              |
	// | name                       | See note on names vs IDs in the migration guide |
	// | platforms                  | `supportedPlatforms`                            |
	// | previousNames              | `previousNames`                                 |
	// | settings                   | `options`                                       |
	// | slug                       | `slug`                                          |
	// | status                     | `status`                                        |
	// | type                       | Not returned                                    |
	// | website                    | `website`                                       |
	//
	// To migrate, replace any use of the Config API endpoints with the Segment Public API counterparts, using the field mappings in the table above.
	//
	Catalog *Catalog
	// The Destination Filters API provides fine-grained controls that allow you to conditionally prevent data delivery to specific destinations. You can filter entire events (for example, selectively drop them) or block/allow individual fields in events before you send them. You can conditionally apply filters to each event based on the contents of that event’s payload. For example, you could apply a filter to Track events with a `plan` property equal to `Professional`, or you could apply a filter to events with a user email address that does not match `*@example.com`.
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
	// {
	//   "event": "Button Clicked",
	//   "type": "track",
	//   "context": {
	//     "library": {
	//       "name": "analytics.js",
	//       "version": "1.0",
	//     }
	//   }
	// }
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
	//
	DestinationFilters *DestinationFilters
	// Destinations receive data _from_ Segment.
	//
	// In the Segment Public API, you can manipulate Destinations and the connections between Sources and Destinations, as well as list and inspect their relationships.
	//
	// ## Migrate from the Config API
	//
	// Like Segment Public API, Config API allows creating, retrieving, updating and deleting Destination objects. See the table below for some key differences:
	//
	// | Config API     | Public API                                      |
	// | -------------- | ----------------------------------------------- |
	// | catalogId      | Not returned                                    |
	// | config         | `settings`                                      |
	// | connectionMode | Not returned                                    |
	// | createTime     | Not returned                                    |
	// | displayName    | `name`                                          |
	// | enabled        | `enabled`                                       |
	// | name           | See note on names vs IDs in the migration guide |
	// | parent         | `sourceId` (prefix removed)                     |
	// | updateTime     | Not returned                                    |
	//
	// To migrate, replace any use of the Config API endpoints with the Segment Public API counterparts, using the field mappings in the table above.
	//
	Destinations *Destinations
	// Functions let you create your own Sources and Destinations directly within your Workspace to bring new types of data into Segment and send data to new tools with JavaScript - no extra infrastructure required.
	//
	// ## Migrate from the Config API
	//
	// The [getFunction endpoint](https://reference.segmentapis.com/#7963d88e-2e1b-41eb-9fa2-66a113c12f1c) returns the following fields:
	//
	// | Config API    | Public API     |
	// | ------------- | -------------- |
	// | `id`          | `id`           |
	// | `type`        | `resourceType` |
	// | `created_at`  | `createdAt`    |
	// | `created_by`  | `createdBy`    |
	// | `description` | `description`  |
	// | `logo_url`    | `logoUrl`      |
	// | `catalog_id`  | `catalogId`    |
	//
	Functions *Functions
	// A User Group is a set of Team Members with a set of shared policies. A Segment Team Member can be a member of one or many Groups. All roles in the Segment App are additive, which means that you can assign group memberships and individual roles to a single team member. For example, a single user can inherit roles from a Group definition AND have access to more resources through individually assigned roles.
	//
	IAMGroups *IAMGroups
	// Segment’s access management tools let Workspace owners manage which users can access different parts of their Segment Workspaces.
	//
	IAMUsers *IAMUsers
	// Workspace owners can use Labels to grant users access to groups of resources. When you add a Label to a Source or Personas Space, any users to whom you assign that Label gain access to those resources.
	//
	// To create or configure labels, go to the Labels tab in your Workspace settings. Only Workspace Owners can manage labels for the entire Workspace.
	//
	Labels *Labels
	// In keeping with Segment's commitment to support GDPR and future privacy regulations such as the CCPA, you can delete and suppress data about end users if you identify that user with a `userId`, should they revoke or alter their consent to data collection. For instance, if an end user in the EU invokes their Right to Object or Right to Erasure under the GDPR, you can use the following features in Segment to block ongoing data collection about the user, and delete all historical data across Segment’s systems, connected S3 buckets and Warehouses, and supported downstream partners.
	//
	// Regulations enable you to issue a single request to delete and suppress data about a user by `userId`. All regulations are by default scoped to your Workspace and target all Sources within the Workspace. This way, you don't need to page over every Source within Segment to delete data about a user across all your users.
	//
	// All deletion and suppression actions within Segment are asynchronous, and fall under the umbrella of what Segment calls "Regulations." Regulations are requests to Segment to impart control over your data flow. You can issue these requests with the Segment Public API using the endpoints below.
	//
	// You can't replay data deleted through Regulations. For standard replay requests, Segment asks that you wait for deletions to complete and not submit any new deletion requests for the period of time that Segment replays data for you.
	//
	// ## Migrate from the Config API
	//
	// Deletion and Suppression got an overhaul in the Segment Public API. They’re now divided into Workspace, Source, and Cloud Source-related endpoints. The Public API simplifies these endpoints: the `attributes` input field is no longer required, and you can now pass an array of IDs to regulate (instead of a `parent`).
	//
	DeletionAndSuppression *DeletionAndSuppression
	// Reverse ETL allows the use of a database (aka: Segment Warehouse) as a source of
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
	// {
	//   "interval": "1h"
	// }
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
	// {
	//   "days": [0, 6],
	//   "hours": [0, 6, 12, 18],
	//   "timezone": "America/Los_Angeles"
	// }
	// ```
	//
	// This would run syncs 4 times per day (midnight, 6am, noon, 6pm) but only on
	// weekends (Saturday and Sunday).
	//
	// [go-time-parse-duration]: https://pkg.go.dev/time#ParseDuration
	// [iana-timezones]: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	//
	ReverseETL *ReverseETL
	// A role gives a user access to resources within a Workspace. Roles are additive, and can combine to configure a custom policy for a Team Member or a Group. A policy is at least one role plus one resource applied to an individual user or group.
	//
	IAMRoles *IAMRoles
	// In Segment, you create a Source (or more than one!) for each website or app you want to track. While it’s not required that you have a single Source for each server, site or app, Segment recommends that you create a Source for each unique source of data.
	//
	// ## Migrate from the Config API
	//
	// Like the Segment Public API, the Config API has one endpoint to retrieve details about a Workspace. The [getSource endpoint](https://reference.segmentapis.com/#5a852761-54d5-46da-8437-6e14e63449f3) returns the following fields:
	//
	// | Config API     | Public API                                      |
	// | -------------- | ----------------------------------------------- |
	// | `name`         | See note on names vs IDs in the migration guide |
	// | `name`         | `slug` (`workspace/` prefix removed)            |
	// | `display_name` | `name`                                          |
	// | `create_time`  | None                                            |
	//
	// To migrate, replace any use of the Config API endpoints with the Segment Public API counterparts, using the field mappings in the table above.
	//
	Sources *Sources
	// Edge functions enables you to write event enrichment and transformation logic in Javascript outside
	// your client codebase and deployed to applications over-the-air dynamically.
	//
	// ## Availability
	//
	// Edge Functions are in Private Alpha testing and available to select customers. To opt in, contact your Customer Success Manager and ensure that you have one of Segment's new Mobile SDKs (Swift or Kotlin) configured.
	//
	// ## Migrate from the Config API
	//
	// | Config API | Public API                                                            |
	// | ---------- | --------------------------------------------------------------------- |
	// | `name`     | Use the Source `id` (See note on names vs IDs in the migration guide) |
	//
	// To migrate, replace any use of the Config API endpoints with the Segment Public API counterparts, using the field mappings in the table above.
	//
	EdgeFunctions *EdgeFunctions
	// A space is a separate Personas environment. Consider the two main reasons you might use spaces:
	//
	// - To separate your development and production environments (highly recommended)
	// - To separate environments for distinct teams or geographical regions
	//
	Spaces *Spaces
	// Audiences play a key role in gaining a deeper understanding of your users. Audiences allow you to group users or profiles based on shared characteristics, behaviors, and attributes. Using events passed into Segment, traits, and computed traits you can create Audiences which can help unlock more relevant engagement and communication.
	//
	// > **Note**: The Audience API is currently in a Private Beta. If you are interested in joining the Private Beta, then please reach out to your customer success manager.
	//
	Audiences *Audiences
	// Computed traits allow you to quickly create new traits for a user or profile based on that user's tracked interactions. Using the events and event properties that you send through page and track calls, Segment will calculate and keep up-to-date, over time, the value for your defined computed trait. These can be computations like the total number of orders a customer has completed, the lifetime revenue of a customer, the most frequent user to determine which user is most active in an account, or the unique visitors count to assess how many visitors from a single domain.
	//
	// > **Note**: The Computed Traits API is currently in a Private Beta. If you are interested in joining the Private Beta, then please reach out to your customer success manager.
	//
	// Note that when using a unique list computed trait, Segment limits the number of Event Properties that can be added to the specific trait to 10,000. If your computed trait exceeds this limit, Segment will not persist any new Event Properties and will drop new trait keys and corresponding values.
	//
	ComputedTraits *ComputedTraits
	// A Profiles Sync Warehouse is a central repository of data collected from your workspace. It is what commonly comes to mind when you think about a relational database: structured data that fits into rows and columns.
	//
	// Using Segment’s Public API, you can create, delete, update, and list Spaces Warehouses connections.
	//
	ProfilesSync *ProfilesSync
	// A Tracking Plan is a data spec outlining the events and properties you intend to collect across your Segment Sources.
	//
	// The Segment Tracking Plan feature allows you to validate your expected events against the live events that Segment receives. Segment generates violations when an event doesn’t match the spec’d event in the Tracking Plan.
	//
	// You can store your Tracking Plans in your Workspace, and connect them to one or more Sources.
	//
	// Using the Segment Public API, you can create, delete, update, list, and connect Tracking Plans and their Rules.
	//
	// ## Migrate from the Config API
	//
	// The [getTrackingPlan endpoint](https://reference.segmentapis.com/#1092fe01-379b-4ca1-8b1d-9f42b33d2899) returns the following fields:
	//
	// | Config API     | Public API  |
	// | -------------- | ----------- |
	// | `display_name` | `name`      |
	// | `name`         | `slug`      |
	// | `updateTime`   | `updatedAt` |
	// | `createTime`   | `createdAt` |
	//
	// To migrate, replace any use of the Config API endpoints with the Segment Public API counterparts, using the field mappings in the table above.
	//
	TrackingPlans *TrackingPlans
	// With Transformations, you can change data as it flows through Segment to:
	// - Correct bad data
	// - Customize data for a specific destination
	// - Align events with your tracking plan
	//
	// With Segment's Public API, you can use Transformations for several use cases:
	// 1. Rename Track events
	// - Use `newEventName` in the `createTransformation` or `updateTransformation` API call.
	// 2. Rename a property (Track events) or a trait (Identify or Group events)
	// - Use `propertyRenames` in the `createTransformation` or `updateTransformation` API call.
	// 3. Update a property value to a different static value or add a new property and set a static value
	// - Use `propertyValueTransformations` in the `createTransformation` or `updateTransformation` API call.
	// - Segment currently supports setting static values for top-level fields with `propertyValueTransformations`. However, Segment doesn't support changing fields outside the properties or traits object with `propertyRenames`.
	//
	// Visit [Segment's Transformations docs](https://segment.com/docs/protocols/transform/) to learn more.
	//
	Transformations *Transformations
	// A Warehouse is a central repository of data collected from one or more Sources. This is what commonly comes to mind when you think about a relational database: structured data that fits into rows and columns.
	//
	// Using the Segment Public API, you can create, delete, update, list, validate and connect Warehouses.
	//
	Warehouses *Warehouses
	// Warehouse Selective Sync allows you to manage the data that you send to your Warehouses. You can use this feature to stop syncing specific events (also known as collections) or properties that aren’t relevant, and may slow down your Warehouse syncs.
	//
	SelectiveSync *SelectiveSync

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*SegmentPublicAPI)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *SegmentPublicAPI) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *SegmentPublicAPI) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *SegmentPublicAPI) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *SegmentPublicAPI) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details
func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *SegmentPublicAPI) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *SegmentPublicAPI) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *SegmentPublicAPI) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *SegmentPublicAPI {
	sdk := &SegmentPublicAPI{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "36.2.0",
			SDKVersion:        "0.7.0",
			GenVersion:        "2.181.1",
			UserAgent:         "speakeasy-sdk/go 0.7.0 2.181.1 36.2.0 segment_public_api",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Workspaces = newWorkspaces(sdk.sdkConfiguration)

	sdk.Catalog = newCatalog(sdk.sdkConfiguration)

	sdk.DestinationFilters = newDestinationFilters(sdk.sdkConfiguration)

	sdk.Destinations = newDestinations(sdk.sdkConfiguration)

	sdk.Functions = newFunctions(sdk.sdkConfiguration)

	sdk.IAMGroups = newIAMGroups(sdk.sdkConfiguration)

	sdk.IAMUsers = newIAMUsers(sdk.sdkConfiguration)

	sdk.Labels = newLabels(sdk.sdkConfiguration)

	sdk.DeletionAndSuppression = newDeletionAndSuppression(sdk.sdkConfiguration)

	sdk.ReverseETL = newReverseETL(sdk.sdkConfiguration)

	sdk.IAMRoles = newIAMRoles(sdk.sdkConfiguration)

	sdk.Sources = newSources(sdk.sdkConfiguration)

	sdk.EdgeFunctions = newEdgeFunctions(sdk.sdkConfiguration)

	sdk.Spaces = newSpaces(sdk.sdkConfiguration)

	sdk.Audiences = newAudiences(sdk.sdkConfiguration)

	sdk.ComputedTraits = newComputedTraits(sdk.sdkConfiguration)

	sdk.ProfilesSync = newProfilesSync(sdk.sdkConfiguration)

	sdk.TrackingPlans = newTrackingPlans(sdk.sdkConfiguration)

	sdk.Transformations = newTransformations(sdk.sdkConfiguration)

	sdk.Warehouses = newWarehouses(sdk.sdkConfiguration)

	sdk.SelectiveSync = newSelectiveSync(sdk.sdkConfiguration)

	return sdk
}

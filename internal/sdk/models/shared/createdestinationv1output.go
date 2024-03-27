// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreateDestinationV1OutputLogos - The Destination's logos.
type CreateDestinationV1OutputLogos struct {
	// The alternative text for this logo.
	Alt *string `json:"alt,omitempty"`
	// The default URL for this logo.
	Default string `json:"default"`
	// The logo mark.
	Mark *string `json:"mark,omitempty"`
}

func (o *CreateDestinationV1OutputLogos) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

func (o *CreateDestinationV1OutputLogos) GetDefault() string {
	if o == nil {
		return ""
	}
	return o.Default
}

func (o *CreateDestinationV1OutputLogos) GetMark() *string {
	if o == nil {
		return nil
	}
	return o.Mark
}

// CreateDestinationV1OutputStatus - Support status of the Destination.
type CreateDestinationV1OutputStatus string

const (
	CreateDestinationV1OutputStatusDeprecated       CreateDestinationV1OutputStatus = "DEPRECATED"
	CreateDestinationV1OutputStatusPrivateBeta      CreateDestinationV1OutputStatus = "PRIVATE_BETA"
	CreateDestinationV1OutputStatusPrivateBuilding  CreateDestinationV1OutputStatus = "PRIVATE_BUILDING"
	CreateDestinationV1OutputStatusPrivateSubmitted CreateDestinationV1OutputStatus = "PRIVATE_SUBMITTED"
	CreateDestinationV1OutputStatusPublic           CreateDestinationV1OutputStatus = "PUBLIC"
	CreateDestinationV1OutputStatusPublicBeta       CreateDestinationV1OutputStatus = "PUBLIC_BETA"
	CreateDestinationV1OutputStatusUnavailable      CreateDestinationV1OutputStatus = "UNAVAILABLE"
)

func (e CreateDestinationV1OutputStatus) ToPointer() *CreateDestinationV1OutputStatus {
	return &e
}

func (e *CreateDestinationV1OutputStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DEPRECATED":
		fallthrough
	case "PRIVATE_BETA":
		fallthrough
	case "PRIVATE_BUILDING":
		fallthrough
	case "PRIVATE_SUBMITTED":
		fallthrough
	case "PUBLIC":
		fallthrough
	case "PUBLIC_BETA":
		fallthrough
	case "UNAVAILABLE":
		*e = CreateDestinationV1OutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDestinationV1OutputStatus: %v", v)
	}
}

// CreateDestinationV1OutputCloudModeInstances - This Destination's support level for cloud mode instances.
// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
type CreateDestinationV1OutputCloudModeInstances string

const (
	CreateDestinationV1OutputCloudModeInstancesZero     CreateDestinationV1OutputCloudModeInstances = "0"
	CreateDestinationV1OutputCloudModeInstancesOne      CreateDestinationV1OutputCloudModeInstances = "1"
	CreateDestinationV1OutputCloudModeInstancesMultiple CreateDestinationV1OutputCloudModeInstances = "MULTIPLE"
	CreateDestinationV1OutputCloudModeInstancesNone     CreateDestinationV1OutputCloudModeInstances = "NONE"
	CreateDestinationV1OutputCloudModeInstancesSingle   CreateDestinationV1OutputCloudModeInstances = "SINGLE"
)

func (e CreateDestinationV1OutputCloudModeInstances) ToPointer() *CreateDestinationV1OutputCloudModeInstances {
	return &e
}

func (e *CreateDestinationV1OutputCloudModeInstances) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "0":
		fallthrough
	case "1":
		fallthrough
	case "MULTIPLE":
		fallthrough
	case "NONE":
		fallthrough
	case "SINGLE":
		*e = CreateDestinationV1OutputCloudModeInstances(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDestinationV1OutputCloudModeInstances: %v", v)
	}
}

// CreateDestinationV1OutputDeviceModeInstances - This Destination's support level for device mode instances.
// Support for multiple device mode instances is currently not planned.
// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
type CreateDestinationV1OutputDeviceModeInstances string

const (
	CreateDestinationV1OutputDeviceModeInstancesZero   CreateDestinationV1OutputDeviceModeInstances = "0"
	CreateDestinationV1OutputDeviceModeInstancesOne    CreateDestinationV1OutputDeviceModeInstances = "1"
	CreateDestinationV1OutputDeviceModeInstancesNone   CreateDestinationV1OutputDeviceModeInstances = "NONE"
	CreateDestinationV1OutputDeviceModeInstancesSingle CreateDestinationV1OutputDeviceModeInstances = "SINGLE"
)

func (e CreateDestinationV1OutputDeviceModeInstances) ToPointer() *CreateDestinationV1OutputDeviceModeInstances {
	return &e
}

func (e *CreateDestinationV1OutputDeviceModeInstances) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "0":
		fallthrough
	case "1":
		fallthrough
	case "NONE":
		fallthrough
	case "SINGLE":
		*e = CreateDestinationV1OutputDeviceModeInstances(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDestinationV1OutputDeviceModeInstances: %v", v)
	}
}

// CreateDestinationV1OutputSupportedFeatures - Features that this Destination supports.
//
// Config API note: holds `browserUnbundling` fields.
type CreateDestinationV1OutputSupportedFeatures struct {
	// Whether this Destination supports browser unbundling.
	BrowserUnbundling *bool `json:"browserUnbundling,omitempty"`
	// Whether this Destination supports public browser unbundling.
	BrowserUnbundlingPublic *bool `json:"browserUnbundlingPublic,omitempty"`
	// This Destination's support level for cloud mode instances.
	// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	CloudModeInstances *CreateDestinationV1OutputCloudModeInstances `json:"cloudModeInstances,omitempty"`
	// This Destination's support level for device mode instances.
	// Support for multiple device mode instances is currently not planned.
	// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	DeviceModeInstances *CreateDestinationV1OutputDeviceModeInstances `json:"deviceModeInstances,omitempty"`
	// Whether this Destination supports replays.
	Replay *bool `json:"replay,omitempty"`
}

func (o *CreateDestinationV1OutputSupportedFeatures) GetBrowserUnbundling() *bool {
	if o == nil {
		return nil
	}
	return o.BrowserUnbundling
}

func (o *CreateDestinationV1OutputSupportedFeatures) GetBrowserUnbundlingPublic() *bool {
	if o == nil {
		return nil
	}
	return o.BrowserUnbundlingPublic
}

func (o *CreateDestinationV1OutputSupportedFeatures) GetCloudModeInstances() *CreateDestinationV1OutputCloudModeInstances {
	if o == nil {
		return nil
	}
	return o.CloudModeInstances
}

func (o *CreateDestinationV1OutputSupportedFeatures) GetDeviceModeInstances() *CreateDestinationV1OutputDeviceModeInstances {
	if o == nil {
		return nil
	}
	return o.DeviceModeInstances
}

func (o *CreateDestinationV1OutputSupportedFeatures) GetReplay() *bool {
	if o == nil {
		return nil
	}
	return o.Replay
}

// CreateDestinationV1OutputSupportedMethods - Methods that this Destination supports.
//
// Config API note: equal to `methods`.
type CreateDestinationV1OutputSupportedMethods struct {
	// Identifies if the Destination supports the `alias` method.
	Alias *bool `json:"alias,omitempty"`
	// Identifies if the Destination supports the `group` method.
	Group *bool `json:"group,omitempty"`
	// Identifies if the Destination supports the `identify` method.
	Identify *bool `json:"identify,omitempty"`
	// Identifies if the Destination supports the `pageview` method.
	Pageview *bool `json:"pageview,omitempty"`
	// Identifies if the Destination supports the `track` method.
	Track *bool `json:"track,omitempty"`
}

func (o *CreateDestinationV1OutputSupportedMethods) GetAlias() *bool {
	if o == nil {
		return nil
	}
	return o.Alias
}

func (o *CreateDestinationV1OutputSupportedMethods) GetGroup() *bool {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *CreateDestinationV1OutputSupportedMethods) GetIdentify() *bool {
	if o == nil {
		return nil
	}
	return o.Identify
}

func (o *CreateDestinationV1OutputSupportedMethods) GetPageview() *bool {
	if o == nil {
		return nil
	}
	return o.Pageview
}

func (o *CreateDestinationV1OutputSupportedMethods) GetTrack() *bool {
	if o == nil {
		return nil
	}
	return o.Track
}

// CreateDestinationV1OutputSupportedPlatforms - Platforms from which the Destination receives events.
//
// Config API note: equal to `platforms`.
type CreateDestinationV1OutputSupportedPlatforms struct {
	// Whether this Destination supports browser events.
	Browser *bool `json:"browser,omitempty"`
	// Whether this Destination supports mobile events.
	Mobile *bool `json:"mobile,omitempty"`
	// Whether this Destination supports server events.
	Server *bool `json:"server,omitempty"`
}

func (o *CreateDestinationV1OutputSupportedPlatforms) GetBrowser() *bool {
	if o == nil {
		return nil
	}
	return o.Browser
}

func (o *CreateDestinationV1OutputSupportedPlatforms) GetMobile() *bool {
	if o == nil {
		return nil
	}
	return o.Mobile
}

func (o *CreateDestinationV1OutputSupportedPlatforms) GetServer() *bool {
	if o == nil {
		return nil
	}
	return o.Server
}

// CreateDestinationV1OutputMetadata - The metadata of the Destination of which this Destination is an instance of. For example, Google Analytics or Amplitude.
type CreateDestinationV1OutputMetadata struct {
	// Actions available for the Destination.
	Actions []DestinationMetadataActionV1 `json:"actions"`
	// A list of categories with which the Destination is associated.
	Categories []string `json:"categories"`
	// A list of components this Destination provides.
	Components []DestinationMetadataComponentV1 `json:"components"`
	// Contact info for Integration Owners.
	Contacts []Contact `json:"contacts,omitempty"`
	// The description of the Destination.
	Description string `json:"description"`
	// The id of the Destination metadata.
	//
	// Config API note: analogous to `name`.
	ID string `json:"id"`
	// The Destination's logos.
	Logos CreateDestinationV1OutputLogos `json:"logos"`
	// The user-friendly name of the Destination.
	//
	// Config API note: equal to `displayName`.
	Name string `json:"name"`
	// Options configured for the Destination.
	Options []IntegrationOptionBeta `json:"options"`
	// Partner Owned flag.
	PartnerOwned *bool `json:"partnerOwned,omitempty"`
	// Predefined Destination subscriptions that can optionally be applied when connecting a new instance of the Destination.
	Presets []DestinationMetadataSubscriptionPresetV1 `json:"presets"`
	// A list of names previously used by the Destination.
	PreviousNames []string `json:"previousNames"`
	// The list of regional endpoints for this Destination.
	RegionEndpoints []string `json:"regionEndpoints,omitempty"`
	// The slug used to identify the Destination in the Segment app.
	Slug string `json:"slug"`
	// Support status of the Destination.
	Status CreateDestinationV1OutputStatus `json:"status"`
	// Features that this Destination supports.
	//
	// Config API note: holds `browserUnbundling` fields.
	SupportedFeatures CreateDestinationV1OutputSupportedFeatures `json:"supportedFeatures"`
	// Methods that this Destination supports.
	//
	// Config API note: equal to `methods`.
	SupportedMethods CreateDestinationV1OutputSupportedMethods `json:"supportedMethods"`
	// Platforms from which the Destination receives events.
	//
	// Config API note: equal to `platforms`.
	SupportedPlatforms CreateDestinationV1OutputSupportedPlatforms `json:"supportedPlatforms"`
	// A list of supported regions for this Destination.
	SupportedRegions []string `json:"supportedRegions,omitempty"`
	// A website URL for this Destination.
	Website string `json:"website"`
}

func (o *CreateDestinationV1OutputMetadata) GetActions() []DestinationMetadataActionV1 {
	if o == nil {
		return []DestinationMetadataActionV1{}
	}
	return o.Actions
}

func (o *CreateDestinationV1OutputMetadata) GetCategories() []string {
	if o == nil {
		return []string{}
	}
	return o.Categories
}

func (o *CreateDestinationV1OutputMetadata) GetComponents() []DestinationMetadataComponentV1 {
	if o == nil {
		return []DestinationMetadataComponentV1{}
	}
	return o.Components
}

func (o *CreateDestinationV1OutputMetadata) GetContacts() []Contact {
	if o == nil {
		return nil
	}
	return o.Contacts
}

func (o *CreateDestinationV1OutputMetadata) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *CreateDestinationV1OutputMetadata) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateDestinationV1OutputMetadata) GetLogos() CreateDestinationV1OutputLogos {
	if o == nil {
		return CreateDestinationV1OutputLogos{}
	}
	return o.Logos
}

func (o *CreateDestinationV1OutputMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateDestinationV1OutputMetadata) GetOptions() []IntegrationOptionBeta {
	if o == nil {
		return []IntegrationOptionBeta{}
	}
	return o.Options
}

func (o *CreateDestinationV1OutputMetadata) GetPartnerOwned() *bool {
	if o == nil {
		return nil
	}
	return o.PartnerOwned
}

func (o *CreateDestinationV1OutputMetadata) GetPresets() []DestinationMetadataSubscriptionPresetV1 {
	if o == nil {
		return []DestinationMetadataSubscriptionPresetV1{}
	}
	return o.Presets
}

func (o *CreateDestinationV1OutputMetadata) GetPreviousNames() []string {
	if o == nil {
		return []string{}
	}
	return o.PreviousNames
}

func (o *CreateDestinationV1OutputMetadata) GetRegionEndpoints() []string {
	if o == nil {
		return nil
	}
	return o.RegionEndpoints
}

func (o *CreateDestinationV1OutputMetadata) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *CreateDestinationV1OutputMetadata) GetStatus() CreateDestinationV1OutputStatus {
	if o == nil {
		return CreateDestinationV1OutputStatus("")
	}
	return o.Status
}

func (o *CreateDestinationV1OutputMetadata) GetSupportedFeatures() CreateDestinationV1OutputSupportedFeatures {
	if o == nil {
		return CreateDestinationV1OutputSupportedFeatures{}
	}
	return o.SupportedFeatures
}

func (o *CreateDestinationV1OutputMetadata) GetSupportedMethods() CreateDestinationV1OutputSupportedMethods {
	if o == nil {
		return CreateDestinationV1OutputSupportedMethods{}
	}
	return o.SupportedMethods
}

func (o *CreateDestinationV1OutputMetadata) GetSupportedPlatforms() CreateDestinationV1OutputSupportedPlatforms {
	if o == nil {
		return CreateDestinationV1OutputSupportedPlatforms{}
	}
	return o.SupportedPlatforms
}

func (o *CreateDestinationV1OutputMetadata) GetSupportedRegions() []string {
	if o == nil {
		return nil
	}
	return o.SupportedRegions
}

func (o *CreateDestinationV1OutputMetadata) GetWebsite() string {
	if o == nil {
		return ""
	}
	return o.Website
}

// Destination - The created Destination.
type Destination struct {
	// Whether this instance of a Destination receives data.
	Enabled bool `json:"enabled"`
	// The unique identifier of this instance of a Destination.
	//
	// Config API note: analogous to `name`.
	ID string `json:"id"`
	// The metadata of the Destination of which this Destination is an instance of. For example, Google Analytics or Amplitude.
	Metadata CreateDestinationV1OutputMetadata `json:"metadata"`
	// The name of this instance of a Destination.
	//
	// Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	// The collection of settings associated with a Destination.
	//
	// Config API note: equal to `config`.
	Settings map[string]interface{} `json:"settings"`
	// The id of a Source connected to this instance of a Destination.
	//
	// Config API note: analogous to `parent`.
	SourceID string `json:"sourceId"`
}

func (o *Destination) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *Destination) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Destination) GetMetadata() CreateDestinationV1OutputMetadata {
	if o == nil {
		return CreateDestinationV1OutputMetadata{}
	}
	return o.Metadata
}

func (o *Destination) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Destination) GetSettings() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Settings
}

func (o *Destination) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

// CreateDestinationV1Output - Creates a new Destination.
type CreateDestinationV1Output struct {
	// The created Destination.
	Destination Destination `json:"destination"`
}

func (o *CreateDestinationV1Output) GetDestination() Destination {
	if o == nil {
		return Destination{}
	}
	return o.Destination
}
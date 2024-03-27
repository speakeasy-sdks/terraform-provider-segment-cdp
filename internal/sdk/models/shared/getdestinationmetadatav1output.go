// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// GetDestinationMetadataV1OutputLogos - The Destination's logos.
type GetDestinationMetadataV1OutputLogos struct {
	// The alternative text for this logo.
	Alt *string `json:"alt,omitempty"`
	// The default URL for this logo.
	Default string `json:"default"`
	// The logo mark.
	Mark *string `json:"mark,omitempty"`
}

func (o *GetDestinationMetadataV1OutputLogos) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

func (o *GetDestinationMetadataV1OutputLogos) GetDefault() string {
	if o == nil {
		return ""
	}
	return o.Default
}

func (o *GetDestinationMetadataV1OutputLogos) GetMark() *string {
	if o == nil {
		return nil
	}
	return o.Mark
}

// GetDestinationMetadataV1OutputStatus - Support status of the Destination.
type GetDestinationMetadataV1OutputStatus string

const (
	GetDestinationMetadataV1OutputStatusDeprecated       GetDestinationMetadataV1OutputStatus = "DEPRECATED"
	GetDestinationMetadataV1OutputStatusPrivateBeta      GetDestinationMetadataV1OutputStatus = "PRIVATE_BETA"
	GetDestinationMetadataV1OutputStatusPrivateBuilding  GetDestinationMetadataV1OutputStatus = "PRIVATE_BUILDING"
	GetDestinationMetadataV1OutputStatusPrivateSubmitted GetDestinationMetadataV1OutputStatus = "PRIVATE_SUBMITTED"
	GetDestinationMetadataV1OutputStatusPublic           GetDestinationMetadataV1OutputStatus = "PUBLIC"
	GetDestinationMetadataV1OutputStatusPublicBeta       GetDestinationMetadataV1OutputStatus = "PUBLIC_BETA"
	GetDestinationMetadataV1OutputStatusUnavailable      GetDestinationMetadataV1OutputStatus = "UNAVAILABLE"
)

func (e GetDestinationMetadataV1OutputStatus) ToPointer() *GetDestinationMetadataV1OutputStatus {
	return &e
}

func (e *GetDestinationMetadataV1OutputStatus) UnmarshalJSON(data []byte) error {
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
		*e = GetDestinationMetadataV1OutputStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDestinationMetadataV1OutputStatus: %v", v)
	}
}

// GetDestinationMetadataV1OutputCloudModeInstances - This Destination's support level for cloud mode instances.
// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
type GetDestinationMetadataV1OutputCloudModeInstances string

const (
	GetDestinationMetadataV1OutputCloudModeInstancesZero     GetDestinationMetadataV1OutputCloudModeInstances = "0"
	GetDestinationMetadataV1OutputCloudModeInstancesOne      GetDestinationMetadataV1OutputCloudModeInstances = "1"
	GetDestinationMetadataV1OutputCloudModeInstancesMultiple GetDestinationMetadataV1OutputCloudModeInstances = "MULTIPLE"
	GetDestinationMetadataV1OutputCloudModeInstancesNone     GetDestinationMetadataV1OutputCloudModeInstances = "NONE"
	GetDestinationMetadataV1OutputCloudModeInstancesSingle   GetDestinationMetadataV1OutputCloudModeInstances = "SINGLE"
)

func (e GetDestinationMetadataV1OutputCloudModeInstances) ToPointer() *GetDestinationMetadataV1OutputCloudModeInstances {
	return &e
}

func (e *GetDestinationMetadataV1OutputCloudModeInstances) UnmarshalJSON(data []byte) error {
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
		*e = GetDestinationMetadataV1OutputCloudModeInstances(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDestinationMetadataV1OutputCloudModeInstances: %v", v)
	}
}

// GetDestinationMetadataV1OutputDeviceModeInstances - This Destination's support level for device mode instances.
// Support for multiple device mode instances is currently not planned.
// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
type GetDestinationMetadataV1OutputDeviceModeInstances string

const (
	GetDestinationMetadataV1OutputDeviceModeInstancesZero   GetDestinationMetadataV1OutputDeviceModeInstances = "0"
	GetDestinationMetadataV1OutputDeviceModeInstancesOne    GetDestinationMetadataV1OutputDeviceModeInstances = "1"
	GetDestinationMetadataV1OutputDeviceModeInstancesNone   GetDestinationMetadataV1OutputDeviceModeInstances = "NONE"
	GetDestinationMetadataV1OutputDeviceModeInstancesSingle GetDestinationMetadataV1OutputDeviceModeInstances = "SINGLE"
)

func (e GetDestinationMetadataV1OutputDeviceModeInstances) ToPointer() *GetDestinationMetadataV1OutputDeviceModeInstances {
	return &e
}

func (e *GetDestinationMetadataV1OutputDeviceModeInstances) UnmarshalJSON(data []byte) error {
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
		*e = GetDestinationMetadataV1OutputDeviceModeInstances(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetDestinationMetadataV1OutputDeviceModeInstances: %v", v)
	}
}

// GetDestinationMetadataV1OutputSupportedFeatures - Features that this Destination supports.
//
// Config API note: holds `browserUnbundling` fields.
type GetDestinationMetadataV1OutputSupportedFeatures struct {
	// Whether this Destination supports browser unbundling.
	BrowserUnbundling *bool `json:"browserUnbundling,omitempty"`
	// Whether this Destination supports public browser unbundling.
	BrowserUnbundlingPublic *bool `json:"browserUnbundlingPublic,omitempty"`
	// This Destination's support level for cloud mode instances.
	// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	CloudModeInstances *GetDestinationMetadataV1OutputCloudModeInstances `json:"cloudModeInstances,omitempty"`
	// This Destination's support level for device mode instances.
	// Support for multiple device mode instances is currently not planned.
	// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	DeviceModeInstances *GetDestinationMetadataV1OutputDeviceModeInstances `json:"deviceModeInstances,omitempty"`
	// Whether this Destination supports replays.
	Replay *bool `json:"replay,omitempty"`
}

func (o *GetDestinationMetadataV1OutputSupportedFeatures) GetBrowserUnbundling() *bool {
	if o == nil {
		return nil
	}
	return o.BrowserUnbundling
}

func (o *GetDestinationMetadataV1OutputSupportedFeatures) GetBrowserUnbundlingPublic() *bool {
	if o == nil {
		return nil
	}
	return o.BrowserUnbundlingPublic
}

func (o *GetDestinationMetadataV1OutputSupportedFeatures) GetCloudModeInstances() *GetDestinationMetadataV1OutputCloudModeInstances {
	if o == nil {
		return nil
	}
	return o.CloudModeInstances
}

func (o *GetDestinationMetadataV1OutputSupportedFeatures) GetDeviceModeInstances() *GetDestinationMetadataV1OutputDeviceModeInstances {
	if o == nil {
		return nil
	}
	return o.DeviceModeInstances
}

func (o *GetDestinationMetadataV1OutputSupportedFeatures) GetReplay() *bool {
	if o == nil {
		return nil
	}
	return o.Replay
}

// GetDestinationMetadataV1OutputSupportedMethods - Methods that this Destination supports.
//
// Config API note: equal to `methods`.
type GetDestinationMetadataV1OutputSupportedMethods struct {
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

func (o *GetDestinationMetadataV1OutputSupportedMethods) GetAlias() *bool {
	if o == nil {
		return nil
	}
	return o.Alias
}

func (o *GetDestinationMetadataV1OutputSupportedMethods) GetGroup() *bool {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *GetDestinationMetadataV1OutputSupportedMethods) GetIdentify() *bool {
	if o == nil {
		return nil
	}
	return o.Identify
}

func (o *GetDestinationMetadataV1OutputSupportedMethods) GetPageview() *bool {
	if o == nil {
		return nil
	}
	return o.Pageview
}

func (o *GetDestinationMetadataV1OutputSupportedMethods) GetTrack() *bool {
	if o == nil {
		return nil
	}
	return o.Track
}

// GetDestinationMetadataV1OutputSupportedPlatforms - Platforms from which the Destination receives events.
//
// Config API note: equal to `platforms`.
type GetDestinationMetadataV1OutputSupportedPlatforms struct {
	// Whether this Destination supports browser events.
	Browser *bool `json:"browser,omitempty"`
	// Whether this Destination supports mobile events.
	Mobile *bool `json:"mobile,omitempty"`
	// Whether this Destination supports server events.
	Server *bool `json:"server,omitempty"`
}

func (o *GetDestinationMetadataV1OutputSupportedPlatforms) GetBrowser() *bool {
	if o == nil {
		return nil
	}
	return o.Browser
}

func (o *GetDestinationMetadataV1OutputSupportedPlatforms) GetMobile() *bool {
	if o == nil {
		return nil
	}
	return o.Mobile
}

func (o *GetDestinationMetadataV1OutputSupportedPlatforms) GetServer() *bool {
	if o == nil {
		return nil
	}
	return o.Server
}

// DestinationMetadata - The catalog item matched by id.
type DestinationMetadata struct {
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
	Logos GetDestinationMetadataV1OutputLogos `json:"logos"`
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
	Status GetDestinationMetadataV1OutputStatus `json:"status"`
	// Features that this Destination supports.
	//
	// Config API note: holds `browserUnbundling` fields.
	SupportedFeatures GetDestinationMetadataV1OutputSupportedFeatures `json:"supportedFeatures"`
	// Methods that this Destination supports.
	//
	// Config API note: equal to `methods`.
	SupportedMethods GetDestinationMetadataV1OutputSupportedMethods `json:"supportedMethods"`
	// Platforms from which the Destination receives events.
	//
	// Config API note: equal to `platforms`.
	SupportedPlatforms GetDestinationMetadataV1OutputSupportedPlatforms `json:"supportedPlatforms"`
	// A list of supported regions for this Destination.
	SupportedRegions []string `json:"supportedRegions,omitempty"`
	// A website URL for this Destination.
	Website string `json:"website"`
}

func (o *DestinationMetadata) GetActions() []DestinationMetadataActionV1 {
	if o == nil {
		return []DestinationMetadataActionV1{}
	}
	return o.Actions
}

func (o *DestinationMetadata) GetCategories() []string {
	if o == nil {
		return []string{}
	}
	return o.Categories
}

func (o *DestinationMetadata) GetComponents() []DestinationMetadataComponentV1 {
	if o == nil {
		return []DestinationMetadataComponentV1{}
	}
	return o.Components
}

func (o *DestinationMetadata) GetContacts() []Contact {
	if o == nil {
		return nil
	}
	return o.Contacts
}

func (o *DestinationMetadata) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *DestinationMetadata) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DestinationMetadata) GetLogos() GetDestinationMetadataV1OutputLogos {
	if o == nil {
		return GetDestinationMetadataV1OutputLogos{}
	}
	return o.Logos
}

func (o *DestinationMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *DestinationMetadata) GetOptions() []IntegrationOptionBeta {
	if o == nil {
		return []IntegrationOptionBeta{}
	}
	return o.Options
}

func (o *DestinationMetadata) GetPartnerOwned() *bool {
	if o == nil {
		return nil
	}
	return o.PartnerOwned
}

func (o *DestinationMetadata) GetPresets() []DestinationMetadataSubscriptionPresetV1 {
	if o == nil {
		return []DestinationMetadataSubscriptionPresetV1{}
	}
	return o.Presets
}

func (o *DestinationMetadata) GetPreviousNames() []string {
	if o == nil {
		return []string{}
	}
	return o.PreviousNames
}

func (o *DestinationMetadata) GetRegionEndpoints() []string {
	if o == nil {
		return nil
	}
	return o.RegionEndpoints
}

func (o *DestinationMetadata) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *DestinationMetadata) GetStatus() GetDestinationMetadataV1OutputStatus {
	if o == nil {
		return GetDestinationMetadataV1OutputStatus("")
	}
	return o.Status
}

func (o *DestinationMetadata) GetSupportedFeatures() GetDestinationMetadataV1OutputSupportedFeatures {
	if o == nil {
		return GetDestinationMetadataV1OutputSupportedFeatures{}
	}
	return o.SupportedFeatures
}

func (o *DestinationMetadata) GetSupportedMethods() GetDestinationMetadataV1OutputSupportedMethods {
	if o == nil {
		return GetDestinationMetadataV1OutputSupportedMethods{}
	}
	return o.SupportedMethods
}

func (o *DestinationMetadata) GetSupportedPlatforms() GetDestinationMetadataV1OutputSupportedPlatforms {
	if o == nil {
		return GetDestinationMetadataV1OutputSupportedPlatforms{}
	}
	return o.SupportedPlatforms
}

func (o *DestinationMetadata) GetSupportedRegions() []string {
	if o == nil {
		return nil
	}
	return o.SupportedRegions
}

func (o *DestinationMetadata) GetWebsite() string {
	if o == nil {
		return ""
	}
	return o.Website
}

// GetDestinationMetadataV1Output - Returns a Destination catalog item.
type GetDestinationMetadataV1Output struct {
	// The catalog item matched by id.
	DestinationMetadata DestinationMetadata `json:"destinationMetadata"`
}

func (o *GetDestinationMetadataV1Output) GetDestinationMetadata() DestinationMetadata {
	if o == nil {
		return DestinationMetadata{}
	}
	return o.DestinationMetadata
}
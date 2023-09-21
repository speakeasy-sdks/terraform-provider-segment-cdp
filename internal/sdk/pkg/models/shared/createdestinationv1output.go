// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// CreateDestinationV1OutputDestinationV1DestinationMetadataV1LogosBeta - Represents a logo.
type CreateDestinationV1OutputDestinationV1DestinationMetadataV1LogosBeta struct {
	// The alternative text for this logo.
	Alt *string `json:"alt,omitempty"`
	// The default URL for this logo.
	Default string `json:"default"`
	// The logo mark.
	Mark *string `json:"mark,omitempty"`
}

// CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status - Support status of the Destination.
type CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status string

const (
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1StatusDeprecated       CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status = "DEPRECATED"
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1StatusPrivateBeta      CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status = "PRIVATE_BETA"
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1StatusPrivateBuilding  CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status = "PRIVATE_BUILDING"
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1StatusPrivateSubmitted CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status = "PRIVATE_SUBMITTED"
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1StatusPublic           CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status = "PUBLIC"
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1StatusPublicBeta       CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status = "PUBLIC_BETA"
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1StatusUnavailable      CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status = "UNAVAILABLE"
)

func (e CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status) ToPointer() *CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status {
	return &e
}

func (e *CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status) UnmarshalJSON(data []byte) error {
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
		*e = CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status: %v", v)
	}
}

// CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances - This Destination's support level for cloud mode instances.
// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
type CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances string

const (
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesZero     CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "0"
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesOne      CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "1"
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesMultiple CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "MULTIPLE"
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesNone     CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "NONE"
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstancesSingle   CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances = "SINGLE"
)

func (e CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances) ToPointer() *CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances {
	return &e
}

func (e *CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances) UnmarshalJSON(data []byte) error {
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
		*e = CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances: %v", v)
	}
}

// CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances - This Destination's support level for device mode instances.
// Support for multiple device mode instances is currently not planned.
// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
type CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances string

const (
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesZero   CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "0"
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesOne    CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "1"
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesNone   CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "NONE"
	CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstancesSingle CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances = "SINGLE"
)

func (e CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances) ToPointer() *CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances {
	return &e
}

func (e *CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances) UnmarshalJSON(data []byte) error {
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
		*e = CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances: %v", v)
	}
}

// CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1 - Represents features that a given Destination supports.
type CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1 struct {
	// Whether this Destination supports browser unbundling.
	BrowserUnbundling *bool `json:"browserUnbundling,omitempty"`
	// Whether this Destination supports public browser unbundling.
	BrowserUnbundlingPublic *bool `json:"browserUnbundlingPublic,omitempty"`
	// This Destination's support level for cloud mode instances.
	// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	CloudModeInstances *CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1CloudModeInstances `json:"cloudModeInstances,omitempty"`
	// This Destination's support level for device mode instances.
	// Support for multiple device mode instances is currently not planned.
	// The values '0' and 'NONE', and '1' and 'SINGLE' are equivalent.
	DeviceModeInstances *CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1DeviceModeInstances `json:"deviceModeInstances,omitempty"`
	// Whether this Destination supports replays.
	Replay *bool `json:"replay,omitempty"`
}

// CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataMethodsV1 - Represents methods that a given Destination supports.
type CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataMethodsV1 struct {
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

// CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataPlatformsV1 - Represents platforms that a given Destination supports.
type CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataPlatformsV1 struct {
	// Whether this Destination supports browser events.
	Browser *bool `json:"browser,omitempty"`
	// Whether this Destination supports mobile events.
	Mobile *bool `json:"mobile,omitempty"`
	// Whether this Destination supports server events.
	Server *bool `json:"server,omitempty"`
}

// CreateDestinationV1OutputDestinationV1DestinationMetadataV1 - Represents a Destination within Segment.
//
// A Destination is a target for Segment to forward data to, and represents a tool or storage Destination.
type CreateDestinationV1OutputDestinationV1DestinationMetadataV1 struct {
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
	Logos CreateDestinationV1OutputDestinationV1DestinationMetadataV1LogosBeta `json:"logos"`
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
	Status CreateDestinationV1OutputDestinationV1DestinationMetadataV1Status `json:"status"`
	// Features that this Destination supports.
	//
	// Config API note: holds `browserUnbundling` fields.
	SupportedFeatures CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataFeaturesV1 `json:"supportedFeatures"`
	// Methods that this Destination supports.
	//
	// Config API note: equal to `methods`.
	SupportedMethods CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataMethodsV1 `json:"supportedMethods"`
	// Platforms from which the Destination receives events.
	//
	// Config API note: equal to `platforms`.
	SupportedPlatforms CreateDestinationV1OutputDestinationV1DestinationMetadataV1DestinationMetadataPlatformsV1 `json:"supportedPlatforms"`
	// A list of supported regions for this Destination.
	SupportedRegions []string `json:"supportedRegions,omitempty"`
	// A website URL for this Destination.
	Website string `json:"website"`
}

// CreateDestinationV1OutputDestinationV1 - Business tools or apps that you can connect to the data flowing through Segment.
//
// This is equal to the Destination object in Config API, with the following fields omitted:
// - catalogId
// - createTime
// - updateTime
// - connectionMode.
type CreateDestinationV1OutputDestinationV1 struct {
	// Whether this instance of a Destination receives data.
	Enabled bool `json:"enabled"`
	// The unique identifier of this instance of a Destination.
	//
	// Config API note: analogous to `name`.
	ID string `json:"id"`
	// The metadata of the Destination of which this Destination is an instance of. For example, Google Analytics or Amplitude.
	Metadata CreateDestinationV1OutputDestinationV1DestinationMetadataV1 `json:"metadata"`
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

// CreateDestinationV1Output - Creates a new Destination.
type CreateDestinationV1Output struct {
	// The created Destination.
	Destination CreateDestinationV1OutputDestinationV1 `json:"destination"`
}

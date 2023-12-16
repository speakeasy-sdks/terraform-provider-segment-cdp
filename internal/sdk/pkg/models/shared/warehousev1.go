// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// WarehouseV1Logos - Logo information for this object.
type WarehouseV1Logos struct {
	// The alternative text for this logo.
	Alt *string `json:"alt,omitempty"`
	// The default URL for this logo.
	Default string `json:"default"`
	// The logo mark.
	Mark *string `json:"mark,omitempty"`
}

func (o *WarehouseV1Logos) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

func (o *WarehouseV1Logos) GetDefault() string {
	if o == nil {
		return ""
	}
	return o.Default
}

func (o *WarehouseV1Logos) GetMark() *string {
	if o == nil {
		return nil
	}
	return o.Mark
}

// WarehouseV1Metadata - The metadata for the Warehouse.
type WarehouseV1Metadata struct {
	// A description, in English, of this object.
	Description string `json:"description"`
	// The id of this object.
	ID string `json:"id"`
	// Logo information for this object.
	Logos WarehouseV1Logos `json:"logos"`
	// The name of this object.
	Name string `json:"name"`
	// The Integration options for this object.
	Options []IntegrationOptionBeta `json:"options"`
	// A human-readable, unique identifier for object.
	Slug string `json:"slug"`
}

func (o *WarehouseV1Metadata) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *WarehouseV1Metadata) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *WarehouseV1Metadata) GetLogos() WarehouseV1Logos {
	if o == nil {
		return WarehouseV1Logos{}
	}
	return o.Logos
}

func (o *WarehouseV1Metadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *WarehouseV1Metadata) GetOptions() []IntegrationOptionBeta {
	if o == nil {
		return []IntegrationOptionBeta{}
	}
	return o.Options
}

func (o *WarehouseV1Metadata) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

// WarehouseV1 - Defines a data Warehouse used as a Destination for Segment data.
type WarehouseV1 struct {
	// When set to true, this Warehouse receives data.
	Enabled bool `json:"enabled"`
	// The id of the Warehouse.
	ID string `json:"id"`
	// The metadata for the Warehouse.
	Metadata WarehouseV1Metadata `json:"metadata"`
	// The settings associated with this Warehouse.
	//
	// Common settings are connection-related configuration used to connect to it, for example host, username, and port.
	Settings map[string]interface{} `json:"settings"`
	// The id of the Workspace that owns this Warehouse.
	WorkspaceID string `json:"workspaceId"`
}

func (o *WarehouseV1) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *WarehouseV1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *WarehouseV1) GetMetadata() WarehouseV1Metadata {
	if o == nil {
		return WarehouseV1Metadata{}
	}
	return o.Metadata
}

func (o *WarehouseV1) GetSettings() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Settings
}

func (o *WarehouseV1) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

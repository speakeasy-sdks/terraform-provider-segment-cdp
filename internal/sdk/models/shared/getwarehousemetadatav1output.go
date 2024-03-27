// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetWarehouseMetadataV1OutputLogos - Logo information for this object.
type GetWarehouseMetadataV1OutputLogos struct {
	// The alternative text for this logo.
	Alt *string `json:"alt,omitempty"`
	// The default URL for this logo.
	Default string `json:"default"`
	// The logo mark.
	Mark *string `json:"mark,omitempty"`
}

func (o *GetWarehouseMetadataV1OutputLogos) GetAlt() *string {
	if o == nil {
		return nil
	}
	return o.Alt
}

func (o *GetWarehouseMetadataV1OutputLogos) GetDefault() string {
	if o == nil {
		return ""
	}
	return o.Default
}

func (o *GetWarehouseMetadataV1OutputLogos) GetMark() *string {
	if o == nil {
		return nil
	}
	return o.Mark
}

// WarehouseMetadata - The Warehouse catalog item.
type WarehouseMetadata struct {
	// A description, in English, of this object.
	Description string `json:"description"`
	// The id of this object.
	ID string `json:"id"`
	// Logo information for this object.
	Logos GetWarehouseMetadataV1OutputLogos `json:"logos"`
	// The name of this object.
	Name string `json:"name"`
	// The Integration options for this object.
	Options []IntegrationOptionBeta `json:"options"`
	// A human-readable, unique identifier for object.
	Slug string `json:"slug"`
}

func (o *WarehouseMetadata) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *WarehouseMetadata) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *WarehouseMetadata) GetLogos() GetWarehouseMetadataV1OutputLogos {
	if o == nil {
		return GetWarehouseMetadataV1OutputLogos{}
	}
	return o.Logos
}

func (o *WarehouseMetadata) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *WarehouseMetadata) GetOptions() []IntegrationOptionBeta {
	if o == nil {
		return []IntegrationOptionBeta{}
	}
	return o.Options
}

func (o *WarehouseMetadata) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

// GetWarehouseMetadataV1Output - Returns a Warehouse catalog item looked up by id.
type GetWarehouseMetadataV1Output struct {
	// The Warehouse catalog item.
	WarehouseMetadata WarehouseMetadata `json:"warehouseMetadata"`
}

func (o *GetWarehouseMetadataV1Output) GetWarehouseMetadata() WarehouseMetadata {
	if o == nil {
		return WarehouseMetadata{}
	}
	return o.WarehouseMetadata
}
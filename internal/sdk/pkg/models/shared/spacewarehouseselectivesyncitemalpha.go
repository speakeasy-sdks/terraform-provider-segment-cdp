// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SpaceWarehouseSelectiveSyncItemAlpha - Represents an entry in the Space Warehouse Selective Sync schema for a Warehouse and Space pair.
type SpaceWarehouseSelectiveSyncItemAlpha struct {
	// The collection within the Source.
	Collection string `json:"collection"`
	// The Enabled flag ok telling whether the Collection is enabled or not.
	Enabled bool `json:"enabled"`
	// A map that contains the properties within the collection to which the Warehouse should sync.
	Properties map[string]interface{} `json:"properties"`
	// The id of the Warehouse this sync belongs to.
	WarehouseID string `json:"warehouseId"`
}

func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetCollection() string {
	if o == nil {
		return ""
	}
	return o.Collection
}

func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetProperties() map[string]interface{} {
	if o == nil {
		return map[string]interface{}{}
	}
	return o.Properties
}

func (o *SpaceWarehouseSelectiveSyncItemAlpha) GetWarehouseID() string {
	if o == nil {
		return ""
	}
	return o.WarehouseID
}

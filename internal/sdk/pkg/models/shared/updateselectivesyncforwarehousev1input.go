// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateSelectiveSyncForWarehouseV1Input - Updates the schema for a Warehouse/sources pair.
type UpdateSelectiveSyncForWarehouseV1Input struct {
	// A list of sync schema overrides to apply to this Warehouse.
	SyncOverrides []WarehouseSyncOverrideV1 `json:"syncOverrides"`
}
// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetWarehousesCatalogV1OutputPagination - Information about the pagination of this response.
type GetWarehousesCatalogV1OutputPagination struct {
	// The current cursor within a collection.
	//
	// Consumers of the API must treat this value as opaque.
	Current string `json:"current"`
	// A pointer to the next page.
	//
	// This does not return when you retrieve the last page.
	//
	// Consumers of the API must treat this value as opaque.
	Next *string `json:"next,omitempty"`
	// A pointer to the previous page.
	//
	// This does not return when you retrieve the first page.
	//
	// Consumers of the API must treat this value as opaque.
	Previous *string `json:"previous,omitempty"`
	// The total number of entries available in the collection.
	//
	// If calculating it impacts performance, the response may omit this field.
	TotalEntries *float64 `json:"totalEntries,omitempty"`
}

func (o *GetWarehousesCatalogV1OutputPagination) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *GetWarehousesCatalogV1OutputPagination) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *GetWarehousesCatalogV1OutputPagination) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *GetWarehousesCatalogV1OutputPagination) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// GetWarehousesCatalogV1Output - Returns a list of all Warehouse catalog items contained within a given page.
type GetWarehousesCatalogV1Output struct {
	// Information about the pagination of this response.
	Pagination GetWarehousesCatalogV1OutputPagination `json:"pagination"`
	// All Warehouse catalog items contained within the requested page.
	WarehousesCatalog []WarehouseMetadataV1 `json:"warehousesCatalog"`
}

func (o *GetWarehousesCatalogV1Output) GetPagination() GetWarehousesCatalogV1OutputPagination {
	if o == nil {
		return GetWarehousesCatalogV1OutputPagination{}
	}
	return o.Pagination
}

func (o *GetWarehousesCatalogV1Output) GetWarehousesCatalog() []WarehouseMetadataV1 {
	if o == nil {
		return []WarehouseMetadataV1{}
	}
	return o.WarehousesCatalog
}
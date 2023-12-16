// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListFiltersFromDestinationV1OutputPagination - Information about the pagination of this response.
type ListFiltersFromDestinationV1OutputPagination struct {
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

func (o *ListFiltersFromDestinationV1OutputPagination) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *ListFiltersFromDestinationV1OutputPagination) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListFiltersFromDestinationV1OutputPagination) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ListFiltersFromDestinationV1OutputPagination) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// ListFiltersFromDestinationV1Output - Output for ListDestinationFiltersV1.
type ListFiltersFromDestinationV1Output struct {
	// A list of the filters that belong to the specified Destination instance.
	Filters []DestinationFilterV1 `json:"filters"`
	// Information about the pagination of this response.
	Pagination ListFiltersFromDestinationV1OutputPagination `json:"pagination"`
}

func (o *ListFiltersFromDestinationV1Output) GetFilters() []DestinationFilterV1 {
	if o == nil {
		return []DestinationFilterV1{}
	}
	return o.Filters
}

func (o *ListFiltersFromDestinationV1Output) GetPagination() ListFiltersFromDestinationV1OutputPagination {
	if o == nil {
		return ListFiltersFromDestinationV1OutputPagination{}
	}
	return o.Pagination
}

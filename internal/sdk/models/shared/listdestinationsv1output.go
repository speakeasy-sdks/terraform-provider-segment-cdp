// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListDestinationsV1OutputPagination - Information about the pagination of this response.
type ListDestinationsV1OutputPagination struct {
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

func (o *ListDestinationsV1OutputPagination) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *ListDestinationsV1OutputPagination) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListDestinationsV1OutputPagination) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ListDestinationsV1OutputPagination) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// ListDestinationsV1Output - Returns all Destinations connected to the current Workspace.
type ListDestinationsV1Output struct {
	// The list that contains the Destinations connected to the Workspace.
	Destinations []DestinationV1 `json:"destinations"`
	// Information about the pagination of this response.
	Pagination ListDestinationsV1OutputPagination `json:"pagination"`
}

func (o *ListDestinationsV1Output) GetDestinations() []DestinationV1 {
	if o == nil {
		return []DestinationV1{}
	}
	return o.Destinations
}

func (o *ListDestinationsV1Output) GetPagination() ListDestinationsV1OutputPagination {
	if o == nil {
		return ListDestinationsV1OutputPagination{}
	}
	return o.Pagination
}
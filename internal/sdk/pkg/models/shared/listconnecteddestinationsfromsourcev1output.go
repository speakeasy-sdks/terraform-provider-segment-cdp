// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListConnectedDestinationsFromSourceV1OutputPagination - Information about the pagination of this response.
type ListConnectedDestinationsFromSourceV1OutputPagination struct {
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

func (o *ListConnectedDestinationsFromSourceV1OutputPagination) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *ListConnectedDestinationsFromSourceV1OutputPagination) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListConnectedDestinationsFromSourceV1OutputPagination) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ListConnectedDestinationsFromSourceV1OutputPagination) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// ListConnectedDestinationsFromSourceV1Output - Returns a list of Destinations connected to a Source.
type ListConnectedDestinationsFromSourceV1Output struct {
	// A list that contains the Destinations connected to the Source.
	Destinations []DestinationV1 `json:"destinations"`
	// Information about the pagination of this response.
	Pagination ListConnectedDestinationsFromSourceV1OutputPagination `json:"pagination"`
}

func (o *ListConnectedDestinationsFromSourceV1Output) GetDestinations() []DestinationV1 {
	if o == nil {
		return []DestinationV1{}
	}
	return o.Destinations
}

func (o *ListConnectedDestinationsFromSourceV1Output) GetPagination() ListConnectedDestinationsFromSourceV1OutputPagination {
	if o == nil {
		return ListConnectedDestinationsFromSourceV1OutputPagination{}
	}
	return o.Pagination
}

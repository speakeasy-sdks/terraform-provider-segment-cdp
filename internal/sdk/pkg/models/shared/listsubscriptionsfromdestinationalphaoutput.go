// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListSubscriptionsFromDestinationAlphaOutputPagination - Information about the pagination of this response.
type ListSubscriptionsFromDestinationAlphaOutputPagination struct {
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

func (o *ListSubscriptionsFromDestinationAlphaOutputPagination) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *ListSubscriptionsFromDestinationAlphaOutputPagination) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListSubscriptionsFromDestinationAlphaOutputPagination) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ListSubscriptionsFromDestinationAlphaOutputPagination) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// ListSubscriptionsFromDestinationAlphaOutput - Output for ListDestinationSubscriptionsAlpha.
type ListSubscriptionsFromDestinationAlphaOutput struct {
	// Information about the pagination of this response.
	Pagination *ListSubscriptionsFromDestinationAlphaOutputPagination `json:"pagination,omitempty"`
	// A list of Destination subscriptions.
	Subscriptions []DestinationSubscription `json:"subscriptions"`
}

func (o *ListSubscriptionsFromDestinationAlphaOutput) GetPagination() *ListSubscriptionsFromDestinationAlphaOutputPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

func (o *ListSubscriptionsFromDestinationAlphaOutput) GetSubscriptions() []DestinationSubscription {
	if o == nil {
		return []DestinationSubscription{}
	}
	return o.Subscriptions
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListRulesFromTrackingPlanV1OutputPagination - Information about the pagination of this response.
type ListRulesFromTrackingPlanV1OutputPagination struct {
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

func (o *ListRulesFromTrackingPlanV1OutputPagination) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *ListRulesFromTrackingPlanV1OutputPagination) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListRulesFromTrackingPlanV1OutputPagination) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ListRulesFromTrackingPlanV1OutputPagination) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// ListRulesFromTrackingPlanV1Output - Lists a Tracking Plan's rules.
type ListRulesFromTrackingPlanV1Output struct {
	// Information about the pagination of this response.
	Pagination ListRulesFromTrackingPlanV1OutputPagination `json:"pagination"`
	// Rules associated with the given Tracking Plan.
	Rules []RuleV1 `json:"rules"`
}

func (o *ListRulesFromTrackingPlanV1Output) GetPagination() ListRulesFromTrackingPlanV1OutputPagination {
	if o == nil {
		return ListRulesFromTrackingPlanV1OutputPagination{}
	}
	return o.Pagination
}

func (o *ListRulesFromTrackingPlanV1Output) GetRules() []RuleV1 {
	if o == nil {
		return []RuleV1{}
	}
	return o.Rules
}

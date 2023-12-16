// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListReverseEtlModelsOutputPagination - Information about the pagination of this response.
type ListReverseEtlModelsOutputPagination struct {
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

func (o *ListReverseEtlModelsOutputPagination) GetCurrent() string {
	if o == nil {
		return ""
	}
	return o.Current
}

func (o *ListReverseEtlModelsOutputPagination) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *ListReverseEtlModelsOutputPagination) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *ListReverseEtlModelsOutputPagination) GetTotalEntries() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalEntries
}

// ListReverseEtlModelsOutput - Defines the result of listing Models.
type ListReverseEtlModelsOutput struct {
	// A list of Models that belong to the Workspace.
	Models []ReverseEtlModel `json:"models"`
	// Information about the pagination of this response.
	Pagination ListReverseEtlModelsOutputPagination `json:"pagination"`
}

func (o *ListReverseEtlModelsOutput) GetModels() []ReverseEtlModel {
	if o == nil {
		return []ReverseEtlModel{}
	}
	return o.Models
}

func (o *ListReverseEtlModelsOutput) GetPagination() ListReverseEtlModelsOutputPagination {
	if o == nil {
		return ListReverseEtlModelsOutputPagination{}
	}
	return o.Pagination
}

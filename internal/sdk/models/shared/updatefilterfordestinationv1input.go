// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateFilterForDestinationV1Input - Input for UpdateDestinationFilterV1.
type UpdateFilterForDestinationV1Input struct {
	// Actions for this Destination filter.
	Actions []DestinationFilterActionV1 `json:"actions,omitempty"`
	// The description of this filter.
	Description *string `json:"description,omitempty"`
	// When set to true, this Destination filter is active.
	Enabled *bool `json:"enabled,omitempty"`
	// The FQL if condition to update.
	If *string `json:"if,omitempty"`
	// The title to update.
	Title *string `json:"title,omitempty"`
}

func (o *UpdateFilterForDestinationV1Input) GetActions() []DestinationFilterActionV1 {
	if o == nil {
		return nil
	}
	return o.Actions
}

func (o *UpdateFilterForDestinationV1Input) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *UpdateFilterForDestinationV1Input) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *UpdateFilterForDestinationV1Input) GetIf() *string {
	if o == nil {
		return nil
	}
	return o.If
}

func (o *UpdateFilterForDestinationV1Input) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}
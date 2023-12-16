// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ComputedTrait - The computed trait summary output.
type ComputedTrait struct {
	// The timestamp of the computed trait's creation.
	CreatedAt string `json:"createdAt"`
	// User id who created the computed trait.
	CreatedBy string `json:"createdBy"`
	// Description of the computed trait.
	Description string `json:"description"`
	// Enabled/disabled status for the computed trait.
	Enabled bool `json:"enabled"`
	// Computed trait id.
	ID string `json:"id"`
	// Key for the computed trait.
	Key string `json:"key"`
	// Name of the computed trait.
	Name string `json:"name"`
	// Space id for the computed trait.
	SpaceID string `json:"spaceId"`
	// The timestamp of the computed trait's last change.
	UpdatedAt string `json:"updatedAt"`
	// User id who last updated the computed trait.
	UpdatedBy string `json:"updatedBy"`
}

func (o *ComputedTrait) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *ComputedTrait) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *ComputedTrait) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *ComputedTrait) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *ComputedTrait) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ComputedTrait) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *ComputedTrait) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ComputedTrait) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

func (o *ComputedTrait) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *ComputedTrait) GetUpdatedBy() string {
	if o == nil {
		return ""
	}
	return o.UpdatedBy
}

// GetComputedTraitAlphaOutput - Computed Trait output for get and update.
type GetComputedTraitAlphaOutput struct {
	// The computed trait summary output.
	ComputedTrait ComputedTrait `json:"computedTrait"`
}

func (o *GetComputedTraitAlphaOutput) GetComputedTrait() ComputedTrait {
	if o == nil {
		return ComputedTrait{}
	}
	return o.ComputedTrait
}

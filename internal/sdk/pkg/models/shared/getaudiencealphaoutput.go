// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Audience - The audience summary output.
type Audience struct {
	// Date the audience was created.
	CreatedAt string `json:"createdAt"`
	// User id who created the audience.
	CreatedBy string `json:"createdBy"`
	// Description of the audience.
	Description string `json:"description"`
	// Enabled/disabled status for the audience.
	Enabled bool `json:"enabled"`
	// Audience id.
	ID string `json:"id"`
	// Key for the audience.
	Key string `json:"key"`
	// Name of the audience.
	Name string `json:"name"`
	// Space id for the audience.
	SpaceID string `json:"spaceId"`
	// Date the audience was last updated.
	UpdatedAt string `json:"updatedAt"`
	// User id who last updated the audience.
	UpdatedBy string `json:"updatedBy"`
}

func (o *Audience) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *Audience) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *Audience) GetDescription() string {
	if o == nil {
		return ""
	}
	return o.Description
}

func (o *Audience) GetEnabled() bool {
	if o == nil {
		return false
	}
	return o.Enabled
}

func (o *Audience) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Audience) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *Audience) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Audience) GetSpaceID() string {
	if o == nil {
		return ""
	}
	return o.SpaceID
}

func (o *Audience) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *Audience) GetUpdatedBy() string {
	if o == nil {
		return ""
	}
	return o.UpdatedBy
}

// GetAudienceAlphaOutput - Audience output for update.
type GetAudienceAlphaOutput struct {
	// The audience summary output.
	Audience Audience `json:"audience"`
}

func (o *GetAudienceAlphaOutput) GetAudience() Audience {
	if o == nil {
		return Audience{}
	}
	return o.Audience
}

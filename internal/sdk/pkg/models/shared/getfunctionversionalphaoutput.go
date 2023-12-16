// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetFunctionVersionAlphaOutputVersion - Functions version.
type GetFunctionVersionAlphaOutputVersion struct {
	// Author of this version.
	Author *string `json:"author,omitempty"`
	// Source code of this version.
	Code string `json:"code"`
	// The time of this Version's creation.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The time of this Version's last deployment.
	DeployedAt *string `json:"deployedAt,omitempty"`
	// An identifier for this version.
	ID string `json:"id"`
	// The deployed status of this version.
	IsDeployed *bool `json:"isDeployed,omitempty"`
	// The time of this Version's latest update.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

func (o *GetFunctionVersionAlphaOutputVersion) GetAuthor() *string {
	if o == nil {
		return nil
	}
	return o.Author
}

func (o *GetFunctionVersionAlphaOutputVersion) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *GetFunctionVersionAlphaOutputVersion) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GetFunctionVersionAlphaOutputVersion) GetDeployedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeployedAt
}

func (o *GetFunctionVersionAlphaOutputVersion) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetFunctionVersionAlphaOutputVersion) GetIsDeployed() *bool {
	if o == nil {
		return nil
	}
	return o.IsDeployed
}

func (o *GetFunctionVersionAlphaOutputVersion) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// GetFunctionVersionAlphaOutput - Get Function version output.
type GetFunctionVersionAlphaOutput struct {
	// Functions version.
	Version GetFunctionVersionAlphaOutputVersion `json:"version"`
}

func (o *GetFunctionVersionAlphaOutput) GetVersion() GetFunctionVersionAlphaOutputVersion {
	if o == nil {
		return GetFunctionVersionAlphaOutputVersion{}
	}
	return o.Version
}

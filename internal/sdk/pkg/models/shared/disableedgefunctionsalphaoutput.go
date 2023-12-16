// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DisableEdgeFunctionsAlphaOutputEdgeFunctions - The latest version of Edge Function bundle.
type DisableEdgeFunctionsAlphaOutputEdgeFunctions struct {
	// Creation date.
	CreatedAt string `json:"createdAt"`
	// Creating user's id.
	CreatedBy string `json:"createdBy"`
	// The CDN URL that can be used to fetch your latest EdgeFunctions bundle.
	DownloadURL string `json:"downloadURL"`
	// The Edge Function id.
	ID string `json:"id"`
	// The Source id.
	SourceID string `json:"sourceId"`
	// Revision number associated with the latest Edge Function.
	Version float64 `json:"version"`
}

func (o *DisableEdgeFunctionsAlphaOutputEdgeFunctions) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *DisableEdgeFunctionsAlphaOutputEdgeFunctions) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *DisableEdgeFunctionsAlphaOutputEdgeFunctions) GetDownloadURL() string {
	if o == nil {
		return ""
	}
	return o.DownloadURL
}

func (o *DisableEdgeFunctionsAlphaOutputEdgeFunctions) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DisableEdgeFunctionsAlphaOutputEdgeFunctions) GetSourceID() string {
	if o == nil {
		return ""
	}
	return o.SourceID
}

func (o *DisableEdgeFunctionsAlphaOutputEdgeFunctions) GetVersion() float64 {
	if o == nil {
		return 0.0
	}
	return o.Version
}

// DisableEdgeFunctionsAlphaOutput - Output for DisableEdgeFunctions.
type DisableEdgeFunctionsAlphaOutput struct {
	// The latest version of Edge Function bundle.
	EdgeFunctions DisableEdgeFunctionsAlphaOutputEdgeFunctions `json:"edgeFunctions"`
}

func (o *DisableEdgeFunctionsAlphaOutput) GetEdgeFunctions() DisableEdgeFunctionsAlphaOutputEdgeFunctions {
	if o == nil {
		return DisableEdgeFunctionsAlphaOutputEdgeFunctions{}
	}
	return o.EdgeFunctions
}

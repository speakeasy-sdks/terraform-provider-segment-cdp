// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateEdgeFunctionsAlphaOutputEdgeFunctionsAlpha - Represents an Edge Function bundle.
type CreateEdgeFunctionsAlphaOutputEdgeFunctionsAlpha struct {
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

// CreateEdgeFunctionsAlphaOutput - Output for CreateEdgeFunctions.
type CreateEdgeFunctionsAlphaOutput struct {
	// The created Edge Function.
	EdgeFunctions CreateEdgeFunctionsAlphaOutputEdgeFunctionsAlpha `json:"edgeFunctions"`
}

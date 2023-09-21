// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateInsertFunctionInstanceAlphaInput - Creates an insert Function instance.
type CreateInsertFunctionInstanceAlphaInput struct {
	// Whether this insert Function instance should be enabled for the Destination.
	Enabled *bool `json:"enabled,omitempty"`
	// Insert Function id to which this instance is associated.
	FunctionID string `json:"functionId"`
	// The Source or Destination id to be connected.
	IntegrationID string `json:"integrationId"`
	// Defines the display name of the insert Function instance.
	Name string `json:"name"`
	// An object that contains settings for this insert Function instance based on the settings present in the
	// insert Function class.
	Settings map[string]interface{} `json:"settings"`
}

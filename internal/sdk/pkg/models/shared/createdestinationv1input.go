// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateDestinationV1Input - Creates a new Destination.
type CreateDestinationV1Input struct {
	// Whether this Destination should receive data.
	Enabled *bool `json:"enabled,omitempty"`
	// The id of the metadata to link to the new Destination.
	MetadataID string `json:"metadataId"`
	// Defines the display name of the Destination.
	//
	// Config API note: equal to `displayName`.
	Name *string `json:"name,omitempty"`
	// An object that contains settings for the Destination based on the "required" and "advanced" settings present in the
	// Destination metadata.
	//
	// Config API note: equal to `config`.
	Settings map[string]interface{} `json:"settings"`
	// The id of the Source to connect to this Destination instance.
	//
	// Config API note: analogous to `parent`.
	SourceID string `json:"sourceId"`
}
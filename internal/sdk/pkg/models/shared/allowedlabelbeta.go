// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// AllowedLabelBeta - Defines a label that you may apply to resources within a Workspace.
type AllowedLabelBeta struct {
	// A description of what this label represents.
	Description *string `json:"description,omitempty"`
	// The key identifier for this label.
	Key string `json:"key"`
	// The value of this label.
	Value string `json:"value"`
}

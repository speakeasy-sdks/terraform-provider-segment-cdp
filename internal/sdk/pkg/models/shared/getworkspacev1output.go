// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GetWorkspaceV1OutputWorkspaceV1 - An organized group of Sources and Destinations managed by a team.
type GetWorkspaceV1OutputWorkspaceV1 struct {
	// The unique identifier.
	ID string `json:"id"`
	// The human-readable name.
	Name string `json:"name"`
	// The URL-friendly slug.
	Slug string `json:"slug"`
}

// GetWorkspaceV1Output - Represents the output of loading the Workspace.
type GetWorkspaceV1Output struct {
	// The Workspace.
	Workspace GetWorkspaceV1OutputWorkspaceV1 `json:"workspace"`
}
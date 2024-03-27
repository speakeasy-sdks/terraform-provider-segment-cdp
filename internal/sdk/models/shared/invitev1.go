// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// InviteV1 - Defines an invitation to join a Workspace.
type InviteV1 struct {
	// The invited user's email to attach the permissions to.
	Email string `json:"email"`
	// The permissions to attach to the invited user.
	Permissions []InvitePermissionV1 `json:"permissions,omitempty"`
}

func (o *InviteV1) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *InviteV1) GetPermissions() []InvitePermissionV1 {
	if o == nil {
		return nil
	}
	return o.Permissions
}
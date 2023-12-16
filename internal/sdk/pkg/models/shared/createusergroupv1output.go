// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateUserGroupV1OutputUserGroup - The newly created user group.
type CreateUserGroupV1OutputUserGroup struct {
	// The id of the user group.
	ID string `json:"id"`
	// The number of members in the user group.
	MemberCount float64 `json:"memberCount"`
	// The name of the user group.
	Name string `json:"name"`
	// The permissions associated with the user group.
	Permissions []PermissionV1 `json:"permissions,omitempty"`
}

func (o *CreateUserGroupV1OutputUserGroup) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateUserGroupV1OutputUserGroup) GetMemberCount() float64 {
	if o == nil {
		return 0.0
	}
	return o.MemberCount
}

func (o *CreateUserGroupV1OutputUserGroup) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateUserGroupV1OutputUserGroup) GetPermissions() []PermissionV1 {
	if o == nil {
		return nil
	}
	return o.Permissions
}

// CreateUserGroupV1Output - Returns the newly created user group.
type CreateUserGroupV1Output struct {
	// The newly created user group.
	UserGroup CreateUserGroupV1OutputUserGroup `json:"userGroup"`
}

func (o *CreateUserGroupV1Output) GetUserGroup() CreateUserGroupV1OutputUserGroup {
	if o == nil {
		return CreateUserGroupV1OutputUserGroup{}
	}
	return o.UserGroup
}

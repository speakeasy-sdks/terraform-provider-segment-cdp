// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Contact - The contact info for Integration Owners.
type Contact struct {
	// Email of this contact.
	Email string `json:"email"`
	// Whether this is a primary contact.
	IsPrimary *bool `json:"isPrimary,omitempty"`
	// Name of this contact.
	Name *string `json:"name,omitempty"`
	// Role of this contact.
	Role *string `json:"role,omitempty"`
}

func (o *Contact) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *Contact) GetIsPrimary() *bool {
	if o == nil {
		return nil
	}
	return o.IsPrimary
}

func (o *Contact) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Contact) GetRole() *string {
	if o == nil {
		return nil
	}
	return o.Role
}

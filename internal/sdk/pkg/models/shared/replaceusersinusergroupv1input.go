// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ReplaceUsersInUserGroupV1Input - Replace a user group's list of users and invites with a new one.
type ReplaceUsersInUserGroupV1Input struct {
	// The email addresses of the users and invites to replace.
	Emails []string `json:"emails"`
}
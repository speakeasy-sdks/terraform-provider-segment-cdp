// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Security struct {
	Token string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}
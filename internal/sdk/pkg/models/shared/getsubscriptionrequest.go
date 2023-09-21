// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// GetSubscriptionRequestType - Type is communication medium used.
type GetSubscriptionRequestType string

const (
	GetSubscriptionRequestTypeAndroidPush GetSubscriptionRequestType = "ANDROID_PUSH"
	GetSubscriptionRequestTypeEmail       GetSubscriptionRequestType = "EMAIL"
	GetSubscriptionRequestTypeIosPush     GetSubscriptionRequestType = "IOS_PUSH"
	GetSubscriptionRequestTypeSms         GetSubscriptionRequestType = "SMS"
	GetSubscriptionRequestTypeWhatsapp    GetSubscriptionRequestType = "WHATSAPP"
)

func (e GetSubscriptionRequestType) ToPointer() *GetSubscriptionRequestType {
	return &e
}

func (e *GetSubscriptionRequestType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ANDROID_PUSH":
		fallthrough
	case "EMAIL":
		fallthrough
	case "IOS_PUSH":
		fallthrough
	case "SMS":
		fallthrough
	case "WHATSAPP":
		*e = GetSubscriptionRequestType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetSubscriptionRequestType: %v", v)
	}
}

type GetSubscriptionRequest struct {
	// Key is the phone number or email.
	Key string `json:"key"`
	// Type is communication medium used.
	Type GetSubscriptionRequestType `json:"type"`
}
